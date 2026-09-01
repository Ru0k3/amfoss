# Pirate King's Scheduler: CPU Scheduling Simulator

## Project Overview

**Pirate King's Scheduler** is a terminal-based CPU scheduling simulator implemented in Golang. It allows users to simulate and visualize how different CPU scheduling algorithms manage process execution. The simulator calculates key performance metrics such as Waiting Time and Turnaround Time, and presents a text-based Gantt chart for a clear visual representation of process execution over time. This tool is designed for students and developers to understand the fundamental concepts of operating system CPU scheduling.

## How to Run

To run the Pirate King's Scheduler, follow these steps:

1.  **Ensure Go is installed:** Make sure you have Go (Golang) installed on your system. You can download it from [golang.org](https://golang.org/dl/).
2.  **Navigate to the project directory:** Open your terminal or command prompt and change your directory to where `main.go` and other source files are located.
    ```bash
    cd /path/to/pirate-scheduler
    ```
3.  **Run the application:** Execute the following command:
    ```bash
    go run main.go models.go fcfs.go sjf.go rr.go output.go
    ```
4.  **Follow the prompts:** The application will present a menu to select a scheduling algorithm (FCFS, SJF, Round Robin). You will then be prompted to enter details for each process (Process ID, Arrival Time, Burst Time) and, for Round Robin, a Time Quantum.

## Approach

The project is structured into several Go files, each responsible for a specific part of the simulation:

*   **`models.go`**: Defines the core data structures used throughout the simulator:
    *   `Process` struct: Represents a CPU process with attributes like `ID`, `ArrivalTime`, `BurstTime`, `RemainingTime`, `CompletionTime`, `WaitingTime`, and `TurnaroundTime`.
    *   `ExecutionSlice` struct: Captures segments of process execution for the Gantt chart, including `ProcessID`, `StartTime`, and `EndTime`.

*   **`main.go`**: Handles the main program flow, including displaying the menu, collecting user input for processes and algorithm choice, and orchestrating the simulation by calling the appropriate scheduling functions.

*   **`fcfs.go`**: Implements the First Come First Serve (FCFS) scheduling algorithm. Processes are executed in the order of their arrival. If processes arrive at the same time, their initial order in the input array determines their execution order.

*   **`sjf.go`**: Implements the Shortest Job First (SJF) non-preemptive scheduling algorithm. At any given time, the process with the smallest burst time among the arrived processes is executed to completion. Ties are broken by arrival time.

*   **`rr.go`**: Implements the Round Robin (RR) scheduling algorithm. Processes are executed in a cyclic manner, each for a fixed time slice called a 'Time Quantum'. A queue is used to manage processes, and context switching occurs after each quantum or when a process completes.

*   **`output.go`**: Contains functions for formatting and printing the simulation results, including the text-based Gantt chart and the detailed results table with average waiting and turnaround times.

## Resources Used

*   **Golang Documentation:** Official Go language documentation for syntax, standard library usage (`fmt`, `sort`, `os`).
*   **Operating System Concepts Textbooks:** Reference materials for understanding CPU scheduling algorithms (FCFS, SJF, Round Robin) and their performance metrics.

## New Concepts Learned

*   **Go Structs and Slices:** Practical application of Go's composite data types for modeling real-world entities (processes, execution slices).
*   **Sorting in Go:** Utilizing the `sort` package with custom comparison functions for scheduling algorithms.
*   **Command Line Interface (CLI) Development:** Building interactive terminal applications using `fmt.Print`, `fmt.Scan`, and `os.Exit`.
*   **CPU Scheduling Algorithm Logic:** Deepened understanding of the implementation details and nuances of FCFS, SJF, and Round Robin, especially handling arrival times, remaining times, and context switching.
*   **Queue Implementation:** Using Go slices as a basic queue structure for the Round Robin algorithm.
