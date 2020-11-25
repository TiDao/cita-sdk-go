/***********************************************************************
# File Name: cita_method_test.go
# Author: TiDao
# mail: tidao2049@gmail.com
# Created Time: 2020-09-11 14:10:19
*********************************************************************/
package cita

import "testing"
import "log"
import "strconv"
//import "time"

func hex2int(val string) int{
    n,err := strconv.ParseUint(val,16,32)
    if err != nil{
        log.Fatal(err)
    }

    return int(n)
}

func TestPeerCount(t *testing.T){

    var url = "http://192.168.1.65:1337"
    req := &Request{
         Jsonrpc: "2.0",
         Method: "peerCount",
         Params: []interface{}{},
         Id: 1,
    }

    var result =new(string)
    _,err := PeerCount(req,result,url)
    if err != nil{
        t.Error(err)
    }
    log.Println(*result)
}

func TestPeersInfo(t *testing.T){

    var url = "http://192.168.1.65:1337"
    req := &Request{
         Jsonrpc: "2.0",
         Method: "peersInfo",
         Params: []interface{}{},
         Id: 1,
    }

    var result =new(ResultPeerInfo)
    _,err := PeersInfo(req,result,url)
    if err != nil{
        t.Error(err)
    }
    log.Println(result)
}


func TestGetLogs(t *testing.T){

    var url = "http://192.168.1.65:1337"
    req := &Request{
         Jsonrpc: "2.0",
         Method: "getLogs",
         Params: []interface{}{},
         Id: 1,
    }

    var result =new([]ResultLogs)
    _,err := GetLogs(req,result,url)
    if err != nil{
        t.Error(err)
    }
    log.Println(result)
}

func TestBlockNumber(t *testing.T){

    var url = "http://192.168.1.65:1337"
    req := &Request{
         Jsonrpc: "2.0",
         Method: "blockNumber",
         Params: []interface{}{},
         Id: 1,
    }

    var result =new(string)
    _,err := BlockNumber(req,result,url)
    if err != nil{
        t.Error(err)
    }
    log.Println(hex2int((*result)[3:len(*result)-1]))
}

func TestCall(t *testing.T) {
    var url = "http://192.168.1.65:1337"
    req := &Request{
         Jsonrpc: "2.0",
         Method: "blockNumber",
         Params: []interface{}{
             CallRequest{
                 From:"0xca35b7d915458ef540ade6068dfe2f44e8fa733c",
                 To:"0xea4f6bc98b456ef085da5c424db710489848cab5",
                 Data:"0x6d4ce63c",
             },
             "latest"},
         Id: 1,
    }
    result := new(string)
    _,err := Call(req,result,url)
    if err != nil{
        t.Error(err)
    }
    log.Println(*result)
}

func TestGetBlock(t *testing.T){
    var url = "http://192.168.1.65:1337"
    req := &Request{
        Jsonrpc: "2.0",
        Method: "getBlockByNumber",
        Params: []interface{}{
            "",
            true,
        },
        Id: 1,
    }

    result := &ResultBlock{}
    //response := Error{}
    for i := 0x0;i<100000;i++{
        req.Params[0] = "0x"+strconv.FormatInt(int64(i),16) 
        log.Printf("the block height is %s",req.Params[0])
        response,err := GetBlock(req,result,url)
        if err != nil{
            t.Error(err)
        }

        if len(response.Message) == 0{
            if len(result.Body.Transactions) != 0 {
                log.Println(*result)
            }
        }else{
            log.Println(response)
        }
    }
}

