# Object Oriented Programming

## Introduction

Before the creation of programming, we planned, built, and executed our business models with pen and paper. There, to make our business system easier to maintain, we designed the business in a way that we could find everything in a little time and do everything easily. Also, we planned in a way that used less paper to handle the business with less context.

After the creation of programming, we could use virtual paper (RAM/ROM) of a computer instead of paper files. Whatever the notebook (paper/computer) is, our main focus is to plan, build, and execute the business first. Now, same as pen and paper, we need to design our business model in a computer system using programming languages. Plus, design the business in a way that we can find everything in a little time and do everything easily. So we need to write programming language code (instead of English/Bengali) in less memory (instead of pages) to handle the business with less context.

The great Object-Oriented Programming helps us to achieve the goal of making a system. This is not a programming language stuff. This is a philosophical way of thinking about a system. Different programming languages capture its diagram in their own way. The core logic of use cases is based on the business domain. The business domain is the actual scenario behind every system logic. Plus, it is the connection between the business team, development team, and more. For this, the business domain has an architecture named Domain Driven Architecture.

Now let's come back to the context! OOP introduced itself by its name. An Object is an instance of a Class. And a Class serves as the template of an object's structure. Wait... don't make an infinite recursion!! A class is a logical design that defines what data (property/attribute) and what behavior (method/function) the object will have. You can think of a class as a blueprint/template/recipe which describes how an object will be made. But a class itself does not exist in memory until we create an object from it (if a class has a static variable/method, then after creating one or multiple objects, only one instance will be created in the class' memory space). When we create an object, it takes memory space and becomes a real thing we can use.

This way, we separate the design (class) and the real-world instance (object). This separation helps us make our system more organized, reusable, and easy to maintain. Also, it saves memory because we store the blueprint one time but can create many objects from it whenever we need. If our business model/logic changes, we only need to change the class, not all objects/instances.

**Definition:** Object-Oriented Programming (OOP) is a programming paradigm that structures software around real-world "objects" rather than just functions and logic. Objects combine data (attributes) and the functions (methods) that operate on that data into a single, self-contained unit.

OOP is based on 4 pillars. Those pillars are class-oriented, where the class is for the object. So the main thing is that this philosophical diagram is object-oriented. Its four pillars are:

1. **Inheritance:** In business, some work or property is common across departments. Instead of redefining them every time, we keep them in one place (parent class) and share with others (child classes).  
2. **Polymorphism:** In business, one action can have different ways to do it based on context. For example, "approve()" can mean leave approval, salary approval, or invoice approval — but the action name stays same.  
3. **Abstraction:** In business, we focus only on what we need to do, not how it is done internally. Like we use an ATM — we only see buttons, not the machine’s wiring.  
4. **Encapsulation:** In business, some information is private (like salary data), some are shared partially, and some are public. We hide internal details and expose only what is needed.

---

## Class & Object

* **Class:**  
A class is the design, blueprint, or recipe for creating objects.  
It defines:  
  - **Data (fields/attributes):** what information an object will hold.  
  - **Behavior (methods/functions):** what actions an object can perform.  
A class itself does not take memory space for its fields (except static fields). It only takes memory when we create objects. You can think of a class as the design of a house — just the drawing, no bricks.  

* **Constructor:**  
A constructor is a special method that runs automatically when an object is created.  
It is used to set initial values of the object’s data.  

* **Destructor:**  
A destructor is a special method that runs when the object is destroyed (memory is freed).

* **Object:**  
An object is the real, usable thing created from a class.  
When we create an object:  
  - It gets its own memory space for fields/attributes.  
  - It can call methods (behaviors) defined in the class.  
Many objects can be created from one class, each having its own data but sharing the same behavior logic.  
You can think of an object as the actual house built from the design — now you can live in it, open its door, use its rooms.

---

### Note on Java (class-based) vs Go (func-based)

Java has **class** keyword and works with class-object model.  
Class has **data and methods together**, so data and behavior stay in one place.  

