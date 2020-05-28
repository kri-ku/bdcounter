// https://yourbasic.org/golang/format-parse-string-time-date-example/
// https://yourbasic.org/golang/day-month-year-from-time/
// https://gobyexample.com/time
// https://teamtreehouse.com/community/how-to-return-a-map
// https://golang.org/pkg/time/#example_Time_Sub
// https://yourbasic.org/golang/day-month-year-from-time/
// https://ednsquare.com/story/date-and-time-manipulation-golang-with-examples------cU1FjK
// https://yourbasic.org/golang/day-month-year-from-time/


package main

import (
	"fmt"
	"time"
	s "strings"
	"strconv"
)


func makeDate(year, month, date int) time.Time {
// how about error??
	return time.Date(year,time.Month(month), date, 0,0,0,0, time.UTC)

}

func convertToNumber(word string)int {
	number, err := strconv.Atoi(word)

	if err != nil {
		return -1
	}

	return number
}

/*func daysBetween(date1 Date, date2 Date)int {

//jos halutaan tietää montako päivää vanha henkilö on, kesken!!!
	days := date2.Sub(date1).Hours() // 24
	return days
}*/

/*func daysTill(x* Date, y* Date) int {
	calendar := makeCalendar()
	return 6

}*/

func main() {

	fmt.Println("Let's find out how many dates you have till your birthday. First, tell me your birthday:\nyear-month-date")
	var bdaystr string
	fmt.Scanln(&bdaystr)

	sdate := s.Split(bdaystr, "-")

	year := convertToNumber(sdate[0])
	month := convertToNumber(sdate[1])
	day := convertToNumber(sdate[2])

	bdDate := makeDate(year, month, day)
	fmt.Println("BD date:", bdDate)

	now := time.Now()
	fmt.Println("Now", now)

	// make function from this?
	datesfrombd := now.Sub(bdDate).Hours()/24
	fmt.Println("You are",datesfrombd, "days old") //round this
	fmt.Println("It's still", "days to your birthday")


//	var nextbd Date
//	var datesBetween

	if  bdDate.Month() < now.Month() {

		fmt.Println("Your next birthday is next year!")
		dyear := int(now.Year()) +1
		next_bd := makeDate(dyear, int(bdDate.Month()), int(bdDate.Day()))
		datesBetween := now.Sub(next_bd).Hours()/24
		fmt.Println("You have about", datesBetween, "dates till your next birthday")

	} 

	fmt.Println("Your next birthday is this year!")
	next_bd := makeDate(int(now.Year()), int(bdDate.Month()), int(bdDate.Day()))
	datesbetween := next_bd.Sub(now).Hours()/24
	fmt.Println("You have about", datesbetween, "dates till your next birthday")


}

