# TimeFlow Wallpaper Sync

TimeFlow Wallpaper Sync is a Python-based desktop utility that transforms your wallpaper into a live, productivity-focused dashboard.

## Features

- **Live Content Rendering**: Displays content from a text file directly on your desktop.
- **Real-time Clock**: Shows a live digital clock with seconds.
- **Auto-Refresh**: Automatically updates the wallpaper when the source file is modified.
- **Cross-Platform**: Supports Windows, macOS, and Linux.
- **Auto-Scaling**: Gracefully handles long text by scaling fonts or truncating.

## Installation

1. Install dependencies:
   ```bash
   pip install -r requirements.txt
   ```

2. (Linux only) Ensure `gsettings` or `feh` is installed for wallpaper management.

## Usage

Run the application:
```bash
python3 -m timeflow.main
```

By default, it monitors `~/notes/today.txt`. You can configure this in `~/.timeflow/config.json`.

## Configuration

Example `config.json`:
```json
{
  "file_path": "~/notes/today.txt",
  "theme": "dark",
  "font_family": "DejaVuSans",
  "font_size": 28,
  "bg_color": "#0F172A",
  "text_color": "#E2E8F0",
  "clock_color": "#38BDF8",
  "time_format": "24h",
  "layout": "right-sidebar"
}
```

## Running Tests

To run the test suite:
```bash
export PYTHONPATH=$PYTHONPATH:.
python3 -m unittest discover tests
```
