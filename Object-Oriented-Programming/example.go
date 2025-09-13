// Complete Object-Oriented Programming Examples in Go
// Covering: Structs, Methods, Interfaces, Embedding, Composition, Polymorphism

package main

import (
	"fmt"
	"strings"
)

// ============================================================================
// 1. ENCAPSULATION - Data hiding using package visibility
// ============================================================================

// Shared variable (like static in Java)
var totalAccounts int

type BankAccount struct {
	accountNumber string  // unexported (private)
	balance       float64 // unexported (private)
	OwnerName     string  // exported (public)
}

// Constructor function (Go doesn't have constructors)
func NewBankAccount(accountNumber, ownerName string, initialBalance float64) *BankAccount {
	if initialBalance < 0 {
		initialBalance = 0
	}
	totalAccounts++
	return &BankAccount{
		accountNumber: accountNumber,
		balance:       initialBalance,
		OwnerName:     ownerName,
	}
}

// Getter methods (controlled access to private fields)
func (b *BankAccount) GetAccountNumber() string { return b.accountNumber }
func (b *BankAccount) GetBalance() float64      { return b.balance }
func GetTotalAccounts() int                     { return totalAccounts }

// Behavior methods with validation
func (b *BankAccount) Deposit(amount float64) bool {
	if amount > 0 {
		b.balance += amount
		return true
	}
	return false
}

func (b *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && amount <= b.balance {
		b.balance -= amount
		return true
	}
	return false
}

func (b *BankAccount) DisplayInfo() {
	fmt.Printf("Account: %s, Owner: %s, Balance: $%.2f\n",
		b.accountNumber, b.OwnerName, b.balance)
}

// ============================================================================
// 2. ABSTRACTION - Interfaces (contracts)
// ============================================================================

// Interface for vehicle operations (abstraction)
type Vehicular interface {
	Start()
	Stop()
	DisplayBasicInfo()
	CalculateFuelEfficiency() float64
}

// Interface for driving behavior
type Drivable interface {
	Accelerate()
	Brake()
	Turn(direction string)
}

// Interface for maintenance
type Maintainable interface {
	PerformMaintenance()
	NeedsMaintenance() bool
}

// ============================================================================
// 3. BASE STRUCT (like abstract class) & INHERITANCE via EMBEDDING
// ============================================================================

// Base Vehicle struct (like abstract class in Java)
type Vehicle struct {
	Brand string
	Model string
	Year  int
}

// Methods for base Vehicle
func (v *Vehicle) DisplayBasicInfo() {
	fmt.Printf("%d %s %s\n", v.Year, v.Brand, v.Model)
}

// Constructor for Vehicle
func NewVehicle(brand, model string, year int) Vehicle {
	return Vehicle{Brand: brand, Model: model, Year: year}
}

// ============================================================================
// 4. INHERITANCE via EMBEDDING - Car "inherits" from Vehicle
// ============================================================================

type Car struct {
	Vehicle  // embedded struct (inheritance)
	doors    int
	fuelType string
	mileage  float64
}

// Constructor for Car
func NewCar(brand, model string, year, doors int, fuelType string) *Car {
	return &Car{
		Vehicle:  NewVehicle(brand, model, year),
		doors:    doors,
		fuelType: fuelType,
		mileage:  0,
	}
}

// Implementing Vehicular interface
func (c *Car) Start() {
	fmt.Println("Car engine started with key ignition")
}

func (c *Car) Stop() {
	fmt.Println("Car engine stopped")
}

func (c *Car) CalculateFuelEfficiency() float64 {
	if c.mileage > 0 {
		return c.mileage / 100 // km per liter
	}
	return 0
}

// Implementing Drivable interface
func (c *Car) Accelerate() {
	fmt.Println("Car is accelerating smoothly")
}

func (c *Car) Brake() {
	fmt.Println("Car brakes applied")
}

func (c *Car) Turn(direction string) {
	fmt.Printf("Car turning %s\n", direction)
}

// Implementing Maintainable interface
func (c *Car) PerformMaintenance() {
	fmt.Println("Performing car maintenance: oil change, tire check")
}

func (c *Car) NeedsMaintenance() bool {
	return c.mileage > 10000 // needs maintenance after 10000 km
}