OOP can also be covered in Go, even though Go is not a class-based language.  
Go has **struct** (holds data) and allows binding methods using **receiver functions**.  
Go also has **interface**, which defines method sets — any struct that implements those methods is said to implement the interface (no explicit `implements` keyword required).  
This combination of **struct + receiver method + interface** allows Go to follow OOP principles like polymorphism and abstraction.

---

### Example in Java

```java
class Employee {
    String name;
    double salary;

    Employee(String name, double salary) { /* constructor code */ }

    void showDetails() { /* method code */ }
}

public class Main {
    public static void main(String[] args) {
        Employee e1 = new Employee("Abul", 50000);
        e1.showDetails();
    }
}

```

### Example in Go
```go
type Employee struct {
    Name   string
    Salary float64
}

func (e Employee) ShowDetails() { /* receiver method code */ }

type DetailsPrinter interface {
    ShowDetails()
}

func NewEmployee(name string, salary float64) Employee {
    return Employee{Name: name, Salary: salary}
}

func main() {
    e1 := NewEmployee("Abul", 50000)
    var printer DetailsPrinter = e1
    printer.ShowDetails()
}
```
### Access Modifier

**Definition:** Access modifiers control visibility and accessibility of class members (where they can be used).

Class data and methods can be set to control from where they can be used. In Java, there are 3 main types of access modifiers:  
1. **Public:** accessible from anywhere -> `public String name;`  
2. **Protected:** accessible from the same package and subclasses -> `protected String role;`  
3. **Private:** accessible only from within the class itself -> `private double salary;`  

In Go, there are two ways to express data accessibility:  
1. **Exported (Public):** accessible from anywhere (including other packages) if the name starts with a capital letter -> `Name int`  
2. **Unexported (private)** accessible only inside the package if the name starts with a lowercase letter -> `salary int`  

* In Java, there is a keyword named `static`. It is used for data or methods that belong to the class itself rather than any object. Static members are stored once in memory, shared across all objects of the class.

#### Example in Java
```java
class Employee {
    public String name;       // accessible anywhere
    protected String role;    // same package + subclasses
    private double salary;    // only inside class

    static int companyCode;   // shared across all objects

    void show() {
        System.out.println(name + ", " + role);
    }
}
```

#### Example in Go
```go
var CompanyCode int // shared across all structs

type Employee struct {
    Name   string // public
    role   string // private
    salary int    // private
}
```

## Inheritance

**Definition:** Inheritance is a core concept where a new class (the subclass or child class) derives properties (attributes) and behaviors (methods) from an existing class (the superclass or parent class).

In business domain logic, we often find different models (classes) that share common properties and behaviors. For example, different types of vehicles (car, bike, truck, microbus) have some features in common, and some that are different. Instead of repeating the same code, we can create a base class called `Vehicle` and then create `Car`, `Bike`, `Truck`, etc. by inheriting from it. Example in Java:
```java
class Vehicle {
    // common data and methods
}
class Bike extends Vehicle {
    // extra data and methods for Bike
}
```

In OOP, this is called inheritance. A class can inherit data and behavior from a parent (super) class and reuse them as defined. Here, the parent class is `Vehicle` and the child class is `Bike`. The parent class is like a father, telling his son (child class): "Take my properties and make what you want." After this, `Bike` is a `Vehicle` with more data and methods. This is called an `is-a` relationship: a bike is a vehicle.

### Types of Inheritance
Depending on the language, there are different types of inheritance:
1. **Single Inheritance:** One class inherits from one parent class. Java example:
    ```java
        class A extends B {}
    ```
2. **Multilevel Inheritance:** A class inherits from a class, which itself inherits from another class. Java example:
    ```java
        class A {}
        class B extends A {}
        class C extends B {}
    ```
3. **Hierarchical Inheritance:** Multiple classes inherit from the same parent class. Java example: 
    ```java
        class A {}
        class B extends A {}
        class C extends A {}
    ```
4. **Multiple Inheritance:** One class inherits from more than one parent class.  
   C++ supports this, but Java and Go do not (Java uses interfaces instead). C++ example:
   ```cpp
        class A {/*data and methods*/};
        class B {/*data and methods*/};
        class C : public A, public B {/*data and methods*/};
   ```
