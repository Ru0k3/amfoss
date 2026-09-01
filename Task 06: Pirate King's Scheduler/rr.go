package main

import "sort"

func CalculateRR(processes []Process, quantum int) ([]ExecutionSlice, []Process) {
	if quantum <= 0 {
		quantum = 1
	}

	// Stable sort by arrival time to preserve input order on ties
	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	var completed []Process
	var slices []ExecutionSlice
	var queue []int

	n := len(processes)
	currentTime := 0
	inQueue := make([]bool, n)
	completedCount := 0

	// Enqueue processes arriving at time 0
	for i, p := range processes {
		if p.ArrivalTime <= currentTime {
			queue = append(queue, i)
			inQueue[i] = true
		}
	}

	for completedCount < n {
		if len(queue) == 0 {
			nextArrival := -1
			for i, p := range processes {
				if inQueue[i] || p.RemainingTime <= 0 {
					continue
				}
				if nextArrival == -1 || p.ArrivalTime < processes[nextArrival].ArrivalTime {
					nextArrival = i
				}
			}

			if nextArrival == -1 {
				break
			}

			arrivalTime := processes[nextArrival].ArrivalTime
			if arrivalTime > currentTime {
				slices = append(slices, ExecutionSlice{
					ProcessID: "IDLE",
					StartTime: currentTime,
					EndTime:   arrivalTime,
				})
				currentTime = arrivalTime
			}

			for i, p := range processes {
				if !inQueue[i] && p.ArrivalTime <= currentTime && p.RemainingTime > 0 {
					queue = append(queue, i)
					inQueue[i] = true
				}
			}
		}

		if len(queue) == 0 {
			break
		}

		idx := queue[0]
		queue = queue[1:]
		p := &processes[idx]

		execTime := quantum
		if p.RemainingTime < quantum {
			execTime = p.RemainingTime
		}

		startTime := currentTime
		endTime := currentTime + execTime

		if startTime < endTime {
			slices = append(slices, ExecutionSlice{
				ProcessID: p.ID,
				StartTime: startTime,
				EndTime:   endTime,
			})
		}

		p.RemainingTime -= execTime
		currentTime = endTime

		// 1. Enqueue newly arrived processes during execution (OS Rule: New arrivals before preempted)
		for i, pOther := range processes {
			if !inQueue[i] && pOther.ArrivalTime <= currentTime && pOther.RemainingTime > 0 {
				queue = append(queue, i)
				inQueue[i] = true
			}
		}

		// 2. Re-enqueue preempted process if it still has remaining burst time
		if p.RemainingTime > 0 {
			queue = append(queue, idx)
			continue
		}

		// Process finished
		p.CompletionTime = currentTime
		p.TurnaroundTime = p.CompletionTime - p.ArrivalTime
		p.WaitingTime = p.TurnaroundTime - p.BurstTime
		completed = append(completed, *p)
		completedCount++
	}

	sort.SliceStable(completed, func(i, j int) bool {
		return completed[i].ID < completed[j].ID
	})

	return slices, completed
}
