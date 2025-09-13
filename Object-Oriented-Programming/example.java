// Complete Object-Oriented Programming Examples in Java
// Covering: Classes, Objects, Inheritance, Polymorphism, Abstraction, Encapsulation

// ============================================================================
// 1. ENCAPSULATION - Data hiding and controlled access
// ============================================================================

class BankAccount {
    // Private fields - encapsulated data
    private String accountNumber;
    private double balance;
    private String ownerName;
    
    // Static field - belongs to class, shared across all instances
    private static int totalAccounts = 0;
    
    // Constructor
    public BankAccount(String accountNumber, String ownerName, double initialBalance) {
        this.accountNumber = accountNumber;
        this.ownerName = ownerName;
        this.balance = initialBalance >= 0 ? initialBalance : 0;
        totalAccounts++;
    }
    
    // Getter methods (controlled access)
    public String getAccountNumber() { return accountNumber; }
    public String getOwnerName() { return ownerName; }
    public double getBalance() { return balance; }
    public static int getTotalAccounts() { return totalAccounts; }
    
    // Controlled behavior methods
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
    
    public void displayInfo() {
        System.out.println("Account: " + accountNumber + ", Owner: " + ownerName + ", Balance: $" + balance);
    }
}

// ============================================================================
// 2. ABSTRACTION - Abstract classes and interfaces
// ============================================================================

// Abstract class (partial abstraction)
abstract class Vehicle {
    protected String brand;
    protected String model;
    protected int year;
    
    // Constructor
    public Vehicle(String brand, String model, int year) {
        this.brand = brand;
        this.model = model;
        this.year = year;
    }
    
    // Concrete method (implemented)
    public void displayBasicInfo() {
        System.out.println(year + " " + brand + " " + model);
    }
    
    // Abstract methods (must be implemented by subclasses)
    public abstract void start();
    public abstract void stop();
    public abstract double calculateFuelEfficiency();
}

// Interface (pure abstraction/contract)
interface Drivable {
    void accelerate();
    void brake();
    void turn(String direction);
}

interface Maintainable {
    void performMaintenance();
    boolean needsMaintenance();
}

// ============================================================================
// 3. INHERITANCE - Single, Multilevel, Hierarchical
// ============================================================================

// Single inheritance: Car extends Vehicle
class Car extends Vehicle implements Drivable, Maintainable {
    private int doors;
    private String fuelType;
    private double mileage;
    
    public Car(String brand, String model, int year, int doors, String fuelType) {
        super(brand, model, year); // Call parent constructor
        this.doors = doors;
        this.fuelType = fuelType;
        this.mileage = 0;
    }
    
    // Implementing abstract methods from Vehicle
    @Override
    public void start() {
        System.out.println("Car engine started with key ignition");
    }
    
    @Override
    public void stop() {
        System.out.println("Car engine stopped");
    }
    
    @Override
    public double calculateFuelEfficiency() {
        return mileage > 0 ? mileage / 100 : 0; // km per liter
    }
    
    // Implementing Drivable interface
    @Override
    public void accelerate() {
        System.out.println("Car is accelerating smoothly");
    }
    
    @Override
    public void brake() {
        System.out.println("Car brakes applied");
    }
    
    @Override
    public void turn(String direction) {
        System.out.println("Car turning " + direction);
    }
    
    // Implementing Maintainable interface
    @Override
    public void performMaintenance() {
        System.out.println("Performing car maintenance: oil change, tire check");
    }
    
    @Override
    public boolean needsMaintenance() {
        return mileage > 10000; // needs maintenance after 10000 km
    }
    
    public void addMileage(double miles) {
        this.mileage += miles;
    }
}

// Hierarchical inheritance: Multiple classes inherit from Vehicle
class Motorcycle extends Vehicle implements Drivable {
    private int engineCC;
    
    public Motorcycle(String brand, String model, int year, int engineCC) {
        super(brand, model, year);
        this.engineCC = engineCC;
    }
    
    @Override
    public void start() {
        System.out.println("Motorcycle started with kick/electric start");
    }
    
    @Override
    public void stop() {
        System.out.println("Motorcycle engine stopped");
    }
    
