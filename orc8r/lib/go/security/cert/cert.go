/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// Package cert implements some common tools for certification related functionalities
package cert

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
)

// load certificate and private key given certificate file and key file
func LoadCertAndPrivKey(certFile, keyFile string) (cert *x509.Certificate, privKey interface{}, err error) {
	tlsCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		err = fmt.Errorf("Failed to load certificate (%s) & key (%s): %s\n",
			certFile, keyFile, err)
		return
	}
	cert, err = x509.ParseCertificate(tlsCert.Certificate[0])
	if err != nil {
		err = fmt.Errorf("Failed to parse cert (%s): %s\n", certFile, err)
		return
	}
	privKey = tlsCert.PrivateKey
	return
}

// SerialToString converts big.Int to hexadecimal string with uppercace letters
// (A,B,C,D,E,F), without base prefix ("0x") and without leading zeros
func SerialToString(certSerialNumber *big.Int) string {
	return strings.ToUpper(certSerialNumber.Text(16))
}

// LoadCert loads, parses and returns certificate from a given file
func LoadCert(certFile string) (*x509.Certificate, error) {
	certPEMBlock, err := ioutil.ReadFile(certFile)
	if err != nil {
		return nil, fmt.Errorf("Cannot read certificate file '%s': %s", certFile, err)
	}
	for {
		var certDERBlock *pem.Block
		certDERBlock, certPEMBlock = pem.Decode(certPEMBlock)
		if certDERBlock == nil {
			break
		}
		if certDERBlock.Type == "CERTIFICATE" {
			x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
			if err != nil {
				return nil, fmt.Errorf("Cannot parse certificate from '%s': %s", certFile, err)
			}
			return x509Cert, nil
		}
	}
	return nil, fmt.Errorf("No Valid certificates found in %s", certFile)
}
