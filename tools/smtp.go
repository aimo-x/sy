package tools

import (
	"strconv"

	"gopkg.in/gomail.v2"
)

// SMTP ...
type SMTP struct {
	MailTo   []string // 发送目标
	Subject  string   // 邮寄主题
	Title    string   // 邮寄标题
	Desc     string   // 邮件描述
	Name     string   // 名字
	Email    string   // 邮件
	Position string   // 职位
	Company  string   // 公司
	BtnLink  string   // 按钮链接
	BinText  string   // 按钮文本

}

// Send use smtp 邮件
func (sm *SMTP) Send() (err error) {
	body := sm.TEM()
	err = sm.SendMail(sm.MailTo, sm.Subject, string(body))
	if err != nil {
		return
	}
	return
}

// SendMail 发送邮件
func (sm *SMTP) SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	/*
		mailConn := map[string]string{
			"user": "noreply@oovmi.com",
			"pass": "oysrefmlxnihbhjf",
			"host": "smtp.qq.com",
			"port": "465",
		}
	*/
	mailConn := map[string]string{
		"user": "noreply@oovmi.com",
		"pass": "oysrefmlxnihbhjf",
		"host": "smtp.qq.com",
		"port": "465",
	}
	/*
		mailConn := map[string]string{
			"user": "mukong@oovmi.com",
			"pass": "S526@qba$",
			"host": "smtp.office365.com",
			"port": "587",
		}
	*/
	port, err := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	if err != nil {
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "XD Game"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                            //发送给多个用户
	m.SetHeader("Subject", subject)                         //设置邮件主题
	m.SetBody("text/html", body)                            //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err = d.DialAndSend(m)
	return err

}

// TEM  邮件模板
func (sm *SMTP) TEM() (tem []byte) {
	tem = []byte(`<!--<![endif]-->
	<table class="wrapper" style="border-collapse: collapse;table-layout: fixed;min-width: 320px;width: 100%;background-color: #fff;">
			<tbody>
					<tr class="firstRow">
							<td style="word-break: break-all;">
									<p>
											<!--[if(mso)|(IE)]><table align="center"class="preheader"cellpadding="0"cellspacing="0"role="presentation"><tr><td style="width: 280px"valign="top"><![endif]--><!--[if(mso)|(IE)]></td><td style="width: 280px"valign="top"><![endif]--><!--[if(mso)|(IE)]></td></tr></table><![endif]-->
									</p>
									<p>
											<!--[if(mso)|(IE)]><table align="center"cellpadding="0"cellspacing="0"role="presentation"><tr class="layout-fixed-width"style="background-color: #ffffff;"><td style="width: 600px"class="w560"><![endif]-->
									</p>
									<h1 style="Margin-top: 0;Margin-bottom: 20px;font-style: normal;font-weight: normal;color: #404040;font-size: 28px;line-height: 36px;text-align: center;">
											` + sm.Subject + `
									</h1>
									<p>
											&nbsp;
									</p>
									<h2 style="Margin-top: 0;Margin-bottom: 0;font-style: normal;font-weight: normal;color: #706f70;font-size: 18px;line-height: 26px;font-family: Cabin,Avenir,sans-serif;">
											<strong>` + sm.Title + `</strong>
									</h2>
									<p style="Margin-top: 16px;Margin-bottom: 20px;">
											` + sm.Desc + `
									</p>
									<p>
											&nbsp;
									</p>
									<h5 style="Margin-top: 0;Margin-bottom: 0;font-style: normal;font-weight: normal;color: #706f70;font-size: 12px;line-height: 26px;font-family: Cabin,Avenir,sans-serif;">
											<strong>姓名：` + sm.Name + `</strong>
									</h5>
									<h5 style="Margin-top: 0;Margin-bottom: 0;font-style: normal;font-weight: normal;color: #706f70;font-size: 12px;line-height: 26px;font-family: Cabin,Avenir,sans-serif;">
											<strong>职位名称：` + sm.Position + `</strong>
									</h5>
									<h5 style="Margin-top: 0;Margin-bottom: 0;font-style: normal;font-weight: normal;color: #706f70;font-size: 12px;line-height: 26px;font-family: Cabin,Avenir,sans-serif;">
											<strong>所在公司：` + sm.Company + `</strong>
									</h5>
									<p>
											&nbsp;
									</p>
									<h5 style="Margin-top: 0;Margin-bottom: 0;font-style: normal;font-weight: normal;color: #706f70;font-size: 12px;line-height: 26px;font-family: Cabin,Avenir,sans-serif;">
											<strong>Email：` + sm.Email + `</strong>
									</h5>
									<p>
											<br/>
									</p>
									<div class="" style="text-align:center;">
											<a href="` + sm.BtnLink + `" id="` + sm.BtnLink + `" target="_blank" rel="noopener" style="border-radius: 4px;display: inline-block;font-size: 14px;font-weight: bold;line-height: 24px;padding: 6px 24px;text-align: center;text-decoration: none !important;transition: opacity 0.1s ease-in;color: #ffffff !important;background-color: #e45d6b;font-family: Open Sans, sans-serif;">` + sm.BinText + `</a>
                      <!--[if mso]><p style="line-height:0;margin:0;">&nbsp;</p><v:roundrect xmlns:v="urn:schemas-microsoft-com:vml" href="http://test.com"style="width:160px"arcsize="9%"fillcolor="#E45D6B"stroke="f"><v:textbox style="mso-fit-shape-to-text:t"inset="0px,11px,0px,11px"><center style="font-size:14px;line-height:24px;color:#FFFFFF;font-family:Open Sans,sans-serif;font-weight:bold;mso-line-height-rule:exactly;mso-text-raise:4px">Take our survey</center></v:textbox></v:roundrect><![endif]-->
                    </div>
									<p>
											<!--[if(mso)|(IE)]><table align="center"cellpadding="0"cellspacing="0"role="presentation"><tr class="layout-email-footer"><td style="width: 600px;"class="w560"><![endif]-->
									</p>
									<p>
											<br/>
									</p>
									<p>
											<!--[if(mso)|(IE)]></td></tr></table><![endif]-->
									</p>
									<p>
											&nbsp;
									</p>
							</td>
					</tr>
			</tbody>
	</table>
	<p>
			<br/>
	</p>`)
	return
}
