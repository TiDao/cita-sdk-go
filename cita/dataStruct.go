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
type ResultBlockTrue struct {
	Version int16  `json:"version"`
	Hash    string `json:"hash"`
	Header  Header `json:"header"`
	Body    BodyTrue   `json:"body"`
}

//easyjson
type ResultBlockFalse struct {
	Version int16  `json:"version"`
	Hash    string `json:"hash"`
	Header  Header `json:"header"`
	Body    BodyFalse   `json:"body"`
}

//easyjson
type Header struct {
	Timestamp        time.Duration `json:"timestamp"`
	PrevHash         string        `json:"prevHash"`
	StateRoot        string        `json:"stateRoot"`
	TransactionsRoot string        `json:"transactionsRoot"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	QuotaUsed        string        `json:"quotaUsed"`
	Number           string        `json:"number"`
	Proposer         string        `json:"proposer"`
	Proof            Proof         `json:"proof"`
}

//easyjson
type Proof struct {
	Bft Bft `json:"Bft"`
}

//easyjson
type Bft struct {
	Proposal string            `json:"proposal"`
	Height   int32             `json:"height"`
	Round    int32             `json:"round"`
	Commits  map[string]string `json:"commits"`
}

//easyjson
type BodyTrue struct {
	Transactions []ResultTransaction `json:"transactions"`
}

//easyjson
type BodyFalse struct {
	Transactions []string `json:"transactions"`
}

//easyjson
type Transactions struct {
	Hash    string `json:"hash"`
	Content string `json:"content"`
}

//Object - 回执对象:
//transactionHash: Data32 - 交易哈希
//transactionIndex: Quantity - 交易 index
//blockHash: Data32 - 交易所在块的块哈希
//blockNumber: Quantity - 交易所在块的块高度
//cumulativeQuotaUsed: Quantity - 块中该交易之前(包含该交易)的所有交易消耗的 quota 总量
//quotaUsed: Quantity - 交易消耗的 quota 数量
//contractAddress: Data20 - 如果是部署合约, 这个地址指的是新创建出来的合约地址. 否则为空
//logs: Array - 交易产生的日志集合
//root: Data32 - 状态树根
//errorMessage: String 错误信息
//
//回执错误：
//No transaction permission - 没有发交易权限
//No contract permission - 没有创建合约权限
//Not enough base quota - 基础配额 不够
//Block quota limit reached - 达到块配额限制
//Account quota limit reached - 达到账户配额限制
//Out of quota - 配额不够
//Jump position wasn't marked with JUMPDEST instruction - EVM 内部错误
//Instruction is not supported - EVM 内部错误
//Not enough stack elements to execute instruction - EVM 内部错误
//Execution would exceed defined Stack Limit - EVM 内部错误
//EVM internal error - EVM 内部错误
//Mutable call in static context - EVM 内部错误
//Out of bounds - EVM 内部错误
//Reverted - EVM 内部错误，REVERTED instruction
//
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
	Id      int32        `json:"id"`
	Jsonrpc string       `json:"jsonrpc"`
	Result  []ResultLogs `json:"result,omitempty"`
	Error   Error        `json:"error,omitempty"`
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
	Hash        string `json:"hash,omitempty"`
	Content     string `json:"content,omitempty"`
	From        string `json:"from,omitempty"`
	BlockNumber string `json:"blockNumber,omitempty"`
	BlockHash   string `json:"blockHash,omitempty"`
	Index       string `json:"index,omitempty'`
}

//chainId, Integer - version < 1 时的 chain_id, 用来防止重放攻击
//chainIdV1, Quantity - version > 1 时的 chain_id
//chainName, String - 链名称
//operator, String - 链的运营者
//genesisTimestamp, Integer - 创世块时间戳
//validators, [Data20] - 验证者地址集合
//blockInterval Integer - 出块间隔
//tokenName, String - Token 名称
//tokenSymbol, String - Token 标识
//tokenAvatar, String - Token 标志
//version, Integer - 链版本
//economicalModel, EconomicalModel - 链经济模型
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
