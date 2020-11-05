package znet

import (
	"boom.com/bzinx/ziface"
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
		
		//启动server网络连接业务
		for  {
			
			//阻塞等待客户端建立连接请求
			conn, err := listenner.Accept()
			if err!=nil{
				fmt.Println("Accept err ",err)
				return
			}

			//1、TODO Server.start()设置服务器最大连接控制，如果超过最大连接，那么则关闭此新的连接

			//2、TODO Server.start()处理该连接请求的业务方法，此时应该有handler和conn绑定的

			//我们暂时做一个最大512字节的回显服务
			go func() {
				//不断的循环从客户端获取数据
				for   {
				   buf :=make([]byte,512)
					cnt, err := conn.Read(buf)
					if err!=nil{
						fmt.Println("rec buf err",err)
						continue
					}
					if _,err:=conn.Write(buf[:cnt]);err!=nil{
						fmt.Println("write back buf err",err)
						continue
					}
				}
			}()
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
/*
 创建一个服务句柄
*/

func NewServe (name string ) ziface.IServer  {
	s:=&Server{
		Name:name,
		IPVersion:"tcp4",
		Ip:"0.0.0.0",
		Port:7777,
	}
	return s

}