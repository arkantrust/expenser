# Expenser

Expenser is a simple expense tracker that allows you to keep track of your expenses built with [Go](https://go.dev). Inspired by [Expense Tracker](https://roadmap.sh/projects/expense-tracker).

## Features

- [x] Add an expense with a description and amount.
- [x] Update an expense.
- [x] Delete an expense.
- [x] View all expenses.
- [x] View a summary of all expenses.
- [ ] View a summary of expenses for a specific month (of current year).
- [ ] Filter expenses by categories created by the user.
- [ ] Set a budget for each month and show a warning when the budget is exceeded.
- [ ] Export expenses to a CSV or JSON file.

## Getting Started

To run this project, you only need to have Go installed on your machine. You can download it from the [official website](https://go.dev/dl/).

Then build the project with the following command:

```bash
go build -o expenser
```

And you're good to go! You can start adding expenses:

```bash
./expenser add --description "Lunch" --amount 15000 # prices are set as int because it's more intended for colombian🇨🇴 pesos than dollars
```

## Contributing

If you want to contribute to this project, feel free to fork it and submit a pull request. You can also open an issue if you find a bug or have a feature request.