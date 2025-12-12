// Go
package main

import (
	"strconv"
	"strings"
	"sort"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	// group events by timestamp
	byTime := make(map[int][][]string)
	for _, ev := range events {
		t, _ := strconv.Atoi(ev[1])
		byTime[t] = append(byTime[t], ev)
	}

	// gather sorted timestamps
	timestamps := make([]int, 0, len(byTime))
	for t := range byTime {
		timestamps = append(timestamps, t)
	}
	sort.Ints(timestamps)

	mentions := make([]int, numberOfUsers)
	isOnline := make([]bool, numberOfUsers)
	offlineUntil := make([]int, numberOfUsers)
	for i := 0; i < numberOfUsers; i++ {
		isOnline[i] = true
	}

	for _, t := range timestamps {
		evs := byTime[t]

		// 1) expirations
		for i := 0; i < numberOfUsers; i++ {
			if !isOnline[i] && offlineUntil[i] <= t {
				isOnline[i] = true
				offlineUntil[i] = 0
			}
		}

		// 2) OFFLINE events first
		for _, ev := range evs {
			if ev[0] == "OFFLINE" {
				id, _ := strconv.Atoi(ev[2])
				isOnline[id] = false
				offlineUntil[id] = t + 60
			}
		}

		// 3) MESSAGE events
		for _, ev := range evs {
			if ev[0] != "MESSAGE" {
				continue
			}
			tokens := strings.Fields(ev[2])
			for _, token := range tokens {
				if token == "ALL" {
					for i := 0; i < numberOfUsers; i++ {
						mentions[i]++
					}
				} else if token == "HERE" {
					for i := 0; i < numberOfUsers; i++ {
						if isOnline[i] {
							mentions[i]++
						}
					}
				} else if strings.HasPrefix(token, "id") {
					id, _ := strconv.Atoi(token[2:])
					if id >= 0 && id < numberOfUsers {
						mentions[id]++
					}
				}
			}
		}
	}

	return mentions
}
