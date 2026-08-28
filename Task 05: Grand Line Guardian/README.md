# ⚓ Grand Line Guardian (Terminal System Monitor)

> *"As the navigator of the Straw Hat Pirates, you must keep track of every ship sailing across the Grand Line. Each ship represents a running process in the system, and your mission is to monitor their activity in real time."*

**Grand Line Guardian** is a lightweight, flicker-free terminal system monitoring utility built in Python with `psutil` and `rich`. It features **PID Cursor Locking** and **Stable Sorting** so you can navigate and terminate specific processes effortlessly.

---

## 📌 Features

- **PID Cursor Locking:** When you move your cursor (`▶`) to a specific process (e.g. `spotify` or `antigravity-ide`), the selection lock stays glued to that process PID even as metrics update every 0.5s.
- **Stable Sorting:** Processes are sorted stably by **Memory Usage % first, then CPU Usage %**, keeping application row positions steady on your screen.
- **Flicker-Free Framed View:** Framed within a single clean `Panel` with live System CPU %, RAM %, rounded-box table, and status messages.
- **Interactive Keyboard Controls:**
  - `↑` / `↓` or `W` / `S` : Move process selection indicator up or down.
  - `K` : Terminate the selected process (Ship).
  - `Q` / `Ctrl+C` : Exit cleanly and restore terminal settings.

---

## 🚀 Setup & Run

```bash
# 1. Navigate to project directory
cd "Task 05: Grand Line Guardian"

# 2. Setup virtual environment & dependencies
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt

# 3. Launch System Monitor
python3 monitor.py
```