5. **Hybrid Inheritance:** Combination of two or more types above.  
   (Supported in C++, not directly in Java). C++ example:
   ```cpp
        class A {/*data and methods*/};
        class B : public A {/*data and methods*/};
        class C {/*data and methods*/};
        class D : public B, public C {/*data and methods*/};

   ```

### The Diamond Problem
The diamond problem happens in multiple inheritance. Suppose class `A` has a method `abc()`. Classes `B` and `C` both inherit from `A`, so they both have `abc()`. Now, class `D` inherits from both `B` and `C`. If you create an object of `D` and call `abc()`, which version should run? This is a conflict.

#### Example in C++ (with diamond problem):
```cpp
class A {
public:
    void abc() { /* ... */ }
};
class B : public A {};
class C : public A {};
class D : public B, public C {};

int main() {
    D obj;
    obj.abc(); // Error: which abc()?
}
```

To solve this, C++ provides virtual inheritance:
```cpp
class A {
public:
    void abc() { /* ... */ }
};
class B : virtual public A {};
class C : virtual public A {};
class D : public B, public C {};

int main() {
    D obj;
    obj.abc(); // Now works, only one copy of A
}
```

#### Example in Java (no multiple inheritance for classes):
Java does not allow multiple inheritance for classes, so the diamond problem does not occur. Instead, Java uses interfaces to achieve similar behavior:
```java
interface A {
    void abc();
}
class B implements A {
    public void abc() { /* ... */ }
}
class C implements A {
    public void abc() { /* ... */ }
}
class D extends B implements A {
    // If needed, override abc() or use super.abc()
}
```
If both B and C have their own implementation, you must resolve which one to use in D.

### Composition (Has-A Relationship)

**Definition:** Composition is a design principle where a class is built by combining instances of other, simpler classes, establishing a "has-a" relationship.

There is another way to reuse code: composition. Creating an object of another class inside a class is called composition. By creating an object, the class gets all properties and methods from the object's class. This is a `has-a` relationship. Like, there are one more class named `Vehicle` and  `Engine`. Now, creating another class `class Car extends Vehicle{Engine engine = new Engine}`

#### Example in Java (composition):
```java
class Vehicle {
    String type;
    Vehicle(String type) { this.type = type; }
    void showType() { System.out.println("Vehicle type: " + type); }
}

class Engine {
    int horsepower;
    Engine(int horsepower) { this.horsepower = horsepower; }
    void start() { System.out.println("Engine started with " + horsepower + " HP"); }
}

class Car extends Vehicle {
    Engine engine; // Car HAS an Engine

    Car(String type, int horsepower) {
        super(type);
        engine = new Engine(horsepower);
    }

    void drive() {
        engine.start();
        System.out.println("Car is driving...");
    }
}

public class Main {
    public static void main(String[] args) {
        Car car = new Car("Sedan", 120);
        car.showType(); // from Vehicle
        car.drive();    // uses Engine
    }
}

```

#### Example in Go (composition):
```go
package main

import "fmt"

type Vehicle struct {
	Type string
}

func (v Vehicle) ShowType() {
	fmt.Println("Vehicle type:", v.Type)
}

type Engine struct {
	Horsepower int
}

func (e Engine) Start() {
	fmt.Println("Engine started with", e.Horsepower, "HP")
}

type Car struct {
	Vehicle // Car HAS a Vehicle (embedded)
	Engine  // Car HAS an Engine (embedded)
}

func (c Car) Drive() {
	c.Engine.Start()
	fmt.Println("Car is driving...")
}

func main() {
	car := Car{
		Vehicle: Vehicle{Type: "Sedan"},
		Engine:  Engine{Horsepower: 120},
	}
	car.ShowType() // from Vehicle
	car.Drive()    // uses Engine
}

```

## Polymorpism

**Definition:** Polymorphism is the ability of an object, method, or function to take on multiple forms or behave differently in different situations, allowing a single interface to represent different behaviors.

