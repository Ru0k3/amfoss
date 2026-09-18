package main

func CalculateFCFS(processes []Process) ([]ExecutionSlice, []Process) {
    var slices []ExecutionSlice
    cpuTime := 0

    for i := 0; i < len(processes); i++ {
        for j := i + 1; j < len(processes); j++ {
            if processes[i].ArrivalTime > processes[j].ArrivalTime {
                processes[i], processes[j] = processes[j], processes[i]
            }
        }
    }

    for i := 0; i < len(processes); i++ {
        p := processes[i]

        if cpuTime < p.ArrivalTime {
            cpuTime = p.ArrivalTime
        }

        startTime := cpuTime
        endTime := cpuTime + p.BurstTime

        slices = append(slices, ExecutionSlice{
            ProcessID: p.ID,
            StartTime: startTime,
            EndTime:   endTime,
        })

        p.TurnaroundTime = endTime - p.ArrivalTime
        p.WaitingTime = p.TurnaroundTime - p.BurstTime
        
        processes[i] = p
        cpuTime = endTime
    }

    return slices, processes
}
