package get_user_data

import (
	"fmt"
	"oneiro_task/loan"
	"time"
)

func GetConsoleLoanDataAndValidate() loan.Loan {

	loan := loan.Loan{}

	// Could refactor out some funcs like "getDate, getString, getFloat" for reusability

	var startDateStr string
	fmt.Print("Enter loan start date (DDMMYYYY): ")
	fmt.Scanln(&startDateStr)
	for false == isDateStrValid(startDateStr) {
		fmt.Print("Invalid format, enter loan start date (DDMMYYYY): ")
		fmt.Scanln(&startDateStr)
	}
	loan.StartDate, _ = time.Parse("02012006", startDateStr)

	var endDateStr string
	fmt.Print("Enter loan start date (DDMMYYYY): ")
	fmt.Scanln(&endDateStr)
	for false == isDateStrValid(endDateStr) {
		fmt.Print("Invalid format, enter loan start date (DDMMYYYY): ")
		fmt.Scanln(&endDateStr)
	}
	loan.EndDate, _ = time.Parse("02012006", endDateStr)

	// Validate that it's a number with 2dp
	fmt.Print("Enter loan amount without currency (eg 123.45 or 123.00): ")
	fmt.Scanln(&loan.Amount)

	// Validate against a whitelist of allowed symbols
	fmt.Print("Enter loan currency symbol (eg £ or $): ")
	fmt.Scanln(&loan.Currency)

	// Validation just requires its a number not a string
	fmt.Print("Enter base interest rate for the loan (eg 2.51 or 1): ")
	fmt.Scanln(&loan.BaseInterestRate)

	// Validation just requires its a number not a string
	fmt.Print("Enter margin rate for the loan (eg 2.51 or 1): ")
	fmt.Scanln(&loan.Margin)

	return loan
}
