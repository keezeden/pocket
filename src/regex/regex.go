package regex

import (
	"errors"
	"regexp"
	"strings"
)

func VersionNumberRegex(line string) (string, error) {
	re := regexp.MustCompile(`\d+\.\d+\.\d+`)
	matched := re.Match([]byte(line))

	if !matched {
		return "", errors.New("No match")
	}

	match := re.FindSubmatch([]byte(line))

	return string(match[0]), nil
}

func FeatureNumberRegex(line string) (string, error) {
	re := regexp.MustCompile(`.\d+\.`)
	match := re.FindSubmatch([]byte(line))

	version := strings.Trim(string(match[0]), ".")

	return version, nil
}
