# ExpenseTracker CLI 💰

A fast and powerful command-line tool for managing your personal expenses, built with Go.

[![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8.svg)](https://golang.org/)  
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)  
[![GitHub Actions CI](https://github.com/parrothacker1/expense-tracker-cli/actions/workflows/release.yml/badge.svg)](https://github.com/parrothacker1/expense-tracker-cli/actions/workflows/release.yml)  
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/parrothacker1/expense-tracker-cli)](https://github.com/parrothacker1/expense-tracker-cli/releases)

ExpenseTracker is a self-contained, cross-platform CLI application that helps you track your spending right from the terminal. It's designed to be fast, simple, and safe, using a local SQLite database to store your data.

---

## ✨ Announcing v1.0.0!

This `v1.0.0` release marks the first stable version of ExpenseTracker. The core features for adding, listing, updating, and deleting expenses are complete, robustly tested, and ready for use.

---

## Features

- **📝 Full CRUD Operations:** Add, update, and delete expenses with simple commands.  
- **📊 Powerful Filtering:** List and view expenses with filters for category, month, or a specific date range.  
- **🛡️ Safe By Default:** Deletions are "soft" by default (records are hidden, not erased). Use the `--permanent` flag for irreversible hard deletes.  
- **Interactive Confirmation:** Destructive actions (like deleting multiple items) require confirmation, preventing accidental data loss.  
- **💻 Cross-Platform:** A single binary works on Windows, macOS, and Linux, thanks to Go.  
- **📦 Self-Contained:** No external dependencies needed to run; just the binary and your database file.  

---

## 🚀 Installation

The easiest way to get started is by downloading a pre-compiled binary from our latest release.

### From GitHub Releases (Recommended)

1. Go to the [**Releases Page**](https://github.com/parrothacker1/expense-tracker-cli/releases/latest).  
2. Find the asset for your operating system and architecture (e.g., `expensetracker-linux-amd64`, `expensetracker-windows-amd64.exe`).  
3. Download the binary.  
4. (Optional but recommended) Place the binary in a directory that is in your system's `PATH` (like `/usr/local/bin` on Linux/macOS) and rename it to `expensetracker`. This allows you to run it from anywhere.  

### From Source

If you have Go installed, you can build it from the source code.

```sh
# Clone the repository
git clone https://github.com/parrothacker1/expense-tracker-cli.git
cd expense-tracker-cli

# Build the binary
go build -o expensetracker .

# Now you can run it
./expensetracker --help
```

---

## 💻 Usage

Here are some examples of the most common commands.

### Add an Expense

```sh
# Add a new expense with all details
expensetracker add --amount 550.75 --category "Groceries" --note "Weekly shopping" --date "2025-10-04"
```

### List Expenses

```sh
# List all expenses
expensetracker list

# List only expenses from the "Food" category
expensetracker list --category Food

# List expenses from October 2025
expensetracker list --month 2025-10

# List expenses within a specific date range
expensetracker list --from 2025-10-01 --to 2025-10-15
```

### Update an Expense

```sh
# Update the amount and note for the expense with ID 1
expensetracker update 1 --amount 600.00 --note "Updated weekly shopping total"
```

### Delete an Expense

```sh
# Soft-delete the expense with ID 3
expensetracker delete 3

# Delete all expenses in the "Temporary" category without a confirmation prompt
expensetracker delete --category "Temporary" --force

# PERMANENTLY delete the expense with ID 5 (this cannot be undone)
expensetracker delete 5 --permanent
```

---

## 🛠️ Development & Testing

Contributions are welcome!

### Prerequisites

- Go (version 1.21 or newer)  
- `bats-core` for running the test suite.  

### Testing

This project uses **`bats-core`** for robust end-to-end integration testing. This approach was chosen because it allows us to test the compiled application binary exactly as a user would. The test suite simulates real command-line usage and verifies the application's output, exit codes, and database state.

To run the full test suite:

```sh
bats test.bats
```