    @Override
    public double calculateFuelEfficiency() {
        return engineCC < 200 ? 40 : 25; // approximate km/l based on engine size
    }
    
    @Override
    public void accelerate() {
        System.out.println("Motorcycle accelerating quickly");
    }
    
    @Override
    public void brake() {
        System.out.println("Motorcycle brakes applied");
    }
    
    @Override
    public void turn(String direction) {
        System.out.println("Motorcycle leaning " + direction + " to turn");
    }
}

// Multilevel inheritance: SportsCar extends Car
class SportsCar extends Car {
    private int horsepower;
    private boolean turbocharged;
    
    public SportsCar(String brand, String model, int year, int doors, int horsepower, boolean turbocharged) {
        super(brand, model, year, doors, "Petrol");
        this.horsepower = horsepower;
        this.turbocharged = turbocharged;
    }
    
    // Method overriding (Runtime polymorphism)
    @Override
    public void start() {
        System.out.println("Sports car engine roars to life!");
    }
    
    @Override
    public void accelerate() {
        System.out.println("Sports car accelerating with " + horsepower + " HP!");
    }
    
    // Method overriding with different behavior
    @Override
    public double calculateFuelEfficiency() {
        return turbocharged ? 8 : 12; // sports cars are less fuel efficient
    }
    
    // Additional method specific to sports car
    public void activateSportMode() {
        System.out.println("Sport mode activated! Maximum performance!");
    }
}

// ============================================================================
// 4. COMPOSITION - Has-A Relationship
// ============================================================================

class Engine {
    private String type;
    private int horsepower;
    private boolean isRunning;
    
    public Engine(String type, int horsepower) {
        this.type = type;
        this.horsepower = horsepower;
        this.isRunning = false;
    }
    
    public void start() {
        isRunning = true;
        System.out.println(type + " engine started (" + horsepower + " HP)");
    }
    
    public void stop() {
        isRunning = false;
        System.out.println(type + " engine stopped");
    }
    
    public boolean isRunning() { return isRunning; }
    public int getHorsepower() { return horsepower; }
}

class GPS {
    private String currentLocation;
    
    public GPS() {
        this.currentLocation = "Unknown";
    }
    
    public void updateLocation(String location) {
        this.currentLocation = location;
        System.out.println("GPS updated: Current location - " + location);
    }
    
    public String getCurrentLocation() { return currentLocation; }
    
    public void navigate(String destination) {
        System.out.println("GPS navigating from " + currentLocation + " to " + destination);
    }
}

// Advanced car with composition
class AdvancedCar extends Vehicle {
    private Engine engine;  // Car HAS an Engine
    private GPS gps;       // Car HAS a GPS
    private String[] features;
    
    public AdvancedCar(String brand, String model, int year, Engine engine) {
        super(brand, model, year);
        this.engine = engine;
        this.gps = new GPS();
        this.features = new String[]{"Air Conditioning", "Power Steering", "ABS"};
    }
    
    @Override
    public void start() {
        engine.start();
        gps.updateLocation("Starting point");
        System.out.println("Advanced car is ready to drive");
    }
    
    @Override
    public void stop() {
        engine.stop();
        System.out.println("Advanced car stopped");
    }
    
    @Override
    public double calculateFuelEfficiency() {
        return 15; // km/l
    }
    
    public void navigate(String destination) {
        if (engine.isRunning()) {
            gps.navigate(destination);
        } else {
            System.out.println("Please start the car first");
        }
    }
    
    public void showFeatures() {
        System.out.print("Car features: ");
        for (String feature : features) {
            System.out.print(feature + " ");
        }
        System.out.println();
    }
}

// ============================================================================
// 5. POLYMORPHISM - Method Overloading and Overriding
// ============================================================================

class Calculator {
    // Method overloading (Compile-time polymorphism)
    public int add(int a, int b) {
        return a + b;
    }
    
    public double add(double a, double b) {
        return a + b;
    }
    
    public int add(int a, int b, int c) {
        return a + b + c;
    }
    
    public String add(String a, String b) {
        return a + " " + b;
    }
}

