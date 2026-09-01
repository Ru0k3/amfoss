package main

import (
	"fmt"
	"testing"
)

func TestSimulation(t *testing.T) {
	fmt.Println("==================================================")
	fmt.Println(" RUNNING 4 CRITICAL AUDIT TEST SCENARIOS")
	fmt.Println("==================================================")

	// Scenario 1: CPU Idle Gap Test
	fmt.Println("\n>>> Scenario 1: CPU Idle Gap Test (All Algorithms)")
	sc1 := []Process{
		{ID: "1", ArrivalTime: 0, BurstTime: 3},
		{ID: "2", ArrivalTime: 10, BurstTime: 4},
	}
	s1_fcfs, c1_fcfs := CalculateFCFS(append([]Process{}, sc1...))
	printTestResult("FCFS (Scenario 1)", s1_fcfs, c1_fcfs)

	// Scenario 2: Round Robin Simultaneous Arrival Tie-Breaker
	fmt.Println("\n>>> Scenario 2: Round Robin Simultaneous Arrival Tie-Breaker")
	sc2 := []Process{
		{ID: "1", ArrivalTime: 0, BurstTime: 5},
		{ID: "2", ArrivalTime: 3, BurstTime: 2},
	}
	s2_rr, c2_rr := CalculateRR(append([]Process{}, sc2...), 3)
	printTestResult("RR Quantum=3 (Scenario 2)", s2_rr, c2_rr)

	// Scenario 3: Non-Preemptive SJF Starvation & Future Arrival
	fmt.Println("\n>>> Scenario 3: Non-Preemptive SJF Starvation & Future Arrival")
	sc3 := []Process{
		{ID: "1", ArrivalTime: 0, BurstTime: 10},
		{ID: "2", ArrivalTime: 2, BurstTime: 2},
		{ID: "3", ArrivalTime: 3, BurstTime: 1},
	}
	s3_sjf, c3_sjf := CalculateSJF(append([]Process{}, sc3...))
	printTestResult("SJF (Scenario 3)", s3_sjf, c3_sjf)
}

func printTestResult(name string, slices []ExecutionSlice, completed []Process) {
	fmt.Printf("\n--- %s ---\n", name)
	PrintGanttChart(slices)
	PrintResults(completed)
}
