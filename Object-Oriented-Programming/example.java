// Complete OOP Demo - Java
// Flow: Class -> Access Modifiers -> Static -> Constructor/Destructor -> Inheritance -> Composition -> Diamond Problem -> Polymorphism -> Abstraction -> Encapsulation

import java.util.*;

// ============================================================================
// 1. CLASS with ACCESS MODIFIERS, STATIC, CONSTRUCTOR/DESTRUCTOR
// ============================================================================
class Employee {
    // Access modifiers: private, protected, public
    private String name;           // private: only within class
    protected double salary;       // protected: same package + subclasses  
    public String department;      // public: accessible anywhere
    
    // Static: belongs to the class itself, not an object.
    // There is only one copy of this variable, shared among all instances.
    private static int totalEmployees = 0;
    
    // Constructor: runs when object is created
    public Employee(String name, double salary, String department) {
        this.name = name;
        this.salary = salary;
        this.department = department;
        totalEmployees++;
        System.out.println("Employee created: " + name);
    }
    
    // Java doesn't have explicit destructors, but finalize() is called by GC (Garbage Collector)
    protected void finalize() {
        System.out.println("Employee " + name + " being destroyed");
    }
    
    // Getter for private field (encapsulation)
    public String getName() { return name; }
    public static int getTotalEmployees() { return totalEmployees; }
    
    public void work() { System.out.println(name + " is working"); }
}

// ============================================================================
// 2. INHERITANCE - Single, Multilevel, Hierarchical
// ============================================================================

// Single inheritance
class Manager extends Employee {
    private String teamName;
    
    public Manager(String name, double salary, String department, String teamName) {
        super(name, salary, department); // Call parent constructor
        this.teamName = teamName;
    }
    
    // Method overriding (Runtime polymorphism)
    @Override
    public void work() { 
        System.out.println(getName() + " is managing team: " + teamName); 
    }
}

// Hierarchical inheritance: Developer also extends Employee
class Developer extends Employee {
    private String programmingLanguage;
    
    public Developer(String name, double salary, String department, String language) {
        super(name, salary, department);
        this.programmingLanguage = language;
    }
    
    @Override
    public void work() { 
        System.out.println(getName() + " is coding in " + programmingLanguage); 
    }
}

// Multilevel inheritance: SeniorDeveloper extends Developer
class SeniorDeveloper extends Developer {
    private int yearsExperience;
    
    public SeniorDeveloper(String name, double salary, String department, String language, int years) {
        super(name, salary, department, language);
        this.yearsExperience = years;
    }
    
    @Override
    public void work() {
        System.out.println(getName() + " is leading development with " + yearsExperience + " years exp");
    }
}

// ============================================================================
// 3. COMPOSITION - Has-A relationship
// ============================================================================
class Computer {
    private String model;
    public Computer(String model) { this.model = model; }
    public void start() { System.out.println(model + " computer started"); }
}

class Office {
    private String location;
    public Office(String location) { this.location = location; }
    public void openOffice() { System.out.println(location + " office opened"); }
}

class WorkStation extends Employee {
    private Computer computer; // HAS-A relationship
    private Office office;     // HAS-A relationship
    
    public WorkStation(String name, double salary, String department, Computer computer, Office office) {
        super(name, salary, department);
        this.computer = computer; // Composition
        this.office = office;     // Composition
    }
    
    @Override
    public void work() {
        office.openOffice();
        computer.start();
        System.out.println(getName() + " started working at workstation");
    }
}

// ============================================================================
// 4. DIAMOND PROBLEM SIMULATION (Java doesn't allow multiple inheritance)
// ============================================================================

// Interface A with default method (Java 8+)
interface Workable {
    default void doWork() { System.out.println("Workable: doing work"); }
}

interface Manageable {
    default void doWork() { System.out.println("Manageable: doing management work"); }
}

// Diamond problem: TeamLead implements both interfaces with same method name
class TeamLead extends Employee implements Workable, Manageable {
    public TeamLead(String name, double salary, String department) {
        super(name, salary, department);
    }
    
    // CONFLICT: Which doWork() to inherit? Solution: Override explicitly
    @Override
    public void doWork() {
        // Solution 1: Choose one
        Workable.super.doWork();
        // Solution 2: Or create own implementation
        System.out.println("TeamLead: doing both work and management");
    }
}

// ============================================================================
// 5. POLYMORPHISM - Method Overloading (Compile-time) & Overriding (Runtime)
// ============================================================================

