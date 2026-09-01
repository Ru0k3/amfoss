package main

import "sort"

func CalculateSJF(processes []Process) ([]ExecutionSlice, []Process) {
	var completed []Process
	var slices []ExecutionSlice
	currentTime := 0
	n := len(processes)
	isCompleted := make([]bool, n)
	completedCount := 0

	for completedCount < n {
		minIndex := -1

		for i, p := range processes {
			if isCompleted[i] || p.ArrivalTime > currentTime {
				continue
			}

			if minIndex == -1 || p.BurstTime < processes[minIndex].BurstTime {
				minIndex = i
				continue
			}

			// Tie-breaker: earlier arrival time; if arrival times match, input order is preserved
			if p.BurstTime == processes[minIndex].BurstTime && p.ArrivalTime < processes[minIndex].ArrivalTime {
				minIndex = i
			}
		}

		if minIndex == -1 {
			// CPU is idle; find the next earliest process arrival
			nextArrival := -1
			for i, p := range processes {
				if isCompleted[i] {
					continue
				}
				if nextArrival == -1 || p.ArrivalTime < processes[nextArrival].ArrivalTime {
					nextArrival = i
				}
			}

			if nextArrival != -1 && processes[nextArrival].ArrivalTime > currentTime {
				slices = append(slices, ExecutionSlice{
					ProcessID: "IDLE",
					StartTime: currentTime,
					EndTime:   processes[nextArrival].ArrivalTime,
				})
				currentTime = processes[nextArrival].ArrivalTime
			}
			continue
		}

		p := &processes[minIndex]
		startTime := currentTime
		endTime := currentTime + p.BurstTime

		if startTime < endTime {
			slices = append(slices, ExecutionSlice{
				ProcessID: p.ID,
				StartTime: startTime,
				EndTime:   endTime,
			})
		}

		p.CompletionTime = endTime
		p.TurnaroundTime = p.CompletionTime - p.ArrivalTime
		p.WaitingTime = p.TurnaroundTime - p.BurstTime

		isCompleted[minIndex] = true
		completedCount++
		completed = append(completed, *p)
		currentTime = endTime
	}

	sort.SliceStable(completed, func(i, j int) bool {
		return completed[i].ID < completed[j].ID
	})

	return slices, completed
}
