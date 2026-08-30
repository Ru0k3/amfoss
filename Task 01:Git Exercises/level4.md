## Exercise 4: split-commit

Instructions: first.txt and second.txt are grouped into one commit. Goal is to split them into two separate commits.

Ran:
```bash
git status
```
Output:
```
On branch split-commit
Your branch is ahead of 'origin/split-commit' by 1 commit.
nothing to commit, working tree clean
```
Both files were already committed together in a single commit, so there was nothing to add or split yet at this point.

Tried:
```bash
git reset HEAD~1
```
This undid the combined commit and left first.txt and second.txt as untracked files.

Ran:
```bash
git status
```
Output:
```
Untracked files:
  first.txt
  second.txt
```

Ran:
```bash
git add first.txt
git commit -m "First"
git add second.txt
git commit -m "Second"
git verify
```
Passed.