// https://yourbasic.org/golang/format-parse-string-time-date-example/
// https://yourbasic.org/golang/day-month-year-from-time/
// https://gobyexample.com/time
// https://teamtreehouse.com/community/how-to-return-a-map
// https://golang.org/pkg/time/#example_Time_Sub
// https://yourbasic.org/golang/day-month-year-from-time/
// https://ednsquare.com/story/date-and-time-manipulation-golang-with-examples------cU1FjK
// https://yourbasic.org/golang/day-month-year-from-time/
// https://golang.org/pkg/fmt/

package main

import (
	"fmt"
	"time"
	s "strings"
	"strconv"
)


func makeDate(year, month, date int) time.Time {

	return time.Date(year,time.Month(month), date, 0,0,0,0, time.UTC)
}

func convertToNumber(word string)int {
	number, err := strconv.Atoi(word)

	if err != nil {
		return -1
	}

	return number
}

func main() {

	fmt.Println("Let's find out how many days you have till your birthday. First, tell me your birthday:\nyear-month-date")
	var bdaystr string
	fmt.Scanln(&bdaystr)

	sdate := s.Split(bdaystr, "-")

	year := convertToNumber(sdate[0])
	month := convertToNumber(sdate[1])
	day := convertToNumber(sdate[2])

	bdDate := makeDate(year, month, day)
	now := time.Now()

	datesfrombd := now.Sub(bdDate).Hours()/24
	fmt.Printf("You are %.2f days old.\n", datesfrombd)


	if bdDate.Month() == now.Month() {

		if bdDate.Day() < now.Day() {

			fmt.Println("Your next birthday is next year!")
                	dyear := int(now.Year()) +1
                	next_bd := makeDate(dyear, int(bdDate.Month()), int(bdDate.Day()))
              	 	datesBetween := next_bd.Sub(now).Hours()/24
                	fmt.Printf("You have about %.2f days till your next birthday.\n", datesBetween)

		} else {

			fmt.Println("Your next birthday is this year!")
        		next_bd := makeDate(int(now.Year()), int(bdDate.Month()), int(bdDate.Day()))
        		datesBetween := next_bd.Sub(now).Hours()/24
        		fmt.Printf("You have about %.2f days till your next birthday.\n", datesBetween)
		}

	} else if bdDate.Month() < now.Month() {

		fmt.Println("Your next birthday is next year!")
		dyear := int(now.Year()) +1
		next_bd := makeDate(dyear, int(bdDate.Month()), int(bdDate.Day()))
		datesBetween := next_bd.Sub(now).Hours()/24
		fmt.Printf("You have about %.2f days till your next birthday.\n", datesBetween)

	} else { 

	fmt.Println("Your next birthday is this year!")
	next_bd := makeDate(int(now.Year()), int(bdDate.Month()), int(bdDate.Day()))
	datesBetween := next_bd.Sub(now).Hours()/24
	fmt.Printf("You have about %.2f days till your next birthday.\n", datesBetween)

	}
}
