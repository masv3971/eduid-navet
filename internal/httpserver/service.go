package httpserver

import (
	"context"
	"eduid-navet/internal/apiv2"
	"eduid-navet/pkg/helpers"
	"eduid-navet/pkg/logger"
	"eduid-navet/pkg/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Service is the service object for httpserver
type Service struct {
	config *model.Cfg
	logger *logger.Logger
	server *http.Server
	apiv1  Apiv1
	gin    *gin.Engine
}

// New creates a new httpserver service
func New(ctx context.Context, config *model.Cfg, api *apiv2.Client, logger *logger.Logger) (*Service, error) {
	s := &Service{
		config: config,
		logger: logger,
		apiv1:  api,
		server: &http.Server{Addr: config.APIServer.Addr},
	}

	switch s.config.Production {
	case true:
		gin.SetMode(gin.ReleaseMode)
	case false:
		gin.SetMode(gin.DebugMode)
	}

	apiValidator := validator.New()
	binding.Validator = &defaultValidator{
		Validate: apiValidator,
	}

	s.gin = gin.New()
	s.server.Handler = s.gin
	s.server.ReadTimeout = time.Second * 5
	s.server.WriteTimeout = time.Second * 30
	s.server.IdleTimeout = time.Second * 90

	// Middlewares
	s.gin.Use(s.middlewareTraceID(ctx))
	s.gin.Use(s.middlewareDuration(ctx))
	s.gin.Use(s.middlewareLogger(ctx))
	s.gin.Use(s.middlewareCrash(ctx))
	s.gin.NoRoute(func(c *gin.Context) {
		status := http.StatusNotFound
		p := helpers.Problem404()
		c.JSON(status, gin.H{"error": p, "data": nil})
	})

	s.regEndpoint(ctx, "api/v2/namnsokning", "POST", s.endpointNamnSokning)

	s.regEndpoint(ctx, "/health", "GET", s.endpointStatus)

	// Run http server
	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			s.logger.New("http").Fatal("listen_error", "error", err)
		}
	}()

	s.logger.Info("started")

	return s, nil
}

func (s *Service) regEndpoint(ctx context.Context, path, method string, handler func(context.Context, *gin.Context) (interface{}, error)) {
	s.gin.Handle(method, path, func(c *gin.Context) {
		res, err := handler(ctx, c)
		status := 200

		if err != nil {
			status = 400
		}
		renderContent(c, status, gin.H{"data": res, "error": helpers.NewErrorFromError(err)})
	})
}

func renderContent(c *gin.Context, code int, data interface{}) {
	switch c.NegotiateFormat(gin.MIMEJSON, "*/*") {
	case gin.MIMEJSON:
		c.JSON(code, data)
	case "*/*": // curl
		c.JSON(code, data)
	default:
		c.JSON(406, gin.H{"data": nil, "error": helpers.NewErrorDetails("not_acceptable", "Accept header is invalid. It should be \"application/json\".")})
	}
}

// Close closing httpserver
func (s *Service) Close(ctx context.Context) error {
	s.logger.Info("Quit")
	return nil
}
