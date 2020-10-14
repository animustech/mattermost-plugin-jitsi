package main

import (
	"strconv"
	"time"

	"crypto/rand"
	"math/big"
	"strings"

	"github.com/google/uuid"
	"github.com/mattermost/mattermost-server/v5/mlog"
)

var LETTERS = []rune("abcdefghijklmnopqrstuvwxyz")

func randomInt(max int) int {
	value, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		mlog.Error("Error generating random number", mlog.Err(err))
		panic(err.Error())
	}
	return int(value.Int64())
}

func randomString(runes []rune, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[randomInt(len(runes))]
	}
	return string(b)
}

func generateEnglishName(delimiter string) string {
	var components = []string{
		"Animus",
		strconv.FormatInt(time.Now().UnixNano()/1000000, 10),
	}
	return strings.Join(components, delimiter)
}

func generateEnglishTitleName() string {
	return generateEnglishName("")
}

func generateUUIDName() string {
	id := uuid.New()
	return (id.String())
}

func generateTeamChannelName(teamName string, channelName string) string {
	name := teamName
	if name != "" {
		name += "-"
	}
	name += channelName
	name += "-" + randomString(LETTERS, 10)

	return name
}

func generatePersonalMeetingName(username string) string {
	return username + "-" + randomString(LETTERS, 20)
}
