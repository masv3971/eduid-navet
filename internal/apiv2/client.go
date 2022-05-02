package apiv2

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"eduid-navet/pkg/logger"
	"eduid-navet/pkg/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/masv3971/gonavet2"
)

type Client struct {
	cfg   *model.Cfg
	log   *logger.Logger
	navet *gonavet2.Client
}

func New(ctx context.Context, cfg *model.Cfg, log *logger.Logger) (*Client, error) {
	c := &Client{
		cfg: cfg,
		log: log,
	}

	var err error
	c.navet, err = gonavet2.New(gonavet2.Config{
		URL:                   "",
		Certificate:           &x509.Certificate{},
		CertificatePEM:        []byte{},
		PrivateKey:            &rsa.PrivateKey{},
		PrivateKeyPEM:         []byte{},
		ProxyURL:              "",
		ClientID:              "",
		SkatteverketJWTConfig: jwt.Claims,
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}