Now we know about class, object, and inheritance. Next comes one of the most powerful concepts in OOP: polymorphism. Let's break down the word: polymorphism = poly (many) + morph (form). So, polymorphism means "many forms." In programming, it means one thing (like a method or object) can take many forms or do different things based on context.

### Why, When, and How to Use Polymorphism
Polymorphism helps us write flexible and reusable code. In business, sometimes we have the same action, but it works differently for different departments or objects. For example, the action `approve()` can mean leave approval, salary approval, or invoice approval. The action name is the same, but the logic is different. Polymorphism lets us use the same interface or method name, but with different implementations.

**When to use:**
- When you want to perform the same action in different ways for different objects.
- When you want to write code that works with superclasses or interfaces, but can handle any subclass or implementation.

**Why use:**
- Makes code more flexible and maintainable.
- Reduces code duplication.
- Makes it easy to add new features (just add a new subclass or method).

**How to use:**
- By using method overloading (same method name, different parameters) or method overriding (subclass provides its own version of a method).
- By using interfaces or abstract classes.

---

### Types of Polymorphism
There are two main types:
1. **Compile-time (Static) Polymorphism:** Decided at compile time. Achieved by method overloading.
2. **Runtime (Dynamic) Polymorphism:** Decided at runtime. Achieved by method overriding.

---

### Method Overloading (Compile-time Polymorphism)

**Definition:** Method overloading in OOP is a feature that allows multiple methods within the same class to share the same name but have different parameter lists.

Method overloading means having multiple methods with the same name but different parameter lists (type, number, or order) in the same class (method signatures are different). The compiler decides which method to call based on the arguments.

**Example in Java:**
```java
class Calculator {
    int add(int a, int b) {
        return a + b;
    }
    double add(double a, double b) {
        return a + b;
    }
    int add(int a, int b, int c) {
        return a + b + c;
    }
}

public class Main {
    public static void main(String[] args) {
        Calculator calc = new Calculator();
        System.out.println(calc.add(2, 3));        // calls int add(int, int)
        System.out.println(calc.add(2.5, 3.5));    // calls double add(double, double)
        System.out.println(calc.add(1, 2, 3));     // calls int add(int, int, int)
    }
}
```

---

### Method Overriding (Runtime Polymorphism)

**Definition:** Method overriding in Object-Oriented Programming (OOP) is a mechanism that allows a subclass (child class) to provide a specific implementation for a method that is already defined in its superclass (parent class).

Method overriding means a subclass provides its own implementation of a method that is already defined in its parent class. The method signature must be the same. At runtime, programming languages like Java decides which method to call based on the object's actual type.

**Example in Java:**
```java
class Animal {
    void sound() {
        System.out.println("Animal makes a sound");
    }
}
class Dog extends Animal {
    void sound() {
        System.out.println("Dog barks");
    }
}
class Cat extends Animal {
    void sound() {
        System.out.println("Cat meows");
    }
}

public class Main {
    public static void main(String[] args) {
        Animal a1 = new Dog();
        Animal a2 = new Cat();
        a1.sound(); // Dog barks
        a2.sound(); // Cat meows
    }
}
```

---

### Compile-time vs Runtime Polymorphism
- **Compile-time (Static):** Method to call is decided at compile time (method overloading).
- **Runtime (Dynamic):** Method to call is decided at runtime (method overriding, interfaces, abstract classes).

**Example:**
```java
// Compile-time
class Demo {
    void show(int a) { System.out.println(a); }
    void show(String s) { System.out.println(s); }
}
// Runtime
class Parent {
    void display() { System.out.println("Parent"); }
}
class Child extends Parent {
    void display() { System.out.println("Child"); }
}

public class Main {
    public static void main(String[] args) {
        Demo d = new Demo();
        d.show(10);         // compile-time
        d.show("hello");   // compile-time

        Parent p = new Child();
        p.display();        // runtime
    }
}
```

---

Polymorphism is a key reason why OOP is so powerful. It lets us write code that is easy to extend and maintain, and that can handle new requirements with minimal changes.

## Abstraction

**Definition:** Abstraction is the process of hiding complex internal details and showing only the essential, relevant features of an object or system to the user

