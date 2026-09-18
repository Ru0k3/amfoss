package main

func CalculateRR(processes []Process, quantum int) ([]ExecutionSlice, []Process) {
    n := len(processes)
    remTime := make([]int, n)
    
    // Sort by arrival time first
    for i := 0; i < n; i++ {
        remTime[i] = processes[i].BurstTime
        for j := i + 1; j < n; j++ {
            if processes[i].ArrivalTime > processes[j].ArrivalTime {
                processes[i], processes[j] = processes[j], processes[i]
                remTime[i], remTime[j] = remTime[j], remTime[i]
            }
        }
    }

    var slices []ExecutionSlice
    var completed []Process
    var queue []int
    currentTime := 0
    completedCount := 0

    // Start by putting the first arrived process(es) into the queue
    for i := 0; i < n; i++ {
        if processes[i].ArrivalTime <= currentTime {
            queue = append(queue, i)
        }
    }

    for completedCount < n {
        // STEP 2: If no process is in the queue, CPU sits idle.
        if len(queue) == 0 {
            nextProc := -1
            for i := 0; i < n; i++ {
                if remTime[i] > 0 {
                    if nextProc == -1 || processes[i].ArrivalTime < processes[nextProc].ArrivalTime {
                        nextProc = i
                    }
                }
            }
            currentTime = processes[nextProc].ArrivalTime // Fast forward time
            queue = append(queue, nextProc)
        }

        // Pop the first element from the queue
        idx := queue[0]
        queue = queue[1:]

        // STEP 3: If rr_time is bigger than burst time, only execute for remaining burst time
        execTime := quantum
        if remTime[idx] < quantum {
            execTime = remTime[idx]
        }

        startTime := currentTime
        endTime := currentTime + execTime

        slices = append(slices, ExecutionSlice{
            ProcessID: processes[idx].ID,
            StartTime: startTime,
            EndTime:   endTime,
        })

        // STEP 1: Subtract burst time - rr time
        remTime[idx] -= execTime
        currentTime = endTime

        // Check if any new process arrived while we were busy
        for i := 0; i < n; i++ {
            if processes[i].ArrivalTime <= currentTime && remTime[i] > 0 && i != idx {
                // Make sure it's not already in the queue
                alreadyIn := false
                for _, q := range queue {
                    if q == i {
                        alreadyIn = true
                    }
                }
                if !alreadyIn {
                    queue = append(queue, i)
                }
            }
        }

        // STEP 1 (continued): Push it back to end of queue if not finished
        if remTime[idx] > 0 {
            queue = append(queue, idx)
        } else {
            // STEP 3 (continued): Process is finished, calculate math
            processes[idx].TurnaroundTime = currentTime - processes[idx].ArrivalTime
            processes[idx].WaitingTime = processes[idx].TurnaroundTime - processes[idx].BurstTime
            completed = append(completed, processes[idx])
            completedCount++
        }
    }

    // Sort the final results by ID for the table
    for i := 0; i < len(completed); i++ {
        for j := i + 1; j < len(completed); j++ {
            if completed[i].ID > completed[j].ID {
                completed[i], completed[j] = completed[j], completed[i]
            }
        }
    }

    return slices, completed
}
