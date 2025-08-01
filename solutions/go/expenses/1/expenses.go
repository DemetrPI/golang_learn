package expenses

import "fmt"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var results = make([]Record, 0) 
	fmt.Printf("Debug message: results variable === %v \n", results)
	for _, record := range in {
		predication := predicate(record)
		fmt.Printf("Debug message: Predication %v \n", predication)
		if predication {
			results = append(results, record)
			fmt.Printf("Debug message: Results %v \n", results)
		}
	}
	return results
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	total := 0.0
	period := ByDaysPeriod(p)
	filter := Filter(in, period)
	for _, record := range filter {
		total += record.Amount
	}
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	records:= Filter(in, ByCategory(c))
	if len(records) == 0  {
		return 0, fmt.Errorf("unknown category %q", c)
	}
	total := TotalByPeriod(records, p)
	return total, nil
}

