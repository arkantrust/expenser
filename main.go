package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var es ExpenseService

func AddExpense() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
    description := addCmd.String("description", "", "What did you spend your money on?")
    amount := addCmd.Int("amount", 0, "How much did you spend?")
	addCmd.Parse(os.Args[2:])

	id, err := es.Add(*description, *amount)
	if err != nil {
		fmt.Printf("Couldn't add expense: %v\n", err.Error())
		return
	}

	fmt.Printf("Expense added successfully (ID: %v)\n", id)
}

func PrintAllExpenses() {
	expenses, err := es.GetAll()
	if err != nil {
		fmt.Printf("Couldn't get expenses: %v\n", err.Error())
		return
	}

	fmt.Println("ID   Date              Description              Amount")
	for _, ex := range expenses {
		date := time.Unix(ex.Date, 0)
		fmt.Printf("%v    %v              %v              %v\n", ex.ID, date.Local().Format("02-01-2006"), ex.Description, ex.Amount)
	}
}

func PrintExpense() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
    id := getCmd.Int("id", 0, "ID of the expense you're looking up")
	getCmd.Parse(os.Args[2:])

	ex, err := es.Get(*id)
	if err != nil {
		fmt.Printf("Couldn't get expense: %v\n", err.Error())
		return
	}

	fmt.Println("ID   Date              Description              Amount")
	fmt.Printf("%v    %v              %v              %v\n", ex.ID, ex.Date, ex.Description, ex.Amount)
}

func PrintSummary() {
	s, err := es.GetTotalCost()
	if err != nil {
		fmt.Printf("Couldn't show summary: %v\n", err.Error())
		return
	}
	fmt.Printf("Total expenses: $%v\n", s)
}

func DeleteExpense() {
	deleteCmd := flag.NewFlagSet("get", flag.ExitOnError)
    id := deleteCmd.Int("id", 0, "ID of the expense you're looking up")
	deleteCmd.Parse(os.Args[2:])

	err := es.Delete(*id)
	if err != nil {
		fmt.Printf("Couldn't delete expense: %v\n", err.Error())
		return
	}
	fmt.Println("Expense deleted successfully")
}

func PrintHelp() {
	fmt.Printf("Usage: %s [command]\n", os.Args[0])
	fmt.Println("Available commands: add, get, delete, list, summary")
}

func main() {
	db, err := NewJSONDB("./db.json")
	if err != nil {
		fmt.Printf("Couldn't start the database: %v\n", err.Error())
		return
	}

	es = ExpenseService{ db: db }

	command := os.Args[1]
	if command == "" {
		PrintHelp()
	}
	command = strings.ToLower(command)

	switch command {
	case "add":
		AddExpense()
	case "get":
		PrintExpense()
	case "delete":
		DeleteExpense()
	case "list":
		PrintAllExpenses()
	case "summary":
		PrintSummary()
	default:
		PrintHelp()
	}

	fmt.Println()
}