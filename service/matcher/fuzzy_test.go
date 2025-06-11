package matcher_test

import (
	"testing"

	"github.com/marioanchevski/docu-reach/service/matcher"
	"github.com/marioanchevski/docu-reach/types"
)

func TestFuzzyMatcher_DocumentSatisfiesFilter(t *testing.T) {
	m := matcher.NewFuzzyMatcher()

	tests := []struct {
		name     string
		doc      types.Document
		filter   types.DocumentFilter
		expected bool
	}{

		{
			name: "Include terms only, in title",
			doc: types.Document{
				Title:       "document1",
				Description: "description1 for document1",
			},
			filter: types.DocumentFilter{
				TitleInclude: []string{"document1"},
			},
			expected: true,
		},
		{
			name: "Exclude terms only, in title",
			doc: types.Document{
				Title:       "document1",
				Description: "description1 for document1",
			},
			filter: types.DocumentFilter{
				TitleExclude: []string{"document1"},
			},
			expected: false,
		},
		{
			name: "Include terms only, in description",
			doc: types.Document{
				Title:       "document1",
				Description: "description1 for document1",
			},
			filter: types.DocumentFilter{
				DescInclude: []string{"document1"},
			},
			expected: true,
		},
		{
			name: "Exclude terms only, in description",
			doc: types.Document{
				Title:       "document1",
				Description: "description1 for document1",
			},
			filter: types.DocumentFilter{
				DescExclude: []string{"document1"},
			},
			expected: false,
		},
		{
			name: "Include terms only, in title and description, and operator",
			doc: types.Document{
				Title:       "document1",
				Description: "description1 for document1",
			},
			filter: types.DocumentFilter{
				TitleInclude: []string{"document1"},
				DescInclude:  []string{"description1"},
				Operator:     "and",
			},
			expected: true,
		},
		{
			name: "Include terms only, in title and description, or operator",
			doc: types.Document{
				Title:       "document1",
				Description: "description1 for document1",
			},
			filter: types.DocumentFilter{
				TitleInclude: []string{"document3"},
				DescInclude:  []string{"description1"},
				Operator:     "or",
			},
			expected: true,
		},
		{
			name: "Include terms for title, exclude terms for desc, or operator",
			doc: types.Document{
				Title:       "document1",
				Description: "description1 for document1",
			},
			filter: types.DocumentFilter{
				TitleInclude: []string{"document3"},
				DescExclude:  []string{"description1"},
				Operator:     "or",
			},
			expected: false,
		},
		{
			name: "Empty filter matches everything",
			doc: types.Document{
				Title:       "Some Title",
				Description: "Some Description",
			},
			filter:   types.DocumentFilter{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := m.DocumentSatisfiesFilter(&tt.doc, tt.filter)
			if result != tt.expected {
				t.Errorf("DocumentSatisfiesFilter() = %v, want %v", result, tt.expected)
			}
		})
	}
}
