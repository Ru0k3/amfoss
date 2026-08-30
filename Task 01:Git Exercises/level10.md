## Exercise 10: case-sensitive

Instructions: File.txt was committed but the filename should be all lowercase (file.txt). Rename it. Tricky on filesystems that treat File.txt and file.txt as the same file.

Ran:
```bash
git mv File.txt file.txt
git commit -m "Rename File.txt to lowercase"
git verify
```
Passed. Used `git mv` instead of the regular OS `mv` command, since on case-insensitive filesystems a plain `mv` can make Git miss the rename entirely.