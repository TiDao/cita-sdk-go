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
    err := PeerCount(req,result,url)
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
    err := PeersInfo(req,result,url)
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
    err := GetLogs(req,result,url)
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
    err := BlockNumber(req,result,url)
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
    err := Call(req,result,url)
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
            "0x9999",
            true,
        },
        Id: 1,
    }

    result := new(ResultBlock)
    err := GetBlock(req,result,url)
    if err != nil{
        t.Error(err)
    }
    log.Println(*result)
}

//func TestBool(t *testing.T){
//    s1 := make([]byte,0)
//    s1 = nil
//    test := Bool(s1)
//
//    log.Println(test)
//}

//func BenchmarkPeerCount(b *testing.B){
//    var url = "http://192.168.1.65:1337"
//    req := &Request{
//         Jsonrpc: "2.0",
//         Method: "peerCount",
//         Params: []interface{}{},
//         Id: 1,
//    }
//
//    var result = new(string)
//    for i:=0;i<b.N;i++{
//        _ = PeerCount(req,result,url) 
//    }
//
//}
