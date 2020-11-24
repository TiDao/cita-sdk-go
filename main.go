/***********************************************************************
# File Name: main.go
# Author: TiDao
# mail: tidao2049@gmail.com
# Created Time: 2020-11-24 12:28:58
*********************************************************************/
package main
import "log"
import "./cita"
import "strconv"
import "sync"

func TestGetBlock(){
    var url = "http://192.168.1.65:1337"

    var wg sync.WaitGroup
    ch := make(chan string,20)
    TransactionCh := make(chan string,20)
    go func(){
        for i := 0x0;i<1000000000;i++{
            ch <- "0x"+strconv.FormatInt(int64(i),16)
        }
    }()
    for i:=0;i<20;i++{
        wg.Add(1)
        req := &cita.Request{
            Jsonrpc: "2.0",
            Method: "getBlockByNumber",
            Params: []interface{}{
                "",
                true,
            },
            Id: int32(i),
        }
        result := new(cita.ResultBlock)
        go func(){
            for {
                req.Params[0] = <-ch 
                //log.Printf("the block height is %s",req.Params[0])
                err := cita.GetBlock(req,result,url)
                if err != nil{
                    log.Panic(err)
                }
                if len(result.Body.Transactions) != 0 {
                    for _,Transaction := range result.Body.Transactions{
                        TransactionCh <- Transaction.Hash
                    }
                }
            }
        }()
    }

    go func(){
        req := &cita.Request{
            Jsonrpc: "2.0",
            Method: "getTransaction",
            Params: []interface{}{
                "",
            },
            Id: 1,
        }

        result := &cita.ResultTransaction{}
        for {
            req.Params[0] = <-TransactionCh
            err := cita.GetTransaction(req,result,url)
            if err != nil{
                log.Panic(err)
            }
            log.Printf("hash: %s",result.Hash)
            log.Printf("content: %s",result.Content)
            log.Printf("from: %s\n",result.From)
            log.Printf("blockNumber: %s\n",result.BlockNumber)
            log.Printf("blockHash: %s\n",result.BlockHash)
            log.Printf("index: %s\n\n\n\n",result.Index)
        }
    }()


    wg.Wait()

}

func main(){
    TestGetBlock()
}
