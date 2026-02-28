package swiftscheduling

import (
	"strconv"
	"strings"
	"time"
)

func DeliveryDate(start, delivery string) string {
	layout := "2006-01-02T15:04:05"
	startTime, _ := time.Parse(layout, start)
	var deliveryTime time.Time
	if delivery == "NOW" {
		deliveryTime = startTime.Add(2 * time.Hour)
	} else if delivery == "ASAP" {
		if startTime.Hour() < 13 {
			deliveryTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 17, 0, 0, 0, &time.Location{})
		} else {
			deliveryTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 13, 0, 0, 0, &time.Location{}).Add(24 * time.Hour)
		}
	} else if delivery == "EOW" {
		weekday := startTime.Weekday()
		if weekday >= time.Monday && weekday <= time.Wednesday {
			deliveryTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 17, 0, 0, 0, &time.Location{})
			for deliveryTime.Weekday() != time.Friday {
				deliveryTime = deliveryTime.Add(24 * time.Hour)
			}
		} else {
			deliveryTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 20, 0, 0, 0, &time.Location{})
			for deliveryTime.Weekday() != time.Sunday {
				deliveryTime = deliveryTime.Add(24 * time.Hour)
			}
		}
	} else if strings.HasSuffix(delivery, "M") {
		m, _ := strconv.Atoi(strings.TrimSuffix(delivery, "M"))
		yearOffset := 0
		if startTime.Month() >= time.Month(m) {
			yearOffset += 1
		}
		deliveryTime = time.Date(startTime.Year()+yearOffset, time.Month(m), 1, 8, 0, 0, 0, &time.Location{})
		for deliveryTime.Weekday() == time.Sunday || deliveryTime.Weekday() > time.Friday {
			deliveryTime = deliveryTime.Add(24 * time.Hour)
		}
	} else if strings.HasPrefix(delivery, "Q") {
		quarter, _ := strconv.Atoi(strings.TrimPrefix(delivery, "Q"))
		current := (startTime.Month()-1)/3 + 1
		yearOffset := 0
		if quarter == 4 || current > time.Month(quarter) {
			yearOffset += 1
		}
		month := []time.Month{time.April, time.July, time.October, time.January}[quarter-1]
		deliveryTime = time.Date(startTime.Year()+yearOffset, month, 1, 8, 0, 0, 0, &time.Location{})
		deliveryTime = deliveryTime.Add(-24 * time.Hour)
		for deliveryTime.Weekday() == time.Sunday || deliveryTime.Weekday() > time.Friday {
			deliveryTime = deliveryTime.Add(-24 * time.Hour)
		}
	}

	return deliveryTime.Format(layout)
}
