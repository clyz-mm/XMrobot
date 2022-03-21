package main

import (
	"XMrobot/mybotgo"
	"XMrobot/mybotgo/mydto"
	"XMrobot/mybotgo/mydto/message"
	"XMrobot/mybotgo/websocket"
	"XMrobot/mytoken"
	"XMrobot/process"
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"strings"
	"syscall"
	"time"
)

var processor process.Processor

// 入口
func main() {
	configName := "robot-config.yaml"
	// 获取配置文件中的 appId 和 token 信息
	appId, token, err := getConfigInfo(configName)
	if err != nil {
		log.Fatal(err)
	}

	botToken := mytoken.BotToken(appId, token)

	// 沙箱
	//api := mybotgo.NewSandboxOpenAPI(botToken).WithTimeout(3 * time.Second)
	// 正式
	api := mybotgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)

	ctx := context.Background()
	// 获取 websocket 信息
	wsInfo, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatal(err)
	}

	processor = process.Processor{Api: api}

	websocket.RegisterResumeSignal(syscall.SIGUSR1)
	// 根据不同的回调，生成 intents
	intent := websocket.RegisterHandlers(
		// at 机器人事件
		ATMessageEventHandler(),
	)

	err = mybotgo.NewSessionManager().Start(wsInfo, botToken, &intent)
	if err != nil {
		log.Fatal(err)
	}

}

// 获取配置文件中的信息
func getConfigInfo(fileName string) (uint64, string, error) {
	// 获取当前go程调用栈所执行的函数的文件和行号信息
	// 忽略pc和line
	_, filePath, _, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("runtime.Caller(1) 读取失败")
	}
	file := fmt.Sprintf("%s/%s", path.Dir(filePath), fileName)
	var conf struct {
		AppID uint64 `yaml:"appid"`
		Token string `yaml:"token"`
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Print("ioutil.ReadFile() 读取失败")
		return 0, "", err
	}
	if err = yaml.Unmarshal(content, &conf); err != nil {
		log.Print("yaml.Unmarshal(content, &conf) 读取失败")
		return 0, "", err
	}
	return conf.AppID, conf.Token, nil
}

// ATMessageEventHandler 实现处理 at 消息的回调
func ATMessageEventHandler() websocket.ATMessageEventHandler {
	return func(event *mydto.WSPayload, data *mydto.WSATMessageData) error {
		input := strings.ToLower(message.ETLInput(data.Content))
		return processor.ProcessMessage(input, data)
	}
}
