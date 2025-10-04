package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/parrothacker1/expense-tracker/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		expenses, err := utils.ListExpenses()
		if err != nil {
			fmt.Println("Error listing expenses:", err)
			return
		}
		if len(expenses) == 0 {
			fmt.Println("No expenses found.")
			return
		}
		table := tablewriter.NewWriter(os.Stdout)

		for _, e := range expenses {
			row := []string{
				strconv.FormatInt(int64(e.ID), 10),
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
	// Add flags for filtering here later (e.g., --category, --month)
}
