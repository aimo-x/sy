package tools

import (
	"testing"
)

func TestSmtp(T *testing.T) {
	var sm SMTP
	sm.Subject = "主题sss"
	sm.Title = "副主题sss"
	sm.Desc = "描述sss"
	sm.Name = "姓名sss"
	sm.Email = "494745sss409@qq.com"
	sm.Position = "技术顾sss问"
	sm.Company = "googlssse"
	sm.BtnLink = "https://iuu.pub/sssdsadfs"
	sm.BinText = "查看更多"
	sm.MailTo = []string{"494745409@qq.com"}
	err := sm.Send()
	if err != nil {
		T.Error(err)
		// panic(err)
	}
}
