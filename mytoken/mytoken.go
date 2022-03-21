package mytoken

import "fmt"

const (
	TypeBot    string = "Bot"
	TypeNormal string = "Bearer"
)

type MyToken struct {
	AppID       uint64
	AccessToken string
	Type        string
}

func BotToken(appID uint64, accessToken string) *MyToken {
	return &MyToken{
		AppID:       appID,
		AccessToken: accessToken,
		Type:        TypeBot,
	}
}

// New 返回一个新myToken对象
func New(tokenType string) *MyToken {
	return &MyToken{
		Type: tokenType,
	}
}

func (t *MyToken) GetString() string {
	if t.Type == TypeNormal {
		return t.AccessToken
	}
	return fmt.Sprintf("%v.%s", t.AppID, t.AccessToken)
}
