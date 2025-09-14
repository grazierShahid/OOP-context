// =============================================
// SOLID PRINCIPLES EXAMPLE - PAYMENT PROCESSING
// =============================================

// 1. SINGLE RESPONSIBILITY PRINCIPLE (SRP)
// Each class should have only one reason to change

// Payment data class - only handles payment information
class Payment {
    private String id;
    private double amount;
    private String currency;
    
    public Payment(String id, double amount, String currency) {
        this.id = id;
        this.amount = amount;
        this.currency = currency;
    }
    
    // Getters only - data class shouldn't modify its own state
    public String getId() { return id; }
    public double getAmount() { return amount; }
    public String getCurrency() { return currency; }
}

// 2. OPEN/CLOSED PRINCIPLE (OCP)
// Open for extension, closed for modification

// Payment processor interface - abstraction for different payment methods
interface PaymentProcessor {
    boolean processPayment(Payment payment);
}

// Credit card processor implementation
class CreditCardProcessor implements PaymentProcessor {
    @Override
    public boolean processPayment(Payment payment) {
        System.out.println("Processing credit card payment: " + payment.getId());
        // Actual credit card processing logic
        return true;
    }
}

// PayPal processor implementation - added without modifying existing code
class PayPalProcessor implements PaymentProcessor {
    @Override
    public boolean processPayment(Payment payment) {
        System.out.println("Processing PayPal payment: " + payment.getId());
        // Actual PayPal processing logic
        return true;
    }
}

// 3. LISKOV SUBSTITUTION PRINCIPLE (LSP)
// Subtypes must be substitutable for their base types

// Refund processor interface
interface RefundProcessor {
    boolean processRefund(Payment payment);
}

// Credit card refund processor - properly implements refund behavior
class CreditCardRefundProcessor implements RefundProcessor {
    @Override
    public boolean processRefund(Payment payment) {
        System.out.println("Processing credit card refund: " + payment.getId());
        // Actual refund logic
        return true;
    }
}

// 4. INTERFACE SEGREGATION PRINCIPLE (ISP)
// Clients shouldn't be forced to depend on interfaces they don't use

// Notification interface - focused only on notification concerns
interface Notifier {
    void sendNotification(String message);
}

// Email notifier implementation
class EmailNotifier implements Notifier {
    @Override
    public void sendNotification(String message) {
        System.out.println("Sending email: " + message);
    }
}

// SMS notifier implementation
class SMSNotifier implements Notifier {
    @Override
    public void sendNotification(String message) {
        System.out.println("Sending SMS: " + message);
    }
}

// 5. DEPENDENCY INVERSION PRINCIPLE (DIP)
// Depend on abstractions, not concretions

// Payment service - high-level module depends on abstractions
class PaymentService {
    private final PaymentProcessor paymentProcessor;
    private final Notifier notifier;
    
    // Constructor injection - dependencies provided externally
    public PaymentService(PaymentProcessor paymentProcessor, Notifier notifier) {
        this.paymentProcessor = paymentProcessor;
        this.notifier = notifier;
    }
    
    public boolean executePayment(Payment payment) {
        boolean success = paymentProcessor.processPayment(payment);
        
        if (success) {
            notifier.sendNotification("Payment successful: " + payment.getId());
        } else {
            notifier.sendNotification("Payment failed: " + payment.getId());
        }
        
        return success;
    }
}

// =============================================
// USAGE EXAMPLE - DEMONSTRATING SOLID PRINCIPLES
// =============================================

public class SOLIDExample {
    public static void main(String[] args) {
        // Create payment
        Payment payment = new Payment("PAY-001", 100.0, "USD");
        
        // Create processors and notifiers (could be injected via DI framework)
        PaymentProcessor creditCardProcessor = new CreditCardProcessor();
        Notifier emailNotifier = new EmailNotifier();
        
        // Create payment service with dependencies
        PaymentService paymentService = new PaymentService(creditCardProcessor, emailNotifier);
        
        // Execute payment - all SOLID principles in action
        boolean result = paymentService.executePayment(payment);
        System.out.println("Payment result: " + result);
        
        // Demonstrate OCP - easily switch to PayPal without changing PaymentService
        PaymentProcessor paypalProcessor = new PayPalProcessor();
        Notifier smsNotifier = new SMSNotifier();
        
        PaymentService paypalService = new PaymentService(paypalProcessor, smsNotifier);
        paypalService.executePayment(payment);
        
        // Demonstrate LSP - refund processors are substitutable
        RefundProcessor refundProcessor = new CreditCardRefundProcessor();
        refundProcessor.processRefund(payment);
    }
}

// =============================================
// ADDITIONAL SOLID COMPLIANT COMPONENTS
// =============================================

// Logger interface - following ISP (focused responsibility)
interface Logger {
    void logInfo(String message);
    void logError(String message);
}

// Database service abstraction - following DIP
interface PaymentRepository {
    void savePayment(Payment payment);
    Payment findPaymentById(String id);
}

// File logger implementation
class FileLogger implements Logger {
    @Override
    public void logInfo(String message) {
        System.out.println("INFO: " + message);
    }
    
    @Override
    public void logError(String message) {
        System.out.println("ERROR: " + message);
    }
}

// Enhanced payment service with logging and repository
class EnhancedPaymentService {
    private final PaymentProcessor paymentProcessor;
    private final Notifier notifier;
    private final Logger logger;
    private final PaymentRepository paymentRepository;
    
    public EnhancedPaymentService(PaymentProcessor paymentProcessor, 
                                 Notifier notifier, 
                                 Logger logger,
                                 PaymentRepository paymentRepository) {
        this.paymentProcessor = paymentProcessor;
        this.notifier = notifier;
        this.logger = logger;
        this.paymentRepository = paymentRepository;
    }
    
    public boolean executePayment(Payment payment) {
        logger.logInfo("Processing payment: " + payment.getId());
        
        boolean success = paymentProcessor.processPayment(payment);
        
        if (success) {
            paymentRepository.savePayment(payment);
            notifier.sendNotification("Payment successful: " + payment.getId());
            logger.logInfo("Payment completed: " + payment.getId());
        } else {
            logger.logError("Payment failed: " + payment.getId());
            notifier.sendNotification("Payment failed: " + payment.getId());
        }
        
        return success;
    }
}

// =============================================
// OUTPUT EXPECTED:
// =============================================
// Processing credit card payment: PAY-001
// Sending email: Payment successful: PAY-001
// Payment result: true
// Processing PayPal payment: PAY-001
// Sending SMS: Payment successful: PAY-001
// Processing credit card refund: PAY-001