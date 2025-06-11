package document_test

import (
	"testing"

	"github.com/marioanchevski/docu-reach/repository/document"
	"github.com/marioanchevski/docu-reach/types"
)

type mockMatcher struct {
	match bool
}

func (m mockMatcher) DocumentSatisfiesFilter(doc *types.Document, filter types.DocumentFilter) bool {
	return m.match
}

func getMockMatcher(match bool) mockMatcher {
	return mockMatcher{match}
}

func TestCreateAndFindById(t *testing.T) {
	m := getMockMatcher(true)
	store := document.NewInMemoryDocumentStore(m)

	req := types.CreateDocumentRequest{
		Title:       "Test Title",
		Description: "Test Description",
	}

	doc := store.Create(req)

	if doc.Id != 1 {
		t.Errorf("Expected ID to be 1, got %d", doc.Id)
	}
	if doc.Title != req.Title {
		t.Errorf("Expected title to be %q, got %q", req.Title, doc.Title)
	}

	found, err := store.FindById(1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if found.Title != "Test Title" {
		t.Errorf("Expected title %q, got %q", "Test Title", found.Title)
	}
}

func TestFindAll(t *testing.T) {
	m := getMockMatcher(true)
	store := document.NewInMemoryDocumentStore(m)

	docs := store.FindAll()
	if len(docs) != 0 {
		t.Errorf("Expected 0 documents, got %d", len(docs))
	}

	store.Create(types.CreateDocumentRequest{Title: "Doc 1", Description: "Desc 1"})
	store.Create(types.CreateDocumentRequest{Title: "Doc 2", Description: "Desc 2"})

	docs = store.FindAll()
	if len(docs) != 2 {
		t.Errorf("Expected 2 documents, got %d", len(docs))
	}
}

func TestFindById_NotFound(t *testing.T) {
	m := getMockMatcher(true)
	store := document.NewInMemoryDocumentStore(m)

	_, err := store.FindById(999)
	if err == nil {
		t.Fatal("Expected error when document not found, got nil")
	}
}

func TestFilter_MatchesAll(t *testing.T) {
	mock := getMockMatcher(true)
	store := document.NewInMemoryDocumentStore(mock)

	store.Create(types.CreateDocumentRequest{
		Title:       "Doc 1",
		Description: "Desc 1",
	})
	store.Create(types.CreateDocumentRequest{
		Title:       "Doc 2",
		Description: "Desc 2",
	})

	filter := types.DocumentFilter{}

	results := store.Filter(filter)

	if len(results) != 2 {
		t.Errorf("Expected 2 matching documents, got %d", len(results))
	}
}

func TestFilter_MatchesNone(t *testing.T) {
	mock := getMockMatcher(false)
	store := document.NewInMemoryDocumentStore(mock)

	store.Create(types.CreateDocumentRequest{
		Title:       "Doc 1",
		Description: "Desc 1",
	})
	store.Create(types.CreateDocumentRequest{
		Title:       "Doc 2",
		Description: "Desc 2",
	})

	filter := types.DocumentFilter{}

	results := store.Filter(filter)

	if len(results) != 0 {
		t.Errorf("Expected 0 matching documents, got %d", len(results))
	}
}

func TestDeleteSuccess(t *testing.T) {
	mock := getMockMatcher(true)
	store := document.NewInMemoryDocumentStore(mock)

	doc := store.Create(types.CreateDocumentRequest{Title: "Test", Description: "Delete me"})

	err := store.DeleteById(doc.Id)
	if err != nil {
		t.Fatalf("Expected no error deleting existing doc, got: %v", err)
	}

	_, err = store.FindById(doc.Id)
	if err == nil {
		t.Fatalf("expected error finding deleted document, got nil")
	}
}

func TestDeleteFail(t *testing.T) {

	mock := getMockMatcher(true)
	store := document.NewInMemoryDocumentStore(mock)
	err := store.DeleteById(9999)
	if err == nil {
		t.Fatalf("expected error deleting non-existing document, got nil")
	}
}
