package main

type Process struct {
	ID             string
	ArrivalTime    int
	BurstTime      int
	WaitingTime    int
	TurnaroundTime int
}

type ExecutionSlice struct {
	ProcessID string
	StartTime int
	EndTime   int
}
