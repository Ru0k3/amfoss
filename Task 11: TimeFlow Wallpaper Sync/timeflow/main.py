import time
import os
import sys
import threading
from timeflow.config import Config
from timeflow.renderer import Renderer
from timeflow.wallpaper_setter import WallpaperSetter
from timeflow.file_monitor import FileMonitor

class TimeFlowApp:
    def __init__(self, config_path=None):
        self.config = Config(config_path)
        self.renderer = Renderer(self.config)
        self.setter = WallpaperSetter()
        self.cache_dir = os.path.expanduser("~/.timeflow/cache")
        os.makedirs(self.cache_dir, exist_ok=True)
        self.wallpaper_path = os.path.join(self.cache_dir, "wallpaper.png")
        
        self.running = True

    def get_content(self):
        path = self.config.get("file_path")
        if not os.path.exists(path):
            return f"File not found: {path}\nPlease create this file to show notes."
        
        try:
            with open(path, 'r', encoding='utf-8') as f:
                content = f.read()
                return content if content.strip() else "No content yet — add notes to your file."
        except Exception as e:
            return f"Error reading file: {e}"

    def update_wallpaper(self):
        content = self.get_content()
        img = self.renderer.render(content)
        self.renderer.save(img, self.wallpaper_path)
        self.setter.set_wallpaper(self.wallpaper_path)

    def clock_loop(self):
        while self.running:
            self.update_wallpaper()
            time.sleep(1)

    def run(self):
        print("TimeFlow Wallpaper Sync starting...")
        
        # Initial render
        self.update_wallpaper()
        
        # Start file monitor
        file_path = self.config.get("file_path")
        self.monitor = FileMonitor(file_path, self.update_wallpaper)
        self.monitor.start()
        
        # Start clock thread
        self.clock_thread = threading.Thread(target=self.clock_loop, daemon=True)
        self.clock_thread.start()
        
        try:
            while True:
                time.sleep(1)
        except KeyboardInterrupt:
            print("\nShutting down...")
            self.running = False
            self.monitor.stop()

if __name__ == "__main__":
    app = TimeFlowApp()
    app.run()
