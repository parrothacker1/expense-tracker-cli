package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/parrothacker1/expense-tracker/utils"
	"github.com/spf13/cobra"
)

var (
	delCategory  string
	delDate      string
	delStartDate string
	delEndDate   string
	delForce     bool
	delPermanent bool
)

func confirmAction(cmd *cobra.Command, prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	cmd.Printf("%s [y/N]: ", prompt)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}

var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Delete an expense by ID or by filters",
	Long: `Delete an expense. By default, this is a "soft delete".
Use the --permanent flag to remove records from the database permanently.

Examples:
  expensetracker delete 1
  expensetracker delete --category Food --force
  expensetracker delete --from 2025-10-01 --permanent`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid ID provided: %w", err)
			}
			if err := utils.DeleteExpense(uint(id), delPermanent); err != nil {
				return err
			}
			if delPermanent {
				cmd.Printf("Successfully PERMANENTLY deleted expense with ID: %d\n", id)
			} else {
				cmd.Printf("Successfully deleted expense with ID: %d\n", id)
			}
			return nil
		}
		if !cmd.Flags().Changed("category") && !cmd.Flags().Changed("date") && !cmd.Flags().Changed("from") && !cmd.Flags().Changed("to") {
			return errors.New("you must specify an ID or at least one filter flag")
		}
		count, err := utils.CountExpenses(delCategory, delDate, delStartDate, delEndDate)
		if err != nil {
			return fmt.Errorf("error finding expenses to delete: %w", err)
		}
		if count == 0 {
			cmd.Println("No expenses found matching the criteria.")
			return nil
		}
		confirmed := delForce
		if !confirmed {
			var prompt string
			if delPermanent {
				prompt = fmt.Sprintf("This will PERMANENTLY delete %d expense(s). This action cannot be undone. Are you sure?", count)
			} else {
				prompt = fmt.Sprintf("This will soft-delete %d expense(s) (they can be recovered). Are you sure?", count)
			}
			confirmed = confirmAction(cmd, prompt)
		}
		if !confirmed {
			cmd.Println("Deletion cancelled.")
			return nil
		}
		deletedCount, err := utils.DeleteExpensesByFilter(delCategory, delDate, delStartDate, delEndDate, delPermanent)
		if err != nil {
			return fmt.Errorf("error deleting expenses: %w", err)
		}
		if delPermanent {
			cmd.Printf("Successfully PERMANENTLY deleted %d expense(s).\n", deletedCount)
		} else {
			cmd.Printf("Successfully soft-deleted %d expense(s).\n", deletedCount)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&delCategory, "category", "c", "", "Delete by category")
	deleteCmd.Flags().StringVarP(&delDate, "date", "d", "", "Delete by specific date (YYYY-MM-DD)")
	deleteCmd.Flags().StringVar(&delStartDate, "from", "", "Delete from date (YYYY-MM-DD)")
	deleteCmd.Flags().StringVar(&delEndDate, "to", "", "Delete to date (YYYY-MM-DD)")
	deleteCmd.Flags().BoolVar(&delForce, "force", false, "Force delete without confirmation")
	deleteCmd.Flags().BoolVar(&delPermanent, "permanent", false, "Permanently delete records (hard delete)")
}
