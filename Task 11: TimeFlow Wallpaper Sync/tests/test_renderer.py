import unittest
from timeflow.config import Config
from timeflow.renderer import Renderer
from PIL import Image

class TestRenderer(unittest.TestCase):
    def setUp(self):
        self.config = Config()
        self.renderer = Renderer(self.config)

    def test_render_returns_image(self):
        content = "Test Content"
        img = self.renderer.render(content)
        self.assertIsInstance(img, Image.Image)
        self.assertEqual(img.size, (1920, 1080))

    def test_render_empty_content(self):
        img = self.renderer.render("")
        self.assertIsInstance(img, Image.Image)

    def test_render_long_content(self):
        content = "\n".join([f"Line {i}" for i in range(200)])
        img = self.renderer.render(content)
        self.assertIsInstance(img, Image.Image)

if __name__ == '__main__':
    unittest.main()
