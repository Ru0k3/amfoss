package main

import "fmt"

func PrintGanttChart(slices []ExecutionSlice) {
	fmt.Println("\n=========================================")
	fmt.Println("          Simulation Results")
	fmt.Println("=========================================")
	fmt.Println("Gantt Chart:")

	for _, s := range slices {
		fmt.Printf("| %s (%d-%d) ", s.ProcessID, s.StartTime, s.EndTime)
	}
	fmt.Println("|")
}

func PrintResults(processes []Process) {
	fmt.Println("\n--------------------------------------------------")
	fmt.Printf("%-10s %-10s %-10s %-10s %-10s\n", "PID", "Arrival", "Burst", "Waiting", "Turnaround")
	fmt.Println("--------------------------------------------------")

	var totalWaiting, totalTurnaround float64
	for _, p := range processes {
		fmt.Printf("%-10s %-10d %-10d %-10d %-10d\n", p.ID, p.ArrivalTime, p.BurstTime, p.WaitingTime, p.TurnaroundTime)
		totalWaiting += float64(p.WaitingTime)
		totalTurnaround += float64(p.TurnaroundTime)
	}
	fmt.Println("--------------------------------------------------")

	if n := float64(len(processes)); n > 0 {
		fmt.Printf("Average Waiting Time: %.2f\n", totalWaiting/n)
		fmt.Printf("Average Turnaround Time: %.2f\n", totalTurnaround/n)
	}
	fmt.Println("=========================================")
}
