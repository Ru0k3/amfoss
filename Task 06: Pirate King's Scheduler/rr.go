package main

import "sort"

func CalculateRR(processes []Process, quantum int) ([]ExecutionSlice, []Process) {
	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	var slices []ExecutionSlice
	var completed []Process
	queue := []int{}
	currentTime := 0
	n := len(processes)
	remainingTime := make([]int, n)
	for i := range processes {
		remainingTime[i] = processes[i].BurstTime
	}

	for i := 0; i < n; i++ {
		if processes[i].ArrivalTime <= currentTime {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		idx := queue[0]
		queue = queue[1:]

		execTime := quantum
		if remainingTime[idx] < quantum {
			execTime = remainingTime[idx]
		}

		startTime := currentTime
		endTime := currentTime + execTime

		slices = append(slices, ExecutionSlice{
			ProcessID: processes[idx].ID,
			StartTime: startTime,
			EndTime:   endTime,
		})

		remainingTime[idx] -= execTime
		currentTime = endTime

		for i := 0; i < n; i++ {
			if i != idx && processes[i].ArrivalTime <= currentTime && remainingTime[i] > 0 {
				alreadyInQueue := false
				for _, q := range queue {
					if q == i {
						alreadyInQueue = true
						break
					}
				}
				if !alreadyInQueue {
					queue = append(queue, i)
				}
			}
		}

		if remainingTime[idx] > 0 {
			queue = append(queue, idx)
		} else {
			processes[idx].TurnaroundTime = currentTime - processes[idx].ArrivalTime
			processes[idx].WaitingTime = processes[idx].TurnaroundTime - processes[idx].BurstTime
			completed = append(completed, processes[idx])
		}
	}

	sort.SliceStable(completed, func(i, j int) bool {
		return completed[i].ID < completed[j].ID
	})

	return slices, completed
}
