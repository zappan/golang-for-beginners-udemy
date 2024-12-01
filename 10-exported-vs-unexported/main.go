package main

import (
	"fmt"
	"myapp/staff"
)

// visible from the imported package are stuff starting with Uppercase letter
var employees = []staff.Employee {
	{ FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true },
	{ FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true },
	{ FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true },
	{ FirstName: "Allan", LastName: "Anderson", Salary: 15000, FullTime: false },
	{ FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: false },
}

func main() {
	myStaff := staff.Office {
		AllStaff: employees,
	}


	fmt.Println("All staff:", myStaff.All())

	// staff.OverpaidLimit = 10000  // PUBLIC so it can be changed
	fmt.Println("Overpaid staff:", myStaff.Overpaid())
	fmt.Println("Underpaid staff:", myStaff.Underpaid())

	// LOWERCASE ARE PRIVATE
	// myStaff.notVisible()
	// fmt.Println("Not Visible:", myStaff.notVisible())

	staff.MyFunction()
}
