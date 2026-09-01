package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		printMenu()
		choice := getAlgorithmChoice()
		if choice == 4 {
			fmt.Println("Fair winds and following seas, Pirate King!")
			os.Exit(0)
		}
		if choice < 1 || choice > 4 {
			fmt.Println("Invalid choice. Try again.")
			continue
		}

		processes := getProcessInput()
		var slices []ExecutionSlice
		var completed []Process

		switch choice {
		case 1:
			slices, completed = CalculateFCFS(processes)
		case 2:
			slices, completed = CalculateSJF(processes)
		case 3:
			quantum := getTimeQuantum()
			slices, completed = CalculateRR(processes, quantum)
		}

		PrintGanttChart(slices)
		PrintResults(completed)

		var again string
		fmt.Print("\nRun another simulation? (y/n): ")
		fmt.Scan(&again)
		if again != "y" && again != "Y" {
			fmt.Println("Fair winds and following seas, Pirate King!")
			break
		}
	}
}

func printMenu() {
	fmt.Println("=========================================")
	fmt.Println(" Pirate King's CPU Scheduler Simulator")
	fmt.Println("=========================================")
	fmt.Println("Select Scheduling Algorithm:")
	fmt.Println("1. First Come First Serve (FCFS)")
	fmt.Println("2. Shortest Job First (SJF - Non-Preemptive)")
	fmt.Println("3. Round Robin (RR)")
	fmt.Println("4. Exit")
	fmt.Print("Enter choice: ")
}

func getAlgorithmChoice() int {
	var choice int
	fmt.Scan(&choice)
	return choice
}

func getProcessInput() []Process {
	var n int
	fmt.Print("\nEnter number of processes (> 0): ")
	fmt.Scan(&n)

	processes := make([]Process, n)
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Process %d:\n", i+1)
		fmt.Print(" Process ID: ")
		fmt.Scan(&processes[i].ID)
		fmt.Print(" Arrival Time (>= 0): ")
		fmt.Scan(&processes[i].ArrivalTime)
		fmt.Print(" Burst Time (>= 0): ")
		fmt.Scan(&processes[i].BurstTime)
	}
	return processes
}

func getTimeQuantum() int {
	var q int
	fmt.Print("\nEnter Time Quantum (> 0): ")
	fmt.Scan(&q)
	return q
}
