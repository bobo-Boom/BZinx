package znet

import "boom.com/bzinx/ziface"

//实现router时，先继承这个基类，然后根据需要对这个基类的方法进行重写
type BaseRouter struct {}

//这里之所以BaseRouter的方法都为空
//是因为Router不希望有PreHandle或者PostHandle
//所以Router全部继承BaseRouter的好处是，不需要实现PreHandle
//和PostHandle也可以实例化
func (b *BaseRouter) PreHandle(request ziface.IRequest) {}

func (b *BaseRouter) Handle(request ziface.IRequest) {}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {}