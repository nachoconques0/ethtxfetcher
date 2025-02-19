package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nachoconques0/ethtxfetcher/internal/model"
)

type Client struct {
	nodeURL    string
	subscribed map[string]bool
	txHistory  map[string][]model.Transaction
}

// NewParser initializes a new Ethereum parser
func NewParser(nodeURL string) *Client {
	return &Client{
		nodeURL:    nodeURL,
		subscribed: make(map[string]bool),
		txHistory:  make(map[string][]model.Transaction),
	}
}

// GetCurrentBlock returns the last parsed block number
func (p *Client) GetCurrentBlock() (string, error) {
	requestBody := model.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		ID:      1,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(p.nodeURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("error post http request: %s\n", err)
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading body response: %s\n", err)
		return "", err
	}

	var blockRes *model.GetBlockResponse
	err = json.Unmarshal(responseBody, &blockRes)
	if err != nil {
		fmt.Printf("client: error decoding unmarshaling: %s\n", err)
		return "", err
	}
	return blockRes.Result, nil
}

// Subscribe adds an address to the observer list
func (p *Client) Subscribe(address string) bool {
	if _, exists := p.subscribed[address]; exists {
		return false
	}
	p.subscribed[address] = true
	p.txHistory[address] = []model.Transaction{}
	return true
}

// GetTransactions returns the list of transactions for an address
func (p *Client) GetTransactions(address string) ([]model.Transaction, error) {
	currentBlockNr, err := p.GetCurrentBlock()
	if err != nil {
		fmt.Printf("error reading body response: %s\n", err)
		return []model.Transaction{}, err
	}
	requestBody := model.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{currentBlockNr, true},
		ID:      0,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(p.nodeURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rpcResponse model.GetBlockTransactionsResponse
	err = json.Unmarshal(body, &rpcResponse)
	if err != nil {
		return nil, err
	}

	var block model.Block
	err = json.Unmarshal(rpcResponse.Result, &block)
	if err != nil {
		return nil, err
	}

	for _, tx := range block.Transactions {
		transaction := &model.Transaction{
			Hash:  tx.Hash,
			From:  tx.From,
			To:    tx.To,
			Value: tx.Value,
		}

		if p.subscribed[tx.From] {
			p.txHistory[tx.From] = append(p.txHistory[tx.From], *transaction)
		}
		if p.subscribed[tx.To] {
			p.txHistory[tx.To] = append(p.txHistory[tx.To], *transaction)
		}
	}
	return p.txHistory[address], nil
}
