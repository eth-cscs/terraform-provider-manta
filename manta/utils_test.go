package manta

import "github.com/r3labs/diff/v3"
import "fmt"

func GetDiff[T comparable](a, b T) string {
	changelog, _ := diff.Diff(a, b)

	var diff string
	for _, change := range changelog {
		if change.Type == "update" {
			diff += fmt.Sprintf("%s: \"%s\" -> \"%s\"\n",
				change.Path,
				change.From,
				change.To)
		}
	}

	return diff
}
