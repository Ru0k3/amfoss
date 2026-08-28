import unittest
import os
import json
from pathlib import Path
from timeflow.config import Config

class TestConfig(unittest.TestCase):
    def setUp(self):
        self.test_config_path = Path("test_config.json")
        if self.test_config_path.exists():
            os.remove(self.test_config_path)

    def tearDown(self):
        if self.test_config_path.exists():
            os.remove(self.test_config_path)

    def test_default_config(self):
        config = Config(self.test_config_path)
        self.assertEqual(config.get("theme"), "dark")
        self.assertEqual(config.get("font_size"), 28)

    def test_load_user_config(self):
        user_data = {"theme": "light", "font_size": 32}
        with open(self.test_config_path, 'w') as f:
            json.dump(user_data, f)
        
        config = Config(self.test_config_path)
        self.assertEqual(config.get("theme"), "light")
        self.assertEqual(config.get("font_size"), 32)
        # Check that defaults still exist
        self.assertEqual(config.get("time_format"), "24h")

if __name__ == '__main__':
    unittest.main()
