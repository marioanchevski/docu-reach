package parser_test

import (
	"reflect"
	"testing"

	"github.com/marioanchevski/docu-reach/service/parser"
)

func TestSimpleSignParser_ParseSearchTerms(t *testing.T) {
	parser := parser.NewSimpleSignParser()

	tests := []struct {
		name        string
		input       string
		wantInclude []string
		wantExclude []string
	}{
		{
			name:        "only include terms",
			input:       "title1,title2,title3",
			wantInclude: []string{"title1", "title2", "title3"},
			wantExclude: []string{},
		},
		{
			name:        "include and exclude terms",
			input:       "title1,-title2,title3,-title4",
			wantInclude: []string{"title1", "title3"},
			wantExclude: []string{"title2", "title4"},
		},
		{
			name:        "empty input",
			input:       "",
			wantInclude: []string{},
			wantExclude: []string{},
		},
		{
			name:        "terms with spaces and empty terms",
			input:       "  title1 , , -title2 , title3  ",
			wantInclude: []string{"title1", "title3"},
			wantExclude: []string{"title2"},
		},
		{
			name:        "single dash",
			input:       "-",
			wantInclude: []string{"-"},
			wantExclude: []string{},
		},
		{
			name:        "double dash",
			input:       "--",
			wantInclude: []string{},
			wantExclude: []string{"-"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInclude, gotExclude := parser.ParseSearchTerms(tt.input)
			if !reflect.DeepEqual(gotInclude, tt.wantInclude) {
				t.Errorf("Include terms = %v; want %v", gotInclude, tt.wantInclude)
			}
			if !reflect.DeepEqual(gotExclude, tt.wantExclude) {
				t.Errorf("Exclude terms = %v; want %v", gotExclude, tt.wantExclude)
			}
		})
	}
}
