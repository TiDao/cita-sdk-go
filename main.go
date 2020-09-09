/***********************************************************************
# File Name: main.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2020-09-09 15:09:28
*********************************************************************/
package main

import(
    "encoding/json"
    "./HttpMethod"
    "./cita"
    "fmt"
    "os"
    "bytes"
)

func main(){
    url := "http://192.168.1.65:1337"
    request := &cita.Request{
        Jsonrpc: "2.0",
        Method: "getBlockByNumber",
        Params: ["0xF9",true],
        Id: 1,
    }

    requestJson,err := json.Marshal(request)
    if err != nil{
        fmt.Println(err)
    }
    responseJson := &[]byte{}

    HttpMethod.Post(url,requestJson,responseJson)

    var response = &cita.Response{}

    err = json.Unmarshal(*responseJson,response)
    if err != nil{
        fmt.Println(err)
    }

    fmt.Println(response)

    var out bytes.Buffer

    err = json.Indent(&out,*responseJson,"","\t")
    out.WriteTo(os.Stdout)
}
