package person

import "fmt"

type employee struct {
	salary float32
}

// 按比例发放工资
func (e *employee) giveRaise(pt float32) float32 {
	salary := e.salary * pt
	return salary
}

func main() {
	xiaoming := &employee{1123467.43}
	salary := xiaoming.giveRaise(0.95)
	fmt.Printf("xiaoming salary is : %.2f\n", salary)

	xiaowang := &employee{9999999.43}
	salary = xiaowang.giveRaise(1.1)
	fmt.Printf("xiaowang salary is : %.2f\n", salary)
}
