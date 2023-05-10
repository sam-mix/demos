package dasanyuan_test

import "testing"

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

// Dragon 判断是否符合龙牌型
func Dragon(pais []uint8) bool {
	if len(pais) != 3 {
		return false
	}
	return pais[0] == Zhong || pais[0] == Fa || pais[0] == Bai
}

// BigThreeDragon 判断是否符合大三元牌型
func BigThreeDragon(pais []uint8) bool {
	countDragon := 0
	for _, pai := range pais {
		if Dragon([]uint8{pai}) {
			countDragon++
		}
	}
	return countDragon == 3
}

// TestBigThreeDragon 单元测试
func TestBigThreeDragon(t *testing.T) {
	testCases := []struct {
		desc         string
		pais         []uint8
		expectedBool bool
	}{
		{
			desc:         "测试1",
			pais:         []uint8{Zhong, Zhong, Zhong, Man1, Man2, Man3, Sou7, Sou8, Sou9},
			expectedBool: false,
		},
		{
			desc:         "测试2",
			pais:         []uint8{Zhong, Zhong, Bai, Man1, Man2, Man3, Sou7, Sou8, Sou9},
			expectedBool: false,
		},
		{
			desc:         "测试3",
			pais:         []uint8{Zhong, Zhong, Fa, Man1, Man2, Man3, Sou7, Sou8, Sou9},
			expectedBool: false,
		},
		{
			desc:         "测试4",
			pais:         []uint8{Zhong, Bai, Fa, Man1, Man2, Man3, Sou7, Sou8, Sou9},
			expectedBool: false,
		},
	}

	for _, tC := range testCases {
		if result := BigThreeDragon(tC.pais); result != tC.expectedBool {
			t.Errorf("错误信息[%s]: BigThreeDragon(%v)={%v}, 期望值={%v}", tC.desc, tC.pais, result, tC.expectedBool)
		}
	}
}
