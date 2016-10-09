package yunpian

import (
	"fmt"
	"github.com/gogather/com/log"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test_Com(t *testing.T) {
	Convey("Send Message Test", t, func() {
		sms := NewSMS("https://sms.yunpian.com/v1/sms/send.json", "********************************")
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		event := "测试短信"
		text := fmt.Sprintf("【elapen】事件提醒：您将于%s进行%s", timeStr, event)
		log.Greenln(text)
		info, err := sms.Send("15989553147", text)
		if err != nil {
			log.Println(err)
		} else {
			log.Redln(info)
		}

	})
}
