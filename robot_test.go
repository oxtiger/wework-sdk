package wework

import "testing"

// https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=15be2cab-6fc2-42e9-a135-99595acf19e3

const robot = "xx1xx2xx-xxx3-xxx4-xxx5-99595acf19e3"

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
