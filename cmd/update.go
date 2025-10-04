package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/parrothacker1/expense-tracker/utils"
	"github.com/spf13/cobra"
)

var (
	updateAmount   float64
	updateDate     string
	updateCategory string
	updateNote     string
)

var updateCmd = &cobra.Command{
	Use:   "update [ID]",
	Short: "Update an existing expense",
	Long: `Update the details of an expense by providing its ID and the fields to modify.
At least one flag must be provided.

Examples:
  expensetracker update 1 --amount 50.50
  expensetracker update 2 --category "Utilities" --note "Monthly electricity bill"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !cmd.Flags().Changed("amount") && !cmd.Flags().Changed("date") && !cmd.Flags().Changed("category") && !cmd.Flags().Changed("note") {
			return errors.New("at least one flag (--amount, --date, --category, --note) must be provided to update an expense")
		}
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid ID provided: %w", err)
		}
		expense, err := utils.GetExpenseByID(uint(id))
		if err != nil {
			return err
		}
		if cmd.Flags().Changed("amount") {
			expense.Amount = updateAmount
		}
		if cmd.Flags().Changed("date") {
			parsedDate, err := time.Parse("2006-01-02", updateDate)
			if err != nil {
				return fmt.Errorf("invalid date format (use YYYY-MM-DD): %w", err)
			}
			expense.Date = parsedDate
		}
		if cmd.Flags().Changed("category") {
			expense.Category = updateCategory
		}
		if cmd.Flags().Changed("note") {
			expense.Note = updateNote
		}
		if err := utils.UpdateExpense(&expense); err != nil {
			return fmt.Errorf("error updating expense: %w", err)
		}
		cmd.Printf("Successfully updated expense with ID: %d\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().Float64VarP(&updateAmount, "amount", "a", 0, "New amount for the expense")
	updateCmd.Flags().StringVarP(&updateDate, "date", "d", "", "New date for the expense (YYYY-MM-DD)")
	updateCmd.Flags().StringVarP(&updateCategory, "category", "c", "", "New category for the expense")
	updateCmd.Flags().StringVarP(&updateNote, "note", "n", "", "New note for the expense")
}
