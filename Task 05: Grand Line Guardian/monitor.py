import select, sys, termios, time, tty, psutil
from rich import box
from rich.console import Console
from rich.live import Live
from rich.panel import Panel
from rich.table import Table

console = Console()


def get_processes():
    """Fetch active processes sorted stably by Memory usage first, then CPU usage."""
    procs = []
    for p in psutil.process_iter(['pid', 'name', 'cpu_percent', 'memory_percent']):
        try:
            info = p.info
            procs.append({
                'pid': info['pid'],
                'name': info['name'] or 'unknown',
                'cpu_percent': info['cpu_percent'] or 0.0,
                'memory_percent': info['memory_percent'] or 0.0,
            })
        except (psutil.NoSuchProcess, psutil.AccessDenied, psutil.ZombieProcess):
            pass
    procs.sort(key=lambda x: (x['memory_percent'], x['cpu_percent']), reverse=True)
    return procs


def build_view(processes, selected_index, message, sys_cpu, sys_mem):
    """Build system monitor UI inside a framed panel."""
    table = Table(box=box.ROUNDED, expand=True, border_style="cyan", header_style="bold cyan")
    table.add_column(" ", width=2, justify="center")
    table.add_column("PID", style="cyan", width=8, no_wrap=True)
    table.add_column("Ship Name (Process)", style="bold white", no_wrap=True)
    table.add_column("CPU (%)", justify="right", width=10)
    table.add_column("MEM (%)", justify="right", width=10)

    max_rows = max(5, (console.height or 24) - 6)
    start_idx = max(0, min(selected_index - max_rows // 2, max(0, len(processes) - max_rows)))

    for i, proc in enumerate(processes[start_idx:start_idx + max_rows]):
        is_sel = (start_idx + i == selected_index)
        marker = "▶" if is_sel else " "
        style = "bold white on blue" if is_sel else ""

        cpu, mem = proc['cpu_percent'], proc['memory_percent']
        if is_sel:
            cpu_s, mem_s = f"{cpu:.1f}%", f"{mem:.1f}%"
        else:
            c_col = "bold red" if cpu > 50 else ("yellow" if cpu > 15 else "green")
            m_col = "bold red" if mem > 50 else ("yellow" if mem > 15 else "blue")
            cpu_s = f"[{c_col}]{cpu:.1f}%[/{c_col}]"
            mem_s = f"[{m_col}]{mem:.1f}%[/{m_col}]"

        table.add_row(marker, str(proc['pid']), str(proc['name']), cpu_s, mem_s, style=style)

    status = f"  │  [STATUS] {message}" if message else f"  │  Active Fleet: {len(processes)} ships"
    return Panel(
        table,
        title=f"[bold cyan]⚓ GRAND LINE GUARDIAN  │  System CPU: {sys_cpu:.1f}%  │  RAM: {sys_mem:.1f}% ⚓[/bold cyan]",
        subtitle=f"[bold yellow][↑/↓ or W/S] Navigate  │  [K] Terminate Ship  │  [Q] Quit{status}[/bold yellow]",
        border_style="cyan",
        padding=(0, 0),
    )


def read_key():
    """Non-blocking keyboard reader (letters only)."""
    if not select.select([sys.stdin], [], [], 0.0)[0]:
        return None
    
    ch = sys.stdin.read(1)
    key = ch.lower()
    
    if key in ('q', '\x03'): return 'q'  # 'q' or Ctrl+C
    if key == 'w': return 'up'
    if key == 's': return 'down'
    if key == 'k': return 'k'
    
    return None


def kill_process(pid):
    """Terminate process by PID."""
    try:
        psutil.Process(pid).terminate()
        return f"Process {pid} terminated."
    except psutil.NoSuchProcess:
        return f"PID {pid} no longer exists."
    except psutil.AccessDenied:
        return f"Access denied for PID {pid}. Use sudo."
    except Exception as e:
        return f"Error: {e}"


def main():
    selected_index, selected_pid, message = 0, None, ""
    old_settings = termios.tcgetattr(sys.stdin.fileno())

    psutil.cpu_percent(interval=None)
    processes = get_processes()
    last_update = time.time()

    try:
        tty.setraw(sys.stdin.fileno())
        with Live(console=console, auto_refresh=False, screen=True) as live:
            while True:
                now = time.time()
                if now - last_update > 0.5:
                    processes = get_processes()
                    last_update = now

                if not processes:
                    time.sleep(0.05)
                    continue

                pids = [p['pid'] for p in processes]
                if selected_pid in pids:
                    selected_index = pids.index(selected_pid)
                else:
                    selected_index = max(0, min(selected_index, len(processes) - 1))
                    selected_pid = processes[selected_index]['pid']

                key = read_key()
                if key == 'q':
                    break
                elif key == 'up':
                    selected_index = max(0, selected_index - 1)
                    selected_pid = processes[selected_index]['pid']
                elif key == 'down':
                    selected_index = min(len(processes) - 1, selected_index + 1)
                    selected_pid = processes[selected_index]['pid']
                elif key == 'k' and selected_pid:
                    message = kill_process(selected_pid)

                view = build_view(processes, selected_index, message, psutil.cpu_percent(interval=None), psutil.virtual_memory().percent)
                live.update(view, refresh=True)
                time.sleep(0.05)
    finally:
        termios.tcsetattr(sys.stdin.fileno(), termios.TCSADRAIN, old_settings)
        console.print("\n[bold yellow]Exiting Grand Line Guardian. Goodbye![/bold yellow]")


if __name__ == "__main__":
    main()
