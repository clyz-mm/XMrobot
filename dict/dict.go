package dict

import "fmt"

// GetDictAll 获取所有成语字典
func GetDictAll() [][]string {
	filename := "cyyy.csv"
	return ReadCsv(filename, "idiom")
}

func GetJokeAll() [][]string {
	filename := "xh.csv"
	return ReadCsv(filename, "joke")
}

// ArrayToMap list转map 方便后续取数据
func ArrayToMap(array [][]string) (map[string]string, map[string][]string, map[string]string) {
	map1 := make(map[string]string)
	map2 := make(map[string][]string)
	map3 := make(map[string]string)

	for i := 1; i < len(array); i++ {
		map1[array[i][0]] = fmt.Sprintf("%s,%s", array[i][6], array[i][7])

		map2[array[i][6]] = append(map2[array[i][6]], array[i][0])

		// 成语解释map
		map3[array[i][0]] = fmt.Sprintf("拼音：%s\n成语解释：%s\n出自：%s", array[i][2], array[i][3], array[i][4])
	}
	return map1, map2, map3
}
