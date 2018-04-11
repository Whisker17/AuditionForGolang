package __defer

import "fmt"

func main() {
	a := 1
	b := 2
	defer cal("1", a, cal("10", a, b))
	a = 0
	defer cal("2", a, cal("20", a, b))
	b = 1
}

func cal(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
