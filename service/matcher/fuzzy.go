package matcher

import (
	"strings"

	"github.com/marioanchevski/docu-reach/types"
)

type FuzzyMatcher struct {
}

func NewFuzzyMatcher() FuzzyMatcher {
	return FuzzyMatcher{}
}

func (FuzzyMatcher) DocumentSatisfiesFilter(doc *types.Document, df types.DocumentFilter) bool {
	titleMatch := matchField(doc.Title, df.TitleInclude, df.TitleExclude)
	descMatch := matchField(doc.Description, df.DescInclude, df.DescExclude)

	if df.Operator == "or" {
		return titleMatch || descMatch
	}
	return titleMatch && descMatch
}

func matchField(text string, include, exclude []string) bool {
	text = strings.ToLower(text)

	for _, word := range include {
		if !strings.Contains(text, strings.ToLower(word)) {
			return false
		}
	}
	for _, word := range exclude {
		if strings.Contains(text, strings.ToLower(word)) {
			return false
		}
	}
	return true
}
