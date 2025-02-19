package model

import "encoding/json"

// JSON-RPC request structure
type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type GetBlockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type GetBlockTransactionsResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
}

// Transaction represents an Ethereum transaction
type Transaction struct {
	Hash  string `json:"hash"`
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
}

// Block represents an Ethereum block with transactions
type Block struct {
	Number       string        `json:"number"`
	Hash         string        `json:"hash"`
	Transactions []Transaction `json:"transactions"`
}
