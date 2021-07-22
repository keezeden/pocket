package changelog

import (
	"bufio"
	"errors"
	"log"
	"os"

	"github.com/keezeden/pocket/src/regex"
	"github.com/keezeden/pocket/src/utilities"
)

func GetVersion(filepath *string) (string, error) {
	file, err := os.Open(*filepath)
    
	utilities.Check(err)

    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line := scanner.Text()
		matched, err := regex.VersionNumberRegex(line)
		

		if err != nil {
			continue
		}
        
		return string(matched), nil
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return "", errors.New("No changelog version found.")
}