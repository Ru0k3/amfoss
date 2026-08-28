import json
import os
from pathlib import Path

DEFAULT_CONFIG = {
    "file_path": "~/notes/today.txt",
    "theme": "dark",
    "font_family": "DejaVuSans",  # Default Linux font
    "font_size": 28,
    "bg_color": "#0F172A",
    "text_color": "#E2E8F0",
    "clock_color": "#38BDF8",
    "time_format": "24h",
    "layout": "right-sidebar",
    "polling_fallback_seconds": 5,
    "restore_on_exit": False
}

class Config:
    def __init__(self, config_path=None):
        self.config_path = config_path or Path.home() / ".timeflow" / "config.json"
        self.data = DEFAULT_CONFIG.copy()
        self.load()

    def load(self):
        if os.path.exists(self.config_path):
            try:
                with open(self.config_path, 'r') as f:
                    user_config = json.load(f)
                    self.data.update(user_config)
            except Exception as e:
                print(f"Error loading config: {e}")
        
        # Expand user path
        self.data['file_path'] = os.path.expanduser(self.data['file_path'])

    def get(self, key):
        return self.data.get(key)
