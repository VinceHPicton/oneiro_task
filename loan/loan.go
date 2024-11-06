package loan

import (
	"fmt"
	"time"
)

type Loan struct {
	StartDate        time.Time
	EndDate          time.Time
	Amount           float64
	Currency         string
	BaseInterestRate float64
	Margin           float64
}

func (loan *Loan) DailyInterestWithoutMargin() float64 {
	return loan.Amount * loan.BaseInterestRate / 100 / 365
}

func (loan *Loan) DailyInterestAccrued() float64 {
	return loan.Amount * (loan.BaseInterestRate + loan.Margin) / 100 / 365
}

func (loan *Loan) LoanPeriodDays() int {
	loanPeriodDuration := loan.EndDate.Sub(loan.StartDate)
	loanPeriodDays := int(loanPeriodDuration.Hours() / 24)
	return loanPeriodDays
}

func (loan *Loan) TotalInterest() float64 {
	return loan.DailyInterestAccrued() * float64(loan.LoanPeriodDays())
}

func (loan *Loan) PrintDailyAccruals() {
	fmt.Printf("\nDaily Interest Accruals (%s):\n", loan.Currency)
	fmt.Printf("%-15s %-20s %-25s %-25s\n", "Accrual Date", "Elapsed Days", "Daily Interest w/o Margin", "Daily Interest w/ Margin")
	dailyInterest := loan.DailyInterestAccrued()
	dailyInterestWOMargin := loan.DailyInterestWithoutMargin()

	for i := 1; i <= loan.LoanPeriodDays(); i++ {
		fmt.Printf("%-15s %-20d %-25.2f %-25.2f\n",
			loan.StartDate.AddDate(0, 0, i).Format("2006-01-02"),
			i,
			dailyInterestWOMargin,
			dailyInterest)
	}
}

func (loan *Loan) PrintTotalInterest() {
	fmt.Printf("\nTotal Interest Over Period (%s): %.2f\n", loan.Currency, loan.TotalInterest())
}
