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
import "time"

//import "strings"

//func getNumber() 获取链上的最新块高，
//比对当前读取到额块高和最新块高，
//如果有差值就将其传递到channel中
//宕机自动恢复
//startNumber 从数据库中获取的最后的块高
//ch 传递块高
//url 链rpc地址
func getNumber(ch chan string, url string, startNumber string) {
	req := &cita.Request{
		Jsonrpc: "2.0",
		Method:  "blockNumber",
		Params:  []interface{}{},
		Id:      1,
	}
	var Error cita.Error //增加了prefix的error类型
	var err error
	var latestNumber string //获取的最新块高

	//将startNumber转换为Int64类型
	startNumberInt, err := strconv.ParseInt(startNumber[2:], 16, 64)
	if err != nil {
		log.Println("change startNumber to Int64 failed,", err)
	}
	startNumberInt = startNumberInt + 1

	for {
		Error, err = cita.BlockNumber(req, &latestNumber, url)
		if err != nil || Error.Code != 0 {
            log.Printf("get blockNumber the Error is %v,the err is %s",Error,err.Error())
		}

		latestNumberInt, err := strconv.ParseInt(latestNumber[2:], 16, 64)
		if err != nil {
			log.Println("change latestNumber to Int64 failed,", err)
		}

		//尽快传递与最新块高间的差值
		for startNumberInt < latestNumberInt-1{
			ch <- "0x" + strconv.FormatInt(startNumberInt, 16)
			startNumberInt = startNumberInt + 1
		}

		//控制链请求次数查询次数
		time.Sleep(time.Second * 1)
	}
}
