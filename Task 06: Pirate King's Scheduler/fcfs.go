package main

import "sort"

func CalculateFCFS(processes []Process) ([]ExecutionSlice, []Process) {
	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	var slices []ExecutionSlice
	currentTime := 0

	for i := range processes {
		if currentTime < processes[i].ArrivalTime {
			currentTime = processes[i].ArrivalTime
		}

		startTime := currentTime
		endTime := currentTime + processes[i].BurstTime

		slices = append(slices, ExecutionSlice{
			ProcessID: processes[i].ID,
			StartTime: startTime,
			EndTime:   endTime,
		})

		processes[i].TurnaroundTime = endTime - processes[i].ArrivalTime
		processes[i].WaitingTime = processes[i].TurnaroundTime - processes[i].BurstTime
		currentTime = endTime
	}

	return slices, processes
}
