package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Tile 麻将牌
type Tile struct {
	Type   string // 花色
	Number int    // 数字
}

// Tiles 麻将牌组
type Tiles []Tile

// NewTiles 创建麻将牌组
func NewTiles() Tiles {
	tiles := Tiles{}
	for i := 0; i < 4; i++ {
		for j := 1; j <= 9; j++ {
			// 万
			tiles = append(tiles, Tile{"wan", j})
			// 索
			tiles = append(tiles, Tile{"suo", j})
			// 条
			tiles = append(tiles, Tile{"tiao", j})
		}
	}
	for i := 1; i <= 4; i++ {
		// 风牌: 东南西北
		tiles = append(tiles, Tile{"zi", i})
	}
	for i := 1; i <= 8; i++ {
		// 花牌: 春夏秋冬梅兰竹菊
		tiles = append(tiles, Tile{"hua", i})
	}
	for i := 1; i <= 3; i++ {
		// 箭牌: 中发白
		tiles = append(tiles, Tile{"zi", i})
	}
	return tiles
}

// RandTiles 不洗牌获取麻将牌组
func RandTiles(tiles Tiles, n int) Tiles {
	var (
		result Tiles
		mu     sync.Mutex
	)
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				mu.Lock()
				if len(tiles) == 0 {
					mu.Unlock()
					break
					// 麻将牌已经全部被取完，退出循环
				}
				idx := rand.Intn(len(tiles))
				tile := tiles[idx]
				tiles[idx] = tiles[len(tiles)-1]
				tiles = tiles[:len(tiles)-1]
				mu.Unlock()
				result = append(result, tile)
			}
		}()
	}
	wg.Wait()
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())
	tiles := NewTiles()
	fmt.Println("麻将牌总数：", len(tiles))
	players := 4
	s := "{\n"
	for i := 0; i < players; i++ {
		playerTiles := RandTiles(tiles, len(tiles)/players)
		fmt.Printf("玩家%d的麻将牌数量：%d\n", i+1, len(playerTiles))

	}
	fmt.Println(s + "}")
}
