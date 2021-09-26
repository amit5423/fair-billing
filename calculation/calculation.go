package calculation

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"fair-billing/stack"
)

var layout string = "15:04:05"

type billingReport struct {
	stack         *stack.Stack
	TotalSession  int
	TotalDuration float64
}

func validLog(str string) bool {
	logLine := strings.Split(str, " ")
	if len(logLine) == 3 {
		_, err := time.Parse(layout, logLine[0])
		if err != nil {
			return false
		}
		if !regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString(logLine[1]) {
			return false
		}
		status := []string{"Start", "End"}
		for _, item := range status {
			if item == logLine[2] {
				return true
			}
		}
	}
	return false
}

func sortedKeys(m map[string]*billingReport) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

func Billing(file *os.File) ([]string, map[string]*billingReport, error) {
	var err error
	var earliestTime, latestTime, sessionTime time.Time

	scanner := bufio.NewScanner(file)

	report := make(map[string]*billingReport)

	index := 1

	for scanner.Scan() {
		if validLog(strings.Trim(scanner.Text(), " ")) {

			logLine := strings.Split(strings.Trim(scanner.Text(), " "), " ")
			sessionTime, err = time.Parse(layout, logLine[0])
			if err != nil {
				return nil, nil, err
			}

			user := logLine[1]
			status := logLine[2]

			if index == 1 {
				earliestTime = sessionTime
			}

			if _, ok := report[user]; !ok {
				s := stack.New()
				rep := &billingReport{
					stack: s,
				}
				report[user] = rep
				if status == "Start" {
					rep.stack.Push(status, sessionTime)
				} else if status == "End" {
					rep.TotalDuration += sessionTime.Sub(earliestTime).Seconds()
					rep.TotalSession += 1
				}
			} else {
				rep := report[user]
				if status == "Start" {
					rep.stack.Push(status, sessionTime)
				} else if status == "End" {
					if rep.stack.Len() == 0 {
						rep.TotalDuration += sessionTime.Sub(earliestTime).Seconds()
						rep.TotalSession += 1
					} else {
						_, time_stamp := rep.stack.Pop()
						rep.TotalDuration += sessionTime.Sub(time_stamp).Seconds()
						rep.TotalSession += 1
					}
				}
			}
		}

		index++
	}

	latestTime = sessionTime
	for _, value := range report {
		for value.stack.Len() != 0 {
			_, time_stamp := value.stack.Pop()
			value.TotalDuration += latestTime.Sub(time_stamp).Seconds()
			value.TotalSession += 1
		}
	}
	keys := sortedKeys(report)
	return keys, report, nil
}
