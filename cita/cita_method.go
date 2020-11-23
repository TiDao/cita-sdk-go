/***********************************************************************
# File Name: cita_method.go
# Author: TiDao
# mail: tidao2049@gmail.com
# Created Time: 2020-09-11 13:19:39
*********************************************************************/
package cita


import(
    "unsafe"
    "errors"
)

func byteToString(b []byte) string{
    return *(*string)(unsafe.Pointer(&b))
}

func byteToBool(b []byte) bool{
    return *(*bool)(unsafe.Pointer(&b))
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

    *Result = byteToString(response.Result)

    return nil
}

func Call(req *Request,Result *string,url string) error {

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

    *Result = byteToString(response.Result)

    return nil
}
func BlockNumber(req *Request,Result *string,url string) error {

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

    *Result = byteToString(response.Result)

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


func GetVersion(req *Request,Result *ResultVersion,url string) error {

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

func GetBlock(req *Request,Result *ResultBlock,url string) error {

    reqJson,err := req.MarshalJSON()
    if err != nil{
        return errors.New("request error:")
    }

    respJson,err := Post(url,reqJson)
    if err != nil{
        return errors.New("Post request error:")
    }

    var response = &Response{}
    err = response.UnmarshalJSON(respJson)
    if err != nil{
        return errors.New("Unmarshal response result error:")
    }

    err = Result.UnmarshalJSON(response.Result)
    if err != nil{
        return err
    }

    return nil
}

func GetLogs(req *Request,Result *[]ResultLogs,url string) error{

    reqJson,err := req.MarshalJSON()
    if err != nil{
        return err
    }

    respJson,err := Post(url,reqJson)
    if err != nil{
        return err
    }

    var response = &Logs{}
    err = response.UnmarshalJSON(respJson)
    if err != nil{
        return err
    }

    //err = Result.UnmarshalJSON(response.Result)
    //if err != nil{
    //    return err
    //}
    *Result = response.Result 

    return nil
}

func GetTransaction(req *Request,Result *ResultTransaction,url string) error{

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

func GetBalance(req *Request,Result *string,url string) error {

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

    *Result = byteToString(response.Result)

    return nil
}

func NewFilter(req *Request,Result *string,url string) error {

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

    *Result = byteToString(response.Result)

    return nil
}

func NewBlockFilter(req *Request,Result *string,url string) error {

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

    *Result = byteToString(response.Result)

    return nil
}

func UninstallFilter(req *Request,Result *bool,url string) error {

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

    *Result = byteToBool(response.Result)

    return nil
}

func GetFilterChanges(req *Request,Result *[]string,url string) error{

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

    //*Result = byteToString(response.Result)
    for i,results := range response.Result {
        (*Result)[i] = *(*string)(unsafe.Pointer(&results))
    }
    return nil
}

func GetFilterLogs(req *Request,Result *[]ResultLogs,url string) error{

    reqJson,err := req.MarshalJSON()
    if err != nil{
        return err
    }

    respJson,err := Post(url,reqJson)
    if err != nil{
        return err
    }

    var response = &Logs{}
    err = response.UnmarshalJSON(respJson)
    if err != nil{
        return err
    }

    *Result = response.Result
    return nil
}

func GetTransactionProof(req *Request,Result *string,url string) error {

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

    *Result = byteToString(response.Result)

    return nil
}

func GetBlockHeader(req *Request,Result *string,url string) error {

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

    *Result = byteToString(response.Result)

    return nil
}

func GetStateProof(req *Request,Result *string,url string) error {

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

    *Result = byteToString(response.Result)

    return nil
}

func GetStorageAt(req *Request,Result *string,url string) error {

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

    *Result = byteToString(response.Result)

    return nil
}

func GetMetaData(req *Request,Result *ResultMetaData,url string) error{

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
