# Grand Line Guardian (Terminal System Monitor)

## Project Overview

**Grand Line Guardian** is a terminal-based, real-time system monitoring tool inspired by the One Piece universe. It displays running system processes as "ships sailing across the Grand Line," providing a live, continuously updating view of system resources. Users can monitor active processes and terminate unresponsive ones directly from the terminal.

## Approach

The application is built using Python 3.x, leveraging the `psutil` library for cross-platform process and system utilization, and the `rich` library for rendering the live-updating terminal UI. The UI is presented as a formatted table, continuously refreshing to show real-time CPU and memory usage for each process. Processes are sorted by CPU usage in descending order.

Interactive features include navigation using arrow keys to select processes and the ability to terminate a selected process by pressing 'k'. Graceful handling of permission errors during process termination is also implemented. The terminal state is restored upon exit (e.g., via CTRL+C).

## Linux Kernel Interface

The tool primarily interacts with the Linux kernel through the `/proc` virtual filesystem, abstracted by the `psutil` library. `psutil` acts as a wrapper, reading information from various files within `/proc` to gather process and system statistics. Specifically:

*   **CPU data:** While `psutil` handles the specifics, it ultimately derives CPU usage information from files like `/proc/[pid]/stat`, which contains process-specific CPU time statistics (user time, system time).
*   **Memory data:** Memory usage, such as Resident Set Size (RSS), is obtained from files like `/proc/[pid]/status`, which provides detailed memory information for each process.

By abstracting these interactions, `psutil` allows the application to remain cross-platform while still providing detailed system insights.

## Resources Used

*   **`psutil` library:** [https://psutil.readthedocs.io/en/latest/](https://psutil.readthedocs.io/en/latest/)
*   **`rich` library:** [https://rich.readthedocs.io/en/stable/](https://rich.readthedocs.io/en/stable/)
*   **Python `termios` and `tty` modules:** For handling non-blocking terminal input.

## New Concepts Learned

During the development of Grand Line Guardian, several key concepts were explored and implemented:

*   **Non-blocking terminal inputs:** Implementing real-time interactivity in a terminal application requires handling keyboard input without blocking the main rendering loop. This was achieved using Python's `termios` and `tty` modules in conjunction with `select.select` to check for available input.
*   **Reading `/proc` stats:** Gaining a deeper understanding of how system monitoring tools interact with the Linux kernel's `/proc` virtual filesystem to retrieve process-specific CPU and memory statistics.
*   **Calculating CPU deltas:** While `psutil` simplifies this, the underlying principle involves calculating the difference in CPU times between two measurement points to get an accurate percentage of CPU utilization over an interval.
*   **Rich library for TUI:** Effectively utilizing the `rich` library to create a dynamic, visually appealing, and responsive terminal user interface with live updates and styled output.
