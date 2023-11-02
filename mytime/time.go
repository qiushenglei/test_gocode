package mytime

import (
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func Timestamp() {
	fmt.Println(time.Now())
	a := time.Now().AddDate(-1, 2, 3) //2023-09-18 16:03:21.6799254 +0800 CST m=+0.009288501
	fmt.Println(a)
	fmt.Println(a.Format(timeFormat))
	return
	if a, err := time.Parse(timeFormat, "2023-09-18 16:03:21"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(a)
	}

	t := time.Date(0, 0, 0, 12, 31, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}
	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}

}

func Timer() {
	a := func() {
		fmt.Println("到期")
	}
	time.AfterFunc(10*time.Second, a)
	fmt.Println("start")

}

func Ticker() {
	t := time.NewTicker(time.Second * 5)
	for now := range t.C {
		fmt.Println(now)
	}
}
