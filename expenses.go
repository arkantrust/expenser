package main

import (
	"errors"
	"time"
)

type ExpenseService struct {
	db DB
}

func (es ExpenseService) GetAll() ([]Expense, error) {
	expenses, err := es.db.Load()
	if err != nil {
		return nil, err
	}

	nonDeleted := make([]Expense, 0)
	for _, e := range expenses {
		if !e.Deleted {
			nonDeleted = append(nonDeleted, e)
		}
	}
	return nonDeleted, nil
}

func (es ExpenseService) Add(description string, amount int) (int, error) {
	expenses, err := es.db.Load()
	if err != nil {
		return 0, err
	}

	if len(description) == 0 {
		return 0, errors.New("description can't be empty")
	}

	if amount <= 0 {
		return 0, errors.New("amount must be a positive number")
	}

	id := len(expenses) + 1
	newExpense := Expense{
		ID:          id,
		Amount:      amount,
		Description: description,
		Date:        time.Now().Unix(),
	}

	expenses = append(expenses, newExpense)

	err = es.db.Save(expenses)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (es ExpenseService) Get(id int) (Expense, error) {
	expenses, err := es.db.Load()
	if err != nil {
		return *new(Expense), err
	}

	e := expenses[id-1]
	if e.Deleted {
		return *new(Expense), errors.New("expense not found")
	}
	return e, nil
}

func (es ExpenseService) Delete(id int) error {
	expenses, err := es.db.Load()
	if err != nil {
		return err
	}

	e := &expenses[id-1]
	if e.Deleted {
		return errors.New("expense not found")
	}
	e.Deleted = true

	err = es.db.Save(expenses)
	if err != nil {
		return err
	}

	return nil
}

func (es ExpenseService) GetTotalCost() (int, error) {
	expenses, err := es.db.Load()
	if err != nil {
		return 0, err
	}
	
	sum := 0
	for _, e := range expenses {
		if !e.Deleted {
			sum += e.Amount
		}
	}
	return sum, nil
}
