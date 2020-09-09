/***********************************************************************
# File Name: main.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2020-09-09 11:38:46
*********************************************************************/
package cita

import (
	"encoding/json"
)

type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type Response struct {
	Id      int             `json:"id"`
	Jsonrpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   Error           `json:"error,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResultPeerCount struct {
	Result string
}

type ResultPeerInfo struct {
	Amount       int               `json:"amount"`
	Peers        map[string]string `json:"peers"`
	ErrorMessage string            `json:"errorMessage"`
}

type ResultRawTransaction struct {
	Hash   string `json:"hash"`
	Status string `json:"status"`
}

type ResultVersion struct {
	SoftwareVersion string `json:"softwareVersion"`
}

//type ResultBlockByHash struct {
//	Version int    `json:"version"`
//	Hash    string `json:"hash"`
//	Header
//}
