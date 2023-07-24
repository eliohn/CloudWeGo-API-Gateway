package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	Areq_URL = "localhost:8888/A-req"
	Breq_URL = "localhost:8888/B-req"
	Creq_URL = "localhost:8888/C-req"
	Dreq_URL = "localhost:8888/D-req"
)

var httpCli = &http.Client{Timeout: 3 * time.Second}

// Test测试用例，但是test主要是针对逻辑的测试，咱需要的是性能数据方面，所以我感觉用处不大
func TestAService(t *testing.T) {
	for i := 1; i <= 100; i++ {

	}
}

// Benchmarck性能基准测试，主要靠此产生数据
func BenchmarkAService(b *testing.B) {
	for i := 1; i < b.N; i++ {

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
