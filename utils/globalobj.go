package utils

import (
	"boom.com/bzinx/ziface"
	"encoding/json"
	"io/ioutil"
)

/*
 存储一切有关BZinx框架的全局参数，供其他模块使用
 一些参数也可以通过 用户根据 BZinx。
*/
type GlobalObj struct {
	//当前服务器名称
	Name string
	//当前Bzinx全局Server对象
	TcpSercer ziface.IServer
	//当前主机服务器IP
	Host string
	//当前服务器主机监听端口
	TcpPort int
	//当前Zinx版本号
	Version string
	//允许数据包最大的值
	MaxPackeitSize uint32
	//当前服务器主机允许的最大链接个数
	MaxConn int
}

//定义一个全局的对象
//目的是让其他模块都能访问里面的参数
var GlobalObject *GlobalObj

//提供init初始化方法
//目的是初始化GlobalObject对象，和加载服务端配置文件conf/bzinx.json

//读取用户的配置文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/bzinx.json")
	if err != nil {
		panic(err)
	}
	//将json数据解析到struct中
	json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

//提供init方法，默认加载
func init() {
	//初始化GlobalObject变量，设置一些值
	GlobalObject = &GlobalObj{
		Name:           "BZinxServerApp",
		Host:           "0.0.0.0",
		TcpPort:        7777,
		Version:        "V0.4",
		MaxPackeitSize: 4096,
		MaxConn:        12000,
	}
	GlobalObject.Reload()
}
