package duiduihu_test

import (
	"sort"
	"testing"
)

// Pair 判断是否符合对子牌型
func Pair(pais []uint8) bool {
	if len(pais) != 2 {
		return false
	}
	return pais[0] == pais[1]
}

// Triple 判断是否符合刻子牌型
func Triple(pais []uint8) bool {
	if len(pais) != 3 {
		return false
	}
	return pais[0] == pais[1] && pais[1] == pais[2]
}

const (
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

// // PungPung 判断是否符合碰碰胡牌型
//
//	func PungPung(pais []uint8) bool {
//		if len(pais)%3 != 2 {
//			return false
//		}
//		sortUint8Slice(pais)
//		for i := 0; i < len(pais); i += 3 {
//			if !Triple(pais[i:i+3]) && !Pair(pais[i:i+2]) {
//				return false
//			}
//		}
//		return true
//	}
//
// PungPung 判断是否符合碰碰胡牌型
func PungPung(pais []uint8) bool {
	if len(pais)%3 != 2 {
		return false
	}
	sortUint8Slice(pais)
	for i := 0; i < len(pais); {
		var n int
		switch {
		case i+2 < len(pais) && Triple(pais[i:i+3]):
			n = 3
		case i+1 < len(pais) && Pair(pais[i:i+2]):
			n = 2
		default:
			return false
		}
		i += n
	}
	return true
}

// TestPungPung 单元测试
func TestPungPung(t *testing.T) {
	testCases := []struct {
		desc         string
		pais         []uint8
		expectedBool bool
	}{
		{
			desc:         "测试1",
			pais:         []uint8{Man1, Man1, Man1, Man2, Man2, Man2, Man3, Man3, Man3, Pin7, Pin7, Pin7, Dong, Dong},
			expectedBool: true,
		},
		{
			desc:         "测试2",
			pais:         []uint8{Man1, Man1, Man1, Man2, Man2, Man2, Man3, Man3, Man3, Pin7, Pin7, Pin8, Dong, Dong},
			expectedBool: false,
		},
		{
			desc:         "测试3",
			pais:         []uint8{Man1, Man1, Man1, Man2, Man2, Man2, Man3, Man3, Man9, Pin7, Pin7, Pin7, Dong, Dong},
			expectedBool: false,
		},
		{
			desc:         "测试4",
			pais:         []uint8{Man1, Man1, Man1, Man2, Man2, Man2, Man3, Man3, Pin7, Pin7, Sou5, Sou5, Dong, Dong},
			expectedBool: true,
		},
	}

	for _, tC := range testCases {
		if result := PungPung(tC.pais); result != tC.expectedBool {
			t.Errorf("错误信息[%s]: PungPung(%v)={%v}, 期望值={%v}", tC.desc, tC.pais, result, tC.expectedBool)
		}
	}
}
