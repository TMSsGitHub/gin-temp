package utils

import "time"

func GetCurrentTime() time.Time {
	return time.Now()
}

func GetCurrentTs() uint64 {
	return uint64(GetCurrentTime().Unix())
}

func GetCurrentMs() uint64 {
	return uint64(GetCurrentTime().UnixMilli())
}
