package types

import (
	"errors"
	"time"
)

type Document struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DocumentFilter struct {
	TitleInclude []string
	TitleExclude []string
	DescInclude  []string
	DescExclude  []string
	Operator     string
}

type CreateDocumentRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r *CreateDocumentRequest) Validate() error {
	if r.Title == "" {
		return errors.New("Document title is required!")
	}
	if r.Description == "" {
		return errors.New("Document description is required!")
	}
	return nil
}

type APIResponse[T any] struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   *string   `json:"message,omitempty"`
	Error     *string   `json:"error,omitempty"`
	Data      T         `json:"data"`
}

type APIError struct {
	Error string `json:"error"`
}

type DocumentStore interface {
	FindById(id int) (*Document, error)
	FindAll() []*Document
	Create(document CreateDocumentRequest) *Document
	DeleteById(id int) error
	Filter(docFilter DocumentFilter) []*Document
}
