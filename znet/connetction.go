package znet

import (
	"boom.com/bzinx/ziface"
	"fmt"
	"net"
)

type Connection struct {
	//当前连接的socket tpc 套接字
	Conn *net.TCPConn
	//当前连接的Id 也可以称作为SessionID ，ID全局唯一
	ConnID uint32
	//当前连接的关闭状态
	isClosed bool
	//该连接的处理方法router
	Router ziface.IRouter
	//告知该连接已经退出/停止channel
	ExitBuffChan chan bool

}

//处理conn读数据的goroutine
func (c *Connection)StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String()," conn reader exit!")
	defer c.Stop()
	for {

		//读取我们最大的数据到buf中
		buf:=make([]byte,512)
		_, err := c.Conn.Read(buf)
		if err!=nil{
			fmt.Println("recv buf err ",err)
			c.ExitBuffChan <- true
			continue
		}

		req:=Request{
			conn:c,
			data:buf,
		}
		//从路由Routers中找到注册绑定conn的对应handle
		go func(request ziface.IRequest) {
			//执行注册的路由方法
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)

		}(&req)
	}

}


//创建连接的方法
func NewConnection(conn *net.TCPConn,connID uint32,router ziface.IRouter )*Connection  {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:   router,
		ExitBuffChan: make(chan bool,1),
	}
	return  c

}
//启动连接，让当前连接开始工作
func (c *Connection) Start() {
	//开启处理该连接读取到客户端数据之后的请求业务
	go c.StartReader()

	for   {
		select {
		case <- c.ExitBuffChan:
			//得到退出消息，不再阻塞
			return

		}

	}

}
//停止连接，结束当前连接状态
func (c *Connection) Stop() {
	//1、如果当前连接已经关闭
	if c.isClosed==true{
		return
	}
	c.isClosed=true

	//TODO Connection Stop（）如果用户注册了该连接的关闭回调业务，那么在此应该显示调用
	//关闭socket连接
	c.Conn.Close()
	//通知从缓冲第队列读取数据的业务
	c.ExitBuffChan <-true

	//关闭连接全部管道
	close(c.ExitBuffChan)
}
//从当连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return  c.Conn
}

//获取当前连接ID
func (c *Connection) GetConnID() uint32 {
	return  c.ConnID
}
//获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}