class Calculator {
    // Method Overloading (Compile-time polymorphism) - same name, different parameters
    public int calculate(int a, int b) { return a + b; }
    public double calculate(double a, double b) { return a + b; }
    public int calculate(int a, int b, int c) { return a + b + c; }
}

// ============================================================================
// 6. ABSTRACTION - Abstract class and Interface
// ============================================================================

// Abstract class (partial abstraction) - cannot be instantiated
abstract class Vehicle {
    protected String brand;
    public Vehicle(String brand) { this.brand = brand; }
    
    // Concrete method
    public void displayInfo() { System.out.println("Vehicle: " + brand); }
    
    // Abstract method - must be implemented by subclasses
    public abstract void start();
}

// Interface (pure abstraction) - contract
interface Drivable {
    void drive(); // implicitly public abstract
}

class Car extends Vehicle implements Drivable {
    public Car(String brand) { super(brand); }
    
    @Override
    public void start() { System.out.println(brand + " car started"); }
    
    @Override
    public void drive() { System.out.println(brand + " car is driving"); }
}

// ============================================================================
// 7. ENCAPSULATION - Data hiding with controlled access
// ============================================================================

class BankAccount {
    // Private data (hidden from outside)
    private double balance;
    private String accountNumber;
    
    public BankAccount(String accountNumber, double initialBalance) {
        this.accountNumber = accountNumber;
        this.balance = initialBalance >= 0 ? initialBalance : 0;
    }
    
    // Controlled access through public methods
    public double getBalance() { return balance; }
    
    public boolean deposit(double amount) {
        if (amount > 0) {
            balance += amount;
            return true;
        }
        return false;
    }
    
    public boolean withdraw(double amount) {
        if (amount > 0 && amount <= balance) {
            balance -= amount;
            return true;
        }
        return false;
    }
}

// ============================================================================
// 8. MAIN CLASS - Demonstrating all concepts
// ============================================================================

public class OOPDemo {
    public static void main(String[] args) {
        System.out.println("=== Complete OOP Demo ===\n");
        
        // 1. CLASS, CONSTRUCTOR, STATIC
        System.out.println("1. Classes & Objects:");
        Employee emp = new Employee("Alice", 50000, "IT");
        Manager mgr = new Manager("Bob", 80000, "IT", "DevTeam");
        System.out.println("Total employees: " + Employee.getTotalEmployees());
        
        // 2. INHERITANCE & RUNTIME POLYMORPHISM
        System.out.println("\n2. Inheritance & Runtime Polymorphism:");
        Employee[] employees = {emp, mgr, new Developer("Charlie", 60000, "IT", "Java"), 
                               new SeniorDeveloper("David", 90000, "IT", "Python", 8)};
        for (Employee e : employees) {
            e.work(); // Different behavior based on actual object type (runtime polymorphism)
        }
        
        // 3. COMPOSITION
        System.out.println("\n3. Composition (Has-A relationship):");
        Computer laptop = new Computer("Dell Laptop");
        Office mainOffice = new Office("New York");
        WorkStation ws = new WorkStation("Eve", 55000, "IT", laptop, mainOffice);
        ws.work();
        
        // 4. DIAMOND PROBLEM & SOLUTION
        System.out.println("\n4. Diamond Problem & Solution:");
        TeamLead lead = new TeamLead("Frank", 75000, "IT");
        lead.doWork(); // Resolved by explicit override
        
        // 5. COMPILE-TIME POLYMORPHISM (Method Overloading)
        System.out.println("\n5. Method Overloading (Compile-time Polymorphism):");
        Calculator calc = new Calculator();
        System.out.println("calculate(5, 3): " + calc.calculate(5, 3));
        System.out.println("calculate(5.5, 3.2): " + calc.calculate(5.5, 3.2));
        System.out.println("calculate(1, 2, 3): " + calc.calculate(1, 2, 3));
        
        // 6. ABSTRACTION
        System.out.println("\n6. Abstraction (Abstract class & Interface):");
        Vehicle car = new Car("Toyota"); // Cannot instantiate Vehicle directly
        car.displayInfo();
        car.start();
        ((Drivable) car).drive(); // Interface method
        
        // 7. ENCAPSULATION
        System.out.println("\n7. Encapsulation (Data hiding & controlled access):");
        BankAccount account = new BankAccount("ACC001", 1000);
        System.out.println("Initial balance: " + account.getBalance());
        account.deposit(500);
        account.withdraw(200);
        System.out.println("Final balance: " + account.getBalance());
        // account.balance = 9999; // Error: cannot access private field directly
        
        System.out.println("\n=== All OOP concepts demonstrated ===");
    }
}