func (c *Car) AddMileage(miles float64) {
	c.mileage += miles
}

// ============================================================================
// 5. HIERARCHICAL INHERITANCE - Another struct inheriting from Vehicle
// ============================================================================

type Motorcycle struct {
	Vehicle  // embedded struct (inheritance)
	engineCC int
}

func NewMotorcycle(brand, model string, year, engineCC int) *Motorcycle {
	return &Motorcycle{
		Vehicle:  NewVehicle(brand, model, year),
		engineCC: engineCC,
	}
}

// Implementing Vehicular interface
func (m *Motorcycle) Start() {
	fmt.Println("Motorcycle started with kick/electric start")
}

func (m *Motorcycle) Stop() {
	fmt.Println("Motorcycle engine stopped")
}

func (m *Motorcycle) CalculateFuelEfficiency() float64 {
	if m.engineCC < 200 {
		return 40 // km/l for smaller engines
	}
	return 25 // km/l for larger engines
}

// Implementing Drivable interface
func (m *Motorcycle) Accelerate() {
	fmt.Println("Motorcycle accelerating quickly")
}

func (m *Motorcycle) Brake() {
	fmt.Println("Motorcycle brakes applied")
}

func (m *Motorcycle) Turn(direction string) {
	fmt.Printf("Motorcycle leaning %s to turn\n", direction)
}

// ============================================================================
// 6. MULTILEVEL INHERITANCE - SportsCar "inherits" from Car
// ============================================================================

type SportsCar struct {
	*Car         // embedded pointer (inheritance from Car)
	horsepower   int
	turbocharged bool
}

func NewSportsCar(brand, model string, year, doors, horsepower int, turbocharged bool) *SportsCar {
	return &SportsCar{
		Car:          NewCar(brand, model, year, doors, "Petrol"),
		horsepower:   horsepower,
		turbocharged: turbocharged,
	}
}

// Method overriding (polymorphism)
func (sc *SportsCar) Start() {
	fmt.Println("Sports car engine roars to life!")
}

func (sc *SportsCar) Accelerate() {
	fmt.Printf("Sports car accelerating with %d HP!\n", sc.horsepower)
}

// Override fuel efficiency calculation
func (sc *SportsCar) CalculateFuelEfficiency() float64 {
	if sc.turbocharged {
		return 8 // less efficient
	}
	return 12
}

// Additional method specific to sports car
func (sc *SportsCar) ActivateSportMode() {
	fmt.Println("Sport mode activated! Maximum performance!")
}

// ============================================================================
// 7. COMPOSITION - Has-A Relationship
// ============================================================================

type Engine struct {
	Type       string
	Horsepower int
	isRunning  bool
}

func NewEngine(engineType string, horsepower int) *Engine {
	return &Engine{
		Type:       engineType,
		Horsepower: horsepower,
		isRunning:  false,
	}
}

func (e *Engine) Start() {
	e.isRunning = true
	fmt.Printf("%s engine started (%d HP)\n", e.Type, e.Horsepower)
}

func (e *Engine) Stop() {
	e.isRunning = false
	fmt.Printf("%s engine stopped\n", e.Type)
}

func (e *Engine) IsRunning() bool { return e.isRunning }

type GPS struct {
	currentLocation string
}

func NewGPS() *GPS {
	return &GPS{currentLocation: "Unknown"}
}

func (g *GPS) UpdateLocation(location string) {
	g.currentLocation = location
	fmt.Printf("GPS updated: Current location - %s\n", location)
}

func (g *GPS) GetCurrentLocation() string { return g.currentLocation }

func (g *GPS) Navigate(destination string) {
	fmt.Printf("GPS navigating from %s to %s\n", g.currentLocation, destination)
}

// Advanced car with composition
type AdvancedCar struct {
	Vehicle          // embedded (inheritance)
	engine   *Engine // HAS-A relationship
	gps      *GPS    // HAS-A relationship
	features []string
}

func NewAdvancedCar(brand, model string, year int, engine *Engine) *AdvancedCar {
	return &AdvancedCar{
		Vehicle:  NewVehicle(brand, model, year),
		engine:   engine,
		gps:      NewGPS(),
		features: []string{"Air Conditioning", "Power Steering", "ABS"},
	}
}

