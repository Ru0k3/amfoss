Here is a simplified, straight-to-the-point version of your README, updated to reflect the removal of arrow keys and focusing purely on what the app does and how to run it:

```markdown
# ⚓ Grand Line Guardian (Terminal System Monitor)

A lightweight, flicker-free terminal system monitor built in Python using `psutil` and `rich`. It features PID cursor locking and stable sorting so you can navigate and terminate processes effortlessly.

## 📌 Features

- **PID Cursor Locking:** Your selection (`▶`) stays glued to a specific process by its PID, even as the list updates every 0.5 seconds.
- **Smart Sorting:** Processes are sorted by **Memory Usage % first**, then **CPU Usage %** (descending), keeping the heaviest processes at the top.
- **Flicker-Free UI:** Displays live System CPU %, RAM %, and a color-coded process table inside a clean, bordered panel.
- **Interactive Controls:**
  - `W` / `S` : Move selection up or down.
  - `K` : Terminate the selected process.
  - `Q` / `Ctrl+C` : Exit cleanly and restore terminal settings.

## 🚀 Setup & Run

```bash
# 1. Navigate to the project directory
cd "Task 05: Grand Line Guardian"

# 2. Create a virtual environment and install dependencies
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt

# 3. Launch the System Monitor
python3 monitor.py
```
```