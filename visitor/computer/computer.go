package computer

const (
	PricePerCore = int(1000)
	PricePerMB   = int(200)
)

type Visitor interface {
	VisitCpu(c Component)
	VisitMem(m Component)
}

type Computer struct {
	Cpu
	Memory
}

type Component interface {
	GetPrice() float64
}

func (c *Computer) Assemble(core, mem int) {
	c.Cpu = Cpu{core: core}
	c.Memory = Memory{size: mem}
}

func (c Computer) Accept(v Visitor) {
	v.VisitCpu(c.Cpu)
	v.VisitMem(c.Memory)
}

/*func (c Computer) GetPrice() float64 {
	return c.Cpu.GetPrice() + c.Memory.GetPrice()
}

func (c Computer) GetInfo() string {
	return c.Memory.GetInfo() + c.Cpu.GetInfo()
}
*/
type Cpu struct {
	core int
	name string
}

func (c Cpu) GetPrice() float64 {
	return float64(c.core * PricePerCore)
}

/*func (c Cpu) GetInfo() string {
	return "core: " + strconv.Itoa(c.core)
}*/

type Memory struct {
	size int
	name string
}

func (m Memory) GetPrice() float64 {
	return float64(m.size * PricePerMB)
}

/*func (m Memory) GetInfo() string {
	return "mem: " + strconv.Itoa(m.size) + "MB"
}
*/
