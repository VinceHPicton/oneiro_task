package main

import (
	"fmt"
	"oneiro_task/get_user_data"
	"oneiro_task/store"
)

func main() {

	fmt.Println("Welcome to the loan calculator")

	store := store.InitStore()

	loan := get_user_data.GetConsoleLoanDataAndValidate()
	store.Add(loan)

	loan.PrintDailyAccruals()

	loan.PrintTotalInterest()

}
