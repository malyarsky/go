package main

import (
	"fmt"
	"time"
)

// TimeIn возвращает время в UTC, если name "" или "UTC".
// Возвращает местное время, если name "Local".
// В противном случае name принимается
// за name местоположения в базе данных часового пояса IANA,
// например, "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func main() {
	for _, name := range []string{
		"",
		"Local",
		"Asia/Shanghai",
		"Europe/Moscow",
	} {
		t, err := TimeIn(time.Now(), name)
		if err == nil {
			fmt.Println(t.Location(), t.Format("15:04"))
		} else {
			fmt.Println(name, "<time unknown>")
		}
	}
}
