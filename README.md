# yunpian 云片网短信平台Golang API

智能模板模式

### Usage

```go
sms := NewSMS("https://sms.yunpian.com/v1/sms/send.json", "********************************")
timeStr := time.Now().Format("2006-01-02 15:04:05")
event := "测试短信"
text := fmt.Sprintf("事件提醒：您将于%s进行%s", timeStr, event)
log.Greenln(text)
info, err := sms.Send("159********", text)
```

### License

MIT License
