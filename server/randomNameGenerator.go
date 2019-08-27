package main

import (
	"strconv"
	"time"
)

func generateRoomWithoutSeparator() string {
	return "Animus" + strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
}
