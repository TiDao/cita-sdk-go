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

//import "strings"

func TestGetBlock() {
	var url = "http://192.168.1.65:1337"

	var wg sync.WaitGroup
	ch := make(chan string, 20)
	TransactionCh := make(chan string, 20)
	go getNumber(ch, url, "0x9d59c")

	for i := 0; i < 20; i++ {
		wg.Add(1)
		req := &cita.Request{
			Jsonrpc: "2.0",
			Method:  "getBlockByNumber",
			Params: []interface{}{
				"",
				true,
			},
			Id: int32(i),
		}
        go func() {
            result := new(cita.ResultBlock)
            var goRun bool = true
            for {
                if goRun {
                    req.Params[0] = <-ch
                }

                Error, err := cita.GetBlock(req, result, url)
				log.Printf("the block height is %s ", req.Params[0])
                //log.Printf("the result is %v\n\n",result)
				if err != nil {
					log.Println("TestGetBlock() error,cita.GetBlock() error,",err)
                    goRun = false
				}else{
                    goRun = true
                }
				if Error.Code == 0 {
					if len(result.Body.Transactions) != 0 {
						for _, Transaction := range result.Body.Transactions {
							TransactionCh <- Transaction.Hash
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
			Method:  "getTransaction",
			Params: []interface{}{
				"",
			},
			Id: 1,
		}

		result := &cita.ResultTransaction{}
        var goRun = true
		for {
            if goRun{
			    req.Params[0] = <-TransactionCh
            }
			_, err := cita.GetTransaction(req, result, url)
			if err != nil {
				log.Println("TestGetBlock() error,cita.GetTransaction() error,",err)
                goRun = false
			}else{
                goRun = true
            }
			log.Printf("hash: %s", result.Hash)
			log.Printf("content: %s", result.Content)
			log.Printf("from: %s\n", result.From)
			log.Printf("blockNumber: %s\n", result.BlockNumber)
			log.Printf("blockHash: %s\n", result.BlockHash)
			log.Printf("index: %s\n\n\n\n", result.Index)
		}
	}()

	wg.Wait()
}

func main() {
	TestGetBlock()
}
