package cmd

import (
	"fmt"
	"os"

	"github.com/parrothacker1/expense-tracker/models"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expensetracker",
	Short: "ExpenseTracker is a CLI tool to manage your personal finances.",
	Long:  `A fast and flexible command-line expense tracker built with Go.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		dbPath := "expenses.db"
		return models.InitDB(dbPath)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
