package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main(){
	http.HandleFunc("/",index)//index为向url发送请求时调用的函数
	//监听8001端口，拉去代码
	http.ListenAndServe("0.0.0.0:8001",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	//自动拉起任务
	autopull()
}

func autopull()  {
	cmd := exec.Command("sh","./deploy.sh")
	err := cmd.Start()
	if err != nil{
		log.Fatal(err)
	}
	err = cmd.Wait()
}




