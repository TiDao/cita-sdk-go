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
//    req := &cita.Request{
//        Jsonrpc: "2.0",
//        Method: "getBlockByNumber",
//        Params: []interface{}{
//            "",
//            true,
//        },
//        Id: 1,
//    }
//    result := new(cita.ResultBlock)

    var wg sync.WaitGroup
    ch := make(chan string,20)
//    for i := 0x0;i<100000000000;i++{
//        req.Params[0] = "0x"+strconv.FormatInt(<-ch,16)
//        log.Printf("the block height is %s",req.Params[0])
//        err := cita.GetBlock(req,result,url)
//        if err != nil{
//            log.Println(err)
//        }
//        if len(result.Body.Transactions) != 0 {
//            log.Println(*result)
//        }
//        time.Sleep(time.Second)
//    }
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
            Id: 1,
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
                        log.Println(Transaction)
                        log.Printf("\n\n\n")
                    }
                }
            }
        }()
    }

    wg.Wait()

}

func main(){
    TestGetBlock()
}
