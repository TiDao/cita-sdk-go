/***********************************************************************
# File Name: main.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2020-09-09 15:09:28
*********************************************************************/
package main

import(
    "./cita"
    "fmt"
//    "os"
//    "bytes"
    "log"
    "net/rpc/jsonrpc"
    "net/rpc"
    "net"
)

func main(){
    request := &cita.Request{
        Jsonrpc: "2.0",
        Method: "peersInfo",
        Params: []interface{}{},
        Id: 1,
    }
    fmt.Println(request)
    reqJson,err := request.MarshalJSON()

    conn,err := net.Dial("tcp","192.168.1.65:1337")
    if err != nil{
        log.Fatal("Dail:",err)
    }

    client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

    var reply string

    err = client.Call("",reqJson,&reply)
    if err != nil{
        log.Fatal("Call:",err)
    }

    fmt.Println(reply)
    //requestJson,err := json.Marshal(request)
    //if err != nil{
    //    fmt.Println(err)
    //}
    //responseJson := &[]byte{}

    //*responseJson,err = cita.Post(url,requestJson)
    //if err != nil{
    //    log.Fatal(err)
    //}

    //var response = &cita.Response{}

    //err = json.Unmarshal(*responseJson,response)
    //if err != nil{
    //    fmt.Println(err)
    //}

    //fmt.Println(string(*responseJson))

    //var out bytes.Buffer

    //err = json.Indent(&out,*responseJson,"","\t")
    //fmt.Println(out.String())
    //out.WriteTo(os.Stdout)
}