In many cases we have the same method in many subclasses, but each subclass has its own logic. For example, in a `Vehicle` class we can declare a `gear()` method, which is overridden by `Car`, `Bike`, `Truck` etc. Each of these subclasses has its own implementation. In the `Vehicle` class we do not need to write the implementation; we can write the method signature only (instead of a concrete implementation). This way we know that anyone who inherits `Vehicle` will have this method, but we don't know or care about the internal implementation — this is abstraction (hiding implementation).

You can apply abstraction partially or fully (0–100%). In Java you have two main tools:

- Abstract class (partial abstraction): use `abstract class ClassName {}` and `abstract` methods. An abstract class can have fields, constructors, concrete (implemented) methods, and abstract methods. It cannot be instantiated. Use abstract classes when subclasses share common state or behavior and you want to provide some reusable code.

Java example (abstract class):
```java
    abstract class Vehicle {
        String name;
        Vehicle(String n){ name = n; }
        abstract void gear();           // abstract method
        void info(){ System.out.println(name); } // concrete method
    }
    class Car extends Vehicle {
        Car(String name){ super(name); }
        void gear(){ System.out.println("Car gear logic"); }
    }
```

- Interface (full/pure abstraction contract): an `interface` declares a set of methods (a contract). In older Java versions interfaces were pure abstraction; since Java 8 interfaces can also have `default` and `static` methods. Use interfaces when you need a contract that many unrelated classes can implement, or when you want multiple-type inheritance of behavior.

Java example (interface):
```java
    interface Drivable { void gear(); }
    class Bike implements Drivable { public void gear(){ System.out.println("Bike gear"); } }
```

Key differences and guidance:
- Abstract class: can have state (fields), constructors, implemented methods; a class can extend only one abstract class. Use when types share code or state.
- Interface: primarily a contract; allows multiple interfaces to be implemented by a class. Use when you need loose coupling and multiple-type capability.
- Since Java 8, interfaces may have `default` methods (small shared behavior) and `static` helpers.

Abstraction in Go:
- Go has no classes, but interfaces provide abstraction. Define a method set in an interface and any type implementing those methods satisfies the interface (implicit).

Go example:
```go
package main
import "fmt"

type Gearer interface { Gear() }

type Car struct {}
func (Car) Gear(){ fmt.Println("Car gear") }

func drive(g Gearer){ g.Gear() }

func main(){ drive(Car{}) }
```

When to choose which:
- Choose abstract class when you want to share implementation and state among related classes.
- Choose interface when you want to define behavior contracts, support multiple-type behavior, and keep implementations independent.

## Encapsulation

**Definition:** Encapsulation is the concept of bundling data (attributes) and the methods that operate on that data into a single unit, like a class, and restricting direct access to the internal data to protect its integrity.

Encapsulation helps us keep important data secret and expose only what is necessary. We use access modifiers to control visibility. In Java we use `private`, `protected`, `public`; in Go we use exported (capitalized) vs unexported (lowercase) names.

In practice we hide fields and provide controlled access via methods (getters/setters or other behavior methods). This prevents invalid states and centralizes validation and side effects.

Java example:
```java
class Account{
    private double balance;
    public Account(double b){ this.balance = b; }
    public double getBalance(){ return balance; }
    public void deposit(double v){ if(v>0) balance += v; }
}
```

Go example:
```go
type Account struct{ balance float64 }
func NewAccount(b float64) *Account{ return &Account{balance:b} }
func (a *Account) Balance() float64{ return a.balance }
func (a *Account) Deposit(v float64){ if v>0 { a.balance += v } }
```

Encapsulation vs Abstraction: Encapsulation is a technique (hide data, expose behavior). Abstraction is a design idea (expose only the essential features). Use encapsulation to implement abstraction.

Benefits:
- Data integrity: prevent invalid changes.
- Controlled access: central place for validation and logging.
- Maintainability: change internals without breaking users.
- Security: hide sensitive fields from direct access.

## What's Next?

Now we have covered everything from OOP. Next, let’s move to FAQ.md to get familiar with common interview questions, real-life scenarios, and understand the differences between various OOP concepts.