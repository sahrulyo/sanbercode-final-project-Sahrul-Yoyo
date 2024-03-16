package service

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// Product represents a product in the system
type Product struct {
	ID       uint
	Name     string
	Quantity uint
	// Add other fields as needed
}

// Order represents an order in the system
type Order struct {
	ID     uint
	UserID uint
	// Add other fields as needed
}

// OrderDetail represents details of an order in the system
type OrderDetail struct {
	ID        uint
	OrderID   uint
	ProductID uint
	Quantity  uint
	// Add other fields as needed
}

// CreateService creates a new order for a user
func CreateService(userID uint, productIDs []uint, db *gorm.DB) error {
	// Check if user exists (omitted for brevity)

	// Start a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Loop through productIDs
	for _, productID := range productIDs {
		// Check product availability and quantity
		var product Product
		if err := tx.First(&product, productID).Error; err != nil {
			tx.Rollback()
			return errors.New("product not found")
		}
		if product.Quantity == 0 {
			tx.Rollback()
			return errors.New("product is out of stock")
		}

		// Create order detail
		orderDetail := OrderDetail{
			ProductID: productID,
			Quantity:  1, // Assuming each product quantity in the order is 1 for simplicity
			// You may also want to retrieve the price from the product and calculate total price here
		}
		if err := tx.Create(&orderDetail).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Deduct product quantity from inventory
		if err := tx.Model(&product).Update("quantity", product.Quantity-1).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Create order
	order := Order{
		UserID: userID,
		// Other fields of the order can be added here
	}
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	// Send order confirmation email
	if err := sendOrderConfirmationEmail(userID); err != nil {
		// Log error but do not fail the order creation process
		fmt.Println("Error sending order confirmation email:", err)
	}

	return nil
}

// sendOrderConfirmationEmail simulates sending an order confirmation email
func sendOrderConfirmationEmail(userID uint) error {
	// Simulated implementation, replace with actual email sending logic
	fmt.Printf("Order confirmation email sent to user %d\n", userID)
	return nil
}
