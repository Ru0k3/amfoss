import os
import shutil
import subprocess
import platform

class WallpaperSetter:
    def __init__(self):
        self.os_type = platform.system().lower()

    def set_wallpaper(self, image_path):
        image_path = os.path.abspath(image_path)
        if self.os_type == "linux":
            self._set_linux(image_path)
        elif self.os_type == "darwin":
            self._set_macos(image_path)
        elif self.os_type == "windows":
            self._set_windows(image_path)
        else:
            print(f"Unsupported OS: {self.os_type}")

    def _set_linux(self, path):
        if os.environ.get("HYPRLAND_INSTANCE_SIGNATURE") and shutil.which("hyprctl"):
            try:
                subprocess.run(["hyprctl", "hyprpaper", "wallpaper", f",{path}"], check=True)
                return
            except Exception:
                pass

        # Try gsettings (GNOME)
        try:
            subprocess.run(["gsettings", "set", "org.gnome.desktop.background", "picture-uri", f"file://{path}"], check=True)
            subprocess.run(["gsettings", "set", "org.gnome.desktop.background", "picture-uri-dark", f"file://{path}"], check=True)
        except Exception:
            # Fallback to feh if available
            try:
                subprocess.run(["feh", "--bg-fill", path], check=True)
            except Exception:
                print("Failed to set Linux wallpaper. Ensure gsettings or feh is installed.")

    def _set_macos(self, path):
        script = f'tell application "Finder" to set desktop picture to POSIX file "{path}"'
        subprocess.run(["osascript", "-e", script])

    def _set_windows(self, path):
        import ctypes
        SPI_SETDESKWALLPAPER = 20
        ctypes.windll.user32.SystemParametersInfoW(SPI_SETDESKWALLPAPER, 0, path, 3)
