package manager

type Manager struct {
    Position string
    Salary   float64
    Address  string
}

func (e *Manager) GetPosition() string {
    return e.Position
}

func (e *Manager) SetPosition(position string) {
    e.Position = position
}

func (e *Manager) GetSalary() float64 {
    return e.Salary
}

func (e *Manager) SetSalary(salary float64) {
    e.Salary = salary
}

func (e *Manager) GetAddress() string {
    return e.Address
}

func (e *Manager) SetAddress(address string) {
    e.Address = address
}
