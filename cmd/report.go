package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/parrothacker1/expense-tracker/utils"
	"github.com/spf13/cobra"
)

var (
	reportByTotal    bool
	reportByCategory bool
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Show summary reports of your expenses",
	Long: `Generate summary reports. You must specify which report to run.

Examples:
  expensetracker report --total
  expensetracker report --by-category`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if reportByTotal {
			total, err := utils.GetTotalExpenses()
			if err != nil {
				return err
			}
			fmt.Printf("Total Expenses: %.2f\n", total)
			return nil
		}

		if reportByCategory {
			reportData, err := utils.GetExpensesByCategory()
			if err != nil {
				return err
			}
			if len(reportData) == 0 {
				cmd.Println("No expenses found to report on.")
				return nil
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"Category", "Total Amount"})
			var total float64
			for category, amount := range reportData {
				table.Append([]string{category, fmt.Sprintf("%.2f", amount)})
				total += amount
			}
			table.Footer([]string{"TOTAL", fmt.Sprintf("%.2f", total)})
			table.Render()
			return nil
		}

		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
	reportCmd.Flags().BoolVar(&reportByTotal, "total", false, "Show the grand total of all expenses")
	reportCmd.Flags().BoolVar(&reportByCategory, "by-category", false, "Show total expenses grouped by category")
}
