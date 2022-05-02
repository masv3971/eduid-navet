package httpserver

import (
	"context"
	"eduid-navet/internal/apiv2"

	"github.com/gin-gonic/gin"
)

func (s *Service) endpointNamnSokning(ctx context.Context, c *gin.Context) (interface{}, error) {
	request := &apiv2.RequestNamnSokning{}
	if err := s.bindRequest(c, request); err != nil {
		return nil, err
	}
	reply, err := s.apiv1.NamnSokning(ctx, request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *Service) endpointStatus(ctx context.Context, c *gin.Context) (interface{}, error) {
	reply, err := s.apiv1.Status(ctx)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
