/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package generate

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"golang.org/x/tools/go/ast/astutil"
)

// RewriteGeneratedRefs rewrites the Go files generated by go-swagger by
// updating the type identifiers for generated types owned by dependency
// swagger specs to the packages where those types are generated into.
//
// This function will add appropriate imports and rewrite all references.
// After rewriting the generated files, all files for types that aren't owned
// by the target swagger spec will be removed from the output directory as
// well.
func RewriteGeneratedRefs(targetFilepath string, rootDir string) error {
	absTargetFilepath, err := filepath.Abs(targetFilepath)
	if err != nil {
		return errors.Wrapf(err, "target filepath %s is invalid", targetFilepath)
	}
	allConfigs, err := ParseSwaggerDependencyTree(targetFilepath, rootDir)
	if err != nil {
		return errors.WithStack(err)
	}
	targetConfig := allConfigs[absTargetFilepath]
	targetOutputDir, err := filepath.Abs(filepath.Join(
		rootDir,
		targetConfig.MagmaGenMeta.OutputDir,
		"models", // go-swagger puts all model definitions under a `models` subpackage
	))
	if err != nil {
		return errors.Wrap(err, "could not compute output dir")
	}

	// Gather owned types: map from ident to package import
	// (name and full path) and filename
	filesToRewrite, err := getFilesToRewrite(targetConfig, targetOutputDir)
	if err != nil {
		return errors.Wrapf(err, "failed to rewrite generated swagger bindings for %s", absTargetFilepath)
	}
	validDependencyTypes := gatherAllValidDependentTypes(absTargetFilepath, allConfigs)

	// rewrite all generated files
	for _, filename := range filesToRewrite {
		err = rewriteGeneratedModelBinding(filename, validDependencyTypes)
		if err != nil {
			return errors.Wrapf(err, "failed to rewrite generated file at %s", filename)
		}
	}

	// delete all files for models that aren't owned by the target swagger spec
	for _, dependency := range validDependencyTypes {
		// ignore errors since not all dependent types are guaranteed to have
		// been generated (only those referenced)
		_ = os.Remove(filepath.Join(targetOutputDir, dependency.filename))
	}
	return nil
}

func getFilesToRewrite(targetConfig MagmaSwaggerConfig, outputDir string) ([]string, error) {
	// this depends on go-swagger generating exactly 1 type into each file
	filesToRewrite := make([]string, 0, len(targetConfig.MagmaGenMeta.Types))
	for _, typeSpec := range targetConfig.MagmaGenMeta.Types {
		filesToRewrite = append(filesToRewrite, filepath.Join(outputDir, typeSpec.Filename))
	}
	return filesToRewrite, nil
}

type swaggerTypeDependency struct {
	goPackage     string
	filename      string
	packageNumber int
}

func gatherAllValidDependentTypes(absTargetFilepath string, allConfigs map[string]MagmaSwaggerConfig) map[string]swaggerTypeDependency {
	ret := map[string]swaggerTypeDependency{}

	sortedPaths := funk.Keys(allConfigs).([]string)
	sort.Strings(sortedPaths)
	for i, path := range sortedPaths {
		if path == absTargetFilepath {
			continue
		}

		for _, t := range allConfigs[path].MagmaGenMeta.Types {
			ret[t.GoStructName] = swaggerTypeDependency{
				goPackage:     allConfigs[path].MagmaGenMeta.GoPackage,
				filename:      t.Filename,
				packageNumber: i + 1,
			}
		}
	}
	return ret
}

func rewriteGeneratedModelBinding(filename string, dependentTypes map[string]swaggerTypeDependency) error {
	fset := token.NewFileSet()
	fileNode, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "could not parse Go file")
	}

	// find the idents to replace, add imports as needed, rewrite the idents
	targetIdents := findIdentsToChange(fset, fileNode, dependentTypes)
	updateImports(fset, fileNode, targetIdents, dependentTypes)
	updatedFileNode := updateIdents(fileNode, targetIdents, dependentTypes)

	err = writeFinalFile(filename, fset, updatedFileNode)
	if err != nil {
		return errors.Wrap(err, "could not write result")
	}

	return nil
}

