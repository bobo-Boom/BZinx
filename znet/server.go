package znet

import (
	"boom.com/bzinx/ziface"
	"errors"
	"fmt"
	"net"
	"time"
)

// iserver接口实现，定义一个server服务类
type Server struct {
	//服务器的名称
	Name string
	//服务器的地址
	Ip string
	//tcp4 or other
	IPVersion string
	//服务器的端口
	Port int
	//当前Server由用户绑定的回调router，也就是Server注册
	//的链接对应的处理业务
	Router ziface.IRouter
}

//定义当前客户端连接的handle api
func CallBackToClient(conn *net.TCPConn,data []byte ,cnt int)error  {
	//回显业务
	fmt.Println("[Conn Handle] CallBackToClient...")
	if _, err := conn.Write(data[:cnt]);err!=nil{
		fmt.Println("write back buf err ",err)
		return errors.New("CallBackToClient error")
	}
	return nil

}

func (s *Server) Start ()  {
	fmt.Printf("[Start] Server listenner ai IP : %s, port : %d, is starting\n",s.Ip,s.Port)

	//开启一个go去做服务器端的listen业务
	go func() {
		//获取一个TCP的addr
		addr,err:=net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.Ip,s.Port))
		if err!=nil{
			fmt.Println("resolve tcp addr err: " ,err)
			return
		}

		//监听服务器地址
		listenner,err:=net.ListenTCP(s.IPVersion,addr)
		if err!=nil{
			fmt.Println("listen",s.IPVersion,"err",err)
			return
		}
		fmt.Println("start BZinx Server    ",s.Name,"success,now listenning")

		//TODO server.go 应该有一个自动生成ID的方法
		var cid uint32
		cid =0

		//启动server网络连接业务
		for  {
			
			//1、阻塞等待客户端建立连接请求
			conn, err := listenner.Accept()
			if err!=nil{
				fmt.Println("Accept err ",err)
				return
			}

			//2、TODO Server.start()设置服务器最大连接控制，如果超过最大连接，那么则关闭此新的连接

			//3、处理该连接请求的业务方法，此时应该有handler和conn绑定的
			dealConn := NewConnection(conn.(*net.TCPConn), cid, s.Router)
			cid++

			go  dealConn.Start()
		}

	}()

}

func (s * Server)Stop()  {
	fmt.Println("[STOP] BZinx server , name ",s.Name)

	//TODO Server.Stop() 将其他需要清理的连接信息或者其他信息，也要一并停止或者清理


}

func (s *Server)Serve()  {
	s.Start()


	//TODO Server.Serve() 是否在启动服务的时候，还要处理其他的事情的 可以在这里添加

	//阻塞，否则主go退出，listenner的go将回退出
	for   {
		time.Sleep(10*time.Second)
	}
}

func (s *Server)AddRouter(router ziface.IRouter){
	s.Router=router
	fmt.Println("Add Router success ! ")
}
/*
 创建一个服务句柄
*/

func NewServe (name string ) ziface.IServer  {
	s:=&Server{
		Name:name,
		IPVersion:"tcp4",
		Ip:"0.0.0.0",
		Port:7777,
		Router:nil,
	}
	return s

}