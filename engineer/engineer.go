package engineer

type Engineer struct {
    Position string
    Salary   float64
    Address  string
}

func (e *Engineer) GetPosition() string {
    return e.Position
}

func (e *Engineer) SetPosition(position string) {
    e.Position = position
}

func (e *Engineer) GetSalary() float64 {
    return e.Salary
}

func (e *Engineer) SetSalary(salary float64) {
    e.Salary = salary
}

func (e *Engineer) GetAddress() string {
    return e.Address
}

func (e *Engineer) SetAddress(address string) {
    e.Address = address
}
