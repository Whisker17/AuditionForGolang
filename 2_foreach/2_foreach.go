package __foreach

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	pase_student()
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{"zhou", 24},
		{"li", 23},
		{"wang", 22},
	}
	//错误示范
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}

	fmt.Println()

	//正确示范
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}
