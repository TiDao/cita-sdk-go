/***********************************************************************
# File Name: main.go
# Author: TiDao
# mail: tidao2049@gmail.com
# Created Time: 2020-09-09 11:38:46
*********************************************************************/
package cita

import (
	"encoding/json"
	"time"
)


type Bool bool
type String string
//requst body type
//easyjson
type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int32         `json:"id"`
}

//easyjson
type Filter struct {
	FromBlock string   `json:"fromBlock,omitempty"`
	ToBlock   string   `json:"toBlock,omitempty"`
	Address   []string `json:"address,omitempty"`
	Topics    []string `json:"topics,omitempty"`
}

//easyjson
type CallRequest struct {
	From string `json:"from,moitempty"`
	To   string `json:"to"`
	Data string `json:"data,moitempt"`
}

//response body type
//easyjson
type Response struct {
	Id      int32           `json:"id"`
	Jsonrpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   Error           `json:"error,omitempty"`
}

//easyjson
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

//easyjson
type ResultBlock struct {
	Version string `json:"version"`
	Hash    string `json:"hash"`
	Header  Header `json:"hearder"`
	Body    Body   `json:"header"`
}

//easyjson
type Header struct {
	Timestamp        time.Duration `json:"timestamp"`
	PreHash          string        `json:"preHash"`
	StateRoot        string        `json:"stateRoot"`
	TransactionsRoot string        `json:"transactionsRoot"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	QuotaUsed        string        `json:"quotaUsed"`
	number           string        `json:"number"`
	Proposer         string        `json:"propose"`
	Proof            Proof         `json:"proof"`
}

//easyjson
type Proof struct {
	Proposal string            `json:"proposal"`
	Height   int32             `json:"height"`
	Round    int32             `json:"round"`
	Commits  map[string]string `json:"commit"`
}

//easyjson
type Body struct {
	Transactions []Transactions `json:"transactions"`
}

//easyjson
type Transactions struct {
	Hash    string `json:"hash"`
	Content string `json:"content"`
}

//easyjson
type ResultTransactionReceipt struct {
	TransactionHash     string       `json:"transactionHash"`
	TransactionIndex    string       `json"transactionIndex"`
	BlockHash           string       `json:"blockHash"`
	BlockNumber         string       `json:"blockNumber"`
	CumulativeQuotaUsed string       `json:"cumulativeQuotaUsed"`
	QuotaUsed           string       `json:"quotaUsed"`
	ConstractAddress    string       `json:"constractAddress"`
	Logs                []ResultLogs `json:"logs"`
	Root                string       `json:"root"`
	LogsBloom           string       `json:"logsBloom"`
}

//easyjson
type Logs struct {
	Id      int32           `json:"id"`
	Jsonrpc string          `json:"jsonrpc"`
	Result  []ResultLogs    `json:"result,omitempty"`
	Error   Error           `json:"error,omitempty"`
}

//easyjson
type ResultLogs struct {
	Address             string   `json:"address"`
	Topic               []string `json:"topic"`
	Data                string   `json:"data"`
	BlockHash           string   `json:"blockHash"`
	BlockNumber         string   `json:"blockNumber"`
	TransactionIndex    string   `json:"transactionIndex"`
	TransactionHash     string   `json:"transactionHash"`
	LogIndex            string   `json:"logIndex"`
	TransactionLogIndex string   `json:"transactionLogIndex"`
}


//easyjson
type ResultPeerInfo struct {
	Amount       int32             `json:"amount"`
	Peers        map[string]string `json:"peers"`
	ErrorMessage string            `json:"errorMessage"`
}

//easyjson
type ResultRawTransaction struct {
	Hash   string `json:"hash"`
	Status string `json:"status"`
}

//easyjson
type ResultVersion struct {
	SoftwareVersion string `json:"softwareVersion"`
}

//easyjson
type ResultTransaction struct {
	Hash        string `json:"hash"`
	Content     string `json:"content"`
	From        string `json:"from"`
	BlockNumber string `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`
	Index       string `json:"index'`
}

//easyjson
type ResultMetaData struct {
	ChainId          int16         `json:"chainId"`
	ChainIdV1        string        `json:"chainIdV1"`
	ChainName        string        `json:"chainName"`
	Operator         string        `json:"operator"`
	GenesisTimestamp time.Duration `json:"genesisTimestamp"`
	Validators       []string      `json:"validators"`
	BlockInterval    time.Duration `json:"blockInterval"`
	TokenName        string        `json:"tokenName"`
	TokenAvatar      string        `json:"tokenAvatar"`
	Version          int32         `json:"version"`
	EconomicalModel  int16         `json:"economicalModel"`
	Website          string        `json:"website,omitempty"`
}
