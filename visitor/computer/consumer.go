package computer

// Customer是一个visitor
// 实现了VisitCpu和VisitMem
type Customer struct {
	CpuRate     float64
	MemRate     float64
	consumption float64
}

func (s *Customer) VisitCpu(c Component) {
	s.consumption += s.CpuRate * c.GetPrice()
}

func (s *Customer) VisitMem(c Component) {
	s.consumption += s.MemRate * c.GetPrice()
}

func (s *Customer) GetConsumption(c Computer) float64 {
	return s.consumption
}
