package main

// import (
// 	"fmt"
// 	"net/http"
// )

// const port = ":9090"

// func main() {
// 	//使用go提供的原生API开启web服务
// 	http.HandleFunc("/hello", helloWorld)
// 	err := http.ListenAndServe(port, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// }

// func helloWorld(respW http.ResponseWriter, req *http.Request) {
// 	n, err := fmt.Fprint(respW, "hello world")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("n: %v\n", n)
// }
