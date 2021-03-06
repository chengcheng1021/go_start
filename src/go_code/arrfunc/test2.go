package main
import (
	"fmt"
)

func fbn(n int) ([]uint64) {
	//声明一个切片，切片大小 n
	fbnSlice := make([]uint64, n)
	//第一个第二个为1
	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i:=2; i<n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}

	return fbnSlice
}

func main() {
	fbnSlice := fbn(10)
	fmt.Println("fbnSlice=", fbnSlice)
}