/**
{
	"created": "10 Feb 2019, Sun",
	"aim": "Working with inetfaces",
	"codedBy": "Rishikesh Agrawani"
}
*/

package main

import "fmt"

// import "reflect"

// Defining an interface
type Person interface {
	getFullname() string
	getAge() int8
	getInterests() []string
	showDetails()
}

type CommonDetail struct {
	fullname  string
	age       int8
	interests []string
}

// Defining a structure named Student
type Student struct {
	detail  CommonDetail
	college string
}

// Defining a structure named Employee
type Employee struct {
	detail CommonDetail
	salary float64
}

func (student Student) getFullname() string {
	return student.detail.fullname
}

func (employee Employee) getFullname() string {
	return employee.detail.fullname
}

func (student Student) getAge() int8 {
	return student.detail.age
}

func (employee Employee) getAge() int8 {
	return employee.detail.age
}

func (student Student) getInterests() []string {
	return student.detail.interests
}

func (employee Employee) getInterests() []string {
	return employee.detail.interests
}

func showCommonDetails(object interface{}, typ string) {
	// var o interface{}

	if typ == "student" { // do not use type (predefined)
		student, _ := object.(Student)

		fmt.Println("Fullname:", student.detail.fullname)
		fmt.Println("Age:", student.detail.age)
		fmt.Println("Interests:", student.detail.interests)
	} else if typ == "employee" {
		employee, _ := object.(Employee)

		fmt.Println("Fullname:", employee.detail.fullname)
		fmt.Println("Age:", employee.detail.age)
		fmt.Println("Interests:", employee.detail.interests)
	}
}

func (student Student) showDetails() {
	showCommonDetails(student, "student")
	fmt.Println("School: ", student.college, "\n")
}

func (employee Employee) showDetails() {
	showCommonDetails(employee, "employee")
	fmt.Println("Salary: ", employee.salary, "\n")
}

func getDetails(person Person) {
	person.showDetails()
}

func main() {
	// Student
	studentDetail := CommonDetail{
		fullname:  "Malinikesh Agrawani",
		age:       22,
		interests: []string{"Reading Books", "Watching TV", "Cooking"},
	}

	student := Student{
		detail:  studentDetail,
		college: "Govt. Nagarjuna Science college, Raipur",
	}

	fmt.Println(student)
	student.showDetails()

	// Employee
	employeeDetail := CommonDetail{
		fullname: "Raghvendra Thakur",
		age:      26,
		interests: []string{
			"Programming", "Reading Books",
			"Watching animated movies", "Writing hindi/english poems",
		},
	}

	employee := Employee{
		detail: employeeDetail,
		salary: 5000000,
	}

	fmt.Println(employee)
	employee.showDetails()
}
