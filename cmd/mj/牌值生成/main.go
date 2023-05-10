package main

import (
	"fmt"
	"sort"
)

const (
	None  uint8 = 0  // 空牌
	Man1  uint8 = 1  // 一万
	Man2  uint8 = 2  // 二万
	Man3  uint8 = 3  // 三万
	Man4  uint8 = 4  // 四万
	Man5  uint8 = 5  // 五万
	Man6  uint8 = 6  // 六万
	Man7  uint8 = 7  // 七万
	Man8  uint8 = 8  // 八万
	Man9  uint8 = 9  // 九万
	Pin1  uint8 = 11 // 一筒
	Pin2  uint8 = 12 // 二筒
	Pin3  uint8 = 13 // 三筒
	Pin4  uint8 = 14 // 四筒
	Pin5  uint8 = 15 // 五筒
	Pin6  uint8 = 16 // 六筒
	Pin7  uint8 = 17 // 七筒
	Pin8  uint8 = 18 // 八筒
	Pin9  uint8 = 19 // 九筒
	Sou1  uint8 = 21 // 一条
	Sou2  uint8 = 22 // 二条
	Sou3  uint8 = 23 // 三条
	Sou4  uint8 = 24 // 四条
	Sou5  uint8 = 25 // 五条
	Sou6  uint8 = 26 // 六条
	Sou7  uint8 = 27 // 七条
	Sou8  uint8 = 28 // 八条
	Sou9  uint8 = 29 // 九条
	Dong  uint8 = 31 // 东
	Nan   uint8 = 32 // 南
	Xi    uint8 = 33 // 西
	Bei   uint8 = 34 // 北
	Zhong uint8 = 41 // 中
	Fa    uint8 = 42 // 发
	Bai   uint8 = 43 // 白
	Chun  uint8 = 51 // 春
	Xia   uint8 = 52 // 夏
	Qiu   uint8 = 53 // 秋
	Dong2 uint8 = 54 // 冬
	Mei   uint8 = 55 // 梅
	Lan   uint8 = 56 // 兰
	Zhu   uint8 = 57 // 竹
	Ju    uint8 = 58 // 菊
)

// 花色
const (
	Man  = 0 // 万
	Pin  = 1 // 筒
	Sou  = 2 // 条
	Feng = 3 // 风
	Jian = 4 // 箭
	Hua  = 5 // 花
)

func main() {
	var mjz []uint8

	fengNameMap := map[uint8]string{
		1: "东",
		2: "南",
		3: "西",
		4: "北",
	}

	jianNameMap := map[uint8]string{
		1: "中",
		2: "发",
		3: "白",
	}

	huaNameMap := map[uint8]string{
		1: "春",
		2: "夏",
		3: "秋",
		4: "冬",
		5: "梅",
		6: "兰",
		7: "竹",
		8: "菊",
	}

	valueMap := map[uint8]string{
		1: "一",
		2: "二",
		3: "三",
		4: "四",
		5: "五",
		6: "六",
		7: "七",
		8: "八",
		9: "九",
	}

	for i := uint8(1); i <= 9; i++ {
		mjz = append(mjz, i)    // 万
		mjz = append(mjz, 10+i) // 筒
		mjz = append(mjz, 20+i) // 条
	}

	for i := uint8(1); i <= 4; i++ {
		mjz = append(mjz, 30+i) // 风
	}

	for i := uint8(1); i <= 3; i++ {
		mjz = append(mjz, 40+i) // 箭
	}

	for i := uint8(1); i <= 8; i++ {
		mjz = append(mjz, 50+i) // 花
	}
	sort.Slice(mjz, func(i, j int) bool {
		return mjz[i] < mjz[j]
	})
	s := "valueNameMap := map[uint8]string{\n"
	for _, k := range mjz {
		num := ""
		name := ""
		kind := k / 10
		val := k % 10
		switch kind {
		case 0:
			num = valueMap[val]
			name = "万"
		case 1:
			num = valueMap[val]
			name = "筒"
		case 2:
			num = valueMap[val]
			name = "条"
		case 3:
			name = fengNameMap[val]
		case 4:
			name = jianNameMap[val]
		case 5:
			name = huaNameMap[val]
		}
		// fullName := fmt.Sprintf("%d // %s%s", k, num, name)
		// fmt.Println(fullName)
		fullName := fmt.Sprintf("\"%s%s\"", num, name)
		s += fmt.Sprintf("%v", k)
		s += ":"
		s += fullName
		s += ", \n"
	}
	fmt.Println(s + "}")

}
