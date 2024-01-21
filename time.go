package main

import (
	"strconv"
	"time"
)

func GetCurrentTime() string {
	timestamp := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	return timestamp
}

