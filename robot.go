package wework

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	json "github.com/json-iterator/go"
)

const robotApi = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
const robotUploadApi = "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?type=file&key="

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

// RobotPushFile 推送文件消息
// robotId		 机器人id
// mediaId		 文件id
func RobotPushFile(robotId string, mediaId string) error {
	msg := &RobotFileMsg{MsgType: "file"}
	msg.File.MediaId = mediaId
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

// RobotUploadFile 上传文件至机器人空间, 返回media_id
// robotId		机器人id
// filename		文件路径
func RobotUploadFile(robotId, filepath, filename string) (string, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	fileWriter, err := writer.CreateFormFile("media", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return "", err
	}

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file")
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(fileWriter, f)
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", robotUploadApi, robotId), payload)

	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var errRes RobotUploadResponse
	err = json.Unmarshal(body, &errRes)
	if err != nil {
		return "", err
	}
	if errRes.MediaId == "" {
		return "", fmt.Errorf("failed to upload files")
	}
	return errRes.MediaId, nil
}
