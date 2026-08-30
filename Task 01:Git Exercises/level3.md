## Exercise 3: commit-one-file-staged

Instructions: A.txt and B.txt are already staged. Goal is to commit only one of them.

Difference from Level 2: in commit-one-file, nothing was staged and I had to choose what to add. Here, both files are already staged for me, so I have to unstage the one I don't want before committing.

Ran:
```bash
git start next
```
This jumped me into a different exercise (split-commit) instead of commit-one-file-staged. Failed because `next` doesn't target a specific exercise by name, it just moves to whatever's next in sequence, which wasn't what I wanted.

Tried:
```bash
git start commit-one-file-staged
```
This loaded the correct environment with A.txt and B.txt pre-staged.

Ran:
```bash
git reset HEAD A.txt
git commit -m "destage one file"
git verify
```
Passed.