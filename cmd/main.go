package main

import (
	"fixy-finance-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	expenseController := &controllers.ExpenseController{}

	// Register routes
	r.GET("/getAllExpenses", expenseController.GetAllExpenses)
	r.POST("/addNewExpense", expenseController.AddNewExpense)

	r.Run(":8080")
}
