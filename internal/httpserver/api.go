package httpserver

import (
	"context"
	"eduid-navet/internal/apiv2"
	"eduid-navet/pkg/model"
)

// Apiv1 interface
type Apiv1 interface {
	NamnSokning(ctx context.Context, indata *apiv2.RequestNamnSokning) (*apiv2.ReplyNamnSokning, error)
	Status(ctx context.Context) (*model.Status, error)
}
