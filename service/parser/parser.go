package parser

import "strings"

func ParseSearchTerms(searchString string) (include, exclude []string) {
	terms := strings.SplitSeq(searchString, ",")
	for term := range terms {
		term = strings.TrimSpace(term)
		if term == "" {
			continue
		}
		if strings.HasPrefix(term, "-") && len(term) > 1 {
			exclude = append(exclude, strings.TrimSpace(term[1:]))
		} else {
			include = append(include, strings.TrimSpace(term))
		}
	}
	return
}
