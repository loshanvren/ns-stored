package email

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Gssssssssy/ns-stored/internal/task"
	"github.com/Gssssssssy/ns-stored/pkg/config"
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
	"html/template"
	"io/ioutil"
	"path"
)

var cfg config.Provider
var ServicePoint *Sender

func init() {
	cfg = config.Config()
	ServicePoint = NewSender()
}

type Sender struct {
	Conn *gomail.Dialer
}

func NewSender() *Sender {
	host := cfg.GetString("email_sender.smtp.server")
	user := cfg.GetString("email_sender.smtp.usr")
	pass := cfg.GetString("email_sender.smtp.pwd")
	if host == "" || user == "" || pass == "" {
		panic(fmt.Errorf("email sender config args not found"))
	}
	port := 465
	useTLS := cfg.GetBool("email_sender.smtp.tls")
	if useTLS {
		port = 587
	}
	sender := new(Sender)
	sender.Conn = gomail.NewDialer(host, port, user, pass)
	//sender.Conn.TLSConfig = &tls.Config{InsecureSkipVerify: useTLS}
	return sender

}

func (s *Sender) Do(_ context.Context, result *task.Result) error {
	var (
		author    = cfg.GetString("email_sender.smtp.usr")
		addressee = cfg.GetStringSlice("email_list")
		text      string
		err       error
	)
	if author == "" || len(addressee) == 0 {
		return errors.WithStack(fmt.Errorf("email body config not found"))
	}

	text, err = generateTextContent(result)
	if err != nil {
		return errors.WithStack(err)
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", author)
	msg.SetHeader("To", addressee...)
	msg.SetHeader("Subject", "NS Monitor")
	msg.SetBody("text/html", text)

	if err = s.Conn.DialAndSend(msg); err != nil {
		return errors.Wrapf(err, "failed to send email")
	}

	return nil
}

func generateTextContent(result *task.Result) (string, error) {
	rootPath := cfg.GetString("root_path")
	if rootPath == "" {
		rootPath = "/ns"
	}
	buf := bytes.NewBuffer([]byte{})
	tplPath := path.Join(rootPath, "asset", "static", "email_template.xhtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return "", errors.WithStack(err)
	}
	err = tpl.Execute(buf, result)
	if err != nil {
		return "", errors.WithStack(err)
	}
	content, readErr := ioutil.ReadAll(buf)
	if readErr != nil {
		return "", errors.WithStack(readErr)
	}
	return string(content), nil
}
