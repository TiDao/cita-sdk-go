/***********************************************************************
# File Name: http_method.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2020-04-20 16:41:52
*********************************************************************/
package HttpMethod

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
    //"time"
    //"strconv"
)

func Post(url string, input []byte, output *[]byte) error {

	var client = &http.Client{}


	reader := bytes.NewReader(input)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Println(err.Error())
		return err
	}
    

    //t := strconv.FormatInt(time.Now().Unix(),10)
    //request.Header.Set('Accept':'application/json', 'Content-Type':'application/json')
    request.Header.Set("Accept","application/json")
    request.Header.Set("Content-Type","application/json")
    //request.Header.Set("Timestamp",t)

	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	*output = responseBody
	return nil
}

func Get(url string, output *[]byte) error {

	var client = &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return  err
	}

	reponse, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	reponseBody, err := ioutil.ReadAll(reponse.Body)
    if err != nil{
        log.Println(err.Error())
        return err
    }

	*output = reponseBody
    return nil
}
