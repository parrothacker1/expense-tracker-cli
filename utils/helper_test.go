package utils

import (
	"testing"
	"time"

	"github.com/parrothacker1/expense-tracker/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) {
	err := models.InitDB("file::memory:")
	require.NoError(t, err, "Failed to initialize test database")
}

func TestAddExpense(t *testing.T) {
	setupTestDB(t)
	expense := &models.Expense{
		Amount:   100.50,
		Date:     time.Now(),
		Category: "Food",
		Note:     "Test Dinner",
	}
	err := AddExpense(expense)
	assert.NoError(t, err)
	assert.NotZero(t, expense.ID, "Expense ID should not be zero after insertion")
	retrieved, err := GetExpenseByID(expense.ID)
	assert.NoError(t, err)
	assert.Equal(t, expense.Amount, retrieved.Amount)
	assert.Equal(t, expense.Note, retrieved.Note)
}

func TestListExpenses(t *testing.T) {
	setupTestDB(t)
	_ = AddExpense(&models.Expense{Amount: 50, Date: time.Now(), Category: "Misc"})
	_ = AddExpense(&models.Expense{Amount: 150, Date: time.Now(), Category: "Food"})
	expenses, err := ListExpenses("", "", "", "")
	assert.NoError(t, err)
	assert.Len(t, expenses, 2, "Should retrieve two expenses")
	expenses, err = ListExpenses("Food", "", "", "")
	assert.NoError(t, err)
	assert.Len(t, expenses, 1, "Should retrieve one expense for Food")
}

func TestDeleteExpense(t *testing.T) {
	setupTestDB(t)
	expense := &models.Expense{
		Amount:   200,
		Date:     time.Now(),
		Category: "Transport",
	}
	_ = AddExpense(expense)
	require.NotZero(t, expense.ID)
	err := DeleteExpense(expense.ID, false)
	assert.NoError(t, err)
	_, err = GetExpenseByID(expense.ID)
	assert.Error(t, err, "Expense should not be found after soft deletion")
	expense2 := &models.Expense{
		Amount:   300,
		Date:     time.Now(),
		Category: "Bills",
	}
	_ = AddExpense(expense2)
	require.NotZero(t, expense2.ID)
	err = DeleteExpense(expense2.ID, true)
	assert.NoError(t, err)
	_, err = GetExpenseByID(expense2.ID)
	assert.Error(t, err, "Expense should not be found after hard deletion")
}

func TestDeleteExpensesByFilter(t *testing.T) {
	setupTestDB(t)
	date := time.Now()
	_ = AddExpense(&models.Expense{Amount: 100, Date: date, Category: "Food"})
	_ = AddExpense(&models.Expense{Amount: 200, Date: date, Category: "Transport"})
	_ = AddExpense(&models.Expense{Amount: 300, Date: date, Category: "Food"})
	rows, err := DeleteExpensesByFilter("Food", "", "", "", false)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), rows, "Should delete 2 food expenses")
	expenses, err := ListExpenses("", "", "", "")
	assert.NoError(t, err)
	assert.Len(t, expenses, 1, "Only one expense should remain")
	assert.Equal(t, "Transport", expenses[0].Category)
}

func TestCountExpenses(t *testing.T) {
	setupTestDB(t)
	_ = AddExpense(&models.Expense{Amount: 100, Date: time.Now(), Category: "Food"})
	_ = AddExpense(&models.Expense{Amount: 200, Date: time.Now(), Category: "Transport"})
	count, err := CountExpenses("Food", "", "", "")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count, "Should count 1 food expense")
	count, err = CountExpenses("", "", "", "")
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count, "Should count all expenses")
}
