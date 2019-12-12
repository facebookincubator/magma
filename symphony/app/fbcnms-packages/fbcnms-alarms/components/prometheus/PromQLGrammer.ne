@{%
const {lexer} = require('../PromQLTokenizer');
const {AggregationOperation, BinaryOperation, Clause, Function, InstantSelector, Label, Labels, RangeSelector, Scalar} = require('../PromQL');
%}

@lexer lexer

EXPRESSION -> METRIC_SELECTOR {% id %}
            | FUNCTION {% id %}
            | AGGREGATION {% id %}
            | BINARY_OPERATION {% id %}
            | SCALAR {% id %}

METRIC_SELECTOR -> INSTANT_SELECTOR {% id %}
                 | RANGE_SELECTOR {% id %}

INSTANT_SELECTOR -> IDENTIFIER label_selector {% ([id, labels]) => new InstantSelector(id, labels) %}
                  | IDENTIFIER {% ([id]) => new InstantSelector(id) %}
                  | label_selector {% ([labels]) => new InstantSelector("", labels) %}

IDENTIFIER -> %word {% ([id]) => id.text %}

RANGE_SELECTOR -> INSTANT_SELECTOR duration {% ([selector, duration]) => new RangeSelector(selector, duration)%}

BINARY_OPERATION -> EXPRESSION bin_op EXPRESSION {% ([lh, op, rh]) => new BinaryOperation(lh, rh, op) %}

FUNCTION -> func_name %lParen func_params %rParen {% ([funcName, _lParen, params, _rParen]) => new Function(funcName, params) %}

AGGREGATION -> agg_op %lParen func_params %rParen {% ([aggOp, _lParen, params, _rParen]) => new AggregationOperation(aggOp, params) %}
             | agg_op %lParen func_params %rParen dimensionClause {% ([aggOp, _lParen, params, _rParen, clause]) => new AggregationOperation(aggOp, params, clause) %}

dimensionClause -> clause_op labelList {% ([op, labelList]) => new Clause(op, labelList) %}

labelList -> %lParen label_name_list %rParen {% ([_lParen, labels, _rParen]) => labels %}

label_name_list -> label_name_list %comma IDENTIFIER {% ([existingLabels, _, newLabel]) => [...existingLabels, newLabel] %}
                 | IDENTIFIER {% d => [d[0]] %}

clause_op -> %clauseOp {% d => d[0].value %}

func_params -> func_params %comma parameter {% ([existingParams, _comma, newParam]) => [...existingParams, newParam] %}
             | parameter {% d => [d[0]] %}

parameter -> SCALAR {% id %}
           | string {% id %}
           | EXPRESSION {% id %}

label_selector -> %lBrace label_match_list %rBrace {% ([_lBrace, labels, _rBrace]) => {const ret = new Labels(); labels.forEach(l => ret.addLabel(l.name, l.value, l.operator)); return ret} %}
                | %lBrace %rBrace {% d => new Labels() %}

label_match_list -> label_match_list %comma label_matcher {% ([existingLabels, _, newLabel]) => [...existingLabels, newLabel] %}
                  | label_matcher {% d => [d[0]] %}

label_matcher -> IDENTIFIER label_op string {% ([name, op, value]) => new Label(name, value, op) %}

string -> %string {% d => d[0].value %}

label_op -> %labelOp {% d => d[0].value %}

bin_op => %binOp {% d => d[0].value %}

agg_op => %aggOp {% d => d[0].value %}

func_name => %functionName {% d => d[0].value %}

duration => %lBracket range %rBracket {% ([_lBracket, range, _rBracket]) => range %}

range => %range {% d => d[0].value %}

SCALAR -> %scalar {% d => new Scalar(d[0].value) %}
