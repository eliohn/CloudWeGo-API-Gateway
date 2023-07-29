package main

import (
	"fmt"
	"net/http"

	/* 	"net/http/httptest" */
	"bytes"
	"encoding/json"
	"io/ioutil"
	_ "net/http/pprof"
	"testing"

	"math/rand"
	"strconv"
	"strings"
	//"unsafe"
)

const (
	aUrlAdd  = "http://127.0.0.1:8888/gateway/FirstLevelCalService/add"
	aUrlSub  = "http://127.0.0.1:8888/gateway/FirstLevelCalService/sub"
	bUrlMul  = "http://127.0.0.1:8888/gateway/SecondLevelCalService/mul"
	bUrlDiv  = "http://127.0.0.1:8888/gateway/SecondLevelCalService/div"
	cUrlFact = "http://127.0.0.1:8888/gateway/AdvancedCalService/fact"
	cUrlFib  = "http://127.0.0.1:8888/gateway/AdvancedCalService/fib"
)

// Benchmarck性能基准测试，主要靠此产生数据
func BenchmarkFirstLevelServiceAdd(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(100)
		b := rand.Intn(100)

		info["operand_1"] = a
		info["operand_2"] = b

		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", aUrlAdd, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("send fail", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Read rep fail", err.Error())
			return
		}
		//对返回的信息进行处理
		str := (string)(respBytes)

		var theInd1 = strings.Index(str, "data")
		var theInd2 = strings.Index(str, "message")
		var str0 = str[theInd1+6 : theInd2-2]

		resultNumber1 := a + b
		resultNumber2, _ := strconv.Atoi(str0)

		as := Assert(resultNumber1, int(resultNumber2))
		if as {
			fmt.Printf("The value deliver to FirstLevelCalservice/Add are %d and %d ,expect %d, get %d,test passed\n", a, b, a+b, int(resultNumber2))
		} else {
			fmt.Printf("The value deliver to FirstLevelCalservice/Add are %d and %d ,expect %d, get %d,test failed\n", a, b, a+b, int(resultNumber2))
		}
		request.Body.Close()
	}
}

func BenchmarkFirstLevelServiceSub(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(100)
		b := rand.Intn(100)

		info["operand_1"] = a
		info["operand_2"] = b

		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", aUrlSub, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("send fail", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Read rep fail", err.Error())
			return
		}
		//对返回的信息进行处理
		str := (string)(respBytes)
		var theInd1 = strings.Index(str, "data")
		var theInd2 = strings.Index(str, "message")
		var str0 = str[theInd1+6 : theInd2-2]

		resultNumber1 := a - b
		resultNumber2, _ := strconv.Atoi(str0)

		as := Assert(resultNumber1, int(resultNumber2))
		if as {
			fmt.Printf("The value deliver to FirstLevelCalservice/Sub are %d and %d ,expect %d, get %d,test passed\n", a, b, a-b, int(resultNumber2))
		} else {
			fmt.Printf("The value deliver to FirstLevelCalservice/Sub are %d and %d ,expect %d, get %d,test failed\n", a, b, a-b, int(resultNumber2))
		}
		request.Body.Close()
	}
}

func BenchmarkSecondLevelServiceMul(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(10)
		b := rand.Intn(10)

		info["operand_1"] = a
		info["operand_2"] = b

		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", bUrlMul, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("send fail", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Read rep fail", err.Error())
			return
		}
		//对返回的信息进行处理
		str := (string)(respBytes)

		var theInd1 = strings.Index(str, "data")
		var theInd2 = strings.Index(str, "message")
		var str0 = str[theInd1+6 : theInd2-2]

		resultNumber1 := a * b
		resultNumber2, _ := strconv.Atoi(str0)

		as := Assert(resultNumber1, int(resultNumber2))
		if as {
			fmt.Printf("The value deliver to SecondLevelCalservice/Mul are %d and %d ,expect %d, get %d,test passed\n", a, b, a*b, int(resultNumber2))
		} else {
			fmt.Printf("The value deliver to SecondLevelCalservice/Mul are %d and %d ,expect %d, get %d,test failed\n", a, b, a*b, int(resultNumber2))
		}
		request.Body.Close()
	}
}

func BenchmarkSecondLevelServiceDiv(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(9) + 1
		b := rand.Intn(9) + 1

		info["operand_1"] = a
		info["operand_2"] = b

		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", bUrlDiv, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("send fail", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Read rep fail", err.Error())
			return
		}
		//对返回的信息进行处理
		if b == 0 {
			fmt.Println("除数为 0 ")
		} else {
			str := (string)(respBytes)
			var theInd1 = strings.Index(str, "data")
			var theInd2 = strings.Index(str, "message")
			var str0 = str[theInd1+6 : theInd2-2]

			resultNumber1 := a / b
			resultNumber2, _ := strconv.Atoi(str0)

			as := Assert(resultNumber1, int(resultNumber2))
			if as {
				fmt.Printf("The value deliver to SecondLevelCalservice/Div are %d and %d ,expect %d, get %d,test passed\n", a, b, a/b, int(resultNumber2))
			} else {
				fmt.Printf("The value deliver to SecondLevelCalservice/Div are %d and %d ,expect %d, get %d,test failed\n", a, b, a/b, int(resultNumber2))
			}
		}

		request.Body.Close()
	}
}

func BenchmarkAdvancedLevelServiceFact(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(10)
		info["operand"] = a
		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", cUrlFact, reader)
		defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复

		if err != nil {
			fmt.Println("send fail", err.Error())
			return
		}

		respBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Read rep fail", err.Error())
			return
		}
		//对返回的信息进行处理

		str := (string)(respBytes)
		var theInd1 = strings.Index(str, "data")
		var theInd2 = strings.Index(str, "message")
		var str0 = str[theInd1+6 : theInd2-2]

		resultNumber1 := Fact(a)
		resultNumber2, _ := strconv.Atoi(str0)

		as := Assert(resultNumber1, int(resultNumber2))
		if as {
			fmt.Printf("The value deliver to AdvancedCalservice/Fact is %d ,expect %d, get %d, test passed\n", a, resultNumber1, int(resultNumber2))
		} else {
			fmt.Printf("The value deliver to AdvancedCalservice/Fact is %d ,expect %d, get %d, test failed\n", a, resultNumber1, int(resultNumber2))
		}

		request.Body.Close()
	}
}

func BenchmarkAdvancedLevelServiceFib(b *testing.B) {
	for i := 1; i < b.N; i++ {
		info := make(map[string]int)

		a := rand.Intn(5)
		info["operand"] = a
		bytesData, err := json.Marshal(info)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", cUrlFib, reader)
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
		//对返回的信息进行处理

		str := (string)(respBytes)
		var theInd1 = strings.Index(str, "data")
		var theInd2 = strings.Index(str, "message")
		var str0 = str[theInd1+6 : theInd2-2]

		resultNumber1 := Fib(a)
		resultNumber2, _ := strconv.Atoi(str0)

		as := Assert(resultNumber1, int(resultNumber2))
		if as {
			fmt.Printf("The value deliver to AdvancedCalservice/Fib is %d ,expect %d, get %d, test passed\n", a, resultNumber1, int(resultNumber2))
		} else {
			fmt.Printf("The value deliver to AdvancedCalservice/Fib is %d ,expect %d, get %d, test passed\n", a, resultNumber1, int(resultNumber2))
		}

		request.Body.Close()
	}
}

func Assert(a int, b int) bool {
	return (a == b)
}
func Fact(a int) int {
	b := 1
	for i := 1; i <= a; i++ {
		b = b * i
	}
	return b
}
func Fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
