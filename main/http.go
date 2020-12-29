package main



import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",index)//index为向url发送请求时调用的函数
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hello world")
}