package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Go launched at: %s\n", time.Now())
	fmt.Printf("Local time is: %s\n", time.Now().Local())

	myTime := time.Date(2024, time.July, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("my time defined as: %s\n", myTime)

	timeToAdd := 24 * time.Hour

	newTime := myTime.Add(timeToAdd)
	fmt.Printf("time after 24 hours: %s\n", newTime)

	const shortForm = "2006-Jan-02"
	t, _ := time.Parse(shortForm, "2024-Feb-03")
	fmt.Println("shortForm::", t)

	formattedTime := newTime.Format("Monday, January 2, 2006 15:12:04")
	fmt.Printf("formatted time is %s\n\n", formattedTime)

	t0 := time.Now()
	processTime() //some expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

func processTime() {
	// See the example for Time.Format for a thorough description of how
	// to define the layout string to parse a time.Time value; Parse and
	// Format use the same model to describe their input and output.

	// longForm shows by example how the reference time would be represented in
	// the desired layout.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t1, _ := time.Parse(longForm, "Feb 3, 2024 at 7:54pm (PST)")
	fmt.Println("longForm:", t1)

	// shortForm is another way the reference time would be represented
	// in the desired layout; it has no time zone present.
	// Note: without explicit zone, returns time in UTC.
	const shortForm = "2006-Jan-02"
	t2, _ := time.Parse(shortForm, "2024-Feb-03")
	fmt.Println("shortForm:", t2)

	// Some valid layouts are invalid time values, due to format specifiers
	// such as _ for space padding and Z for zone information.
	// For example the RFC3339 layout 2006-01-02T15:04:05Z07:00
	// contains both Z and a time zone offset in order to handle both valid options:
	// 2006-01-02T15:04:05Z
	// 2006-01-02T15:04:05+07:00
	t3, _ := time.Parse(time.RFC3339, "2024-01-02T15:04:05Z")
	fmt.Println("RFC3339:", t3)

	t4, _ := time.Parse(time.RFC3339, "2024-01-02T15:04:05+07:00")
	fmt.Println("RFC3339:", t4)

	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error", err) // Returns an error as the layout is not a valid time value

	fmt.Println("sleeping for 2 seconds")
	time.Sleep(2 * time.Second)
	fmt.Println("done")
}

/*
const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)
*/
