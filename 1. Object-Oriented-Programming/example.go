// Complete OOP Demo - Go
// Flow: Struct -> Access Control -> Constructor -> Embedding -> Composition -> Polymorphism -> Interface -> Encapsulation

package main

import "fmt"

// ============================================================================
// 1. STRUCT (like class) with ACCESS CONTROL & CONSTRUCTOR
// ============================================================================

// Static-like variable (package level)
var totalEmployees int

type Employee struct {
	name       string  // unexported (private) - lowercase
	salary     float64 // unexported (private)
	Department string  // exported (public) - uppercase
}

// NewEmployee is a constructor-like function that creates a new Employee struct.
func NewEmployee(name string, salary float64, department string) *Employee {
	totalEmployees++ // like static increment
	fmt.Printf("Employee created: %s\n", name)
	return &Employee{name: name, salary: salary, Department: department}
}

// Getter for private field (encapsulation)
func (e *Employee) GetName() string { return e.name }
func GetTotalEmployees() int        { return totalEmployees }

// Method for Employee
func (e *Employee) Work() {
	fmt.Printf("%s is working\n", e.name)
}

// ============================================================================
// 2. INHERITANCE via EMBEDDING - Single, Multilevel, Hierarchical
// ============================================================================

// Single inheritance: Manager embeds Employee
type Manager struct {
	*Employee // embedded pointer (inheritance)
	teamName  string
}

func NewManager(name string, salary float64, department, teamName string) *Manager {
	return &Manager{
		Employee: NewEmployee(name, salary, department),
		teamName: teamName,
	}
}

// Method overriding (Runtime polymorphism)
func (m *Manager) Work() {
	fmt.Printf("%s is managing team: %s\n", m.GetName(), m.teamName)
}

// Hierarchical inheritance: Developer also embeds Employee
type Developer struct {
	*Employee
	programmingLanguage string
}

func NewDeveloper(name string, salary float64, department, language string) *Developer {
	return &Developer{
		Employee:            NewEmployee(name, salary, department),
		programmingLanguage: language,
	}
}

func (d *Developer) Work() {
	fmt.Printf("%s is coding in %s\n", d.GetName(), d.programmingLanguage)
}

// Multilevel inheritance: SeniorDeveloper embeds Developer
type SeniorDeveloper struct {
	*Developer
	yearsExperience int
}

func NewSeniorDeveloper(name string, salary float64, department, language string, years int) *SeniorDeveloper {
	return &SeniorDeveloper{
		Developer:       NewDeveloper(name, salary, department, language),
		yearsExperience: years,
	}
}

func (sd *SeniorDeveloper) Work() {
	fmt.Printf("%s is leading development with %d years exp\n", sd.GetName(), sd.yearsExperience)
}

// ============================================================================
// 3. COMPOSITION - Has-A relationship
// ============================================================================

type Computer struct {
	model string
}

func NewComputer(model string) *Computer {
	return &Computer{model: model}
}

func (c *Computer) Start() {
	fmt.Printf("%s computer started\n", c.model)
}

type Office struct {
	location string
}

func NewOffice(location string) *Office {
	return &Office{location: location}
}

func (o *Office) OpenOffice() {
	fmt.Printf("%s office opened\n", o.location)
}

type WorkStation struct {
	*Employee           // inheritance
	computer  *Computer // HAS-A relationship
	office    *Office   // HAS-A relationship
}

func NewWorkStation(name string, salary float64, department string, computer *Computer, office *Office) *WorkStation {
	return &WorkStation{
		Employee: NewEmployee(name, salary, department),
		computer: computer,
		office:   office,
	}
}

func (ws *WorkStation) Work() {
	ws.office.OpenOffice()
	ws.computer.Start()
	fmt.Printf("%s started working at workstation\n", ws.GetName())
}

// ============================================================================
// 4. DIAMOND PROBLEM SIMULATION & SOLUTION
// ============================================================================

// Interface A
type Workable interface {
	DoWork()
}

// Interface B
type Manageable interface {
	DoWork()
}

// Struct implementing both interfaces (potential diamond problem)
type TeamLead struct {
	*Employee
}

func NewTeamLead(name string, salary float64, department string) *TeamLead {
	return &TeamLead{Employee: NewEmployee(name, salary, department)}
}

// SOLUTION: Single method implementation satisfies both interfaces
func (tl *TeamLead) DoWork() {
	fmt.Printf("TeamLead %s: doing both work and management\n", tl.GetName())
}

// Go doesn't have the classic diamond problem because:
// 1. No multiple struct embedding of same type
// 2. Interface methods are satisfied by single implementation
// 3. If embedding conflicts occur, you must explicitly resolve them

// ============================================================================
// 5. POLYMORPHISM - Method Overloading simulation & Runtime polymorphism
// ============================================================================

type Calculator struct{}

// Go doesn't have method overloading, so we simulate with different names
func (c *Calculator) CalculateInt(a, b int) int           { return a + b }
func (c *Calculator) CalculateFloat(a, b float64) float64 { return a + b }
func (c *Calculator) CalculateThree(a, b, c int) int      { return a + b + c }

// Alternative: using variadic and type assertion
func (c *Calculator) Calculate(values ...interface{}) interface{} {
	if len(values) == 2 {
		switch v1 := values[0].(type) {
		case int:
			if v2, ok := values[1].(int); ok {
				return v1 + v2
			}
		case float64:
			if v2, ok := values[1].(float64); ok {
				return v1 + v2
			}
		}
	}
	return nil
}

