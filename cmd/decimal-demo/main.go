package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// var num float64 = 1
	// for num < 100 {
	// 	rs := num/100 + 1
	// 	decimalRs := decimal.NewFromFloat(rs)
	// 	fmt.Println(rs, decimalRs)
	// 	//fmt.Printf("origin:%v,originType:%T,decimal:%v,decimalType:%T\n",rs,rs,decimalRs,decimalRs)
	// 	num += 1
	// }
	a, _ := decimal.NewFromString("1000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	b, _ := decimal.NewFromString("1000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	c := a.Add(b)
	d := a.Div(b)
	e := a.Mul(b)
	f := a.Sub(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
