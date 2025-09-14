# SOLID Principles

## Introduction

After learning Object-Oriented Programming (OOP), you might think: "Great! Now I can create classes and objects, use inheritance, polymorphism, abstraction, and encapsulation. My code will be perfect!" But wait... just knowing OOP concepts doesn't guarantee good code design. You can still write messy, hard-to-maintain, and fragile code even with OOP.

Think about it like building a house. OOP gives you the materials (bricks, cement, wood), but SOLID principles give you the architectural guidelines on how to arrange those materials properly. Without proper architectural principles, your house might collapse, be hard to renovate, or cost too much to maintain.

Before SOLID principles, developers often faced these problems:
- **Rigid Code:** Hard to change without breaking other parts
- **Fragile Code:** One small change breaks multiple unrelated features
- **Immobile Code:** Hard to reuse components in different contexts
- **Viscous Code:** Easier to do wrong things than right things

SOLID principles were introduced by Robert C. Martin (Uncle Bob) to solve these problems. These are not programming language features - these are design principles that guide us on how to structure our classes and their relationships.

**Definition:** SOLID is an acronym for five design principles that help developers write more maintainable, flexible, and robust object-oriented code. When followed together, these principles lead to software that is easier to understand, extend, and modify.

The five SOLID principles are:
1. **S** - Single Responsibility Principle (SRP)
2. **O** - Open/Closed Principle (OCP)
3. **L** - Liskov Substitution Principle (LSP)
4. **I** - Interface Segregation Principle (ISP)
5. **D** - Dependency Inversion Principle (DIP)

Think of SOLID principles as the "best practices" or "quality guidelines" for OOP design. Just like in business, we have quality standards to ensure our products meet customer expectations and are maintainable over time.

---

## Single Responsibility Principle (SRP)

**Definition:** A class should have only one reason to change, meaning it should have only one job or responsibility.

Imagine you're running a restaurant. You wouldn't want your chef to also handle billing, manage inventory, and clean tables. Each person should focus on one main responsibility. The chef cooks, the cashier handles billing, the manager manages inventory, and the cleaner maintains hygiene. If the chef had to do everything, he would be overwhelmed and nothing would be done properly.

Similarly, in programming, each class should have only one responsibility. When a class has multiple responsibilities, it becomes:
- **Hard to understand:** Too many things happening in one place
- **Hard to test:** Multiple reasons for the class to fail
- **Hard to change:** Changing one feature might break another unrelated feature
- **Hard to reuse:** You can't use just one part without dragging the whole class

### Example: Violating SRP

```java
// BAD: This class has multiple responsibilities
class Employee {
    private String name;
    private double salary;
    private String position;

    public Employee(String name, double salary, String position) {
        this.name = name;
        this.salary = salary;
        this.position = position;
    }

    // Responsibility 1: Manage employee data
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    public double getSalary() { return salary; }
    public void setSalary(double salary) { this.salary = salary; }

    // Responsibility 2: Calculate tax (business logic)
    public double calculateTax() {
        if (salary > 50000) {
            return salary * 0.3;
        }
        return salary * 0.2;
    }

    // Responsibility 3: Save to database (data persistence)
    public void saveToDatabase() {
        // Database connection code
        System.out.println("Connecting to database...");
        System.out.println("Saving " + name + " to database");
        // SQL queries, error handling, etc.
    }

    // Responsibility 4: Generate reports (reporting)
    public void printPaySlip() {
        System.out.println("=== PAY SLIP ===");
        System.out.println("Name: " + name);
        System.out.println("Salary: $" + salary);
        System.out.println("Tax: $" + calculateTax());
        System.out.println("Net: $" + (salary - calculateTax()));
    }
}
```

**Problems with above code:**
- If tax calculation rules change, we need to modify the `Employee` class
- If database structure changes, we need to modify the `Employee` class  
- If report format changes, we need to modify the `Employee` class
- We can't reuse tax calculation logic for other entities
- Testing becomes complex - we need to test data management, tax calculation, database operations, and reporting all together

### Example: Following SRP

```java
// GOOD: Each class has single responsibility

// Responsibility 1: Manage employee data only
class Employee {
    private String name;
    private double salary;
    private String position;

    public Employee(String name, double salary, String position) {
        this.name = name;
        this.salary = salary;
        this.position = position;
    }

    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    public double getSalary() { return salary; }
    public void setSalary(double salary) { this.salary = salary; }
    public String getPosition() { return position; }
    public void setPosition(String position) { this.position = position; }
}

// Responsibility 2: Handle tax calculations only
class TaxCalculator {
    public double calculateTax(Employee employee) {
        double salary = employee.getSalary();
        if (salary > 50000) {
            return salary * 0.3;
        }
        return salary * 0.2;
    }
}

// Responsibility 3: Handle database operations only
class EmployeeRepository {
    public void save(Employee employee) {
        System.out.println("Connecting to database...");
        System.out.println("Saving " + employee.getName() + " to database");
        // Database specific code here
    }

    public Employee findByName(String name) {
        // Database query logic
        System.out.println("Finding employee: " + name);
        return null; // simplified
    }
}

// Responsibility 4: Handle report generation only
class PaySlipGenerator {
    private TaxCalculator taxCalculator;

    public PaySlipGenerator(TaxCalculator taxCalculator) {
        this.taxCalculator = taxCalculator;
    }

    public void generatePaySlip(Employee employee) {
        double tax = taxCalculator.calculateTax(employee);
        double netSalary = employee.getSalary() - tax;

        System.out.println("=== PAY SLIP ===");
        System.out.println("Name: " + employee.getName());
        System.out.println("Salary: $" + employee.getSalary());
        System.out.println("Tax: $" + tax);
        System.out.println("Net: $" + netSalary);
    }
}

// Usage
public class Main {
    public static void main(String[] args) {
        Employee emp = new Employee("John Doe", 60000, "Developer");
        TaxCalculator taxCalc = new TaxCalculator();
        EmployeeRepository repo = new EmployeeRepository();
        PaySlipGenerator paySlipGen = new PaySlipGenerator(taxCalc);

        repo.save(emp);
        paySlipGen.generatePaySlip(emp);
    }
}
```

