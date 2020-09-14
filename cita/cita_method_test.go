/***********************************************************************
# File Name: cita_method_test.go
# Author: TiDao
# mail: tidao2049@gmail.com
# Created Time: 2020-09-11 14:10:19
*********************************************************************/
package cita

import "testing"
import "log"

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
