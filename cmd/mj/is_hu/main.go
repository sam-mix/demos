package main

import (
	"fmt"
	"sort"
)

func main() {

}

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

func sortUint8Slice(slice []uint8) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func IsHu(pais []uint8) bool {
	// 1. 首先将牌按照花色进行分类
	man, pin, sou, feng, hua := classifyPais(pais)
	fmt.Printf("man = [%v], pin = [%v], sou = [%v], feng = [%v], hua = [%v],", man, pin, sou, feng, hua)
	// 2. 将每种花色的牌按照牌值进行排序
	sortUint8Slice(man)
	sortUint8Slice(pin)
	sortUint8Slice(sou)

	// 3. 判断是否存在七对子
	if len(pais) == 14 && isQiDui(man, pin, sou, feng, hua) {
		return true
	}

	// 4. 判断是否存在基本牌型
	return isHu(man, pin, sou, feng, hua)
}

// 将牌按照花色进行分类
func classifyPais(pais []uint8) ([]uint8, []uint8, []uint8, []uint8, []uint8) {
	man := make([]uint8, 0)
	pin := make([]uint8, 0)
	sou := make([]uint8, 0)
	feng := make([]uint8, 0)
	hua := make([]uint8, 0)
	for _, p := range pais {
		switch p {
		case Dong, Nan, Xi, Bei, Zhong, Fa, Bai:
			feng = append(feng, p)
		case Chun, Xia, Qiu, Dong2, Mei, Lan, Zhu, Ju:
			hua = append(hua, p)
		case Man1, Man2, Man3, Man4, Man5, Man6, Man7, Man8, Man9:
			man = append(man, p)
		case Pin1, Pin2, Pin3, Pin4, Pin5, Pin6, Pin7, Pin8, Pin9:
			pin = append(pin, p)
		case Sou1, Sou2, Sou3, Sou4, Sou5, Sou6, Sou7, Sou8, Sou9:
			sou = append(sou, p)
		}
	}
	return man, pin, sou, feng, hua
}

// 判断是否存在七对子
func isQiDui(man, pin, sou, feng, hua []uint8) bool {
	if len(man)+len(pin)+len(sou)+len(feng)+len(hua) != 14 {
		return false
	}

	count := 0
	count += getDuiCount(man)
	count += getDuiCount(pin)
	count += getDuiCount(sou)
	count += getDuiCount(feng)
	count += getDuiCount(hua)

	return count == 7
}

// 计算对子的数量
func getDuiCount(pais []uint8) int {
	count := 0
	for i := 0; i < len(pais)-1; i++ {
		if pais[i] == pais[i+1] {
			count++
			i++
		}
	}
	return count
}

func isHu(man, pin, sou, feng, hua []uint8) bool {
	// 检查是否有字牌
	if len(feng) > 0 || len(hua) > 0 {
		return false
	}
	// 检查是否有刻子
	for _, p := range man {
		if getCount(man, p) == 3 {
			return true
		}
	}
	for _, p := range pin {
		if getCount(pin, p) == 3 {
			return true
		}
	}
	for _, p := range sou {
		if getCount(sou, p) == 3 {
			return true
		}
	}
	// 检查是否有对子
	for _, p := range man {
		if getCount(man, p) == 2 {
			return true
		}
	}
	for _, p := range pin {
		if getCount(pin, p) == 2 {
			return true
		}
	}
	for _, p := range sou {
		if getCount(sou, p) == 2 {
			return true
		}
	}
	return false
}

func getCount(s []uint8, p uint8) int {
	count := 0
	for _, v := range s {
		if v == p {
			count++
		}
	}
	return count
}
