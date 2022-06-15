package wework

import "testing"

// https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=15be2cab-6fc2-42e9-a135-99595acf19e3

const robot = "15be2cab-6fc2-42e9-a135-99595acf19e3"

func TestRobotPushText(t *testing.T) {
	err := RobotPushText(robot, "> test", "xxx@xxxxxx.com")
	if err != nil {
		t.Error(err)
	}
	err = RobotPushText(robot, "> test")
	if err != nil {
		t.Error(err)
	}
}

func TestRobotPushMarkdown(t *testing.T) {
	err := RobotPushMarkdown(robot, "> test")
	if err != nil {
		t.Error(err)
	}
	err = RobotPushMarkdown(robot, "> test")
	if err != nil {
		t.Error(err)
	}
}

func TestRobotPushFile(t *testing.T) {
	mediaId, err := RobotUploadFile(robot, "/Users/twq/Downloads/峰值低于5%的机器总表-0519.xlsx", "test.xlsx")
	if err != nil {
		t.Error(err)
	}
	err = RobotPushFile(robot, mediaId)
	if err != nil {
		t.Error(err)
	}
}
