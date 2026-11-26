package analysis

import "sync"

type State struct {
	mu sync.RWMutex

	// documents is a map of file names to contents
	documents map[string]string
}

func NewState() *State {
	return &State{documents: map[string]string{}}
}

func (s *State) SetDocument(uri, text string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.documents[uri] = text
}

func (s *State) Document(uri string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	doc, ok := s.documents[uri]
	return doc, ok
}

func (s *State) CloseDocument(uri string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.documents, uri)
}
