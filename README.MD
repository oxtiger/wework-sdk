# wework

golang企业微信sdk


## 特性
* wework对象参数均可自定义修改
* access token自动处理 (支持另起goroutine)
* 支持插入handler方法去hook所有请求response(如需在handler内向其他服务发送消息建议使用goroutine)
* 一个wework对象一个应用
* 扩展接口极为方便
* 可以使用IsWeworkError()方法判断返回的err是否为企业微信报错

## 示例

> 注: 使用New初始化时已自带返回错误码为 42001(token过期)/41001(token为空)/40014(token不合法)/-1(系统繁忙) 时的自动重试handler

```golang
package main

import (
	"log"
    
	"github.com/oxtiger/wework-sdk"
)

func exampleHandler(wx *wework.Client, api string, code int, data *[]byte) bool  {
	log.Printf("api_path: %s, errCode: %d, response: %b\n", api, code, data)
	if code != 0 {
		log.Fatalf("api response code: %d != 0\n", code)
		// retry this request
		return true
	}
	if api == wework.ApiUserSend {
		log.Printf("call user push api")
	}
	return false
}

func main()  {
	corpId := "id" // 企业ID
	corpSecret := "secret" // 企业secret
	var agentId int64 = 1 // "应用id"
	wx, err := wework.New(corpId, corpSecret, agentId)
	if err != nil {
		log.Fatalln(err)
	}
	// 增加自定义handler
	wx.Use(exampleHandler)
	// 设置handler返回true后的最大重试次数
	wx.SetRetry(3)

	invalidUser, err := wx.UserTextPush("user_id", "hello world")
	if err != nil {
		log.Fatalln(err)
	}
	if invalidUser != "" {
		log.Fatalf("failed to push message to: %s", invalidUser)
	}

	err = wx.ChatTextPush("chat_id", "hello world")
	if err != nil {
		log.Fatalln(err)
	}

	// 每小时更新一次token
	go wx.LoopUpdateToken()
}

```
