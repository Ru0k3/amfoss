package main

func CalculateSJF(processes []Process) ([]ExecutionSlice, []Process) {
    var completed []Process
    var slices []ExecutionSlice
    currentTime := 0
    n := len(processes)
    isCompleted := make([]bool, n)
    completedCount := 0

    for completedCount < n {
        minIndex := -1
        for j := 0; j < n; j++ {
            if !isCompleted[j] && processes[j].ArrivalTime <= currentTime {
                if minIndex == -1 || processes[j].BurstTime < processes[minIndex].BurstTime {
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
        completedCount++
    }

    for i := 0; i < len(completed); i++ {
        for j := i + 1; j < len(completed); j++ {
            if completed[i].ID > completed[j].ID {
                completed[i], completed[j] = completed[j], completed[i]
            }
        }
    }

    return slices, completed
}
