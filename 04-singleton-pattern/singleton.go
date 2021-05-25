package singleton

import (
	"fmt"
	"sync"
)

/*
	单类模式严格一个类只有一个实例，并提供一个全局的访问接口
	*设计思想
		1.声明一个全局变量
		2.多线程考虑线程安全，引入sync.Once
*/
var once sync.Once
var helper *DocumentHelper

type DocumentHelper struct {
	// ...一些线程安全の成员
}

func GetDocumentHelper() *DocumentHelper {
	if helper != nil {
		fmt.Println("已经存在，直接获取...")
		return helper
	}

	once.Do(func() {
		fmt.Println("第一次初始化...")
		helper = &DocumentHelper{}
	})
	return helper
}
