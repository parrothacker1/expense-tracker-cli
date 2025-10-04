package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/parrothacker1/expense-tracker/models"
	"gorm.io/gorm"
)

func AddExpense(expense *models.Expense) (err error) {
	err = models.DB.Create(expense).Error
	return
}

func ListExpenses(category, month, startDate, endDate string) ([]models.Expense, error) {
	var expenses []models.Expense
	query := models.DB
	if category != "" {
		query = query.Where("category LIKE ?", "%"+category+"%")
	}
	if month != "" {
		startOfMonth, err := time.Parse("2006-01", month)
		if err == nil {
			endOfMonth := startOfMonth.AddDate(0, 1, -1)
			query = query.Where("date BETWEEN ? AND ?", startOfMonth, endOfMonth)
		}
	}
	if startDate != "" {
		start, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			query = query.Where("date >= ?", start)
		}
	}
	if endDate != "" {
		end, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			query = query.Where("date <= ?", end)
		}
	}
	result := query.Order("date desc").Find(&expenses)
	return expenses, result.Error
}

func buildFilterQuery(category, date, startDate, endDate string) *gorm.DB {
	query := models.DB
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if date != "" {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err == nil {
			startOfDay := parsedDate
			endOfDay := parsedDate.AddDate(0, 0, 1).Add(-time.Second)
			query = query.Where("date BETWEEN ? AND ?", startOfDay, endOfDay)
		}
	}
	if startDate != "" {
		parsedStartDate, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			query = query.Where("date >= ?", parsedStartDate)
		}
	}
	if endDate != "" {
		parsedEndDate, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			endOfDay := parsedEndDate.AddDate(0, 0, 1).Add(-time.Second)
			query = query.Where("date <= ?", endOfDay)
		}
	}
	return query
}

func CountExpenses(category, date, startDate, endDate string) (int64, error) {
	var count int64
	query := buildFilterQuery(category, date, startDate, endDate)
	result := query.Model(&models.Expense{}).Count(&count)
	return count, result.Error
}

func DeleteExpensesByFilter(category, date, startDate, endDate string, permanent bool) (int64, error) {
	query := buildFilterQuery(category, date, startDate, endDate)
	db := query
	if permanent {
		db = db.Unscoped()
	}
	result := db.Delete(&models.Expense{})
	return result.RowsAffected, result.Error
}

func DeleteExpense(id uint, permanent bool) error {
	db := models.DB
	if permanent {
		db = db.Unscoped()
	}
	result := db.Delete(&models.Expense{}, id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("no expense found with ID %d", id)
	}
	return result.Error
}

func GetExpenseByID(id uint) (expense models.Expense, err error) {
	result := models.DB.First(&expense, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("no expense found with ID %d", id)
		return
	}
	err = result.Error
	return
}

func UpdateExpense(expense *models.Expense) error {
	result := models.DB.Save(expense)
	return result.Error
}

func GetTotalExpenses() (float64, error) {
	var total float64
	result := models.DB.Model(&models.Expense{}).Select("sum(amount)").Row()
	if err := result.Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

func GetExpensesByCategory() (map[string]float64, error) {
	var results []struct {
		Category string
		Total    float64
	}
	reportMap := make(map[string]float64)
	err := models.DB.Model(&models.Expense{}).
		Select("category, sum(amount) as total").
		Group("category").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	for _, res := range results {
		reportMap[res.Category] = res.Total
	}
	return reportMap, nil
}
