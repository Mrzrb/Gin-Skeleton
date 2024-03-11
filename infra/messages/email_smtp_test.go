package messages

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

const (
	EmailName     = "go_skeleton@163.com"
	EmailPassPop3 = "MFZNIPMCGVLAVUUE"
	EmailServer   = "smtp.163.com"
	EmailPort     = 25
)

func TestSendSmtpEmail(t *testing.T) {
	client := NewSmtpEmail(EmailServer, EmailName, EmailPassPop3, EmailPort)
	convey.Convey("Send basic email test", t, func() {
		m := NewEmail("test send email", "<h1>This</h1> is a test email from go skeleton", []string{"mrzhangrb@outlook.com"})
		err := client.Send(m)
		convey.So(err, convey.ShouldBeNil)
	})
}
