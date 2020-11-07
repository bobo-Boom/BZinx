package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

/*
模拟客户端
*/

func ClientTest()  {
	fmt.Println("Client Test .... start")
	// 3秒后发起测试请求，给服务器开启服务的机会
	time.Sleep(3*time.Second)
	dial, err := net.Dial("tcp", "127.0.0.1:7777")
	if err!=nil{
		fmt.Println("client start err ,exit !")
		return
	}

	for{
		n, err := dial.Write([]byte("Hello BZINX"))
		if err!=nil{
			fmt.Println("write error err",err)
			return
		}
		buf:=make([]byte,512)
		cnt , err := dial.Read(buf)
		if err!=nil{
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("server call back : %s ,cnt =%d\n",buf[:n],cnt)
		time.Sleep(1*time.Second)

	}

}

func TestServer(t *testing.T) {
	/*
	  服务端测试
	*/
	//1\创建一个server句柄 s
	s:=NewServe("[BZinx V0.2]")

	go ClientTest()

	s.Serve()

}
