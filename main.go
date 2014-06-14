package main

import (
	"fmt"
	"safemap/safemap"
)

func main() {
	newMap := safemap.New()
	newMap.Insert("huanganxin", "xiaohuang")
	fmt.Println(newMap.Len())
	fmt.Println(newMap.Find("huanganxin"))
}
