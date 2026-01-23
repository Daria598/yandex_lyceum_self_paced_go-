package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var mtx = sync.Mutex{}
var money = atomic.Int64{}
var bank = atomic.Int64{}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP body request:", err)
		return
	}
	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to integer:", err)
		return
	}

	mtx.Lock()
	if money.Load()-int64(paymentAmount) >= 0 {
		time.Sleep(3 * time.Second)
		money.Add(int64(-paymentAmount))
		fmt.Println("Оплата прошла успешно", money.Load())
	} else {
		fmt.Println("Не хватает денег на проведение оплаты")
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP body request:", err)
		return
	}
	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to integer:", err)
		return
	}

	mtx.Lock()
	if money.Load() >= int64(saveAmount) {
		money.Add(int64(-saveAmount))
		bank.Add(int64(saveAmount))
		fmt.Println("Новое значение переменной money:", money.Load())
		fmt.Println("Новое значение переменной bank:", bank.Load())
	} else {
		fmt.Println("Не хватает денег, чтобы положить в копилку")
		return
	}
	mtx.Unlock()
}

func main() {
	money.Add(1000)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error:", err)
	}
}
