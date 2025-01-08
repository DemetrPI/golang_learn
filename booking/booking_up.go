package booking

import (
    "fmt"
    "time"
)

// Common layout formats for parsing
const (
    layoutDefault  = "1/2/2006 15:04:05"
    layoutExtended = "January 2, 2006 15:04:05"
    layoutFull     = "Monday, January 2, 2006 15:04:05"
)

// parseDateOrZero parses dates and returns a zero value if there's an error.
func parseDateOrZero(date, layout string) time.Time {
    t, err := time.Parse(layout, date)
    if err != nil {
        fmt.Println("Error parsing date:", err)
        return time.Time{} // Return zero value if parsing fails
    }
    return t
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    return parseDateOrZero(date, layoutDefault)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    t := parseDateOrZero(date, layoutExtended)
    return !t.IsZero() && t.Before(time.Now()) // Check if date is valid and in the past
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    t := parseDateOrZero(date, layoutFull)
    return !t.IsZero() && t.Hour() >= 12 && t.Hour() < 18 // Check if date is valid and in the afternoon
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    t := parseDateOrZero(date, layoutDefault)
    if t.IsZero() {
        return ""
    }

    return fmt.Sprintf(
        "You have an appointment on %s, %s %d, %d, at %02d:%02d.",
        t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute(),
    )
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    currentYear := time.Now().Year()
    return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}