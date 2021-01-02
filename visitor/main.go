package main

import (
	"fmt"
	"pioneerlfn/visitor/computer"
)

func main() {
	// 攒机
	c := computer.Computer{}
	c.Assemble(1, 2)

	// 创建visitor 1
	// visitor实现了VisitCpu(Component), VisitMem(Component)
	s := computer.Customer{CpuRate: 0.75, MemRate: 0.8}
	c.Accept(&s)
	fmt.Println("学生攒机,  花费:", s.GetConsumption(c))

	// visitor 2
	e := computer.Customer{CpuRate: 0.8, MemRate: 0.95}
	c.Accept(&e)
	fmt.Println("上班族攒机,花费:", e.GetConsumption(c))
}
