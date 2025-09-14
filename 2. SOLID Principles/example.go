package main

import "fmt"

// 1. SRP: Payment struct only handles payment data
type Payment struct {
	id       string
	amount   float64
	currency string
}

// Constructor function for Payment
func NewPayment(id string, amount float64, currency string) *Payment {
	return &Payment{
		id:       id,
		amount:   amount,
		currency: currency,
	}
}

// 2. OCP: PaymentProcessor interface allows for extension
type PaymentProcessor interface {
	ProcessPayment(payment *Payment) bool
}

// Credit card processor implementation
type CreditCardProcessor struct{}

func (c *CreditCardProcessor) ProcessPayment(payment *Payment) bool {
	fmt.Printf("Processing credit card payment: %s\n", payment.id)
	return true
}

// PayPal processor implementation
type PayPalProcessor struct{}

func (p *PayPalProcessor) ProcessPayment(payment *Payment) bool {
	fmt.Printf("Processing PayPal payment: %s\n", payment.id)
	return true
}

// 3. LSP: RefundProcessor interface
type RefundProcessor interface {
	ProcessRefund(payment *Payment) bool
}

type CreditCardRefundProcessor struct{}

func (c *CreditCardRefundProcessor) ProcessRefund(payment *Payment) bool {
	fmt.Printf("Processing credit card refund: %s\n", payment.id)
	return true
}

// 4. ISP: Separate interfaces for different responsibilities
type Notifier interface {
	SendNotification(message string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) SendNotification(message string) {
	fmt.Printf("Sending email: %s\n", message)
}

type SMSNotifier struct{}

func (s *SMSNotifier) SendNotification(message string) {
	fmt.Printf("Sending SMS: %s\n", message)
}

// Logger interface - following ISP
type Logger interface {
	LogInfo(message string)
	LogError(message string)
}

type FileLogger struct{}

func (f *FileLogger) LogInfo(message string) {
	fmt.Printf("INFO: %s\n", message)
}

func (f *FileLogger) LogError(message string) {
	fmt.Printf("ERROR: %s\n", message)
}

// Repository interface - following DIP
type PaymentRepository interface {
	SavePayment(payment *Payment)
	FindPaymentByID(id string) *Payment
}

// 5. DIP: PaymentService depends on abstractions
type PaymentService struct {
	processor  PaymentProcessor
	notifier   Notifier
	logger     Logger
	repository PaymentRepository
}

// Constructor function for PaymentService
func NewPaymentService(processor PaymentProcessor, notifier Notifier) *PaymentService {
	return &PaymentService{
		processor: processor,
		notifier:  notifier,
	}
}

// ExecutePayment method
func (s *PaymentService) ExecutePayment(payment *Payment) bool {
	success := s.processor.ProcessPayment(payment)

	if success {
		s.notifier.SendNotification("Payment successful: " + payment.id)
	} else {
		s.notifier.SendNotification("Payment failed: " + payment.id)
	}

	return success
}

// Enhanced PaymentService with logging and repository
type EnhancedPaymentService struct {
	PaymentService
	logger     Logger
	repository PaymentRepository
}

func NewEnhancedPaymentService(
	processor PaymentProcessor,
	notifier Notifier,
	logger Logger,
	repository PaymentRepository,
) *EnhancedPaymentService {
	return &EnhancedPaymentService{
		PaymentService: PaymentService{processor: processor, notifier: notifier},
		logger:         logger,
		repository:     repository,
	}
}

func (s *EnhancedPaymentService) ExecutePayment(payment *Payment) bool {
	s.logger.LogInfo("Processing payment: " + payment.id)

	success := s.processor.ProcessPayment(payment)

	if success {
		s.repository.SavePayment(payment)
		s.notifier.SendNotification("Payment successful: " + payment.id)
		s.logger.LogInfo("Payment completed: " + payment.id)
	} else {
		s.logger.LogError("Payment failed: " + payment.id)
		s.notifier.SendNotification("Payment failed: " + payment.id)
	}

	return success
}

func main() {
	// Create payment
	payment := NewPayment("PAY-001", 100.0, "USD")

	// Create processors and notifiers
	creditCardProcessor := &CreditCardProcessor{}
	emailNotifier := &EmailNotifier{}

	// Create payment service
	paymentService := NewPaymentService(creditCardProcessor, emailNotifier)

	// Execute payment
	result := paymentService.ExecutePayment(payment)
	fmt.Printf("Payment result: %v\n", result)

	// Demonstrate OCP - switch to PayPal
	paypalProcessor := &PayPalProcessor{}
	smsNotifier := &SMSNotifier{}

	paypalService := NewPaymentService(paypalProcessor, smsNotifier)
	paypalService.ExecutePayment(payment)

	// Demonstrate LSP
	refundProcessor := &CreditCardRefundProcessor{}
	refundProcessor.ProcessRefund(payment)
}
