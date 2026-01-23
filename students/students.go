// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func payHandler(w http.ResponseWriter, r *http.Request) {
// 	str := "Новый платеж обработан!"
// 	b := []byte(str)

// 	_, err := w.Write(b)
// 	if err != nil {
// 		fmt.Println("Произошла ошибка:", err.Error())
// 	} else {
// 		fmt.Println("Я отменил оплату!")
// 	}
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	str := "Hello, world!"
// 	b := []byte(str)

// 	_, err := w.Write(b)
// 	if err != nil {
// 		fmt.Println("Произошла ошибка:", err.Error())
// 	} else {
// 		fmt.Println("Я обработал HTTP запрос!")
// 	}
// }

// func main() {
// 	http.HandleFunc("/default", handler)
// 	http.HandleFunc("/pay", payHandler)

// 	fmt.Println("Запускаю HTTP сервер")

// 	err := http.ListenAndServe(":9091", nil)
// 	if err != nil {
// 		fmt.Println("Произошла ошибка:", err.Error())
// 	}

// 	fmt.Println("Программа закончила свое выполнение")
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strconv"
// 	"sync"
// )

// var mtx = sync.Mutex{}
// var money = 1000
// var bank = 0

// func payHandler(w http.ResponseWriter, r *http.Request) {
// 	httpRequestBody, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		msg := "fail to read HTTP body request" + err.Error()
// 		fmt.Println(msg)
// 		_, err := w.Write([]byte(msg))
// 		if err != nil {
// 			fmt.Println("Произошла ошибка", err)
// 		}
// 		return
// 	}
// 	httpRequestBodyString := string(httpRequestBody)

// 	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		msg := "fail to convert HTTP body to integer" + err.Error()
// 		fmt.Println(msg)
// 		_, err := w.Write([]byte(msg))
// 		if err != nil {
// 			fmt.Println("Произошла ошибка:", err.Error())
// 		}
// 		return
// 	}

// 	mtx.Lock()
// 	if money-paymentAmount >= 0 {
// 		money -= paymentAmount

// 		msg := "Оплата прошла успешно:" + strconv.Itoa(money)
// 		fmt.Println(msg)
// 		_, err := w.Write([]byte(msg))
// 		if err != nil {
// 			fmt.Println("Произошла ошибка", err)
// 		}
// 	} else {
// 		msg := "Не хватает денег на проведение оплаты"
// 		fmt.Println(msg)
// 		_, err := w.Write([]byte(msg))
// 		if err != nil {
// 			fmt.Println("Произошла ошибка")
// 		}
// 	}
// 	mtx.Unlock()
// }

// func saveHandler(w http.ResponseWriter, r *http.Request) {
// 	httpRequestBody, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		msg := "fail to read HTTP body request:" + err.Error()
// 		fmt.Println(msg)
// 		_, err := w.Write([]byte(msg))
// 		if err != nil {
// 			fmt.Println("Произошла ошибка:", err.Error())
// 			return
// 		}
// 		return
// 	}
// 	httpRequestBodyString := string(httpRequestBody)

// 	saveAmount, err := strconv.Atoi(httpRequestBodyString)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		msg := "fail to convert HTTP body to integer:" + err.Error()
// 		fmt.Println(msg)
// 		_, err := w.Write([]byte(msg))
// 		if err != nil {
// 			fmt.Println("Произошла ошибка:", err.Error())
// 		}
// 		return
// 	}

// 	mtx.Lock()
// 	if money >= saveAmount {
// 		money -= saveAmount
// 		bank += saveAmount

// 		msg := "Новое значение переменной money:" + strconv.Itoa(money)
// 		fmt.Println(msg)
// 		_, err := w.Write([]byte(msg))
// 		if err != nil {
// 			fmt.Println("Произошла ошибка", err)
// 		}

// 		msg1 := "Новое значение переменной bank:" + strconv.Itoa(bank)
// 		fmt.Println(msg1)
// 		_, err1 := w.Write([]byte(msg1))
// 		if err1 != nil {
// 			fmt.Println("Произошла ошибка", err1)
// 		}

// 		fmt.Println("Новое значение переменной money:", money)
// 		fmt.Println("Новое значение переменной bank:", bank)
// 	} else {
// 		msg := "Не хватает денег на проведение оплаты"
// 		fmt.Println(msg)
// 		_, err := w.Write([]byte(msg))
// 		if err != nil {
// 			fmt.Println("Произошла ошибка")
// 		}
// 	}
// 	mtx.Unlock()
// }

// func main() {
// 	http.HandleFunc("/pay", payHandler)
// 	http.HandleFunc("/save", saveHandler)

// 	err := http.ListenAndServe(":9091", nil)
// 	if err != nil {
// 		fmt.Println("HTTP server error:", err)
// 	}
// }

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")
	fmt.Println("foo:", fooParam)
	fmt.Println("boo:", booParam)
}

func main() {
	http.HandleFunc("/default", handler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("failed to run HTTP server:", err)
	}
}
