package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/nachoconques0/ethtxfetcher/internal/api"
	"github.com/nachoconques0/ethtxfetcher/internal/model"
)

const defaultTimeout = 5 * time.Second

type parserCtrl interface {
	// GetCurrentBlock returns current block
	GetCurrentBlock(w http.ResponseWriter, r *http.Request)
	// Subscribe add an address observer
	Subscribe(w http.ResponseWriter, r *http.Request)
	// GetTransactions returns a list of inbound or outboundstx for an address
	GetTransactions(w http.ResponseWriter, r *http.Request)
}

type parserSvc interface {
	// GetCurrentBlock returns current block
	GetCurrentBlock() (string, error)
	// Subscribe add an address observer
	Subscribe(address string) bool
	// GetTransactions returns a list of inbound or outboundstx for an address
	GetTransactions(blockNumber string) ([]model.Transaction, error)
}

// Application holds the basic structure that
// defines the application
type Application struct {
	httpPort string
	nodeURL  string
	server   *api.Server

	// value used for determine the gap of time
	// required for shutdown the application
	timeout time.Duration

	// HTTP Controllers
	parserCtrl parserCtrl

	// Services
	parserSvc parserSvc
}

// New function builds a new application applying
// the given application options
func New(opts ...Option) (*Application, error) {
	a := &Application{
		timeout: defaultTimeout,
	}
	for _, o := range opts {
		o(a)
	}
	err := a.setupDomain()
	if err != nil {
		return nil, fmt.Errorf("error while setting up domains - %w", err)
	}

	api, err := api.NewServer(a.httpPort, a.parserCtrl)
	if err != nil {
		return nil, fmt.Errorf("error while setting up the http server - %w", err)
	}

	a.server = api
	return a, nil
}

func (a *Application) Start() {
	slog.Info(fmt.Sprintf("HTTP server: starting at %s\n", a.server.Addr))
	if err := a.server.Run(); err != nil {
		slog.Error(fmt.Sprintf("Application: error running the http server: %s", err))
	}
}
