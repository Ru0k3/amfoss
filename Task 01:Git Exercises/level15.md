## Exercise 15: split-commit

Instructions: first.txt and second.txt were committed together in one commit. Split them so each file has its own separate commit.

Ran:
```bash
git reset HEAD~1
git status
```
Output:
```
Untracked files:
  first.txt
  second.txt
```
The combined commit was undone, but the file changes stayed in the working directory as untracked files (used `git reset HEAD~1`, not `--hard`, so nothing was lost).

Ran:
```bash
git add first.txt
git commit -m "Add first.txt"
git add second.txt
git commit -m "Add second.txt"
git verify
```
Passed.

---

## Why this works

`git reset HEAD~1` moves the branch pointer back one commit but leaves the actual file changes untouched in the working directory — it just un-commits them, it doesn't delete them. That's the key difference from `git reset --hard HEAD~1`, which would have wiped the changes too.

Once the combined commit was undone, both files sat there as plain uncommitted changes again, which meant they could be staged and committed one at a time instead of together — giving a clean, separate commit for each file.