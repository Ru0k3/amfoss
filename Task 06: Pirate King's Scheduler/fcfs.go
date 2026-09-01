package main

import "sort"

func CalculateFCFS(processes []Process) ([]ExecutionSlice, []Process) {
	// Stable sort by arrival time to preserve input order for ties
	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	var slices []ExecutionSlice
	currentTime := 0

	for i := range processes {
		p := &processes[i]

		if currentTime < p.ArrivalTime {
			slices = append(slices, ExecutionSlice{
				ProcessID: "IDLE",
				StartTime: currentTime,
				EndTime:   p.ArrivalTime,
			})
			currentTime = p.ArrivalTime
		}

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
		currentTime = endTime
	}

	return slices, processes
}
