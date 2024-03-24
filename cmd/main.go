package main

import (
	"fixy-finance-api/internal/controllers"
	"fixy-finance-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	middlewares.SetCorsToAcceptAll(r)

	expenseController := &controllers.ExpenseController{}

	// Register routes
	r.GET("/getAllExpenses", expenseController.GetAllExpenses)
	r.POST("/addNewExpense", expenseController.AddNewExpense)

	r.Run(":8080")
}
