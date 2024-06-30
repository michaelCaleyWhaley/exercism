package booking

import (
	"fmt"
	"strconv"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	scheduledTime, _ := time.Parse("1/02/2006 15:04:05", date)
	return scheduledTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	scheduledTime, _ := time.Parse("January 2, 2006 15:04:05", date)
	return scheduledTime.Compare(time.Now()) < 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	dateTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	dateHour, _, _ := dateTime.Clock()
	return dateHour >= 12 && dateHour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	dateTime, _ := time.Parse("1/2/2006 15:04:05", date)
	weekday := dateTime.Weekday().String()
	month := dateTime.Month().String()
	day := strconv.Itoa(dateTime.Day())
	year := strconv.Itoa(dateTime.Year())
	dateTimeHour, dateTimeMin, _ := dateTime.Clock()

	return fmt.Sprintf("You have an appointment on %s, %s %s, %s, at %d:%d.", weekday, month, day, year, dateTimeHour, dateTimeMin)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	anniTime, _ := time.Parse("January 2, 2006", "September 15, "+strconv.Itoa(time.Now().Year()))
	return anniTime
}
