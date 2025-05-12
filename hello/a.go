package hello

import "fmt"

func GetSay(name string) {
	fmt.Printf("您的姓名是%s,很高兴认识你\n", name)
}
func Add(x int, y int) {
	sum := x + y
	fmt.Printf("和为%d\n", sum)
}
