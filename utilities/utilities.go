package utilities

import (
	"log"
	"strconv"
	"strings"
)
func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func VersionToEntry(version string) int {
	segments := strings.Split(version, ".")

	major, err := strconv.Atoi(segments[0])
	minor, err := strconv.Atoi(segments[1])

	if err != nil {
		log.Fatal(err)
	}

	entry := major + minor

	return entry
}