func findIdentsToChange(fset *token.FileSet, fileNode ast.Node, validDependentTypes map[string]swaggerTypeDependency) map[*ast.Ident]bool {
	// map keys are intentionally pointer types (the AST doesn't change across
	// traversals)
	ret := map[*ast.Ident]bool{}
	ast.Inspect(fileNode, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.TypeSpec:
			typeSpec := n.(*ast.TypeSpec)
			switch typeSpec.Type.(type) {
			case *ast.StructType:
				for _, field := range typeSpec.Type.(*ast.StructType).Fields.List {
					ident := findLeafIdent(fset, field)
					if ident == nil {
						break
					}
					if _, shouldReplace := validDependentTypes[ident.Name]; shouldReplace {
						ret[ident] = true
					}
				}
			case *ast.ArrayType:
				ident := findLeafIdent(fset, typeSpec.Type)
				if ident == nil {
					break
				}
				if _, shouldReplace := validDependentTypes[ident.Name]; shouldReplace {
					ret[ident] = true
				}
			}
		}
		return true
	})
	return ret
}

func findLeafIdent(fset *token.FileSet, n ast.Node) *ast.Ident {
	// We have a simple assumption that go-swagger will not generate types that
	// will be referenced by a selector expression (i.e. all swagger-defined
	// types are local to the package that models are generated into).
	switch t := n.(type) {
	case *ast.Ident:
		return t
	case *ast.Field:
		return findLeafIdent(fset, t.Type)
	case *ast.StarExpr:
		return findLeafIdent(fset, t.X)
	case *ast.ArrayType:
		return findLeafIdent(fset, t.Elt)
	case *ast.MapType:
		// again making a quick hacky assumption that generated types won't be
		// used as map keys
		return findLeafIdent(fset, t.Value)
	case *ast.SelectorExpr, *ast.InterfaceType:
		return nil
	default:
		panic(fmt.Sprintf("Unsupported AST type %T; implement support yourself!", t))
	}
}

func updateImports(fset *token.FileSet, fileNode *ast.File, targetIdents map[*ast.Ident]bool, validDependentTypes map[string]swaggerTypeDependency) {
	sortedIdents := funk.Keys(targetIdents).([]*ast.Ident)
	sort.Slice(sortedIdents, func(i, j int) bool { return sortedIdents[i].Name < sortedIdents[j].Name })
	for _, ident := range sortedIdents {
		dependency := validDependentTypes[ident.Name]
		importName := getUniqueNameForImport(dependency)
		astutil.AddNamedImport(fset, fileNode, importName, dependency.goPackage)
	}
}

func getUniqueNameForImport(dependency swaggerTypeDependency) string {
	return fmt.Sprintf("models%d", dependency.packageNumber)
}

func updateIdents(fileNode *ast.File, targetIdents map[*ast.Ident]bool, validDependentTypes map[string]swaggerTypeDependency) ast.Node {
	return astutil.Apply(
		fileNode,
		func(c *astutil.Cursor) bool {
			switch n := c.Node().(type) {
			case *ast.Ident:
				_, found := targetIdents[n]
				if found {
					actualDependency := validDependentTypes[n.Name]
					c.Replace(
						// The right way to do this is to replace the node with
						// *ast.SelectorExpr, but this is easy and it works
						ast.NewIdent(
							fmt.Sprintf(
								"%s.%s",
								getUniqueNameForImport(actualDependency),
								n.Name,
							),
						),
					)
				}
			}
			return true
		},
		nil)
}

func writeFinalFile(filename string, fset *token.FileSet, fileNode ast.Node) error {
	// we need to write the resulting AST twice due to a latent bug in the
	// ast library - https://github.com/golang/go/issues/23771
	intermediateBuffer := strings.Builder{}
	err := printer.Fprint(&intermediateBuffer, fset, fileNode)
	if err != nil {
		return errors.Wrap(err, "failed to write AST to intermediate buffer")
	}
	intermediateOutput := intermediateBuffer.String()
	intermediateFset := token.NewFileSet()
	intermediateFileNode, err := parser.ParseFile(intermediateFset, filename, intermediateOutput, parser.ParseComments)
	if err != nil {
		return errors.Wrap(err, "failed to parse intermediate output")
	}

	// Fprint again, then go fmt the code
	outputBuffer := &bytes.Buffer{}
	err = printer.Fprint(outputBuffer, intermediateFset, intermediateFileNode)
	if err != nil {
		return errors.Wrap(err, "failed to write final AST to output file")
	}
	formattedOutput, err := format.Source(outputBuffer.Bytes())
	if err != nil {
		return errors.Wrap(err, "failed to gofmt final source code")
	}
	// write out the modified file
	err = ioutil.WriteFile(filename, formattedOutput, 0664)
	if err != nil {
		return errors.Wrap(err, "failed to write formatted source code to output file")
	}

	return nil
}
