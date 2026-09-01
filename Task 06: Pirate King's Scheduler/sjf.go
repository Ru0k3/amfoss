package main

import "sort"

func CalculateSJF(processes []Process) ([]ExecutionSlice, []Process) {
	var completed []Process
	var slices []ExecutionSlice
	currentTime := 0
	n := len(processes)
	isCompleted := make([]bool, n)

	for i := 0; i < n; i++ {
		minIndex := -1
		for j, p := range processes {
			if !isCompleted[j] && p.ArrivalTime <= currentTime {
				if minIndex == -1 || p.BurstTime < processes[minIndex].BurstTime {
					minIndex = j
				}
			}
		}

		if minIndex == -1 {
			currentTime++
			continue
		}

		startTime := currentTime
		endTime := currentTime + processes[minIndex].BurstTime

		slices = append(slices, ExecutionSlice{
			ProcessID: processes[minIndex].ID,
			StartTime: startTime,
			EndTime:   endTime,
		})

		processes[minIndex].TurnaroundTime = endTime - processes[minIndex].ArrivalTime
		processes[minIndex].WaitingTime = processes[minIndex].TurnaroundTime - processes[minIndex].BurstTime
		isCompleted[minIndex] = true
		completed = append(completed, processes[minIndex])
		currentTime = endTime
	}

	sort.SliceStable(completed, func(i, j int) bool {
		return completed[i].ID < completed[j].ID
	})

	return slices, completed
}