// Polymorphism demonstration with vehicles
class VehicleManager {
    // Method that works with any Vehicle (polymorphism)
    public void testVehicle(Vehicle vehicle) {
        vehicle.displayBasicInfo();
        vehicle.start();
        System.out.println("Fuel efficiency: " + vehicle.calculateFuelEfficiency() + " km/l");
        vehicle.stop();
        System.out.println();
    }
    
    // Method that works with any Drivable object
    public void testDriving(Drivable vehicle) {
        vehicle.accelerate();
        vehicle.turn("left");
        vehicle.brake();
        System.out.println();
    }
}

// ============================================================================
// 6. MAIN CLASS - Demonstration of all concepts
// ============================================================================

public class OOPExample {
    public static void main(String[] args) {
        System.out.println("=== Object-Oriented Programming Demo ===\n");
        
        // ========== ENCAPSULATION DEMO ==========
        System.out.println("1. ENCAPSULATION DEMO:");
        BankAccount account1 = new BankAccount("ACC001", "Alice Johnson", 1000);
        BankAccount account2 = new BankAccount("ACC002", "Bob Smith", 500);
        
        account1.displayInfo();
        account1.deposit(200);
        account1.withdraw(150);
        account1.displayInfo();
        System.out.println("Total accounts created: " + BankAccount.getTotalAccounts());
        System.out.println();
        
        // ========== INHERITANCE & POLYMORPHISM DEMO ==========
        System.out.println("2. INHERITANCE & POLYMORPHISM DEMO:");
        VehicleManager manager = new VehicleManager();
        
        // Creating different types of vehicles
        Car sedan = new Car("Toyota", "Camry", 2023, 4, "Petrol");
        SportsCar ferrari = new SportsCar("Ferrari", "F8", 2023, 2, 710, true);
        Motorcycle bike = new Motorcycle("Honda", "CBR600RR", 2023, 600);
        
        // Polymorphism - same method, different behaviors
        manager.testVehicle(sedan);     // Calls Car's implementation
        manager.testVehicle(ferrari);   // Calls SportsCar's implementation  
        manager.testVehicle(bike);      // Calls Motorcycle's implementation
        
        // Interface polymorphism
        System.out.println("3. INTERFACE POLYMORPHISM:");
        manager.testDriving(sedan);
        manager.testDriving(ferrari);
        manager.testDriving(bike);
        
        // ========== METHOD OVERLOADING DEMO ==========
        System.out.println("4. METHOD OVERLOADING DEMO:");
        Calculator calc = new Calculator();
        System.out.println("add(5, 3) = " + calc.add(5, 3));
        System.out.println("add(5.5, 3.2) = " + calc.add(5.5, 3.2));
        System.out.println("add(1, 2, 3) = " + calc.add(1, 2, 3));
        System.out.println("add(\"Hello\", \"World\") = " + calc.add("Hello", "World"));
        System.out.println();
        
        // ========== COMPOSITION DEMO ==========
        System.out.println("5. COMPOSITION DEMO:");
        Engine v8Engine = new Engine("V8", 450);
        AdvancedCar luxuryCar = new AdvancedCar("Mercedes", "S-Class", 2023, v8Engine);
        
        luxuryCar.displayBasicInfo();
        luxuryCar.showFeatures();
        luxuryCar.start();
        luxuryCar.navigate("Downtown Mall");
        luxuryCar.stop();
        System.out.println();
        
        // ========== SPORTS CAR SPECIFIC FEATURES ==========
        System.out.println("6. SPORTS CAR SPECIFIC FEATURES:");
        ferrari.displayBasicInfo();
        ferrari.start();
        ferrari.accelerate();
        ferrari.activateSportMode();
        ferrari.addMileage(5000);
        System.out.println("Needs maintenance: " + ferrari.needsMaintenance());
        System.out.println();
        
        // ========== RUNTIME POLYMORPHISM DEMO ==========
        System.out.println("7. RUNTIME POLYMORPHISM DEMO:");
        Vehicle[] vehicles = {sedan, ferrari, bike};
        
        for (Vehicle v : vehicles) {
            System.out.print("Vehicle: ");
            v.displayBasicInfo();
            v.start(); // Different implementation called based on actual object type
            System.out.println();
        }
        
        System.out.println("=== Demo Complete ===");
    }
}