// ============================================================================
// 6. ABSTRACTION - Interface (pure abstraction/contract)
// ============================================================================

// Interface defines contract (like abstract class methods)
type Vehicular interface {
	Start()
	DisplayInfo()
}

type Drivable interface {
	Drive()
}

// Base struct (like abstract class)
type Vehicle struct {
	brand string
}

func NewVehicle(brand string) Vehicle {
	return Vehicle{brand: brand}
}

func (v *Vehicle) DisplayInfo() {
	fmt.Printf("Vehicle: %s\n", v.brand)
}

// Car implements interfaces
type Car struct {
	Vehicle // embedded (inheritance)
}

func NewCar(brand string) *Car {
	return &Car{Vehicle: NewVehicle(brand)}
}

func (c *Car) Start() {
	fmt.Printf("%s car started\n", c.brand)
}

func (c *Car) Drive() {
	fmt.Printf("%s car is driving\n", c.brand)
}

// ============================================================================
// 7. ENCAPSULATION - Data hiding with controlled access
// ============================================================================

type BankAccount struct {
	balance       float64 // unexported (private)
	accountNumber string  // unexported (private)
}

func NewBankAccount(accountNumber string, initialBalance float64) *BankAccount {
	balance := initialBalance
	if balance < 0 {
		balance = 0
	}
	return &BankAccount{accountNumber: accountNumber, balance: balance}
}

// Controlled access through methods
func (ba *BankAccount) GetBalance() float64 { return ba.balance }

func (ba *BankAccount) Deposit(amount float64) bool {
	if amount > 0 {
		ba.balance += amount
		return true
	}
	return false
}

func (ba *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && amount <= ba.balance {
		ba.balance -= amount
		return true
	}
	return false
}

// ============================================================================
// 8. INTERFACE for POLYMORPHISM DEMO
// ============================================================================

type Worker interface {
	Work()
}

// ============================================================================
// 9. MAIN FUNCTION - Demonstrating all concepts
// ============================================================================

func main() {
	fmt.Println("=== Complete OOP Demo in Go ===\n")

	// 1. STRUCT, CONSTRUCTOR, PACKAGE-LEVEL VARIABLE
	fmt.Println("1. Structs & Objects:")
	emp := NewEmployee("Alice", 50000, "IT")
	mgr := NewManager("Bob", 80000, "IT", "DevTeam")
	fmt.Printf("Total employees: %d\n", GetTotalEmployees())

	// 2. INHERITANCE & RUNTIME POLYMORPHISM
	fmt.Println("\n2. Embedding (Inheritance) & Runtime Polymorphism:")
	workers := []Worker{emp, mgr, NewDeveloper("Charlie", 60000, "IT", "Java"),
		NewSeniorDeveloper("David", 90000, "IT", "Python", 8)}
	for _, w := range workers {
		w.Work() // Different behavior based on actual type (runtime polymorphism)
	}

	// 3. COMPOSITION
	fmt.Println("\n3. Composition (Has-A relationship):")
	laptop := NewComputer("Dell Laptop")
	mainOffice := NewOffice("New York")
	ws := NewWorkStation("Eve", 55000, "IT", laptop, mainOffice)
	ws.Work()

	// 4. DIAMOND PROBLEM SOLUTION
	fmt.Println("\n4. Diamond Problem Solution:")
	lead := NewTeamLead("Frank", 75000, "IT")
	lead.DoWork()
	fmt.Println("// Go solves diamond problem through interface design")
	fmt.Println("// Single method implementation satisfies multiple interfaces")

	// 5. METHOD OVERLOADING SIMULATION (Compile-time like)
	fmt.Println("\n5. Method Overloading Simulation:")
	calc := &Calculator{}
	fmt.Printf("CalculateInt(5, 3): %d\n", calc.CalculateInt(5, 3))
	fmt.Printf("CalculateFloat(5.5, 3.2): %.1f\n", calc.CalculateFloat(5.5, 3.2))
	fmt.Printf("CalculateThree(1, 2, 3): %d\n", calc.CalculateThree(1, 2, 3))
	fmt.Printf("Calculate(10, 20): %v\n", calc.Calculate(10, 20))

	// 6. ABSTRACTION (Interface)
	fmt.Println("\n6. Abstraction (Interface contract):")
	var vehicle Vehicular = NewCar("Toyota") // Interface reference
	vehicle.DisplayInfo()
	vehicle.Start()
	if drivable, ok := vehicle.(Drivable); ok {
		drivable.Drive() // Type assertion for interface
	}

	// 7. ENCAPSULATION
	fmt.Println("\n7. Encapsulation (Data hiding & controlled access):")
	account := NewBankAccount("ACC001", 1000)
	fmt.Printf("Initial balance: %.0f\n", account.GetBalance())
	account.Deposit(500)
	account.Withdraw(200)
	fmt.Printf("Final balance: %.0f\n", account.GetBalance())
	// account.balance = 9999 // Error: cannot access unexported field

	// 8. TYPE ASSERTION & INTERFACE POLYMORPHISM
	fmt.Println("\n8. Type Assertion & Interface Polymorphism:")
	var worker Worker = mgr
	if manager, ok := worker.(*Manager); ok {
		fmt.Printf("Type assertion success: %s is a manager\n", manager.GetName())
	}

	fmt.Println("\n=== All OOP concepts demonstrated ===")
}
