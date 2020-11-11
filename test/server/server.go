package main

import (
	"boom.com/bzinx/ziface"
	"boom.com/bzinx/znet"
	"fmt"
)

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

func (this * PingRouter)PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping..."))
	if err!=nil{
		fmt.Println("call back ping ping ping error")
	}
}

func (this * PingRouter)Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Handle ping..."))
	if err!=nil{
		fmt.Println("call back ping ping ping error")
	}

}
func (this * PingRouter)PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("PostHandle ping..."))
	if err!=nil{
		fmt.Println("call back ping ping ping errot")
	}

}
func main() {
	//创建一个server句柄
	s := znet.NewServe("[BZinx V0.3]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}