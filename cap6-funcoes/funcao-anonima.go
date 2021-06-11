package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	texto := "Anderson tem 21 anos"
	expr := regexp.MustCompile("\\d")

	fmt.Println(expr.ReplaceAllString(texto, "3"))

	expr = regexp.MustCompile("\\b\\w")
	texto = "antonio carlos jobim"

	processado := expr.ReplaceAllStringFunc(texto, func(s string) string {
		return strings.ToUpper(s)
	})

	fmt.Println(processado)

	transformadora := func(s string) string {
		return strings.ToUpper(s)
	}

	fmt.Println()

	texto = "jorge carlos"

	fmt.Println(transformadora(texto))

	nomeProcessado := expr.ReplaceAllStringFunc(texto, transformadora)

	fmt.Println(nomeProcessado)
}
