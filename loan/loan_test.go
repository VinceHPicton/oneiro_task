package loan

import (
	"testing"
	"time"
)

var (
	startDate, _ = time.Parse("02-01-2006", "20-03-2014")
	endDate, _   = time.Parse("02-01-2006", "20-03-2015")
	inputLoan    = Loan{
		StartDate:        startDate,
		EndDate:          endDate,
		Amount:           10000.00,
		Currency:         "£",
		BaseInterestRate: 3.0,
		Margin:           3.0,
	}
)

func Test_DailyInterestWithoutMargin(t *testing.T) {

	expected := 0.821917808219178
	got := inputLoan.DailyInterestWithoutMargin()

	if expected != got {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func Test_TotalInterest(t *testing.T) {

	newEndDate, _ := time.Parse("02-01-2006", "26-03-2015")
	newLoan := inputLoan
	newLoan.EndDate = newEndDate

	expected := float64(609.8630136986301)
	got := newLoan.TotalInterest()

	if expected != got {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func Test_LoanPeriodDays(t *testing.T) {

	expected := 365
	got := inputLoan.LoanPeriodDays()

	if expected != got {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func Test_PrintDailyAccruals(t *testing.T) {

	inputLoan.PrintDailyAccruals()
}
