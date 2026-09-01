package main

import (
	"fmt"
	"os"
)

func printMenu() {
	fmt.Println("=========================================")
	fmt.Println("   Pirate King's CPU Scheduler Simulator")
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

func getProcessInput(isRR bool) ([]Process, int) {
	var numProcesses int
	for {
		fmt.Print("\nEnter number of processes (> 0): ")
		fmt.Scan(&numProcesses)
		if numProcesses > 0 {
			break
		}
		fmt.Println("Invalid input. Number of processes must be at least 1.")
	}

	processes := make([]Process, numProcesses)
	for i := 0; i < numProcesses; i++ {
		fmt.Printf("\nEnter details for Process %d:\n", i+1)
		fmt.Print("  Process ID: ")
		fmt.Scan(&processes[i].ID)

		for {
			fmt.Print("  Arrival Time (>= 0): ")
			fmt.Scan(&processes[i].ArrivalTime)
			if processes[i].ArrivalTime >= 0 {
				break
			}
			fmt.Println("  Invalid input. Arrival time cannot be negative.")
		}

		for {
			fmt.Print("  Burst Time (>= 0): ")
			fmt.Scan(&processes[i].BurstTime)
			if processes[i].BurstTime >= 0 {
				break
			}
			fmt.Println("  Invalid input. Burst time cannot be negative.")
		}

		processes[i].RemainingTime = processes[i].BurstTime
	}

	timeQuantum := 0
	if isRR {
		for {
			fmt.Print("\nEnter Time Quantum (> 0): ")
			fmt.Scan(&timeQuantum)
			if timeQuantum > 0 {
				break
			}
			fmt.Println("Invalid input. Time quantum must be a positive integer.")
		}
	}

	return processes, timeQuantum
}

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

		isRR := (choice == 3)
		processes, quantum := getProcessInput(isRR)

		var slices []ExecutionSlice
		var completed []Process

		switch choice {
		case 1:
			slices, completed = CalculateFCFS(processes)
		case 2:
			slices, completed = CalculateSJF(processes)
		case 3:
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
