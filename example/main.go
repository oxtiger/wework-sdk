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
