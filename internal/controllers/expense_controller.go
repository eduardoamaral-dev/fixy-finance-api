package controllers

import (
	. "fixy-finance-api/internal/models"
	. "fixy-finance-api/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExpenseController struct {
	expenseService ExpenseService
}

func (controller *ExpenseController) GetAllExpenses(context *gin.Context) {
	expenses, err := controller.expenseService.GetAllExpenses()

	if err != nil {
		println(err)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": expenses,
	})
}

func (controller *ExpenseController) AddNewExpense(context *gin.Context) {
	var requestBody NewExpenseRequest

	// Bind request body to struct
	if err := context.BindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entryResult, err := controller.expenseService.SaveNewExpense(requestBody)

	if err != nil {
		fmt.Println(err)
	}

	context.JSON(http.StatusOK, gin.H{
		"entryResult": entryResult,
	})
}
