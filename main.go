/***********************************************************************
# File Name: main.go
# Author: TiDao
# mail: tidao2049@gmail.com
# Created Time: 2020-11-24 12:28:58
*********************************************************************/
package main
import "log"
import "./cita"
//import "strconv"
import "io"
import "sync"
//import "time"
//import "strings"

func TestGetBlock(){
    var url = "http://192.168.1.65:1337"

    var wg sync.WaitGroup
    ch := make(chan string,20)
    TransactionCh := make(chan string,20)

    go func(){
        //for i := 0x1000000;i<1000000000;i++{
        //    ch <- "0x"+strconv.FormatInt(int64(i),16)
        //    log.Println("starting block number grow")
        //}
        req := &cita.Request{
            Jsonrpc: "2.0",
            Method: "blockNumber",
            Params: []interface{}{},
            Id: 1,
        }
        var Error cita.Error
        var err   error
        //var lock sync.Mutex
        var blockNumber string
        var temp string
        for {
            Error,err = cita.BlockNumber(req,&blockNumber,url)

            if err != nil || Error.Code != 0 {
                log.Println(err)
                log.Println(Error)
            }

            if blockNumber != temp{
                temp = blockNumber
                ch <- temp
            }
            //time.Sleep(time.Second*1)
        }
    }()
    for i:=0;i<30;i++{
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
                log.Printf("the block height is %s",req.Params[0])
                //log.Println(req)
                Error,err := cita.GetBlock(req,result,url)
                //log.Println(result)
                if err != nil{
                    if err == io.EOF{
                        log.Printf("the err is EOF")
                    }else{
                        log.Panic(err)
                    }
                }
                if Error.Code == 0 {
                    if len(result.Body.Transactions) != 0 {
                        for _,Transaction := range result.Body.Transactions{
                            TransactionCh <- Transaction.Hash
                        }
                    }
                }else{
                    log.Println("get Block error:",Error)
                }
                //if len(result.Body.Transactions) != 0 {
                //    for _,Transaction := range result.Body.Transactions{
                //        TransactionCh <- Transaction.Hash
                //    }
                //}
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
            _,err := cita.GetTransaction(req,result,url)
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
