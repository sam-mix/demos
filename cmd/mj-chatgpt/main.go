package main

import (
	"math/rand"
)

// 牌型常量
const (
	ManzuStart = 0  // 万
	PinzuStart = 9  // 饼
	SouzuStart = 18 // 条
	JiStart    = 27 // 字牌
	TotalTiles = 34 // 牌种类数
)

// 牌型
type TileType int

const (
	Manzu TileType = iota // 万
	Pinzu                 // 饼
	Souzu                 // 条
	Ji                    // 字牌
)

// 牌
type Tile int

const (
	Man1  Tile                = iota + ManzuStart // 1万
	Man2                                          // 2万
	Man3                                          // 3万
	Man4                                          // 4万
	Man5                                          // 5万
	Man6                                          // 6万
	Man7                                          // 7万
	Man8                                          // 8万
	Man9                                          // 9万
	Pin1  = iota + PinzuStart                     // 1饼
	Pin2                                          // 2饼
	Pin3                                          // 3饼
	Pin4                                          // 4饼
	Pin5                                          // 5饼
	Pin6                                          // 6饼
	Pin7                                          // 7饼
	Pin8                                          // 8饼
	Pin9                                          // 9饼
	Sou1  = iota + SouzuStart                     // 1条
	Sou2                                          // 2条
	Sou3                                          // 3条
	Sou4                                          // 4条
	Sou5                                          // 5条
	Sou6                                          // 6条
	Sou7                                          // 7条
	Sou8                                          // 8条
	Sou9                                          // 9条
	East  = iota + JiStart                        // 东
	South                                         // 南
	West                                          // 西
	North                                         // 北
	Zhong                                         // 中
	Fa                                            // 发
	Bai                                           // 白
)

// 牌组
type TileSet []Tile

// 牌组类型
type SetType int

const (
	Chi    SetType = iota // 顺子
	Pon                   // 刻子
	Kan                   // 暗杠
	Minkan                // 明杠
)

// 牌组信息
type SetInfo struct {
	Tiles TileSet // 牌组
	Type  SetType // 牌组类型
}

// 手牌
type Hand struct {
	Tiles TileSet   // 手牌
	Sets  []SetInfo // 牌组
}

// 初始化牌
func InitTiles() []Tile {
	var tiles []Tile
	for i := Man1; i <= Bai; i++ {
		tiles = append(tiles, i)
	}
	return tiles
}

// 洗牌
func ShuffleTiles(tiles []Tile) {
	rand.Shuffle(len(tiles), func(i, j int) { tiles[i], tiles[j] = tiles[j], tiles[i] })
}

// 摸牌
func DrawTile(tiles []Tile) (Tile, []Tile) {
	tile := tiles[0]
	tiles = tiles[1:]
	return tile, tiles
}
