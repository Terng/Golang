package main

import (
	"fmt"
	"regexp"
	"strings"
)

func keywordToUpper(src string, key ...string) string {
	var re = regexp.MustCompile(`\b(` + strings.Join(key, "|") + `)\b`)
	return re.ReplaceAllStringFunc(src, func(s string) string {
		return strings.ToLower(s)
	})
}

const t = `
Test1 test2 test3
Test1 test2 test3
Test1 test2 test3
`

func main() {
	fmt.Println(keywordToUpper(t, "Test1", "test2", "test3"))

}
