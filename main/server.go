package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrade = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)


//http.Handle回调函数
func wsHandler(w http.ResponseWriter,r *http.Request){

	var (
		conn *websocket.Conn
		err error
		messageType int
		data []byte
	)
	//把resp写回消息
	//w.Write([]byte("hello"))
	//收到http请求 转化协议为websocket
	//将header转化为upgrade,判断返回值,error不为空时，有错误，断开连接return
	if conn,err = upgrade.Upgrade(w,r,nil);err!=nil{
		return
	}

	//收发coon数据
	for{
		//获取conn连接数据
		if messageType,data,err = conn.ReadMessage(); err != nil {
			conn.Close()
		}
		//消息发回
		if err = conn.WriteMessage(messageType,data);err!=nil{
			conn.Close()
		}
	}

}

func main()  {
	//http服务端
	http.HandleFunc("/ws",wsHandler)
	//启动服务端
	http.ListenAndServe("0.0.0.0:7777",nil)

}

