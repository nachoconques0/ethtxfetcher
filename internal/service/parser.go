package service

import (
	"github.com/nachoconques0/ethtxfetcher/internal/model"
)

type parserClient interface {
	// GetCurrentBlock returns current block
	GetCurrentBlock() (string, error)
	// Subscribe add an address observer
	Subscribe(address string) bool
	// GetTransactions returns a list of inbound or outboundstx for an address
	GetTransactions(blockNumber string) ([]model.Transaction, error)
}

type parserService struct {
	parserClient parserClient
}

// NewParserSvc creates a new parser service
func NewParserSvc(c parserClient) (*parserService, error) {
	return &parserService{
		parserClient: c,
	}, nil
}

// GetCurrentBlock returns current block
func (p *parserService) GetCurrentBlock() (string, error) {
	return p.parserClient.GetCurrentBlock()
}

// GetTransactions returns a list of inbound or outbounds txs for an address based on a block number
func (p *parserService) GetTransactions(address string) ([]model.Transaction, error) {
	return p.parserClient.GetTransactions(address)
}

// Subscribe add an address observer
func (p *parserService) Subscribe(address string) bool {
	return p.parserClient.Subscribe(address)
}
