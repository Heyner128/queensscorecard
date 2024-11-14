package messages

import (
	"bufio"
	"bytes"
	"iter"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ScoreDto struct {
	Name           string
	GameNumber     int
	SecondsToSolve int
	Time           time.Time
}

func scanMessages(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n[")); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func isMessageQueensScore(message string) bool {
	return strings.Contains(message, "lnkd.in/queens") &&
		strings.Contains(message, "Queens")
}

func GetScoresFromExport() iter.Seq[ScoreDto] {
	contents, readErr := os.ReadFile("chat.txt")
	if readErr != nil {
		log.Fatal("failed to read the file")
	}
	scanner := bufio.NewScanner(bytes.NewReader(contents))
	scanner.Split(scanMessages)
	return func(yield func(ScoreDto) bool) {
		for scanner.Scan() {
			if !isMessageQueensScore(scanner.Text()) {
				continue
			}
			if !yield(extractScoreFromMessage(scanner.Text())) {
				return
			}
		}
	}
}

func cleanSpaces(message string) string {
	return strings.ReplaceAll(message, "\u202F", " ")
}

func extractNameFromMessage(message string) string {
	re := regexp.MustCompile("]\\W*([A-Za-zÀ-ÖØ-öø-ÿ\\s]*)\\W*:")
	match := re.FindStringSubmatch(message)
	name := match[1]
	return name
}

func extractTimestampFromMessage(message string) time.Time {
	re := regexp.MustCompile("\\[(.+)]")
	match := re.FindStringSubmatch(message)
	t, _ := time.Parse("02/01/2006 15:04:05", match[1])
	log.Println(t)
	return t
}

func extractGameNumberFromMessage(message string) int {
	re := regexp.MustCompile("Queens\\D*(\\d*)")
	match := re.FindStringSubmatch(message)
	gameNumber, _ := strconv.Atoi(match[1])
	return gameNumber
}

func extractDurationInSecondsFromMessage(message string) int {
	re := regexp.MustCompile("][\\s\\S]*(\\d{1,2}:\\d{1,2})")
	match := re.FindStringSubmatch(message)
	fullTime := strings.Split(match[1], ":")
	minutes, _ := strconv.Atoi(fullTime[0])
	seconds, _ := strconv.Atoi(fullTime[1])
	return minutes*60 + seconds
}

func extractScoreFromMessage(message string) ScoreDto {
	message = cleanSpaces(message)
	return ScoreDto{
		Name:           extractNameFromMessage(message),
		GameNumber:     extractGameNumberFromMessage(message),
		SecondsToSolve: extractDurationInSecondsFromMessage(message),
		Time:           extractTimestampFromMessage(message),
	}
}
