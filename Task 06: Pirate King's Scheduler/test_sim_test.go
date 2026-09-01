package main

import (
	"fmt"
	"testing"
)

func printTestResult(name string, slices []ExecutionSlice, completed []Process) {
	fmt.Printf("\n--- %s ---\n", name)
	PrintGanttChart(slices)
	PrintResults(completed)
}

func TestSimulation(t *testing.T) {
	fmt.Println("==================================================")
	fmt.Println("   RUNNING 4 CRITICAL AUDIT TEST SCENARIOS")
	fmt.Println("==================================================")

	// SCENARIO 1: CPU Idle Gap Test
	fmt.Println("\n>>> Scenario 1: CPU Idle Gap Test (All Algorithms)")
	sc1 := []Process{
		{ID: "1", ArrivalTime: 0, BurstTime: 3, RemainingTime: 3},
		{ID: "2", ArrivalTime: 10, BurstTime: 4, RemainingTime: 4},
	}
	s1_fcfs, c1_fcfs := CalculateFCFS(append([]Process{}, sc1...))
	printTestResult("FCFS (Scenario 1)", s1_fcfs, c1_fcfs)
	s1_sjf, c1_sjf := CalculateSJF(append([]Process{}, sc1...))
	printTestResult("SJF (Scenario 1)", s1_sjf, c1_sjf)
	s1_rr, c1_rr := CalculateRR(append([]Process{}, sc1...), 3)
	printTestResult("RR Quantum=3 (Scenario 1)", s1_rr, c1_rr)

	// SCENARIO 2: Round Robin Simultaneous Arrival Tie-Breaker (TQ=3)
	fmt.Println("\n>>> Scenario 2: Round Robin Simultaneous Arrival Tie-Breaker")
	sc2 := []Process{
		{ID: "1", ArrivalTime: 0, BurstTime: 5, RemainingTime: 5},
		{ID: "2", ArrivalTime: 3, BurstTime: 2, RemainingTime: 2},
	}
	s2_rr, c2_rr := CalculateRR(append([]Process{}, sc2...), 3)
	printTestResult("RR Quantum=3 (Scenario 2)", s2_rr, c2_rr)

	// SCENARIO 3: Non-Preemptive SJF Starvation & Future Arrival
	fmt.Println("\n>>> Scenario 3: Non-Preemptive SJF Starvation & Future Arrival")
	sc3 := []Process{
		{ID: "1", ArrivalTime: 0, BurstTime: 10, RemainingTime: 10},
		{ID: "2", ArrivalTime: 2, BurstTime: 2, RemainingTime: 2},
		{ID: "3", ArrivalTime: 3, BurstTime: 1, RemainingTime: 1},
	}
	s3_sjf, c3_sjf := CalculateSJF(append([]Process{}, sc3...))
	printTestResult("SJF (Scenario 3)", s3_sjf, c3_sjf)

	// SCENARIO 4: Zero Burst Time Safeguard
	fmt.Println("\n>>> Scenario 4: Zero Burst Time Safeguard")
	sc4 := []Process{
		{ID: "1", ArrivalTime: 0, BurstTime: 0, RemainingTime: 0},
		{ID: "2", ArrivalTime: 0, BurstTime: 3, RemainingTime: 3},
	}
	s4_fcfs, c4_fcfs := CalculateFCFS(append([]Process{}, sc4...))
	printTestResult("FCFS (Scenario 4)", s4_fcfs, c4_fcfs)
	s4_sjf, c4_sjf := CalculateSJF(append([]Process{}, sc4...))
	printTestResult("SJF (Scenario 4)", s4_sjf, c4_sjf)
	s4_rr, c4_rr := CalculateRR(append([]Process{}, sc4...), 2)
	printTestResult("RR Quantum=2 (Scenario 4)", s4_rr, c4_rr)
}
