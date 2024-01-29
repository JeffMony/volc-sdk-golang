package imagex

// client 这个文件由 Codegen 生成，但后续不再更新（如果没有，则会重新生成），包含配置结构体和创建服务实例的代码
// 开发者可以在这里给服务结构体添加自定义扩展方法

import (
	"fmt"

	common "github.com/volcengine/volc-sdk-golang/base"
)

var DefaultInstance = NewInstance()

type ImageX struct {
	*common.Client
}

func NewInstance() *ImageX {
	return NewInstanceWithRegion("cn-north-1")
}

func NewInstanceWithRegion(region string) *ImageX {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		panic(fmt.Errorf("ImageX not support region %s", region))
	}
	instance := &ImageX{
		Client: common.NewClient(&serviceInfo, ApiListInfo),
	}
	return instance
}
