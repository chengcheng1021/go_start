package main
import (
	"fmt"
)

func main() {
	//打印九九乘法表
	for i:=1; i<=9; i++ {
		for j:=1; j<=i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, i*j)
		}
		fmt.Println()
	}
}