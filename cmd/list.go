package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/parrothacker1/expense-tracker/utils"
	"github.com/spf13/cobra"
)

var (
	filterCategory  string
	filterMonth     string
	filterStartDate string
	filterEndDate   string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses with optional filters",
	Long: `List expenses. You can filter by category, a specific month, or a date range.
Examples:
  expensetracker list --category Food
  expensetracker list --month 2025-10
  expensetracker list --from 2025-10-01 --to 2025-10-15`,
	Run: func(cmd *cobra.Command, args []string) {
		expenses, err := utils.ListExpenses(filterCategory, filterMonth, filterStartDate, filterEndDate)
		if err != nil {
			cmd.Println("Error listing expenses:", err)
			return
		}
		if len(expenses) == 0 {
			cmd.Println("No expenses found matching the criteria.")
			return
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"ID", "Amount", "Date", "Category", "Note"})
		for _, e := range expenses {
			row := []string{
				strconv.FormatUint(uint64(e.ID), 10),
				fmt.Sprintf("%.2f", e.Amount),
				e.Date.Format("2006-01-02"),
				e.Category,
				e.Note,
			}
			table.Append(row)
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&filterCategory, "category", "c", "", "Filter by category")
	listCmd.Flags().StringVarP(&filterMonth, "month", "m", "", "Filter by month (format: YYYY-MM)")
	listCmd.Flags().StringVar(&filterStartDate, "from", "", "Filter from date (format: YYYY-MM-DD)")
	listCmd.Flags().StringVar(&filterEndDate, "to", "", "Filter to date (format: YYYY-MM-DD)")
}
