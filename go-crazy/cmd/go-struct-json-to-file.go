package main

import (
	"encoding/json"
	"os"
)

type Salary struct {
	Basic float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func main() {
	data := Employee{
		FirstName:     "Nicolas",
		LastName:      "Modryzk",
		Email:         "hellonico at gmail.com",
		Age:           24,
		MonthlySalary: []Salary{{Basic: 15000.0}, {Basic: 16000.0}, {Basic: 17000.0}},
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("my_salary.json", file, 0644)
}
