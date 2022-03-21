package process

import (
	"XMrobot/dict"
	"XMrobot/instructions"
	"XMrobot/mybotgo/mydto"
	"XMrobot/mybotgo/mydto/message"
	"XMrobot/myopenapi"
	"context"
	"fmt"
	"log"
	"strings"
)

type Processor struct {
	Api myopenapi.OpenAPI
}

var play = false
var gameName string
var usedMap = make(map[string]string)
var dictList = dict.GetDictAll()
var mapString, mapList, mapDesc = dict.ArrayToMap(dictList)
var currentIdiom string

// ProcessMessage 消息处理
func (p Processor) ProcessMessage(input string, data *mydto.WSATMessageData) error {
	ctx := context.Background()
	// 获取命令
	cmd := strings.Replace(message.ParseCommand(input).Cmd, "/", "", -1)
	toCreate := &mydto.MessageToCreate{
		Content: "你好，有什么吩咐？",
		MessageReference: &mydto.MessageReference{
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
	}
	var temp string
	// 判断当前是否为游玩模式
	if play && cmd != "结束" && cmd != "提示" {
		temp = cmd
		cmd = gameName
	}
	if strings.Contains(cmd, "笑话") {
		cmd = "讲笑话"
	}

	switch cmd {
	case "你好":
		p.sendMsg(ctx, data.ChannelID, toCreate)
	case "菜单":
		toCreate.Content = "成语接龙" + message.Emoji(13) + "\t|\t" +
			"讲笑话" + message.Emoji(13) + "\n" +
			"成语解释" + message.Emoji(13) + "\t|\t"
		p.sendMsg(ctx, data.ChannelID, toCreate)
	case "讲笑话":
		toCreate.Content = instructions.GetOneJoke() + message.Emoji(20)
		p.sendMsg(ctx, data.ChannelID, toCreate)
	case "成语解释":
		if !play {
			play = true
			gameName = "成语解释"
			toCreate.Content = "再次@我并输入想要了解的成语吧"
		} else {
			toCreate.Content = instructions.GetIdiomDesc(temp, mapDesc)
		}
		p.sendMsg(ctx, data.ChannelID, toCreate)
	case "成语接龙":
		if !play {
			play = true
			gameName = "成语接龙"
			currentIdiom = instructions.RandomIdiom(dictList)
			usedMap[currentIdiom] = currentIdiom
			toCreate.Content = currentIdiom
			p.sendMsg(ctx, data.ChannelID, toCreate)
		} else {
			_, contains := usedMap[temp]
			if contains {
				toCreate.Content = fmt.Sprintf("亲，这个成语[%s]已经被使用了哦，换一个吧！", temp)
				p.sendMsg(ctx, data.ChannelID, toCreate)
				break
			}
			idiom := instructions.GetNextIdiom(currentIdiom, temp, usedMap, mapString, mapList)
			if idiom == "-1" {
				toCreate.Content = "呃，不要逗我玩了，这不是一个成语"
			} else if idiom == "0" {
				toCreate.Content = "刚才的成语没有接出来哦"
			} else if idiom == "win" {
				// 用户赢得比赛，并退出游玩模式
				play = false
				gameName = ""
				toCreate.Content = "恭喜你，赢得了这场游戏，真厉害！"
			} else {
				currentIdiom = idiom
				usedMap[currentIdiom] = currentIdiom
				usedMap[temp] = temp
				toCreate.Content = currentIdiom
			}
			p.sendMsg(ctx, data.ChannelID, toCreate)
		}
	case "结束":
		toCreate.Content = "呃，那这次就先玩到这里吧，欢迎下次再来玩：" + gameName
		play = false
		gameName = ""
		p.sendMsg(ctx, data.ChannelID, toCreate)
	case "提示":
		if play && gameName == "成语接龙" {
			toCreate.Content = fmt.Sprintf("小提示：可以跟我说以 %s 开头的成语哦", instructions.Tips(currentIdiom, mapString))
			p.sendMsg(ctx, data.ChannelID, toCreate)
		} else {
			toCreate.Content = "成语接龙游戏还没开始呢"
			p.sendMsg(ctx, data.ChannelID, toCreate)
		}
	default:
		toCreate.Content = "可以@我并输入：菜单 查看所有项目哦"
		p.sendMsg(ctx, data.ChannelID, toCreate)
	}

	return nil
}

// 发送消息
func (p Processor) sendMsg(ctx context.Context, channelID string, toCreate *mydto.MessageToCreate) {
	_, err := p.Api.PostMessage(ctx, channelID, toCreate)
	if err != nil {
		log.Println(err)
	}
}
