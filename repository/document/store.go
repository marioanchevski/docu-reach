package document

import (
	"sync"

	"github.com/marioanchevski/docu-reach/types"
)

type InMemoryDocumentStore struct {
	m         sync.RWMutex
	documents map[int]*types.Document
	idCounter int
}

func NewInMemoryDocumentStore() *InMemoryDocumentStore {
	return &InMemoryDocumentStore{
		m:         sync.RWMutex{},
		documents: make(map[int]*types.Document),
		idCounter: 1,
	}
}

func (ds *InMemoryDocumentStore) FindAll() []*types.Document {
	ds.m.RLock()
	defer ds.m.RUnlock()
	resultSlice := make([]*types.Document, 0, len(ds.documents))
	for _, value := range ds.documents {
		resultSlice = append(resultSlice, value)
	}
	return resultSlice
}

func (ds *InMemoryDocumentStore) FindById(docId int) *types.Document {
	ds.m.RLock()
	defer ds.m.RUnlock()

	value, ok := ds.documents[docId]
	if !ok {
		return nil
	}
	return value
}

func (ds *InMemoryDocumentStore) Create(docRequest types.CreateDocumentRequest) *types.Document {
	ds.m.Lock()
	defer ds.m.Unlock()

	newDoc := &types.Document{
		Id:          ds.idCounter,
		Title:       docRequest.Title,
		Description: docRequest.Description,
	}

	ds.documents[ds.idCounter] = newDoc
	ds.idCounter++
	return newDoc
}

func (ds *InMemoryDocumentStore) DeleteById(id int) bool {
	ds.m.Lock()
	defer ds.m.Unlock()
	_, ok := ds.documents[id]
	if !ok {
		return false
	}

	delete(ds.documents, id)
	return true
}

func (ds *InMemoryDocumentStore) Search(query string) []*types.Document {
	return nil
}
