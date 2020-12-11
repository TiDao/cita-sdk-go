/***********************************************************************
# File Name: http_method.go
# Author: TiDao
# mail: tidao2049@gmail.com
# Created Time: 2020-04-20 16:41:52
*********************************************************************/
package cita

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	//"strconv"
)

func Post(url string, input []byte) ([]byte, error) {

	var client = &http.Client{
        Timeout: 60 * time.Second,
    }
    //defer client.CloseIdleConnections()
	reader := bytes.NewReader(input)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	//t := strconv.FormatInt(time.Now().Unix(),10)
	//request.Header.Set('Accept':'application/json', 'Content-Type':'application/json')
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("Timestamp",t)

    //begin := time.Now()
	response, err := client.Do(request)
    //if int64(time.Since(begin)) > int64(time.Microsecond*100000) {
    //  log.Println(int64(time.Since(begin)))
    //  log.Println(int64(time.Microsecond * 1000))
    //  log.Println("time:", time.Since(begin))
    //}
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
        log.Println("response.Status:",response.Status)
        log.Println("response.Header:",response.Header)
		return nil, err
	}

	return responseBody, nil
}

func Get(url string) ([]byte, error) {

	var client = &http.Client{}
    defer client.CloseIdleConnections()

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	reponse, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	reponseBody, err := ioutil.ReadAll(reponse.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return reponseBody, nil
}
