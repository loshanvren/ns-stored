package alarm

import (
	"context"
	"github.com/Gssssssssy/ns-stored/internal/task"
	"github.com/Gssssssssy/ns-stored/pkg/alarm/email"
)

const (
	TypeEmail    Type = "email"
	TypeTelegram Type = "telegram"
	TypeWeChat   Type = "wechat"
)

type Type string

func (t Type) String() string {
	return string(t)
}

func (t Type) Make() Sender {
	return defaultConfigTool[t]
}

type Sender interface {
	Do(ctx context.Context, result *task.Result) error
}

var defaultConfigTool map[Type]Sender

func init() {
	defaultConfigTool = make(map[Type]Sender, 0)
	defaultConfigTool[TypeEmail] = &email.Sender{}
	//defaultConfigTool[TypeTelegram] = &telegram.Sender{}
	//defaultConfigTool[TypeWeChat] = &wechat.Sender{}
}

