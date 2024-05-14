package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	expr, _ := govaluate.NewEvaluableExpression("foo > 0")
	parameters := make(map[string]interface{})
	parameters["foo"] = 1
	result, err := expr.Evaluate(parameters)
	fmt.Println(result, err)

	expr, _ = govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90")
	parameters = make(map[string]interface{})
	parameters["requests_made"] = 100
	parameters["requests_succeeded"] = 80
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)

	expr, _ = govaluate.NewEvaluableExpression("(mem_used / total_mem) * 100")
	parameters = make(map[string]interface{})
	parameters["total_mem"] = 1024
	parameters["mem_used"] = 512
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)

	functions := map[string]govaluate.ExpressionFunction{
		"strlen": func(args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return length, nil
		},
	}

	exprString := "strlen('aaa')"
	expr, _ = govaluate.NewEvaluableExpressionWithFunctions(exprString, functions)
	result, _ = expr.Evaluate(nil)
	fmt.Println(result)
}
