package translation

import (
	"github.com/domicmeia/gcp_practice/handler/rest"
)

var _ rest.Translator = &RemoteService{}

type RemoteService struct {
	client HelloClient
}

type HelloClient interface {
	Translate(word, language string) (string, error)
}

func NewRemoteService(client HelloClient) *RemoteService {
	return &RemoteService{
		client: client,
	}
}

func (s *RemoteService) Translate(word string, language string) string {
	resp, _ := s.client.Translate(word, language)
	return resp
}
