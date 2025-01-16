package main

import (
	// "bufio"
	"encoding/json"
	// "io"
	"os"
)

type Expenses []Expense

type DB interface {
	Load() (Expenses, error)

	Save(ex Expenses) error
}

type JSONDB struct {
	Filepath string
	expenses Expenses
}

func NewJSONDB(fp string) (DB, error) {

	// Get the directory and filename

	// Ensure the directory exists

	// Ensure the file is a JSON file

	return JSONDB{
		Filepath: fp,
		expenses: make([]Expense, 0),
	}, nil
}

func (db JSONDB) Load() (Expenses, error) {
	j, err := os.ReadFile(db.Filepath)
	if err != nil {
		return nil, err
	}

	var ex Expenses

	err = json.Unmarshal(j, &ex)
	if err != nil {
		return nil, err
	}

	return ex, nil
}

func (db JSONDB) Save(ex Expenses) error {

	j, err := json.Marshal(ex)
	if err != nil {
		return err
	}

	err = os.WriteFile(db.Filepath, j, 0644)
	if err != nil {
		return err
	}

	return nil
}
