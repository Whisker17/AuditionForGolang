package __gooop

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	for i := 0; i < 10; i++ {
		select {
		case int_chan <- 1:
			fmt.Println(<-int_chan)
		case string_chan <- "hello":
			fmt.Println(<-string_chan)
		}
	}
}
