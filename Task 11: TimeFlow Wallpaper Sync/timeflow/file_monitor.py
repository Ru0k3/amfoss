import time
import os
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler

class FileMonitor:
    def __init__(self, file_path, callback):
        self.file_path = os.path.abspath(file_path)
        self.callback = callback
        self.observer = Observer()
        
    def start(self):
        folder = os.path.dirname(self.file_path)
        if not os.path.exists(folder):
            os.makedirs(folder, exist_ok=True)
            
        event_handler = self.ChangeHandler(self.file_path, self.callback)
        self.observer.schedule(event_handler, folder, recursive=False)
        self.observer.start()

    def stop(self):
        self.observer.stop()
        self.observer.join()

    class ChangeHandler(FileSystemEventHandler):
        def __init__(self, file_path, callback):
            self.file_path = file_path
            self.callback = callback
            self.last_triggered = 0
            self.debounce_seconds = 0.5

        def on_modified(self, event):
            if not event.is_directory and os.path.abspath(event.src_path) == self.file_path:
                current_time = time.time()
                if current_time - self.last_triggered > self.debounce_seconds:
                    self.last_triggered = current_time
                    self.callback()