// Implementing Vehicular interface
func (ac *AdvancedCar) Start() {
	ac.engine.Start()
	ac.gps.UpdateLocation("Starting point")
	fmt.Println("Advanced car is ready to drive")
}

func (ac *AdvancedCar) Stop() {
	ac.engine.Stop()
	fmt.Println("Advanced car stopped")
}

func (ac *AdvancedCar) CalculateFuelEfficiency() float64 {
	return 15 // km/l
}

func (ac *AdvancedCar) Navigate(destination string) {
	if ac.engine.IsRunning() {
		ac.gps.Navigate(destination)
	} else {
		fmt.Println("Please start the car first")
	}
}

func (ac *AdvancedCar) ShowFeatures() {
	fmt.Printf("Car features: %s\n", strings.Join(ac.features, ", "))
}

// ============================================================================
// 8. POLYMORPHISM - Method Overloading simulation (Go doesn't have it)
// ============================================================================

type Calculator struct{}

func NewCalculator() *Calculator {
	return &Calculator{}
}

// Go doesn't have method overloading, so we use different names or variadic functions
func (c *Calculator) AddInts(a, b int) int {
	return a + b
}

func (c *Calculator) AddFloats(a, b float64) float64 {
	return a + b
}

func (cal *Calculator) AddThreeInts(a, b, c int) int {
	return a + b + c
}

func (c *Calculator) AddStrings(a, b string) string {
	return a + " " + b
}

// Using variadic function to simulate overloading
func (c *Calculator) Add(values ...interface{}) interface{} {
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
		case string:
			if v2, ok := values[1].(string); ok {
				return v1 + " " + v2
			}
		}
	}
	return nil
}

// ============================================================================
// 9. VEHICLE MANAGER - Demonstrates polymorphism with interfaces
// ============================================================================

type VehicleManager struct{}

func NewVehicleManager() *VehicleManager {
	return &VehicleManager{}
}

// Method that works with any Vehicular interface (polymorphism)
func (vm *VehicleManager) TestVehicle(vehicle Vehicular) {
	vehicle.DisplayBasicInfo()
	vehicle.Start()
	fmt.Printf("Fuel efficiency: %.1f km/l\n", vehicle.CalculateFuelEfficiency())
	vehicle.Stop()
	fmt.Println()
}

// Method that works with any Drivable interface
func (vm *VehicleManager) TestDriving(vehicle Drivable) {
	vehicle.Accelerate()
	vehicle.Turn("left")
	vehicle.Brake()
	fmt.Println()
}

// ============================================================================
// 10. MAIN FUNCTION - Demonstration of all concepts
// ============================================================================

