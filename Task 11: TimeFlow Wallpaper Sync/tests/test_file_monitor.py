import unittest
import os
import time
from timeflow.file_monitor import FileMonitor

class TestFileMonitor(unittest.TestCase):
    def setUp(self):
        self.test_file = "test_watch.txt"
        with open(self.test_file, 'w') as f:
            f.write("initial")
        self.triggered = False

    def tearDown(self):
        if os.path.exists(self.test_file):
            os.remove(self.test_file)

    def callback(self):
        self.triggered = True

    def test_monitor_detects_change(self):
        monitor = FileMonitor(self.test_file, self.callback)
        monitor.start()
        
        # Give it a moment to start
        time.sleep(0.5)
        
        # Modify the file
        with open(self.test_file, 'a') as f:
            f.write("modified")
            
        # Wait for watchdog to detect and debounce
        time.sleep(1.0)
        
        monitor.stop()
        self.assertTrue(self.triggered)

if __name__ == '__main__':
    unittest.main()
