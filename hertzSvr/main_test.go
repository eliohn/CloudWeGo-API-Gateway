package main

import (
	"fmt"
	"net/http"
	/* 	"net/http/httptest" */
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"
	//"time"
	"unsafe"
)

const (
	url = "http://127.0.0.1:8888/request"
)

//var httpCli = &http.Client{Timeout: 3 * time.Second}

// Test测试用例，但是test主要是针对逻辑的测试，咱需要的是性能数据方面，所以我感觉用处不大
func TestAService(t *testing.T) {
	for i := 1; i <= 100; i++ {

	}
}

// Benchmarck性能基准测试，主要靠此产生数据
func BenchmarkAService(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]string)
		info["svrName"] = "BService" //这里只是以B做最初的测试，后续会在循环里向不同方向发送请求的
		info["password"] = "******"

		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(bytesData)

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", url, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("发送请求失败", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Read回复失败", err.Error())
			return
		}

		//byte数组直接转成string，优化内存
		str := (*string)(unsafe.Pointer(&respBytes))
		fmt.Println("testOK", *str)

		request.Body.Close()
	}
}

// Assert asserts cond is true, otherwise fails the test.断言
func Assert(t testingTB, cond bool, val ...interface{}) {
	t.Helper()
	if !cond {
		if len(val) > 0 {
			val = append([]interface{}{"assertion failed:"}, val...)
			t.Fatal(val...)
		} else {
			t.Fatal("assertion failed")
		}
	}
}

// testingTB is a subset of common methods between *testing.T and *testing.B.
type testingTB interface {
	// 等同于 t.Log + t.FailNow
	Fatal(args ...interface{})

	// 等同于 t.Logf + t.FailNow
	Fatalf(format string, args ...interface{})

	// 将调用函数标记标记为测试助手函数
	Helper()
}
