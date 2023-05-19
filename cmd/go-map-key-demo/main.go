package main

import "fmt"

// type kkey struct {
// 	x int
// 	y [3]int
// }

type key struct {
	X int
	Y []int
	N string
}

// 根据key结构体随机生成100个测试用数据

// func main() {
// 	// keys := make([]key, 100)

// 	// for i := 0; i < 100; i++ {
// 	// 	n := rand.Intn(10) + 1 // 随机生成Y数组的长度，范围为[1, 10]
// 	// 	y := make([]int, n)    // 创建一个长度为n的切片
// 	// 	for j := 0; j < n; j++ {
// 	// 		y[j] = rand.Intn(100) - 50 // 随机生成Y数组元素的值，范围为[-50, 49]
// 	// 	}
// 	// 	keys[i] = key{
// 	// 		X: rand.Intn(100) - 50, // 随机生成X的值，范围为[-50, 49]
// 	// 		Y: y,
// 	// 		N: fmt.Sprintf("key%d", i), // 生成名称，格式为"keyX"，其中X为索引
// 	// 	}
// 	// }

// 	// // 打印生成的100个key
// 	// for _, k := range keys {
// 	// 	fmt.Printf("%#v\n", k)
// 	// }

// 	keys := []key{
// 		{X: -45, Y: []int{7, -45, -29, 49, -23}, N: "key0"},
// 		{X: -13, Y: []int{-3, 11}, N: "key1"},
// 		{X: 19, Y: []int{44}, N: "key2"},
// 		{X: -12, Y: []int{-34, 31, 27}, N: "key3"},
// 		{X: 20, Y: []int{28, -29, -33, 7, 35, 10, 12, 26}, N: "key4"},
// 		{X: 2, Y: []int{-49, 36, 33, -22, -48}, N: "key5"},
// 		{X: 1, Y: []int{40, -9, -21, -42, 36, -23, 44, 9, -21, -27}, N: "key6"},
// 		{X: -19, Y: []int{20, -20, 49, -15, -33, -44}, N: "key7"},
// 		{X: 28, Y: []int{-14, 29, 47, 48}, N: "key8"},
// 		{X: 5, Y: []int{5}, N: "key9"},
// 		{X: 12, Y: []int{-3, 22, -2, 44}, N: "key10"},
// 		{X: -1, Y: []int{-30, -23, -13, -30}, N: "key11"},
// 		{X: 20, Y: []int{41, 37, -2, 24}, N: "key12"},
// 		{X: -9, Y: []int{16, 15, 18, 47, 24, 15, -43, -9, 23}, N: "key13"},
// 		{X: -12, Y: []int{-1, 15}, N: "key14"},
// 		{X: -10, Y: []int{-19, -9, -24, 9, 14, 5, 44, 48}, N: "key15"},
// 		{X: 5, Y: []int{-23, 7, 30}, N: "key16"},
// 		{X: 0, Y: []int{-5, -29, 18, -26, 7, -14, 9, -25, -33, 2}, N: "key17"},
// 		{X: -25, Y: []int{-6, -22, -32, -46, 33, 42}, N: "key18"},
// 		{X: 29, Y: []int{-32, -13, -36}, N: "key19"},
// 		{X: -33, Y: []int{11, 13, -6, 40, 4, -15}, N: "key20"},
// 		{X: 18, Y: []int{-37, 23, -44, -7}, N: "key21"},
// 		{X: 24, Y: []int{22, 49}, N: "key22"},
// 		{X: 28, Y: []int{-26, -4, -24, -16, -3, 45, -16, -47}, N: "key23"},
// 		{X: -50, Y: []int{10, 21, -11, -33, 47}, N: "key24"},
// 		{X: -8, Y: []int{17, 34, -10, 37, -42, -11, 45, 47, -46}, N: "key25"},
// 		{X: -2, Y: []int{-4, -17, 0, 10, -14, -2, -14, 22, 30}, N: "key26"},
// 		{X: 34, Y: []int{-49, -18, 28, -2, 26}, N: "key27"},
// 		{X: -48, Y: []int{27}, N: "key28"},
// 		{X: -44, Y: []int{-32, 10, -8, 25, 30, 15, 33}, N: "key29"},
// 		{X: 20, Y: []int{-32, -45, -9, 30, 32, 28, -24, -12, -23}, N: "key30"},
// 		{X: 27, Y: []int{-46, -1, -8}, N: "key31"},
// 		{X: -35, Y: []int{35, -10}, N: "key32"},
// 		{X: -9, Y: []int{47, -46, -34, 48, 9}, N: "key33"},
// 		{X: 33, Y: []int{6, 48, 15, 12, 27, -36}, N: "key34"},
// 		{X: 48, Y: []int{42, -16, -18, -7}, N: "key35"},
// 		{X: 4, Y: []int{-9, 41, -48, 23, -45}, N: "key36"},
// 		{X: -22, Y: []int{-26, -48, -2, 49, -9}, N: "key37"},
// 		{X: -39, Y: []int{37, -46, -9, 36, 29, 40, -7}, N: "key38"},
// 		{X: -46, Y: []int{-41, 6, -19, -39, 9, 29, -5, -34, -40}, N: "key39"},
// 		{X: 49, Y: []int{10, -14, 7, -49, -26, -23}, N: "key40"},
// 		{X: -23, Y: []int{14, 45}, N: "key41"},
// 		{X: -36, Y: []int{9, 17, -28, 25}, N: "key42"},
// 		{X: -38, Y: []int{-19, -44, -15, 15, 40, 27, 19}, N: "key43"},
// 		{X: 30, Y: []int{34, -15, 20, -9, -48, 5, 8, 36}, N: "key44"},
// 		{X: -37, Y: []int{32, 15, -42, -20}, N: "key45"},
// 		{X: -13, Y: []int{-19, -49, 17, 44, -39, 23, 15, 7, -6}, N: "key46"},
// 		{X: -25, Y: []int{-16, 0, -18, 14, 17, -44, 3, 18}, N: "key47"},
// 		{X: -3, Y: []int{-27, 12, -50, 43, -48, 46, -14, -19, 6, 6}, N: "key48"},
// 		{X: -8, Y: []int{-27, 4, 42}, N: "key49"},
// 		{X: 34, Y: []int{7, 37, 43}, N: "key50"},
// 		{X: 9, Y: []int{26, -45}, N: "key51"},
// 		{X: 9, Y: []int{-45, 35, -29, 10, -20, -13, 18, -48}, N: "key52"},
// 		{X: 29, Y: []int{-9, 43, 46, 6}, N: "key53"},
// 		{X: 15, Y: []int{-29, -36, -2, 2, 19, 43, 5, 18}, N: "key54"},
// 		{X: 24, Y: []int{-5, 6, -43, -18, 3, -10}, N: "key55"},
// 		{X: -38, Y: []int{26, -41, 17, 38, 43, -11, -20}, N: "key56"},
// 		{X: -4, Y: []int{40}, N: "key57"},
// 		{X: -36, Y: []int{42, 22, -11, -34, 48, -36, -6, -15}, N: "key58"},
// 		{X: 6, Y: []int{7, -19, -44, 27}, N: "key59"},
// 		{X: -17, Y: []int{-24, 10, -16, -33, 29}, N: "key60"},
// 		{X: -46, Y: []int{-16, -38, -1, 14}, N: "key61"},
// 		{X: -38, Y: []int{2}, N: "key62"},
// 		{X: 30, Y: []int{32, -14, 12, 6, 25, -13}, N: "key63"},
// 		{X: -14, Y: []int{-24, -46, -18, 46, 47, -46}, N: "key64"},
// 		{X: -16, Y: []int{47, 9, -21, 18, -25, 7, 8, -45, 0, -49}, N: "key65"},
// 		{X: 37, Y: []int{35, 10}, N: "key66"},
// 		{X: 25, Y: []int{12, 41, 22, -20}, N: "key67"},
// 		{X: 34, Y: []int{-1, 44, 8, 15, 49, -1, 14, 23}, N: "key68"},
// 		{X: 15, Y: []int{41, -12, -46, 14, 43, 18, -8}, N: "key69"},
// 		{X: 6, Y: []int{-28, 26, 22, 19, 30, -19, -4}, N: "key70"},
// 		{X: -23, Y: []int{-20, -24, -45, 19, 10, -4}, N: "key71"},
// 		{X: 32, Y: []int{-12, -39, -3, -46, 2, -15, -6, 14, 30}, N: "key72"},
// 		{X: 15, Y: []int{-26, -13, -18, 25, -17, 33}, N: "key73"},
// 		{X: -45, Y: []int{11}, N: "key74"},
// 		{X: -20, Y: []int{-12, -11, -2, 45, -31, -41, 4}, N: "key75"},
// 		{X: 28, Y: []int{8, -42, 11, -22, -39, -46, 16, 19, -13, -41}, N: "key76"},
// 		{X: 40, Y: []int{-48, 8}, N: "key77"},
// 		{X: 33, Y: []int{4, -13, 22, 37, -3, 11, -27, 22}, N: "key78"},
// 		{X: -30, Y: []int{-41, 16, 43, 49, -16, 4, 46, 38, -19, 29}, N: "key79"},
// 		{X: 31, Y: []int{40}, N: "key80"},
// 		{X: -6, Y: []int{-20, 40, 1, 5, -26, 49, -17, 23}, N: "key81"},
// 		{X: 36, Y: []int{-4, -14, 29, 22, 33, -24}, N: "key82"},
// 		{X: -5, Y: []int{25, -16}, N: "key83"},
// 		{X: 1, Y: []int{49, 4, -14, 20, 0, 38, -41}, N: "key84"},
// 		{X: -28, Y: []int{29, -49, -50, -1, 9, 32, 44, -4}, N: "key85"},
// 		{X: -24, Y: []int{-21, -28, 43, 48, -47, -50, -32, -42, 31, 39}, N: "key86"},
// 		{X: 16, Y: []int{10, -13, 6, -30}, N: "key87"},
// 		{X: -27, Y: []int{-43, 32, -22, 3, -20, 15}, N: "key88"},
// 		{X: -34, Y: []int{5, -24}, N: "key89"},
// 		{X: 35, Y: []int{-47, -35, -23, 0, -33, 33, -5, 2, 4}, N: "key90"},
// 		{X: -13, Y: []int{-44, 27, 34, -13, 29, 9, 7, 1, 12}, N: "key91"},
// 		{X: 16, Y: []int{3, -19, -33, -36, 43, 16, -35, 2, 43}, N: "key92"},
// 		{X: 0, Y: []int{12, 18, 46, 12, 22, 5, -44}, N: "key93"},
// 		{X: 23, Y: []int{-6, -1, 37, -18, 9, 5, -18, 24}, N: "key94"},
// 		{X: -31, Y: []int{-1, 27, -21}, N: "key95"},
// 		{X: 43, Y: []int{-39, 8, -13, 43}, N: "key96"},
// 		{X: -15, Y: []int{-2, -47, 37, 1, -17, 13}, N: "key97"},
// 		{X: -47, Y: []int{-10, -38}, N: "key98"},
// 		{X: -38, Y: []int{24}, N: "key99"},
// 	}
// 	fmt.Println(keys)
// }

// func main() {
// 	m := map[kkey]int{
// 		{
// 			x: 1,
// 			y: [3]int{1, 3, 4},
// 		}: 1,
// 	}
// 	fmt.Println(m)
// }

func main() {
	key := find(1, 2)
	fmt.Println("find(1,2): ", key.N)
}

func find(a int, b int) *key {
	if a == 1 && b == 2 {
		return &key{
			X: 1,
			Y: []int{2},
			N: "one",
		}
	}
	return nil
}