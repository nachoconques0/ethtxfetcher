// Package controller contain all controller logic & needed interfaces
package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	internalErrors "github.com/nachoconques0/ethtxfetcher/internal/errors"
	"github.com/nachoconques0/ethtxfetcher/internal/model"
)

// NewParserCtrl creates a new HTTP Controller with the given service
func NewParserCtrl(parserSvc parserSvc) *parserCtrl {
	return &parserCtrl{svc: parserSvc}
}

// Parser defines what are the different allowed http
// endpoints that can be consumed
type parserSvc interface {
	// GetCurrentBlock returns current block
	GetCurrentBlock() (string, error)
	// Subscribe add an address observer
	Subscribe(address string) bool
	// GetTransactions returns a list of inbound or outboundstx for an address
	GetTransactions(blockNumber string) ([]model.Transaction, error)
}

type parserCtrl struct {
	svc parserSvc
}

// GetCurrentBlock returns current block
func (p *parserCtrl) GetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	res, err := p.svc.GetCurrentBlock()
	if err != nil {
		slog.Error(fmt.Sprintf("getting current block: %s\n", err))
		responseError(w, r, err)
		return
	}
	encodeResponse(w, res)
}

// Subscribe add an address observer
func (p *parserCtrl) Subscribe(w http.ResponseWriter, r *http.Request) {
	address := r.PathValue("address")
	res := p.svc.Subscribe(address)
	encodeResponse(w, res)
}

// GetTransactions returns a list of inbound or outbounds txs for an address
func (p *parserCtrl) GetTransactions(w http.ResponseWriter, r *http.Request) {
	address := r.PathValue("address")
	res, err := p.svc.GetTransactions(address)
	if err != nil {
		slog.Error(fmt.Sprintf("getting transactions of block: %s\n", err))
		responseError(w, r, err)
		return
	}
	encodeResponse(w, res)
}

func encodeResponse(w http.ResponseWriter, res interface{}) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		slog.Error(fmt.Sprintf("error encoding response: %s", err))
	}
}

func responseError(w http.ResponseWriter, r *http.Request, err error) {
	var internalErr *internalErrors.Error
	if errors.As(err, &internalErr) {
		w.WriteHeader(internalErr.HTTPStatus())
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(internalErr); err != nil {
			slog.Error(fmt.Sprintf("error encoding response: %s", err))
		}
		return
	}
	responseError(w, r, err)
}
