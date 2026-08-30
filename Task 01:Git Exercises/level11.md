## Exercise 11: fix-typo

Instructions: file.txt was committed with a typo (wordl instead of world). Fix it by editing the previous commit so it looks like the typo never happened — including the commit message.

Ran:
```bash
nano file.txt
git add file.txt
git commit --amend
```
This opened an editor for the commit message instead of reusing the old one, and the amend didn't fully overwrite the previous commit as expected.

Ran:
```bash
git verify
```
Output:
```
Exercise: fix-typo
Status: FAILED
Expected number of commits: 1. Received 2.
```
Failed because the amend created an extra commit instead of folding into the original one.

Tried:
```bash
git start fix-typo
nano file.txt
git add file.txt
git commit --amend --no-edit
git verify
```
Passed. `--no-edit` reuses the existing commit message so the amend cleanly replaces the old commit instead of leaving a second one behind.