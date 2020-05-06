package bestbuy_com

import (
	"github.com/Gssssssssy/ns-stored/pkg/config"
	"github.com/Gssssssssy/ns-stored/pkg/log"
	"github.com/pkg/errors"
)

const HOST = "https://api.bestbuy.com"

var Token string

func init() {
	tk, err := token()
	if err != nil {
		log.Errorf(nil, errors.Cause(err).Error())
	}
	Token = tk
}

// 获取开发者 token
func token() (token string, err error) {
	cfg := config.Config()
	token = cfg.GetString("bestbuy_developer_key")
	if token == "" {
		err = errors.New("bestbuy.com api key not found")
		return "", errors.WithStack(err)
	}
	return token, nil
}
