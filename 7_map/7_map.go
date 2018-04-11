package __map

import (
	"sync"
	"fmt"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func main() {
	m :=new(map[string]int)
	m["Amy"] = 10
	ua := UserAges{m}
	ua.Add("Bob",12)
	fmt.Println(ua.Get("Bob"))
}

func (ua *UserAges) Add(name string,age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age,ok := ua.ages[name]; ok {
		return age
	}
	return -1
}