**Benefits of following SRP:**
- **Easy to understand:** Each class has a clear, single purpose
- **Easy to test:** Test each responsibility independently
- **Easy to change:** Changing tax rules only affects `TaxCalculator`
- **Reusable:** Can use `TaxCalculator` for other entities like contractors
- **Maintainable:** Bug in database code won't affect tax calculation

### When to Apply SRP
- When you see a class doing multiple unrelated things
- When you find yourself saying "and" while describing what a class does ("This class manages employee data AND calculates tax AND saves to database")
- When changes to one feature require modifying the same class as another unrelated feature
- When you want to reuse just one part of a class's functionality

---

## Open/Closed Principle (OCP)

**Definition:** Software entities (classes, modules, functions) should be open for extension but closed for modification.

Think about your smartphone. When you want new functionality, you don't break open the phone and rewire its circuits. Instead, you install an app (extension) without modifying the phone's core system. The phone is "closed" for modification (you can't change its internal structure) but "open" for extension (you can add new apps).

In programming, once a class is written, tested, and working correctly, we should be able to extend its behavior without modifying its existing code. This prevents introducing bugs in working code while allowing new features.

**Open for Extension:** We can add new functionality
**Closed for Modification:** We should not change existing, working code

### Example: Violating OCP

```java
// BAD: Violates Open/Closed Principle
class AreaCalculator {
    public double calculateArea(Object shape) {
        if (shape instanceof Rectangle) {
            Rectangle rectangle = (Rectangle) shape;
            return rectangle.getWidth() * rectangle.getHeight();
        } else if (shape instanceof Circle) {
            Circle circle = (Circle) shape;
            return Math.PI * circle.getRadius() * circle.getRadius();
        }
        // What if we want to add Triangle? We need to modify this method!
        // What if we want to add Pentagon? Modify again!
        return 0;
    }
}

class Rectangle {
    private double width, height;
    
    public Rectangle(double width, double height) {
        this.width = width;
        this.height = height;
    }
    
    public double getWidth() { return width; }
    public double getHeight() { return height; }
}

class Circle {
    private double radius;
    
    public Circle(double radius) {
        this.radius = radius;
    }
    
    public double getRadius() { return radius; }
}
```

**Problems with above code:**
- Every time we add a new shape, we must modify `AreaCalculator`
- Risk of introducing bugs in existing, working code
- Violates Single Responsibility - `AreaCalculator` knows about all shapes
- Hard to test - need to test all shape types together

### Example: Following OCP

```java
// GOOD: Follows Open/Closed Principle

// Define an interface for shapes
interface Shape {
    double calculateArea();
}

// Each shape implements its own area calculation
class Rectangle implements Shape {
    private double width, height;
    
    public Rectangle(double width, double height) {
        this.width = width;
        this.height = height;
    }
    
    @Override
    public double calculateArea() {
        return width * height;
    }
}

class Circle implements Shape {
    private double radius;
    
    public Circle(double radius) {
        this.radius = radius;
    }
    
    @Override
    public double calculateArea() {
        return Math.PI * radius * radius;
    }
}

// Adding new shapes without modifying existing code
class Triangle implements Shape {
    private double base, height;
    
    public Triangle(double base, double height) {
        this.base = base;
        this.height = height;
    }
    
    @Override
    public double calculateArea() {
        return 0.5 * base * height;
    }
}

// Calculator is now closed for modification, open for extension
class AreaCalculator {
    public double calculateArea(Shape shape) {
        return shape.calculateArea(); // No need to modify this method!
    }
    
    public double calculateTotalArea(Shape[] shapes) {
        double total = 0;
        for (Shape shape : shapes) {
            total += calculateArea(shape);
        }
        return total;
    }
}

// Usage
public class Main {
    public static void main(String[] args) {
        Shape[] shapes = {
            new Rectangle(5, 4),
            new Circle(3),
            new Triangle(6, 8)  // New shape added without changing AreaCalculator!
        };
        
        AreaCalculator calculator = new AreaCalculator();
        System.out.println("Total area: " + calculator.calculateTotalArea(shapes));
    }
}
```

### Another Example: Notification System

```java
// Following OCP for notification system

interface NotificationSender {
    void send(String message);
}

class EmailSender implements NotificationSender {
    @Override
    public void send(String message) {
        System.out.println("Sending email: " + message);
    }
}

class SMSSender implements NotificationSender {
    @Override
    public void send(String message) {
        System.out.println("Sending SMS: " + message);
    }
}

// Adding new notification types without modifying existing code
class PushNotificationSender implements NotificationSender {
    @Override
    public void send(String message) {
        System.out.println("Sending push notification: " + message);
    }
}

class SlackSender implements NotificationSender {
    @Override
    public void send(String message) {
        System.out.println("Sending Slack message: " + message);
    }
}

class NotificationService {
    public void notify(NotificationSender sender, String message) {
        sender.send(message); // Works with any notification type!
    }
}
```

**Benefits of following OCP:**
- **No risk to existing code:** Adding new features doesn't break working functionality
- **Easy to extend:** Add new shapes/notifications without touching existing classes
- **Better testing:** Test new features independently
- **Follows SRP:** Each shape handles its own area calculation

### How to Achieve OCP
1. **Use abstraction:** Create interfaces or abstract classes
2. **Dependency injection:** Depend on abstractions, not concrete classes
3. **Strategy pattern:** Encapsulate algorithms in separate classes
4. **Template method pattern:** Define algorithm structure, let subclasses fill details

### When to Apply OCP
- When you frequently add new similar functionality (new shapes, new payment methods, new notification types)
- When modification of existing code is risky or expensive
- When you want to make your system extensible for future requirements
- When you're building libraries or frameworks that others will extend

---

## Liskov Substitution Principle (LSP)

**Definition:** Objects of a superclass should be replaceable with objects of its subclasses without breaking the application's functionality.

Imagine you have a universal remote control that works with any TV. If you replace your Samsung TV with an LG TV, the remote should still work the same way - turn on/off, change channels, adjust volume. If the LG TV suddenly required you to press buttons in a different sequence or behaved completely differently, it would violate the "substitution" principle.

In programming, if class `B` is a subtype of class `A`, then we should be able to replace objects of type `A` with objects of type `B` without altering the correctness of the program. The subclass should strengthen, not weaken, the parent class's behavior.

LSP ensures that inheritance is used correctly. Just because you CAN inherit doesn't mean you SHOULD inherit.

### Example: Violating LSP

```java
// BAD: Violates Liskov Substitution Principle

class Rectangle {
    protected double width, height;
    
    public Rectangle(double width, double height) {
        this.width = width;
        this.height = height;
    }
    
    public void setWidth(double width) {
        this.width = width;
    }
    
    public void setHeight(double height) {
        this.height = height;
    }
    
    public double getWidth() { return width; }
    public double getHeight() { return height; }
    
    public double getArea() {
        return width * height;
    }
}

// This seems logical - Square IS-A Rectangle, right?
class Square extends Rectangle {
    public Square(double side) {
        super(side, side);
    }
    
    // Problem: Square changes the behavior of Rectangle!
    @Override
    public void setWidth(double width) {
        this.width = width;
        this.height = width; // Square must keep width = height
    }
    
    @Override
    public void setHeight(double height) {
        this.width = height;
        this.height = height; // Square must keep width = height
    }
}

// This function works correctly with Rectangle
class RectangleUser {
    public void demonstrateProblem(Rectangle rect) {
        rect.setWidth(5);
        rect.setHeight(4);
        
        // For Rectangle: area should be 5 * 4 = 20
        // For Square: area will be 4 * 4 = 16 (unexpected!)
        System.out.println("Expected area: 20, Actual area: " + rect.getArea());
        
        // This assertion will fail for Square!
        assert rect.getArea() == 20 : "Area calculation failed!";
    }
}

public class Main {
    public static void main(String[] args) {
        RectangleUser user = new RectangleUser();
        
        Rectangle rect = new Rectangle(3, 4);
        user.demonstrateProblem(rect); // Works fine
        
        Rectangle square = new Square(3); // Square IS-A Rectangle?
        user.demonstrateProblem(square); // Fails! LSP violated
    }
}
```

**Problems with above code:**
- `Square` changes the expected behavior of `Rectangle`
- Code written for `Rectangle` breaks when `Square` is used
- The "IS-A" relationship doesn't hold behaviorally (only structurally)

### Example: Following LSP

```java
// GOOD: Following Liskov Substitution Principle

// Create a proper abstraction
abstract class Shape {
    public abstract double getArea();
    public abstract double getPerimeter();
}

class Rectangle extends Shape {
    private double width, height;
    
    public Rectangle(double width, double height) {
        this.width = width;
        this.height = height;
    }
    
    public void setWidth(double width) {
        this.width = width;
    }
    
    public void setHeight(double height) {
        this.height = height;
    }
    
    public double getWidth() { return width; }
    public double getHeight() { return height; }
    
    @Override
    public double getArea() {
        return width * height;
    }
    
    @Override
    public double getPerimeter() {
        return 2 * (width + height);
    }
}

class Square extends Shape {
    private double side;
    
    public Square(double side) {
        this.side = side;
    }
    
    public void setSide(double side) {
        this.side = side;
    }
    
    public double getSide() { return side; }
    
    @Override
    public double getArea() {
        return side * side;
    }
    
    @Override
    public double getPerimeter() {
        return 4 * side;
    }
}

// Now both Rectangle and Square can be used interchangeably as Shape
class ShapeCalculator {
    public void printShapeInfo(Shape shape) {
        System.out.println("Area: " + shape.getArea());
        System.out.println("Perimeter: " + shape.getPerimeter());
        // This works correctly for both Rectangle and Square
    }
    
    public double calculateTotalArea(Shape[] shapes) {
        double total = 0;
        for (Shape shape : shapes) {
            total += shape.getArea(); // Works correctly for all shapes
        }
        return total;
    }
}
```

### Another Example: Bird Hierarchy

```java
// BAD: Violating LSP
class Bird {
    public void fly() {
        System.out.println("Flying...");
    }
}

class Penguin extends Bird {
    @Override
    public void fly() {
        throw new UnsupportedOperationException("Penguins can't fly!");
        // This breaks LSP - client code expecting Bird behavior will fail
    }
}

// GOOD: Following LSP
abstract class Bird {
    public abstract void move();
    public abstract void eat();
}

class FlyingBird extends Bird {
    @Override
    public void move() {
        fly();
    }
    
    public void fly() {
        System.out.println("Flying...");
    }
    
    @Override
    public void eat() {
        System.out.println("Eating...");
    }
}

class SwimmingBird extends Bird {
    @Override
    public void move() {
        swim();
    }
    
    public void swim() {
        System.out.println("Swimming...");
    }
    
    @Override
    public void eat() {
        System.out.println("Eating fish...");
    }
}

class Eagle extends FlyingBird {
    // Eagle can fly - no problem
}

class Penguin extends SwimmingBird {
    // Penguin can swim - no problem
}

class BirdWatcher {
    public void observeBird(Bird bird) {
        bird.move(); // Works correctly for all birds
        bird.eat();  // Works correctly for all birds
    }
}
```

### Key Rules for LSP

1. **Preconditions cannot be strengthened:** Subclass cannot demand more restrictive inputs
2. **Postconditions cannot be weakened:** Subclass cannot provide less than what parent promised
3. **Exceptions:** Subclass cannot throw new exceptions that parent doesn't throw
4. **History constraint:** Subclass cannot modify state in ways parent wouldn't allow

### How to Follow LSP

1. **Think about behavior, not just structure:** Just because something "IS-A" doesn't mean it should inherit
2. **Use composition over inheritance:** When LSP is violated, consider composition
3. **Design contracts carefully:** Define what each method should do and ensure subclasses honor it
4. **Use abstract classes or interfaces:** Define common behavior that all subclasses can properly implement

### When LSP is Important

- When building class hierarchies
- When using polymorphism extensively  
- When designing libraries or frameworks
- When you want code to work with parent classes and all their subclasses seamlessly

---

## Interface Segregation Principle (ISP)

**Definition:** A client should not be forced to implement interfaces they don't use. Many client-specific interfaces are better than one general-purpose interface.

Think about a multi-function printer in your office. It can print, scan, fax, and photocopy. Now imagine if every device that wanted to connect to this printer had to support ALL these functions. Your simple printer-only device would be forced to have fax and scan capabilities it doesn't need. ISP says: create separate interfaces for each function (IPrinter, IScanner, IFax) so devices can implement only what they actually use.

In programming, when we create large interfaces with many methods, we force implementing classes to provide code for methods they don't actually need. This leads to:
- **Fat interfaces:** Interfaces doing too many things
- **Unnecessary dependencies:** Classes depending on methods they don't use
- **Violation of SRP:** Interface has multiple responsibilities
- **Hard to maintain:** Changes in one part affect all implementers

### Example: Violating ISP

```java
// BAD: Fat interface violating ISP
interface Worker {
    void work();
    void eat();
    void sleep();
    void attendMeeting();
    void writeCode();
    void managePeople();
    void createDocuments();
    void testSoftware();
}

// Human worker - can do everything
class HumanWorker implements Worker {
    @Override
    public void work() { System.out.println("Human working..."); }
    
    @Override
    public void eat() { System.out.println("Human eating..."); }
    
    @Override
    public void sleep() { System.out.println("Human sleeping..."); }
    
    @Override
    public void attendMeeting() { System.out.println("Attending meeting..."); }
    
    @Override
    public void writeCode() { System.out.println("Writing code..."); }
    
    @Override
    public void managePeople() { System.out.println("Managing people..."); }
    
    @Override
    public void createDocuments() { System.out.println("Creating documents..."); }
    
    @Override
    public void testSoftware() { System.out.println("Testing software..."); }
}

// Robot worker - forced to implement methods it doesn't need!
class RobotWorker implements Worker {
    @Override
    public void work() { System.out.println("Robot working..."); }
    
    @Override
    public void eat() { 
        // Robot doesn't eat! But forced to implement
        throw new UnsupportedOperationException("Robot doesn't eat!");
    }
    
    @Override
    public void sleep() { 
        // Robot doesn't sleep! But forced to implement
        throw new UnsupportedOperationException("Robot doesn't sleep!");
    }
    
    @Override
    public void attendMeeting() { 
        throw new UnsupportedOperationException("Robot doesn't attend meetings!");
    }
    
    @Override
    public void writeCode() { System.out.println("Robot writing code..."); }
    
    @Override
    public void managePeople() { 
        throw new UnsupportedOperationException("Robot doesn't manage people!");
    }
    
    @Override
    public void createDocuments() { System.out.println("Robot creating documents..."); }
    
    @Override
    public void testSoftware() { System.out.println("Robot testing software..."); }
}

// Junior Developer - forced to implement management methods!
class JuniorDeveloper implements Worker {
    @Override
    public void work() { System.out.println("Junior developer working..."); }
    
    @Override
    public void eat() { System.out.println("Junior developer eating..."); }
    
    @Override
    public void sleep() { System.out.println("Junior developer sleeping..."); }
    
    @Override
    public void attendMeeting() { System.out.println("Attending meeting..."); }
    
    @Override
    public void writeCode() { System.out.println("Writing code..."); }
    
    @Override
    public void managePeople() { 
        // Junior doesn't manage people!
        throw new UnsupportedOperationException("Junior doesn't manage people!");
    }
    
    @Override
    public void createDocuments() { System.out.println("Creating simple documents..."); }
    
    @Override
    public void testSoftware() { System.out.println("Basic testing..."); }
}
```

**Problems with above code:**
- Classes implement methods they don't need
- Many methods throw `UnsupportedOperationException`
- Interface changes affect all implementers
- Hard to add new types of workers

### Example: Following ISP

```java
// GOOD: Segregated interfaces following ISP

// Basic work capability
interface Workable {
    void work();
}

// Biological needs
interface Biological {
    void eat();
    void sleep();
}

// Communication capabilities
interface Communicable {
    void attendMeeting();
}

// Technical capabilities  
interface Programmer {
    void writeCode();
    void testSoftware();
}

// Documentation capabilities
interface DocumentCreator {
    void createDocuments();
}

// Management capabilities
interface Manager {
    void managePeople();
}

// Now each class implements only what it needs

class HumanWorker implements Workable, Biological, Communicable {
    @Override
    public void work() { System.out.println("Human working..."); }
    
    @Override
    public void eat() { System.out.println("Human eating..."); }
    
    @Override
    public void sleep() { System.out.println("Human sleeping..."); }
    
    @Override
    public void attendMeeting() { System.out.println("Attending meeting..."); }
}

class RobotWorker implements Workable, Programmer, DocumentCreator {
    @Override
    public void work() { System.out.println("Robot working..."); }
    
    @Override
    public void writeCode() { System.out.println("Robot writing code..."); }
    
    @Override
    public void testSoftware() { System.out.println("Robot testing software..."); }
    
    @Override
    public void createDocuments() { System.out.println("Robot creating documents..."); }
}

class JuniorDeveloper implements Workable, Biological, Communicable, Programmer {
    @Override
    public void work() { System.out.println("Junior developer working..."); }
    
    @Override
    public void eat() { System.out.println("Junior developer eating..."); }
    
    @Override
    public void sleep() { System.out.println("Junior developer sleeping..."); }
    
    @Override
    public void attendMeeting() { System.out.println("Attending meeting..."); }
    
    @Override
    public void writeCode() { System.out.println("Writing code..."); }
    
    @Override
    public void testSoftware() { System.out.println("Basic testing..."); }
}

class SeniorDeveloper implements Workable, Biological, Communicable, Programmer, Manager, DocumentCreator {
    @Override
    public void work() { System.out.println("Senior developer working..."); }
    
    @Override
    public void eat() { System.out.println("Senior developer eating..."); }
    
    @Override
    public void sleep() { System.out.println("Senior developer sleeping..."); }
    
    @Override
    public void attendMeeting() { System.out.println("Leading meeting..."); }
    
    @Override
    public void writeCode() { System.out.println("Writing complex code..."); }
    
    @Override
    public void testSoftware() { System.out.println("Advanced testing..."); }
    
    @Override
    public void managePeople() { System.out.println("Managing team..."); }
    
    @Override
    public void createDocuments() { System.out.println("Creating technical documents..."); }
}

// Usage - each client depends only on what it needs
class WorkManager {
    public void assignWork(Workable worker) {
        worker.work(); // Only needs work capability
    }
}

class MeetingOrganizer {
    public void organizeMeeting(Communicable participant) {
        participant.attendMeeting(); // Only needs communication capability
    }
}

class ProjectManager {
    public void assignCodingTask(Programmer programmer) {
        programmer.writeCode(); // Only needs programming capability
        programmer.testSoftware();
    }
    
    public void assignManagementTask(Manager manager) {
        manager.managePeople(); // Only needs management capability
    }
}
```

### Another Example: Payment Processing

```java
// BAD: Fat interface
interface PaymentProcessor {
    void processPayment(double amount);
    void refundPayment(String transactionId);
    void subscribeMonthly(double amount);
    void cancelSubscription(String subscriptionId);
    void sendInvoice(String email);
    void generateReport();
}

// GOOD: Segregated interfaces
interface PaymentProcessor {
    void processPayment(double amount);
}

interface RefundProcessor {
    void refundPayment(String transactionId);
}

interface SubscriptionProcessor {
    void subscribeMonthly(double amount);
    void cancelSubscription(String subscriptionId);
}

interface InvoiceGenerator {
    void sendInvoice(String email);
}

interface ReportGenerator {
    void generateReport();
}

// Simple payment gateway - only processes payments
class SimplePaymentGateway implements PaymentProcessor {
    @Override
    public void processPayment(double amount) {
        System.out.println("Processing payment: $" + amount);
    }
}

// Full-featured payment system
class AdvancedPaymentGateway implements PaymentProcessor, RefundProcessor, SubscriptionProcessor, InvoiceGenerator {
    @Override
    public void processPayment(double amount) {
        System.out.println("Processing payment: $" + amount);
    }
    
    @Override
    public void refundPayment(String transactionId) {
        System.out.println("Refunding transaction: " + transactionId);
    }
    
    @Override
    public void subscribeMonthly(double amount) {
        System.out.println("Setting up monthly subscription: $" + amount);
    }
    
    @Override
    public void cancelSubscription(String subscriptionId) {
        System.out.println("Cancelling subscription: " + subscriptionId);
    }
    
    @Override
    public void sendInvoice(String email) {
        System.out.println("Sending invoice to: " + email);
    }
}

// Client that only needs basic payment processing
class SimpleShop {
    private PaymentProcessor paymentProcessor;
    
    public SimpleShop(PaymentProcessor paymentProcessor) {
        this.paymentProcessor = paymentProcessor;
    }
    
    public void checkout(double amount) {
        paymentProcessor.processPayment(amount);
        // Only depends on payment processing, nothing else
    }
}
```

### Benefits of Following ISP

- **Reduced coupling:** Clients depend only on methods they use
- **Better maintainability:** Changes in one interface don't affect unrelated clients
- **Easier testing:** Mock only the interfaces you actually need
- **Flexibility:** Classes can implement multiple small interfaces as needed
- **Clear responsibilities:** Each interface has a focused purpose

### How to Apply ISP

1. **Identify fat interfaces:** Look for interfaces with many methods
2. **Group related methods:** Find methods that belong together
3. **Create focused interfaces:** Split large interfaces into smaller, cohesive ones
4. **Use composition:** Classes can implement multiple small interfaces
5. **Think from client perspective:** What does each client actually need?

### When to Apply ISP

- When you see interfaces with many unrelated methods
- When implementing classes throw `UnsupportedOperationException`
- When different clients use different subsets of an interface
- When interface changes affect many unrelated classes
- When you want to reduce coupling between components

---

## Dependency Inversion Principle (DIP)

**Definition:** High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions.

Think about electrical appliances in your home. Your TV, refrigerator, and computer don't depend on a specific type of power plant (coal, nuclear, solar). They depend on the electrical socket (abstraction). The power plant can change from coal to solar, but your appliances keep working. The "high-level" appliances depend on the "abstraction" (electrical socket), not on the "low-level" implementation (specific power plant).

In programming:
- **High-level modules:** Business logic, use cases, core application functionality
- **Low-level modules:** Database access, file systems, external APIs, UI components
- **Abstractions:** Interfaces, abstract classes that define contracts

DIP prevents tight coupling between layers and makes your system more flexible and testable.

### Example: Violating DIP

```java
// BAD: High-level class depends directly on low-level classes

// Low-level modules (concrete implementations)
class MySQLDatabase {
    public void save(String data) {
        System.out.println("Saving to MySQL: " + data);
        // MySQL specific code
    }
    
    public String fetch(String id) {
        System.out.println("Fetching from MySQL: " + id);
        return "MySQL data for " + id;
    }
}

class EmailService {
    public void sendEmail(String to, String message) {
        System.out.println("Sending email to " + to + ": " + message);
        // SMTP specific code
    }
}

class FileLogger {
    public void log(String message) {
        System.out.println("Logging to file: " + message);
        // File I/O specific code
    }
}

// High-level module depending on low-level modules
class OrderService {
    private MySQLDatabase database;      // Depends on concrete MySQL
    private EmailService emailService;  // Depends on concrete Email
    private FileLogger logger;          // Depends on concrete File Logger
    
    public OrderService() {
        this.database = new MySQLDatabase();    // Tight coupling!
        this.emailService = new EmailService(); // Tight coupling!
        this.logger = new FileLogger();         // Tight coupling!
    }
    
    public void processOrder(String orderData) {
        logger.log("Processing order: " + orderData);
        
        // Business logic
        String orderId = "ORDER-" + System.currentTimeMillis();
        
        database.save(orderId + ": " + orderData);
        emailService.sendEmail("customer@email.com", "Order confirmed: " + orderId);
        
        logger.log("Order processed: " + orderId);
    }
}
```

**Problems with above code:**
- `OrderService` is tightly coupled to specific implementations
- Cannot switch to PostgreSQL, SMS, or database logging without modifying `OrderService`
- Hard to test - cannot mock dependencies
- Violates OCP - adding new database types requires modifying `OrderService`
- Business logic mixed with infrastructure concerns

### Example: Following DIP

```java
// GOOD: Following Dependency Inversion Principle

// Abstractions (interfaces)
interface Database {
    void save(String data);
    String fetch(String id);
}

interface NotificationService {
    void sendNotification(String to, String message);
}

interface Logger {
    void log(String message);
}

// Low-level modules implementing abstractions
class MySQLDatabase implements Database {
    @Override
    public void save(String data) {
        System.out.println("Saving to MySQL: " + data);
    }
    
    @Override
    public String fetch(String id) {
        System.out.println("Fetching from MySQL: " + id);
        return "MySQL data for " + id;
    }
}

class PostgreSQLDatabase implements Database {
    @Override
    public void save(String data) {
        System.out.println("Saving to PostgreSQL: " + data);
    }
    
    @Override
    public String fetch(String id) {
        System.out.println("Fetching from PostgreSQL: " + id);
        return "PostgreSQL data for " + id;
    }
}

class EmailService implements NotificationService {
    @Override
    public void sendNotification(String to, String message) {
        System.out.println("Sending email to " + to + ": " + message);
    }
}

class SMSService implements NotificationService {
    @Override
    public void sendNotification(String to, String message) {
        System.out.println("Sending SMS to " + to + ": " + message);
    }
}

class FileLogger implements Logger {
    @Override
    public void log(String message) {
        System.out.println("Logging to file: " + message);
    }
}

class DatabaseLogger implements Logger {
    @Override
    public void log(String message) {
        System.out.println("Logging to database: " + message);
    }
}

// High-level module depending on abstractions
class OrderService {
    private Database database;
    private NotificationService notificationService;
    private Logger logger;
    
    // Dependency injection through constructor
    public OrderService(Database database, NotificationService notificationService, Logger logger) {
        this.database = database;
        this.notificationService = notificationService;
        this.logger = logger;
    }
    
    public void processOrder(String orderData) {
        logger.log("Processing order: " + orderData);
        
        // Pure business logic - no infrastructure concerns
        String orderId = generateOrderId();
        String processedData = processOrderData(orderData);
        
        database.save(orderId + ": " + processedData);
        notificationService.sendNotification("customer@email.com", "Order confirmed: " + orderId);
        
        logger.log("Order processed: " + orderId);
    }
    
    private String generateOrderId() {
        return "ORDER-" + System.currentTimeMillis();
    }
    
    private String processOrderData(String orderData) {
        // Business logic for processing order
        return orderData.toUpperCase();
    }
}

// Configuration/Composition root
class OrderApplicationConfig {
    public OrderService createOrderService() {
        Database database = new MySQLDatabase();
        NotificationService notificationService = new EmailService();
        Logger logger = new FileLogger();
        
        return new OrderService(database, notificationService, logger);
    }
    
    public OrderService createTestOrderService() {
        // Different configuration for testing
        Database database = new PostgreSQLDatabase();
        NotificationService notificationService = new SMSService();
        Logger logger = new DatabaseLogger();
        
        return new OrderService(database, notificationService, logger);
    }
}

// Usage
public class Main {
    public static void main(String[] args) {
        OrderApplicationConfig config = new OrderApplicationConfig();
        
        // Production configuration
        OrderService orderService = config.createOrderService();
        orderService.processOrder("Product A, Quantity: 2");
        
        System.out.println("---");
        
        // Test configuration
        OrderService testOrderService = config.createTestOrderService();
        testOrderService.processOrder("Product B, Quantity: 1");
    }
}
```

### DIP with Dependency Injection Frameworks

```java
// Using annotations for dependency injection (Spring-like)
class OrderService {
    private Database database;
    private NotificationService notificationService;
    private Logger logger;
    
    // Constructor injection
    public OrderService(Database database, 
                       NotificationService notificationService, 
                       Logger logger) {
        this.database = database;
        this.notificationService = notificationService;
        this.logger = logger;
    }
    
    // Business logic remains the same...
}

// Configuration class
class AppConfig {
    public Database database() {
        return new MySQLDatabase();
    }
    
    public NotificationService notificationService() {
        return new EmailService();
    }
    
    public Logger logger() {
        return new FileLogger();
    }
    
    public OrderService orderService() {
        return new OrderService(database(), notificationService(), logger());
    }
}
```

### Testing with DIP

```java
// Easy to test with DIP
class OrderServiceTest {
    
    // Mock implementations for testing
    class MockDatabase implements Database {
        private boolean saveCalled = false;
        
        @Override
        public void save(String data) {
            saveCalled = true;
            System.out.println("Mock save called with: " + data);
        }
        
        @Override
        public String fetch(String id) {
            return "mock data";
        }
        
        public boolean wasSaveCalled() {
            return saveCalled;
        }
    }
    
    class MockNotificationService implements NotificationService {
        private boolean notificationSent = false;
        
        @Override
        public void sendNotification(String to, String message) {
            notificationSent = true;
            System.out.println("Mock notification sent");
        }
        
        public boolean wasNotificationSent() {
            return notificationSent;
        }
    }
    
    class MockLogger implements Logger {
        @Override
        public void log(String message) {
            System.out.println("Mock log: " + message);
        }
    }
    
    public void testProcessOrder() {
        // Arrange
        MockDatabase mockDatabase = new MockDatabase();
        MockNotificationService mockNotificationService = new MockNotificationService();
        MockLogger mockLogger = new MockLogger();
        
        OrderService orderService = new OrderService(mockDatabase, mockNotificationService, mockLogger);
        
        // Act
        orderService.processOrder("Test order");
        
        // Assert
        assert mockDatabase.wasSaveCalled() : "Database save should be called";
        assert mockNotificationService.wasNotificationSent() : "Notification should be sent";
        System.out.println("Test passed!");
    }
}
```

### Benefits of Following DIP

- **Flexibility:** Easy to switch implementations
- **Testability:** Easy to mock dependencies
- **Maintainability:** Changes in low-level modules don't affect high-level modules
- **Reusability:** High-level modules can work with different low-level implementations
- **Separation of concerns:** Business logic separated from infrastructure

### How to Apply DIP

1. **Identify dependencies:** Find where high-level modules depend on low-level modules
2. **Create abstractions:** Define interfaces for the dependencies
3. **Implement abstractions:** Make low-level modules implement the interfaces
4. **Inject dependencies:** Use constructor injection, setter injection, or dependency injection frameworks
5. **Configure at composition root:** Wire up dependencies at application startup

### Dependency Injection Types

1. **Constructor Injection:** Pass dependencies through constructor (recommended)
2. **Setter Injection:** Set dependencies through setter methods
3. **Interface Injection:** Dependencies provide injector method to inject dependency

### When to Apply DIP

- When high-level modules depend directly on low-level modules
- When you want to make code testable
- When you need to switch implementations (databases, services, etc.)
- When building layered architectures
- When using frameworks that support dependency injection

---

## SOLID Principles Working Together

The SOLID principles don't work in isolation - they complement each other to create well-designed, maintainable software. Let's see how they work together in a real-world example.

### Example: E-commerce Order Processing System

```java
// Following all SOLID principles together

// === SRP: Each class has single responsibility ===

class Order {
    private String orderId;
    private String customerId;
    private List<OrderItem> items;
    private OrderStatus status;
    private double totalAmount;
    
    public Order(String orderId, String customerId, List<OrderItem> items) {
        this.orderId = orderId;
        this.customerId = customerId;
        this.items = items;
        this.status = OrderStatus.PENDING;
        this.totalAmount = calculateTotal();
    }
    
    private double calculateTotal() {
        return items.stream().mapToDouble(OrderItem::getPrice).sum();
    }
    
    // Getters and basic methods only
    public String getOrderId() { return orderId; }
    public String getCustomerId() { return customerId; }
    public List<OrderItem> getItems() { return items; }
    public OrderStatus getStatus() { return status; }
    public double getTotalAmount() { return totalAmount; }
    
    public void updateStatus(OrderStatus status) { 
        this.status = status; 
    }
}

class OrderItem {
    private String productId;
    private int quantity;
    private double price;
    
    public OrderItem(String productId, int quantity, double price) {
        this.productId = productId;
        this.quantity = quantity;
        this.price = price;
    }
    
    public double getPrice() { return price * quantity; }
    // Other getters...
}

enum OrderStatus { PENDING, CONFIRMED, SHIPPED, DELIVERED, CANCELLED }

// === ISP: Segregated interfaces ===

interface OrderRepository {
    void save(Order order);
    Order findById(String orderId);
}

interface PaymentProcessor {
    PaymentResult processPayment(String customerId, double amount);
}

interface NotificationSender {
    void sendOrderConfirmation(Order order);
    void sendShippingNotification(Order order);
}

interface InventoryChecker {
    boolean isAvailable(String productId, int quantity);
    void reserveItems(List<OrderItem> items);
}

interface OrderValidator {
    ValidationResult validate(Order order);
}

// === OCP & DIP: Open for extension, depend on abstractions ===

abstract class OrderProcessor {
    protected OrderRepository orderRepository;
    protected PaymentProcessor paymentProcessor;
    protected NotificationSender notificationSender;
    protected InventoryChecker inventoryChecker;
    protected OrderValidator orderValidator;
    
    public OrderProcessor(OrderRepository orderRepository,
                         PaymentProcessor paymentProcessor,
                         NotificationSender notificationSender,
                         InventoryChecker inventoryChecker,
                         OrderValidator orderValidator) {
        this.orderRepository = orderRepository;
        this.paymentProcessor = paymentProcessor;
        this.notificationSender = notificationSender;
        this.inventoryChecker = inventoryChecker;
        this.orderValidator = orderValidator;
    }
    
    // Template method pattern - following OCP
    public final ProcessingResult processOrder(Order order) {
        ValidationResult validation = orderValidator.validate(order);
        if (!validation.isValid()) {
            return ProcessingResult.failure("Validation failed: " + validation.getErrors());
        }
        
        if (!inventoryChecker.isAvailable(order.getItems().get(0).getProductId(), 
                                          order.getItems().get(0).getQuantity())) {
            return ProcessingResult.failure("Insufficient inventory");
        }
        
        PaymentResult paymentResult = paymentProcessor.processPayment(
            order.getCustomerId(), order.getTotalAmount());
        
        if (!paymentResult.isSuccess()) {
            return ProcessingResult.failure("Payment failed: " + paymentResult.getErrorMessage());
        }
        
        return executeSpecificProcessing(order);
    }
    
    // Subclasses provide specific implementation - OCP
    protected abstract ProcessingResult executeSpecificProcessing(Order order);
}

// === LSP: Subclasses can substitute parent ===

class StandardOrderProcessor extends OrderProcessor {
    public StandardOrderProcessor(OrderRepository orderRepository,
                                 PaymentProcessor paymentProcessor,
                                 NotificationSender notificationSender,
                                 InventoryChecker inventoryChecker,
                                 OrderValidator orderValidator) {
        super(orderRepository, paymentProcessor, notificationSender, inventoryChecker, orderValidator);
    }
    
    @Override
    protected ProcessingResult executeSpecificProcessing(Order order) {
        inventoryChecker.reserveItems(order.getItems());
        order.updateStatus(OrderStatus.CONFIRMED);
        orderRepository.save(order);
        notificationSender.sendOrderConfirmation(order);
        
        return ProcessingResult.success("Standard order processed successfully");
    }
}

class ExpressOrderProcessor extends OrderProcessor {
    public ExpressOrderProcessor(OrderRepository orderRepository,
                                PaymentProcessor paymentProcessor,
                                NotificationSender notificationSender,
                                InventoryChecker inventoryChecker,
                                OrderValidator orderValidator) {
        super(orderRepository, paymentProcessor, notificationSender, inventoryChecker, orderValidator);
    }
    
    @Override
    protected ProcessingResult executeSpecificProcessing(Order order) {
        inventoryChecker.reserveItems(order.getItems());
        order.updateStatus(OrderStatus.CONFIRMED);
        orderRepository.save(order);
        
        // Express orders get immediate shipping notification
        notificationSender.sendOrderConfirmation(order);
        notificationSender.sendShippingNotification(order);
        
        return ProcessingResult.success("Express order processed successfully");
    }
}

// Implementation classes (following all principles)
class DatabaseOrderRepository implements OrderRepository {
    @Override
    public void save(Order order) {
        System.out.println("Saving order to database: " + order.getOrderId());
    }
    
    @Override
    public Order findById(String orderId) {
        System.out.println("Finding order: " + orderId);
        return null; // Simplified
    }
}

class CreditCardProcessor implements PaymentProcessor {
    @Override
    public PaymentResult processPayment(String customerId, double amount) {
        System.out.println("Processing credit card payment: $" + amount);
        return PaymentResult.success("Payment processed");
    }
}

class EmailNotificationSender implements NotificationSender {
    @Override
    public void sendOrderConfirmation(Order order) {
        System.out.println("Sending email confirmation for order: " + order.getOrderId());
    }
    
    @Override
    public void sendShippingNotification(Order order) {
        System.out.println("Sending email shipping notification for order: " + order.getOrderId());
    }
}

class WarehouseInventoryChecker implements InventoryChecker {
    @Override
    public boolean isAvailable(String productId, int quantity) {
        System.out.println("Checking inventory for product: " + productId);
        return true; // Simplified
    }
    
    @Override
    public void reserveItems(List<OrderItem> items) {
        System.out.println("Reserving " + items.size() + " items");
    }
}

class BusinessRuleOrderValidator implements OrderValidator {
    @Override
    public ValidationResult validate(Order order) {
        if (order.getItems().isEmpty()) {
            return ValidationResult.invalid("Order must have at least one item");
        }
        if (order.getTotalAmount() <= 0) {
            return ValidationResult.invalid("Order total must be positive");
        }
        return ValidationResult.valid();
    }
}

// Result classes
class ProcessingResult {
    private boolean success;
    private String message;
    
    private ProcessingResult(boolean success, String message) {
        this.success = success;
        this.message = message;
    }
    
    public static ProcessingResult success(String message) {
        return new ProcessingResult(true, message);
    }
    
    public static ProcessingResult failure(String message) {
        return new ProcessingResult(false, message);
    }
    
    public boolean isSuccess() { return success; }
    public String getMessage() { return message; }
}

class PaymentResult {
    private boolean success;
    private String errorMessage;
    
    private PaymentResult(boolean success, String errorMessage) {
        this.success = success;
        this.errorMessage = errorMessage;
    }
    
    public static PaymentResult success(String message) {
        return new PaymentResult(true, null);
    }
    
    public static PaymentResult failure(String errorMessage) {
        return new PaymentResult(false, errorMessage);
    }
    
    public boolean isSuccess() { return success; }
    public String getErrorMessage() { return errorMessage; }
}

class ValidationResult {
    private boolean valid;
    private String errors;
    
    private ValidationResult(boolean valid, String errors) {
        this.valid = valid;
        this.errors = errors;
    }
    
    public static ValidationResult valid() {
        return new ValidationResult(true, null);
    }
    
    public static ValidationResult invalid(String errors) {
        return new ValidationResult(false, errors);
    }
    
    public boolean isValid() { return valid; }
    public String getErrors() { return errors; }
}

// Application configuration (Composition Root)
class OrderApplicationConfig {
    public OrderProcessor createStandardOrderProcessor() {
        return new StandardOrderProcessor(
            new DatabaseOrderRepository(),
            new CreditCardProcessor(),
            new EmailNotificationSender(),
            new WarehouseInventoryChecker(),
            new BusinessRuleOrderValidator()
        );
    }
    
    public OrderProcessor createExpressOrderProcessor() {
        return new ExpressOrderProcessor(
            new DatabaseOrderRepository(),
            new CreditCardProcessor(),
            new EmailNotificationSender(),
            new WarehouseInventoryChecker(),
            new BusinessRuleOrderValidator()
        );
    }
}

// Usage demonstrating all principles working together
public class SOLIDExample {
    public static void main(String[] args) {
        OrderApplicationConfig config = new OrderApplicationConfig();
        
        // Create order
        List<OrderItem> items = List.of(
            new OrderItem("PRODUCT-1", 2, 25.99),
            new OrderItem("PRODUCT-2", 1, 15.50)
        );
        Order order = new Order("ORDER-123", "CUSTOMER-456", items);
        
        // Process with standard processor - LSP in action
        OrderProcessor standardProcessor = config.createStandardOrderProcessor();
        ProcessingResult result1 = standardProcessor.processOrder(order);
        System.out.println("Standard: " + result1.getMessage());
        
        System.out.println("---");
        
        // Process with express processor - LSP in action
        OrderProcessor expressProcessor = config.createExpressOrderProcessor();
        ProcessingResult result2 = expressProcessor.processOrder(order);
        System.out.println("Express: " + result2.getMessage());
    }
}
```

### How SOLID Principles Work Together

1. **SRP + ISP:** Single responsibility leads to focused interfaces
2. **OCP + DIP:** Open/closed is achieved through dependency inversion
3. **LSP + OCP:** Proper substitution enables extension without modification
4. **DIP + ISP:** Depending on small interfaces reduces coupling
5. **All together:** Create flexible, maintainable, and testable systems

### Benefits of Following All SOLID Principles

- **Maintainability:** Easy to understand and modify
- **Extensibility:** Easy to add new features without breaking existing code
- **Testability:** Easy to unit test individual components
- **Flexibility:** Easy to swap implementations
- **Reusability:** Components can be reused in different contexts
- **Reduced coupling:** Components are loosely coupled
- **High cohesion:** Related functionality is grouped together

---

## When NOT to Use SOLID Principles

While SOLID principles are powerful guidelines, they shouldn't be applied blindly. Sometimes, following them can lead to over-engineering.

### Cases to Be Careful

1. **Simple applications:** For small, one-time scripts or simple applications, SOLID might add unnecessary complexity
2. **Performance-critical code:** Sometimes direct dependencies perform better than abstracted ones
3. **Rapid prototyping:** When building proof-of-concepts, flexibility might be less important than speed
4. **Very stable requirements:** If requirements never change, some flexibility might be unnecessary

### Balance is Key

The goal is not to follow SOLID principles religiously, but to use them to create better software. Consider:
- **Project size and complexity**
- **Team experience**  
- **Long-term maintenance needs**
- **Performance requirements**
- **Time constraints**

---

## What's Next?

Now you have a solid understanding of SOLID principles! These principles form the foundation of good object-oriented design. To deepen your knowledge, consider exploring:

1. **Design Patterns:** Learn common solutions to recurring problems (Singleton, Factory, Observer, etc.)
2. **Clean Architecture:** Understand how to structure entire applications
3. **Domain-Driven Design:** Learn how to model complex business domains
4. **Test-Driven Development:** Write tests that guide your design decisions
5. **Refactoring:** Practice improving existing code using SOLID principles

Remember: SOLID principles are guidelines, not laws. Use your judgment to apply them appropriately for your specific context and requirements. The goal is to write code that is maintainable, flexible, and easy to understand - SOLID principles help you achieve that goal.