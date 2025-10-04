#!/usr/bin/env bats

setup() {
    go build -o expensetracker .
    rm -f .expenses.sqlite3
    export EXPENSE_DB_PATH=".expenses.sqlite3"
}

# --- ADD COMMAND TESTS ---

@test "should add an expense successfully" {
    run ./expensetracker add -a 100 -c "Test" -n "Initial item"
    [ "$status" -eq 0 ]
    [[ "$output" == *"Successfully added expense with ID: 1"* ]]
}

@test "should fail to add an expense without required amount" {
    run ./expensetracker add -c "Test"
    [ "$status" -ne 0 ]
    [[ "$output" == *"required flag(s) \"amount\" not set"* ]]
}

# --- LIST COMMAND TESTS ---

@test "should list previously added expenses" {
    run ./expensetracker add -a 50 -n "First Item" -c "Food"
    run ./expensetracker add -a 75 -n "Second Item" -c "Transport"

    run ./expensetracker list
    [ "$status" -eq 0 ]
    [[ "$output" == *"First Item"* ]]
    [[ "$output" == *"Second Item"* ]]
}

@test "should list and filter by category" {
    run ./expensetracker add -a 50 -n "Pizza" -c "Food"
    run ./expensetracker add -a 75 -n "Bus Ticket" -c "Transport"

    run ./expensetracker list --category "Food"
    [ "$status" -eq 0 ]
    [[ "$output" == *"Pizza"* ]]
    [[ "$output" != *"Bus Ticket"* ]]
}

@test "should show a message when listing with no results" {
    run ./expensetracker list --category "Non-Existent"
    [ "$status" -eq 0 ]
    [[ "$output" == *"No expenses found"* ]]
}

# --- UPDATE COMMAND TESTS ---

@test "should update an existing expense" {
    run ./expensetracker add -a 100 -c "Old Category"

    run ./expensetracker update 1 --amount 150.50 --category "New Category"
    [ "$status" -eq 0 ]
    [[ "$output" == *"Successfully updated expense with ID: 1"* ]]

    run ./expensetracker list
    [ "$status" -eq 0 ]
    [[ "$output" == *"150.50"* ]]
    [[ "$output" == *"New Category"* ]]
    [[ "$output" != *"Old Category"* ]]
}

@test "should fail to update a non-existent expense" {
    run ./expensetracker update 99 --amount 100
    [ "$status" -ne 0 ]
    [[ "$output" == *"no expense found with ID 99"* ]]
}

# --- DELETE COMMAND TESTS ---

@test "should soft-delete an expense by ID" {
    run ./expensetracker add -a 100 -n "Item to delete"
    run ./expensetracker add -a 200 -n "Item to keep"

    run ./expensetracker delete 1
    [ "$status" -eq 0 ]
    [[ "$output" == *"Successfully deleted expense with ID: 1"* ]]

    run ./expensetracker list
    [ "$status" -eq 0 ]
    [[ "$output" != *"Item to delete"* ]]
    [[ "$output" == *"Item to keep"* ]]
}

@test "should permanently delete an expense with --permanent flag" {
    run ./expensetracker add -a 100 -n "Delete me permanently"

    run ./expensetracker delete 1 --permanent --force
    [ "$status" -eq 0 ]
    [[ "$output" == *"Successfully PERMANENTLY deleted"* ]]
}

@test "should delete multiple expenses by filter" {
    run ./expensetracker add -a 10 -c "Delete"
    run ./expensetracker add -a 20 -c "Delete"
    run ./expensetracker add -a 30 -c "Keep"

    run ./expensetracker delete --category "Delete" --force
    [ "$status" -eq 0 ]
    [[ "$output" == *"Successfully soft-deleted 2 expense(s)."* ]]

    run ./expensetracker list
    [ "$status" -eq 0 ]
    [[ "$output" != *"Delete"* ]]
    [[ "$output" == *"Keep"* ]]
}

# --- REPORT COMMAND TESTS ---

@test "should report the grand total of all expenses" {
    run ./expensetracker add -a 100.50 -c "Food"
    run ./expensetracker add -a 75 -c "Transport"
    run ./expensetracker add -a 25.25 -c "Food"

    run ./expensetracker report --total
    [ "$status" -eq 0 ]
    [[ "$output" == *"Total Expenses: 200.75"* ]]
}

@test "should report expenses grouped by category" {
    run ./expensetracker add -a 100 -c "Food"
    run ./expensetracker add -a 75 -c "Transport"
    run ./expensetracker add -a 25 -c "Food"
    run ./expensetracker add -a 50 -c "Transport"

    run ./expensetracker report --by-category
    [ "$status" -eq 0 ]
    grep -q "Food" <<< "$output"
    grep -q "Transport" <<< "$output"
    grep -q "125.00" <<< "$output"
    grep -q "250.00" <<< "$output"
}
