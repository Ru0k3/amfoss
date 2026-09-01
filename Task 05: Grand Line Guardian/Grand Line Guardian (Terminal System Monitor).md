```markdown
# Grand Line Guardian — Terminal System Monitor

## Project Overview

**Grand Line Guardian** is a terminal-based system monitor that shows live CPU and memory usage for running processes. 

The process list updates continuously, and processes are sorted primarily by **Memory usage**, and secondarily by **CPU usage** (in descending order), so that the processes consuming the most resources appear at the top. A process can be selected using the `W` and `S` keys and terminated directly from the terminal using `K`.

The project was built to understand how system monitoring works on Linux, how a terminal application can interact with running processes, and how to handle non-blocking terminal input.

## Approach

The application is written in **Python 3** and mainly uses four cooperating layers of libraries:

* `psutil` — The data engine. Fetches process and system information.
* `rich` — The UI layer. Displays the information as a live-updating, color-coded terminal table.
* `select`, `sys`, `tty`, `termios` — The input layer. Captures keystrokes instantly without freezing the UI.
* `time` — Paces the refresh loop.

The main loop repeatedly collects the current process information, updates the CPU and memory values, sorts the processes (highest RAM first, ties broken by CPU), and redraws the table.

For interaction, the user can move through the process list with `W` (up) and `S` (down). Pressing `K` attempts to terminate the currently selected process. Because the list is constantly re-sorting, the application tracks the selected process by its **PID**, ensuring the highlight stays on the correct process even if its position in the list changes.

The terminal settings are safely restored when the program exits using a `try/finally` block, ensuring the user's shell is never left in a broken "raw mode" state, even if the program crashes or is stopped using `CTRL+C`.

## Linux Kernel Interface

On Linux, process information is exposed through the `/proc` virtual filesystem. The application does not read these files directly; instead, `psutil` handles that part.

Some of the information ultimately comes from files such as:

* `/proc/[pid]/stat` — contains process statistics including CPU time.
* `/proc/[pid]/status` — contains information about memory usage and other process details.

`psutil` reads and processes this information and provides it to the Python application in a simpler form.

### CPU Usage
CPU usage is based on the amount of CPU time a process consumes between two measurements. By comparing the CPU time at different points, the application can determine how much CPU the process is using. The first call to `psutil.cpu_percent()` primes the pump, and subsequent calls return the actual percentage.

### Memory Usage
Memory information such as RSS (Resident Set Size) is obtained through the process information exposed by `/proc`, which `psutil` provides through its process APIs.

## Resources Used

* **`psutil`** — Process and system monitoring
  https://psutil.readthedocs.io/en/latest/

* **`rich`** — Terminal formatting and live UI
  https://rich.readthedocs.io/en/stable/

* **Python `termios`, `tty`, and `select`** — Terminal input and terminal mode handling

## New Concepts Learned

### Non-blocking Terminal Input
A normal `input()` call would stop the program until the user enters something. That would not work well for a continuously updating monitor.

I used `termios`, `tty`, and `select.select` to check for keyboard input without stopping the monitoring loop. By using `select` with a `0.0` timeout, the program can "glance" at the keyboard, and if no key is pressed, it immediately continues redrawing the UI. This allows the app to feel responsive while continuously updating in the background. I also learned how to intercept `Ctrl+C` manually since raw mode disables the default terminal behavior.

### Efficient Data Fetching with `psutil`
Instead of manually querying each process attribute one by one, I learned to pass a list of attributes to `psutil.process_iter()`. This prompts `psutil` to fetch the data efficiently and attach it as a cached dictionary (`p.info`) to the process object, making the data collection loop much faster and cleaner.

### Linux `/proc` Filesystem
I learned how Linux exposes process information through the `/proc` virtual filesystem and how tools such as `psutil` use this information to provide process statistics.

### CPU Usage Calculation
CPU usage is not simply a value that can be read once. It is calculated over a period of time by comparing CPU time between two measurements. `psutil` handles most of this calculation, but understanding what happens underneath helped me understand how process monitoring works.

### Building a Terminal UI
Using `rich` made it possible to turn the raw process information into a readable terminal interface. I learned how to use `rich.Live` with an alternate screen buffer (`screen=True`) so the app takes over the terminal like `vim` or `htop`, leaving the shell history clean upon exit. I also learned how to apply dynamic styling (like turning CPU% red if it goes above 50%) and handle pagination/scrolling math when the terminal window is too small to fit all processes.

Overall, this project helped me understand the connection between a Python application, the Linux `/proc` filesystem, running processes, and terminal input handling.
```