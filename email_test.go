package itools

import (
	"fmt"
	"os"
	"testing"
)

func TestSendEmail(t *testing.T) {
	tests := map[string]struct {
		test func(t *testing.T)
	}{
		"Send163Mail":       {testSend163Mail},
		"SendQQMail":        {testSendQQMail},
		"Send126Mail":       {testSend126Mail},
		"SendEmailStartTls": {testSendEmailStartTls},
		"SendGmail":         {testSendGmail},
	}

	t.Parallel()
	for name, tt := range tests {
		t.Run(name, tt.test)
	}
}

func testSend163Mail(t *testing.T) {
	from := os.Getenv("Mail163From")
	to := "" // 接收人的邮箱
	mailPassword := os.Getenv("Mail163Pass")
	mailServer := os.Getenv("Mail163Server")
	mailServerPort := os.Getenv("Mail163ServerPort")
	title := "单元测试，邮件发送"
	content := `能发送过去吗:<h1>" + 123456 + "</h1>`
	err := SendMail(from, to, title, content, mailServer, mailServerPort, mailPassword)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("邮件发送成功...")
}

func testSendQQMail(t *testing.T) {
	from := os.Getenv("MailQQFrom")
	to := "" // 接收人的邮箱
	mailPassword := os.Getenv("MailQQPass")
	mailServer := os.Getenv("MailQQServer")
	mailServerPort := os.Getenv("MailQQServerPort")
	title := "单元测试，邮件发送"
	content := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>IM注册邮件</title>
</head>
<style>
    .mail{
        margin: 0 auto;
        border-radius: 45px;
        height: 400px;
        padding: 10px;
        background-color: #CC9933;
        background: url("http://images.caixiaoxin.cn/note.png") no-repeat;
    }
    .code {
        color: #f6512b;
        font-weight: bold;
        font-size: 30px;
        padding: 2px;
    }
</style>
<body>
<div class="mail">
    <h3>您好:您正在测试qq邮箱发送!</h3>
    <p>下面是您的验证码:</p>
        <p class="code">%s</p>
        <p>请注意查收!谢谢</p>
</div>
<h3>如果可以请给项目点个star～<a target="_blank" href="https://github.com/mazezen/itools">项目地址</a> </h3>
</body>
</html>`, "123456")
	err := SendMail(from, to, title, content, mailServer, mailServerPort, mailPassword)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("邮件发送成功...")
}

func testSend126Mail(t *testing.T) {
	from := os.Getenv("Mail126From")
	to := "1752676696@qq.com" // 接收人的邮箱
	mailPassword := os.Getenv("Mail126Pass")
	mailServer := os.Getenv("Mail126Server")
	mailServerPort := os.Getenv("Mail126ServerPort")
	title := "单元测试，邮件发送"
	content := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>IM注册邮件</title>
</head>
<style>
    .mail{
        margin: 0 auto;
        border-radius: 45px;
        height: 400px;
        padding: 10px;
        background-color: #CC9933;
        background: url("http://images.caixiaoxin.cn/note.png") no-repeat;
    }
    .code {
        color: #f6512b;
        font-weight: bold;
        font-size: 30px;
        padding: 2px;
    }
</style>
<body>
<div class="mail">
    <h3>您好:您正在测试qq邮箱发送!</h3>
    <p>下面是您的验证码:</p>
        <p class="code">%s</p>
        <p>请注意查收!谢谢</p>
</div>
<h3>如果可以请给项目点个star～<a target="_blank" href="https://github.com/mazezen/itools">项目地址</a> </h3>
</body>
</html>`, "123456")
	err := SendMail(from, to, title, content, mailServer, mailServerPort, mailPassword)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("邮件发送成功...")
}

func testSendEmailStartTls(t *testing.T) {
	from := os.Getenv("MailGmailFrom")
	to := "" // 接收人的邮箱
	mailPassword := os.Getenv("MailGmailPass")
	mailServer := os.Getenv("MailGmailServer")
	mailServerPort := os.Getenv("MailGmailServerPort")
	title := "单元测试，邮件发送"
	content := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>IM注册邮件</title>
</head>
<style>
    .mail{
        margin: 0 auto;
        border-radius: 45px;
        height: 400px;
        padding: 10px;
        background-color: #CC9933;
        background: url("http://images.caixiaoxin.cn/note.png") no-repeat;
    }
    .code {
        color: #f6512b;
        font-weight: bold;
        font-size: 30px;
        padding: 2px;
    }
</style>
<body>
<div class="mail">
    <h3>您好:您正在测试qq邮箱发送!</h3>
    <p>下面是您的验证码:</p>
        <p class="code">%s</p>
        <p>请注意查收!谢谢</p>
</div>
<h3>如果可以请给项目点个star～<a target="_blank" href="https://github.com/mazezen/itools">项目地址</a> </h3>
</body>
</html>`, "123456")
	err := SendGmailEmail(from, to, title, content, mailServer, mailServerPort, mailPassword)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("邮件发送成功...")
}

func testSendGmail(t *testing.T) {
	from := os.Getenv("MailGmailFrom")
	to := []string{
		"", // 接收人的邮箱1
		"", // 接收人的邮箱2
	}
	mailPassword := os.Getenv("MailGmailPass")
	mailServer := os.Getenv("MailGmailServer")
	mailServerPort := os.Getenv("MailGmailServerPort")
	content := []byte(fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>IM注册邮件</title>
</head>
<style>
    .mail{
        margin: 0 auto;
        border-radius: 45px;
        height: 400px;
        padding: 10px;
        background-color: #CC9933;
        background: url("http://images.caixiaoxin.cn/note.png") no-repeat;
    }
    .code {
        color: #f6512b;
        font-weight: bold;
        font-size: 30px;
        padding: 2px;
    }
</style>
<body>
<div class="mail">
    <h3>您好:您正在测试qq邮箱发送!</h3>
    <p>下面是您的验证码:</p>
        <p class="code">%s</p>
        <p>请注意查收!谢谢</p>
</div>
<h3>如果可以请给项目点个star～<a target="_blank" href="https://github.com/mazezen/itools">项目地址</a> </h3>
</body>
</html>`, "123456"))

	err := SendGmail(from, to, content, mailServer, mailServerPort, mailPassword)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("邮件发送成功...")
}
