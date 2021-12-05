package main

import (
	"time"
	"fmt"
	"sort"
)

type eventType int

const (
	guardChange eventType = iota
	fallAsleep
	wakeUp
)

type event struct {
	time time.Time
	et eventType
	guardID int
}
	
type sleepSchedule struct {
	schedule map[int]int
	amount int
}

func puzzle1(input []string) int {
	events := make([]event, 0)
	for _, str := range input {
		event := event{}
		runes := []rune(str)
		timesstr := string(runes[1:17])
		rest := string(runes[19:])
		
		timestamp, err := time.Parse("2006-01-02 15:04", timesstr)
		if err != nil {
			panic(err)
		}
		event.time = timestamp
		if rest == "falls asleep" {
			event.et = fallAsleep
		} else if rest == "wakes up" {
			event.et = wakeUp
		} else {
			event.et = guardChange
			_, err = fmt.Sscanf(rest, "Guard #%d begins shift", &event.guardID)
			if err != nil {
				panic(err)
			}
		}
		events = append(events, event)
	}

	sort.Slice(events, func(i, j int) bool {
    return events[i].time.Before(events[j].time)
	})

	guards := make(map[int]*sleepSchedule)
	current := 0
	sleepMin := 0
	for _, ev := range events {
		if ev.et == guardChange {
			current = ev.guardID
			if _, ok := guards[current]; !ok {
				guards[current] = &sleepSchedule{make(map[int]int), 0}
			}
		} else if ev.et == fallAsleep {
			sleepMin = ev.time.Minute()
		} else if ev.et == wakeUp {
			wakeMin := ev.time.Minute()
			for i := sleepMin; i < wakeMin; i++ {
				guards[current].schedule[i]++
				guards[current].amount++
			}
		}
	}

	maxGuard := 0
	maxValue := 0
	for i, guard := range guards {
		if guard.amount > maxValue {
			maxValue = guard.amount
			maxGuard = i
		}
	}

	maxMin := 0
	maxValue = 0
	for i, val := range guards[maxGuard].schedule {
		if val > maxValue {
			maxValue = val
			maxMin = i
		}
	}

	return maxGuard * maxMin
}

func puzzle2(input []string) int {
	events := make([]event, 0)
	for _, str := range input {
		event := event{}
		runes := []rune(str)
		timesstr := string(runes[1:17])
		rest := string(runes[19:])
		
		timestamp, err := time.Parse("2006-01-02 15:04", timesstr)
		if err != nil {
			panic(err)
		}
		event.time = timestamp
		if rest == "falls asleep" {
			event.et = fallAsleep
		} else if rest == "wakes up" {
			event.et = wakeUp
		} else {
			event.et = guardChange
			_, err = fmt.Sscanf(rest, "Guard #%d begins shift", &event.guardID)
			if err != nil {
				panic(err)
			}
		}
		events = append(events, event)
	}

	sort.Slice(events, func(i, j int) bool {
    return events[i].time.Before(events[j].time)
	})

	guards := make(map[int]*sleepSchedule)
	current := 0
	sleepMin := 0
	for _, ev := range events {
		if ev.et == guardChange {
			current = ev.guardID
			if _, ok := guards[current]; !ok {
				guards[current] = &sleepSchedule{make(map[int]int), 0}
			}
		} else if ev.et == fallAsleep {
			sleepMin = ev.time.Minute()
		} else if ev.et == wakeUp {
			wakeMin := ev.time.Minute()
			for i := sleepMin; i < wakeMin; i++ {
				guards[current].schedule[i]++
				guards[current].amount++
			}
		}
	}

	maxGuard := 0
	maxValue := 0
	maxMin := 0
	for i, guard := range guards {
		for j, val := range guard.schedule {
			if val > maxValue {
				maxValue = val
				maxGuard = i
				maxMin = j
			}
		}
	}

	return maxGuard * maxMin


}


