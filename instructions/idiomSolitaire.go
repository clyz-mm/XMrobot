package instructions

import (
	"math/rand"
	"strings"
	"time"
)

// RandomIdiom 随机获取一个成语
func RandomIdiom(dictList [][]string) string {
	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	// 从第 1 行开始到最后一行
	return dictList[rand.Intn(len(dictList))][0]
}

// GetNextIdiom 根据用户输入的成语获取下一个成语接龙的成语
func GetNextIdiom(currentIdiom string, idiom string, usedMap map[string]string, mapString map[string]string, mapList map[string][]string) string {
	// 用户输入的成语 首字拼音和末字拼音
	firstAndLast, contains := mapString[idiom]
	if !contains {
		return "-1"
	}
	// 机器人给的成语 首字拼音和末字拼音
	cuFirstAndLast := mapString[currentIdiom]
	// 判断是否符合接龙规则
	if strings.Split(cuFirstAndLast, ",")[1] != strings.Split(firstAndLast, ",")[0] {
		return "0"
	}
	idiomList, contains := mapList[strings.Split(firstAndLast, ",")[1]]
	if !contains {
		return "win"
	}
	for i := 0; i < len(idiomList); i++ {
		_, contains := usedMap[idiomList[i]]
		if !contains {
			return idiomList[i]
		}
	}

	return "win"
}

// Tips 根据当前的成语给出末拼音提示
func Tips(currentIdiom string, mapString map[string]string) string {
	cuFirstAndLast := mapString[currentIdiom]
	return strings.Split(cuFirstAndLast, ",")[1]
}

func GetIdiomDesc(idiom string, mapDesc map[string]string) string {
	desc, contains := mapDesc[idiom]

	if contains {
		return desc
	}
	return "呃，不要逗我玩了，这不是一个成语，如果不想玩了可以跟我说：结束"
}
