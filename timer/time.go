package timer

import (
	"fmt"
	"time"
)


func example() {
	timestamp1 := GetCurrentTimestamp()
	time.Sleep(time.Second * 2) 
	timestamp2 := GetCurrentTimestamp()

	differenceInMilliseconds := CalculateTimestampDifference(timestamp1, timestamp2)

	fmt.Printf("Difference in milliseconds: %d\n", differenceInMilliseconds)
}

func GetCurrentTimestamp() int64 {
	return time.Now().UTC().UnixNano() / int64(time.Millisecond)
}

func CalculateTimestampDifference(ts1, ts2 int64) int64 {
	return ts2 - ts1
}