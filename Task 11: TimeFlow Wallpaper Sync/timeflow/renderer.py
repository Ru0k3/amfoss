import os
from datetime import datetime
from PIL import Image, ImageDraw, ImageFont
import textwrap

class Renderer:
    def __init__(self, config):
        self.config = config
        self.width = 1920  # Default, should be detected or configurable
        self.height = 1080
        
        # Load fonts
        try:
            # Try common Linux font paths
            font_paths = [
                "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf",
                "/usr/share/fonts/TTF/DejaVuSans.ttf",
                "/usr/share/fonts/truetype/liberation/LiberationSans-Regular.ttf"
            ]
            self.font_path = None
            for p in font_paths:
                if os.path.exists(p):
                    self.font_path = p
                    break
            
            self.base_font_size = self.config.get("font_size")
            self.font = ImageFont.truetype(self.font_path, self.base_font_size) if self.font_path else ImageFont.load_default()
            self.clock_font = ImageFont.truetype(self.font_path, 80) if self.font_path else ImageFont.load_default()
            self.date_font = ImageFont.truetype(self.font_path, 30) if self.font_path else ImageFont.load_default()
        except Exception:
            self.font = ImageFont.load_default()
            self.clock_font = ImageFont.load_default()
            self.date_font = ImageFont.load_default()

    def render(self, content):
        # Create base image
        bg_color = self.config.get("bg_color")
        img = Image.new('RGB', (self.width, self.height), color=bg_color)
        draw = ImageDraw.Draw(img)

        # Draw Clock Panel (Right Sidebar)
        panel_width = 400
        panel_color = "#1E293B"
        draw.rectangle([self.width - panel_width, 0, self.width, self.height], fill=panel_color)

        # Draw Clock
        now = datetime.now()
        time_str = now.strftime("%H:%M:%S")
        date_str = now.strftime("%A, %B %d, %Y")
        
        clock_color = self.config.get("clock_color")
        text_color = self.config.get("text_color")
        
        # Clock position
        draw.text((self.width - panel_width + 20, 50), time_str, font=self.clock_font, fill=clock_color)
        draw.text((self.width - panel_width + 20, 150), date_str, font=self.date_font, fill=text_color)

        # Draw Content
        margin = 50
        max_content_width = self.width - panel_width - (margin * 2)
        
        # Simple wrapping
        lines = []
        avg_char_width = 15 # Rough estimate
        chars_per_line = max(20, int(max_content_width / avg_char_width))
        
        for line in content.splitlines():
            if not line.strip():
                lines.append("")
                continue
            wrapped = textwrap.wrap(line, width=chars_per_line)
            lines.extend(wrapped)

        # Render lines
        y_offset = margin
        line_height = self.base_font_size + 10
        
        # Auto-scaling font if too many lines
        if len(lines) * line_height > self.height - (margin * 2):
            new_size = max(10, int((self.height - margin * 2) / len(lines)) - 2)
            current_font = ImageFont.truetype(self.font_path, new_size) if self.font_path else ImageFont.load_default()
            line_height = new_size + 2
        else:
            current_font = self.font

        for line in lines:
            if y_offset + line_height > self.height - margin:
                draw.text((margin, y_offset), "...", font=current_font, fill=text_color)
                break
            draw.text((margin, y_offset), line, font=current_font, fill=text_color)
            y_offset += line_height

        return img

    def save(self, img, path):
        os.makedirs(os.path.dirname(path), exist_ok=True)
        img.save(path)
