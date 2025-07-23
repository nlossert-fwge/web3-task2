package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 员工ID: %s\n", e.Name, e.Age, e.EmployeeID)
}

func main() {
	e := Employee{
		Person:     Person{Name: "张三", Age: 28},
		EmployeeID: "E1001",
	}
	e.PrintInfo()
}
