// isnotLeapyear
package main

import (
	"fmt"
	"time"
)

func printime() {
	//获取时间戳
	t := time.Now()
	fmt.Printf(t.Format("2006 年 01 月 02 日"))
	switch t.Weekday().String() {
	case "Sunday":
		fmt.Printf("星期天")
	case "Monday":
		fmt.Printf("星期一")
	case "Tuesday":
		fmt.Printf("星期二")
	case "Wednesday":
		fmt.Printf("星期三")
	case "Thursday":
		fmt.Printf("星期四")
	case "Friday":
		fmt.Printf("星期五")
	case "Saturday":
		fmt.Printf("星期六")
	}
	fmt.Println(" ")
}
func main() {
	printime()
	var year int
	fmt.Println("Input a year:")
	fmt.Scanln(&year)
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				// 都能被4、100、400整除，是闰年
				fmt.Printf("% d is a leap year.\n", year)
			} else {
				//能被4、100整出，但不能被400整除，不是闰年
				fmt.Printf("% d is not a leap year.\n", year)
			}
		} else {
			// 能被4整除，但不能被100整除，不是闰年
			fmt.Printf("% d is not a leap year.\n", year)
		}
	} else {
		// 不能被4整除， 不是闰年
		fmt.Printf("% d is not a leap year.\n", year)
	}
}
