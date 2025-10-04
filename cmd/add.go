package cmd

import (
	"fmt"
	"time"

	"github.com/parrothacker1/expense-tracker/models"
	"github.com/parrothacker1/expense-tracker/utils"
	"github.com/spf13/cobra"
)

var (
	amount   float64
	date     string
	category string
	note     string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Long:  `Add a new expense to your tracker.`,
	Run: func(cmd *cobra.Command, args []string) {
		expenseDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			fmt.Println("Error parsing date. Please use YYYY-MM-DD format.", err)
			return
		}
		expense := models.Expense{
			Amount:   amount,
			Date:     expenseDate,
			Category: category,
			Note:     note,
		}
		if err := utils.AddExpense(&expense); err != nil {
			fmt.Println("Error adding expense:", err)
			return
		}
		fmt.Printf("Successfully added expense with ID: %d\n", expense.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0.0, "Amount of the expense (required)")
	addCmd.Flags().StringVarP(&date, "date", "d", time.Now().Format("2006-01-02"), "Date of the expense (YYYY-MM-DD)")
	addCmd.Flags().StringVarP(&category, "category", "c", "General", "Category of the expense")
	addCmd.Flags().StringVarP(&note, "note", "n", "", "A short note for the expense")
	addCmd.MarkFlagRequired("amount")
}
