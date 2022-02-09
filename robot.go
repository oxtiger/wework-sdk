package wework

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	json "github.com/json-iterator/go"
)

const robotApi = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="

// robotPost http
func robotPost(robotId string, data interface{}) (respBytes []byte, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json.Marshal(data)")
		return nil, err
	}
	req, err := http.NewRequest("POST", robotApi+robotId, bytes.NewBuffer(b))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return respBytes, err
}

// RobotPushText 推送文本消息
// robotId		 机器人id
// content		 消息内容
// atUser		 需要@的user_id
func RobotPushText(robotId string, content string, atUser ...string) error {
	msg := &RobotTextMsg{MsgType: "text"}
	if len(atUser) > 0 {
		msg.Text.MentionedList = atUser
	}
	msg.Text.Content = content
	res, err := robotPost(robotId, msg)
	if err != nil {
		return err
	}
	var errRes errResponse
	err = json.Unmarshal(res, &errRes)
	if err != nil {
		return err
	}
	return nil
}

// RobotPushMarkdown 推送markdown消息
// robotId		 机器人id
// content		 消息内容
// atUser		 需要@的user_id
func RobotPushMarkdown(robotId string, content string) error {
	msg := &RobotMarkdownMsg{MsgType: "markdown"}
	msg.Markdown.Content = content
	res, err := robotPost(robotId, msg)
	if err != nil {
		return err
	}
	var errRes errResponse
	err = json.Unmarshal(res, &errRes)
	if err != nil {
		return err
	}
	return nil
}
