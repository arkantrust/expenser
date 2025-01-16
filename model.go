package main

type Expense struct {
	ID          int    `json:"id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Date        int64  `json:"date"`
	Deleted     bool   `json:"deleted"`
}
