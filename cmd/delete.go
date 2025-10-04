package cmd

import (
	"fmt"
	"strconv"

	"github.com/parrothacker1/expense-tracker/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Delete an expense by its ID",
	Args:  cobra.ExactArgs(1), // Requires exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid ID provided.", err)
			return
		}
		err = utils.DeleteExpense(uint(id))
		if err != nil {
			fmt.Println("Error deleting expense:", err)
			return
		}
		fmt.Printf("Successfully deleted expense with ID: %d\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
