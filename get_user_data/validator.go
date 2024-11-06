package get_user_data

import "time"

func isDateStrValid(date string) bool {
	// Could take out a config var of date format for the  format string
	_, err := time.Parse("02012006", date)
	return err == nil
}
