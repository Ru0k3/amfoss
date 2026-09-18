package main

func CalculateRR(processes []Process, quantum int) ([]ExecutionSlice, []Process) {
    n := len(processes)
    remainingTime := make([]int, n)
    
    for i := 0; i < n; i++ {
        remainingTime[i] = processes[i].BurstTime
    }

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if processes[i].ArrivalTime > processes[j].ArrivalTime {
                processes[i], processes[j] = processes[j], processes[i]
                remainingTime[i], remainingTime[j] = remainingTime[j], remainingTime[i]
            }
        }
    }

    var slices []ExecutionSlice
    var completed []Process
    var queue []int
    currentTime := 0
    inQueue := make([]bool, n)
    completedCount := 0

    for i := 0; i < n; i++ {
        if processes[i].ArrivalTime <= currentTime {
            queue = append(queue, i)
            inQueue[i] = true
        }
    }

    for completedCount < n {
        if len(queue) == 0 {
            nextArrival := -1
            for i := 0; i < n; i++ {
                if remainingTime[i] > 0 {
                    if nextArrival == -1 || processes[i].ArrivalTime < processes[nextArrival].ArrivalTime {
                        nextArrival = i
                    }
                }
            }
            currentTime = processes[nextArrival].ArrivalTime
            queue = append(queue, nextArrival)
            inQueue[nextArrival] = true
            continue
        }

        idx := queue[0]
        queue = queue[1:]
        inQueue[idx] = false

        execTime := quantum
        if remainingTime[idx] < quantum {
            execTime = remainingTime[idx]
        }

        startTime := currentTime
        endTime := currentTime + execTime

        slices = append(slices, ExecutionSlice{
            ProcessID: processes[idx].ID,
            StartTime: startTime,
            EndTime:   endTime,
        })

        remainingTime[idx] -= execTime
        currentTime = endTime

        for i := 0; i < n; i++ {
            if i != idx && !inQueue[i] && remainingTime[i] > 0 && processes[i].ArrivalTime <= currentTime {
                queue = append(queue, i)
                inQueue[i] = true
            }
        }

        if remainingTime[idx] > 0 {
            queue = append(queue, idx)
            inQueue[idx] = true
        } else {
            processes[idx].TurnaroundTime = currentTime - processes[idx].ArrivalTime
            processes[idx].WaitingTime = processes[idx].TurnaroundTime - processes[idx].BurstTime
            completed = append(completed, processes[idx])
            completedCount++
        }
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