func main() {
	fmt.Println("=== Object-Oriented Programming Demo in Go ===\n")

	// ========== ENCAPSULATION DEMO ==========
	fmt.Println("1. ENCAPSULATION DEMO:")
	account := NewBankAccount("ACC001", "Alice Johnson", 1000)

	account.DisplayInfo()
	account.Deposit(200)
	account.Withdraw(150)
	account.DisplayInfo()
	fmt.Printf("Total accounts created: %d\n\n", GetTotalAccounts())

	// ========== INHERITANCE & POLYMORPHISM DEMO ==========
	fmt.Println("2. INHERITANCE & POLYMORPHISM DEMO:")
	manager := NewVehicleManager()

	// Creating different types of vehicles
	sedan := NewCar("Toyota", "Camry", 2023, 4, "Petrol")
	ferrari := NewSportsCar("Ferrari", "F8", 2023, 2, 710, true)
	bike := NewMotorcycle("Honda", "CBR600RR", 2023, 600)

	// Polymorphism - same method, different behaviors
	manager.TestVehicle(sedan)   // Calls Car's implementation
	manager.TestVehicle(ferrari) // Calls SportsCar's implementation
	manager.TestVehicle(bike)    // Calls Motorcycle's implementation

	// ========== INTERFACE POLYMORPHISM ==========
	fmt.Println("3. INTERFACE POLYMORPHISM:")
	manager.TestDriving(sedan)
	manager.TestDriving(ferrari)
	manager.TestDriving(bike)

	// ========== CALCULATOR DEMO (Simulated Method Overloading) ==========
	fmt.Println("4. CALCULATOR DEMO (Simulated Method Overloading):")
	calc := NewCalculator()
	fmt.Printf("AddInts(5, 3) = %d\n", calc.AddInts(5, 3))
	fmt.Printf("AddFloats(5.5, 3.2) = %.1f\n", calc.AddFloats(5.5, 3.2))
	fmt.Printf("AddThreeInts(1, 2, 3) = %d\n", calc.AddThreeInts(1, 2, 3))
	fmt.Printf("AddStrings(\"Hello\", \"World\") = %s\n", calc.AddStrings("Hello", "World"))

	// Using variadic approach
	fmt.Printf("Add(10, 20) = %v\n", calc.Add(10, 20))
	fmt.Printf("Add(\"Go\", \"Programming\") = %v\n\n", calc.Add("Go", "Programming"))

	// ========== COMPOSITION DEMO ==========
	fmt.Println("5. COMPOSITION DEMO:")
	v8Engine := NewEngine("V8", 450)
	luxuryCar := NewAdvancedCar("Mercedes", "S-Class", 2023, v8Engine)

	luxuryCar.DisplayBasicInfo()
	luxuryCar.ShowFeatures()
	luxuryCar.Start()
	luxuryCar.Navigate("Downtown Mall")
	luxuryCar.Stop()
	fmt.Println()

	// ========== SPORTS CAR SPECIFIC FEATURES ==========
	fmt.Println("6. SPORTS CAR SPECIFIC FEATURES:")
	ferrari.DisplayBasicInfo()
	ferrari.Start()
	ferrari.Accelerate()
	ferrari.ActivateSportMode()
	ferrari.AddMileage(5000)
	fmt.Printf("Needs maintenance: %t\n\n", ferrari.NeedsMaintenance())

	// ========== RUNTIME POLYMORPHISM DEMO ==========
	fmt.Println("7. RUNTIME POLYMORPHISM DEMO:")
	vehicles := []Vehicular{sedan, ferrari, bike}

	for _, v := range vehicles {
		fmt.Print("Vehicle: ")
		v.DisplayBasicInfo()
		v.Start() // Different implementation called based on actual object type
		fmt.Println()
	}

	// ========== INTERFACE COMPOSITION DEMO ==========
	fmt.Println("8. INTERFACE COMPOSITION DEMO:")

	// Creating slice of different interface types
	drivableVehicles := []Drivable{sedan, ferrari, bike}
	fmt.Println("Testing all drivable vehicles:")
	for i, dv := range drivableVehicles {
		fmt.Printf("Vehicle %d:\n", i+1)
		manager.TestDriving(dv)
	}

	// ========== MAINTENANCE CHECK ==========
	fmt.Println("9. MAINTENANCE CHECK:")
	maintainableVehicles := []Maintainable{sedan, ferrari}

	for i, mv := range maintainableVehicles {
		fmt.Printf("Vehicle %d maintenance status:\n", i+1)
		if mv.NeedsMaintenance() {
			fmt.Println("⚠️  Maintenance required!")
			mv.PerformMaintenance()
		} else {
			fmt.Println("✅ No maintenance needed")
		}
		fmt.Println()
	}

	// ========== TYPE ASSERTION AND INTERFACE CHECKING ==========
	fmt.Println("10. TYPE ASSERTION AND INTERFACE CHECKING:")

	var vehicle Vehicular = ferrari

	// Type assertion
	if sportsCar, ok := vehicle.(*SportsCar); ok {
		fmt.Println("This is a sports car!")
		sportsCar.ActivateSportMode()
	}

	// Interface checking
	if drivable, ok := vehicle.(Drivable); ok {
		fmt.Println("This vehicle is drivable!")
		drivable.Accelerate()
	}

	if maintainable, ok := vehicle.(Maintainable); ok {
		fmt.Println("This vehicle is maintainable!")
		fmt.Printf("Needs maintenance: %t\n", maintainable.NeedsMaintenance())
	}

	fmt.Println("\n=== Demo Complete ===")
}
