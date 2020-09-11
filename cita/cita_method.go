/***********************************************************************
# File Name: cita_method.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2020-09-11 13:19:39
*********************************************************************/
package cita

import(
    "unsafe"
)

func String(b []byte) string{
    return *(*string)(unsafe.Pointer(&b))
}

func PeerCount(req *Request,Result *string,url string) error {

    reqJson,err := req.MarshalJSON()
    if err != nil{
        return err
    }

    respJson,err := Post(url,reqJson)
    if err != nil{
        return err
    }

    var response = &Response{}
    err = response.UnmarshalJSON(respJson)
    if err != nil{
        return err
    }

    *Result = String(response.Result)

    return nil
}


func PeersInfo(req *Request,Result *ResultPeerInfo,url string) error {

    reqJson,err := req.MarshalJSON()
    if err != nil{
        return err
    }

    respJson,err := Post(url,reqJson)
    if err != nil{
        return err
    }

    var response = &Response{}
    err = response.UnmarshalJSON(respJson)
    if err != nil{
        return err
    }

    err = Result.UnmarshalJSON(response.Result)
    if err != nil{
        return err
    }

    return nil
}
