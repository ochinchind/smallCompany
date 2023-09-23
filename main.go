package main

import (
    "fmt"
    "first/engineer"
    "first/manager"
)

func main() {
    eng := &engineer.Engineer{}
    eng.SetPosition("Software Engineer")
    eng.SetSalary(99999999)
    eng.SetAddress("1 Almaty")

    fmt.Printf("Engineer: Position=%s, Salary=%.2f, Address=%s\n", eng.GetPosition(), eng.GetSalary(), eng.GetAddress())

    mgr := &manager.Manager{}
    mgr.SetPosition("Project Manager")
    mgr.SetSalary(9000000000000000000000000)
    mgr.SetAddress("Alm")

    fmt.Printf("Manager: Position=%s, Salary=%.2f, Address=%s\n", mgr.GetPosition(), mgr.GetSalary(), mgr.GetAddress())
}
