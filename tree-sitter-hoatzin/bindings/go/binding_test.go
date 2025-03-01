package tree_sitter_hoatzin_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_hoatzin "github.com/sapoturge/hoatzin/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_hoatzin.Language())
	if language == nil {
		t.Errorf("Error loading Hoatzin grammar")
	}
}
