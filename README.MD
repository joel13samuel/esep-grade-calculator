
# ESEP Grade Calculator

This project is a simple grade calculator written in Go. It calculates a student's final grade using weighted averages from assignments, exams, and essays.

## Requirements

- Go 1.24 or newer

## Grade Weights

- Assignments: 50%
- Exams: 35%
- Essays: 15%

The calculator averages the grades in each category and multiplies those averages by their weights to calculate the final numerical grade.

## Project Files

- `grade_calculator.go` - Contains the main grade calculator implementation.
- `grade_calculator_test.go` - Contains unit tests to check the calculator.
- `go.mod` - Defines the Go module and its version.

## Running Tests

From the repository's root folder, navigate into the Go module:

```bash
cd grade-calculator
```

Run the unit tests:

```bash
go test -v ./...
```

## Checking Code Coverage

To run the tests and generate a coverage report:

```bash
go test -coverprofile=coverage.out ./...
```

To see the total coverage percentage:

```bash
go tool cover -func=coverage.out
```

To view which lines are covered in your browser:

```bash
go tool cover -html=coverage.out
```

## Current Project Status

This repository is part of an automated unit testing exercise. Some of the original tests currently fail, and the calculator implementation requires corrections.
