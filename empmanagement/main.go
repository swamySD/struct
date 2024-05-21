// package main

// import "fmt"

// // Employee struct represents an employee
// type Employee struct {
// 	ID     int
// 	Name   string
// 	Age    int
// 	Salary float64
// }

// // EmployeeManager interface defines common operations for employee management
// type EmployeeManager interface {
// 	Add(Employee)
// 	Update(int, Employee)
// 	Delete(int)
// 	Display()
// }

// // EmployeeService implements EmployeeManager interface using map
// type EmployeeService struct {
// 	employees map[int]Employee
// }

// // NewEmployeeService creates a new instance of EmployeeService
// func NewEmployeeService() *EmployeeService {
// 	return &EmployeeService{
// 		employees: make(map[int]Employee),
// 	}
// }

// // Add adds a new employee
// func (es *EmployeeService) Add(emp Employee) {
// 	es.employees[emp.ID] = emp
// }

// // Update updates an existing employee
// func (es *EmployeeService) Update(id int, emp Employee) {
// 	if _, ok := es.employees[id]; ok {
// 		es.employees[id] = emp
// 	} else {
// 		fmt.Println("Employee not found")
// 	}
// }

// // Delete deletes an employee
// func (es *EmployeeService) Delete(id int) {
// 	if _, ok := es.employees[id]; ok {
// 		delete(es.employees, id)
// 	} else {
// 		fmt.Println("Employee not found")
// 	}
// }

// // Display displays all employees
// func (es *EmployeeService) Display() {
// 	for id, emp := range es.employees {
// 		fmt.Printf("ID: %d, Name: %s, Age: %d, Salary: %.2f\n", id, emp.Name, emp.Age, emp.Salary)
// 	}
// }

// func main() {
// 	// Create new instance of EmployeeService
// 	empService := NewEmployeeService()

// 	// Adding employees
// 	empService.Add(Employee{ID: 1, Name: "John Doe", Age: 30, Salary: 50000})
// 	empService.Add(Employee{ID: 2, Name: "Jane Smith", Age: 25, Salary: 60000})

// 	// Displaying employees
// 	fmt.Println("Employees:")
// 	fmt.Println(empService)
// 	empService.Display()

// 	// Updating an employee
// 	empService.Update(1, Employee{ID: 1, Name: "John Doe", Age: 35, Salary: 55000})

// 	// Displaying employees after update
// 	fmt.Println("\nEmployees after update:")
// 	empService.Display()

// 	// Deleting an employee
// 	empService.Delete(2)

// 	// Displaying employees after delete
// 	fmt.Println("\nEmployees after delete:")
// 	empService.Display()
// }

package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Channels in golang Deadlock")
	myChan:=make(chan int,2)
	wg:=&sync.WaitGroup{}
	// myChan <- 5
	// fmt.Println(<-myChan)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup){
		val,isChannelOpen:=<-myChan
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		fmt.Println(<-myChan)
		fmt.Println(<-myChan)
		 wg.Done()
		
		
	}(myChan,wg)
	go func(ch chan int, wg *sync.WaitGroup){
		 myChan <- 0
		 
		// myChan <- 5
		// myChan <- 6
		
		close(myChan)
		 wg.Done()
		}(myChan,wg)
		wg.Wait()
}
