package utils

import "github.com/parrothacker1/expense-tracker/models"

func AddExpense(expense *models.Expense) (err error) {
	err = models.DB.Create(expense).Error
	return
}

func ListExpenses() ([]models.Expense, error) {
	var expenses []models.Expense
	result := models.DB.Find(&expenses)
	return expenses, result.Error
}

func DeleteExpense(id uint) error {
	result := models.DB.Delete(&models.Expense{}, id)
	return result.Error
}

func GetExpenseByID(id uint) (models.Expense, error) {
	var expense models.Expense
	result := models.DB.First(&expense, id)
	return expense, result.Error
}
