package gonavet2

import (
	"context"
	"net/http"

	"github.com/masv3971/gonavet2/navettypes"
)

type searchService struct {
	client   *Client
	endpoint string
}

type RequestSearchPerson struct {
	GivenName string `json:"given_name" validate:"required_without_all=Surename City"`
	Surname   string `json:"surename" validate:"required_without_all=GivenName City"`
	City      string `json:"city" validate:"required_without_all=GivenName Surename"`
}

type ReplySearchPerson struct {
	*navettypes.OkSokResponse
}

func (s *searchService) Person(ctx context.Context, req *RequestSearchPerson) (*ReplySearchPerson, *http.Response, error) {
	reply := &ReplySearchPerson{}
	resp, err := s.client.call(ctx, http.MethodPost, "/sok", req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}
