package wechat

import "context"

type Sender struct {
}

func (s Sender) Do(ctx context.Context, text string) error {
	return nil
}
