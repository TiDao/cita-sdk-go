/***********************************************************************
# File Name: main.go
# Author: TiDao
# mail: tidao2049@gmail.com
# Created Time: 2020-11-24 12:28:58
*********************************************************************/
package main

import "log"
import "./cita"

//import "io"
import "sync"

import "time"
 
//import "strings"

func TestGetBlock() {
    //var url = "http://172.16.0.49:1337"
    //var url = "http://192.168.1.87:30000"
    //var url = "http://192.168.1.67:8083"
    var url = "http://192.168.1.243:8081" 
    log.Println(url)
	var wg sync.WaitGroup
	ch := make(chan string, 50)
	TransactionCh := make(chan string, 50)
	go getNumber(ch, url, "0x0")

	for i := 0; i < 200; i++ {
		wg.Add(1)
		req := &cita.Request{
			Jsonrpc: "2.0",
			Method:  "getBlockByNumber",
			Params: []interface{}{
				"",
				false,
			},
			Id: int32(i),
		}

		var Error cita.Error
		var err error
		go func() {
			result := new(cita.ResultBlockFalse)
			log.SetFlags(log.Lmicroseconds)
			var goRun bool = true
			for {
                //timeout := time.NewTimer(time.Microsecond*500)
				if goRun {
                    req.Params[0] = <-ch
				    log.Printf("the block height is %s ", req.Params[0])
				}

                begin := time.Now()
				Error, err = cita.GetBlockFalse(req, result, url)
                if int64(time.Since(begin)) > int64(time.Microsecond*10000000) {
                    log.Println(int64(time.Since(begin)))
                    log.Println("time:",time.Since(begin))
                }
				//log.Printf("the result is %v\n\n",result)
				if err != nil {
					log.Println("TestGetBlock() error,cita.GetBlock() error,", err)
					goRun = false
				} else {
					goRun = true
				}
				if Error.Code == 0 {
					if len(result.Body.Transactions) != 0 {
                        for _, Transaction := range result.Body.Transactions {
                            log.Println("transaction is: ",Transaction)
							TransactionCh <- Transaction
						}
					}
				} else {
                    log.Println("get Block error:", Error)
				}
			}
		}()
	}

	go func() {
		req := &cita.Request{
			Jsonrpc: "2.0",
			Method:  "getTransactionReceipt",
			Params: []interface{}{
				"",
			},
			Id: 1,
		}

		result := &cita.ResultTransactionReceipt{}
		var goRun = true
		for {
			if goRun {
				req.Params[0] = <-TransactionCh
			}
			_, err := cita.GetTransactionReceipt(req, result, url)
			if err != nil {
				log.Println("TestGetTransactionReceipt() error,cita.GetTransactionReceipt() error,", err)
				goRun = false
			} else {
				goRun = true
			}

            if len(result.Logs) > 0{
                log.Printf("transactionReceipt.Logs.Address is:%s",result.Logs[0].Address)
            }

			//log.Printf("hash: %s", result.Hash)
			//log.Printf("content: %s", result.Content)
			//log.Printf("from: %s\n", result.From)
			//log.Printf("blockNumber: %s\n", result.BlockNumber)
			//log.Printf("blockHash: %s\n", result.BlockHash)
			//log.Printf("index: %s\n\n\n\n", result.Index)
		}
	}()

	wg.Wait()
}

func main() {
	TestGetBlock()
}
