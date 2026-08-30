## Exercise 7: save-your-work

Instructions: working area is messy and not ready to commit, but a bug needs fixing ASAP. Stash the current work, fix the bug in bug.txt and commit it, then bring the stashed work back and finish it by adding a line "Finally, finished it!" to bug.txt, then commit that too.

Ran:
```bash
git stash
```
Saved the messy working state.

Edited bug.txt, removed the buggy line.

Ran:
```bash
git add bug.txt
git commit -m "removed the bug"
git stash pop
```
Output showed both bug.txt and program.txt as modified after the pop — the stash contained changes to both files, not just bug.txt.

Edited bug.txt again, added the final line "Finally, finished it!".

Ran:
```bash
git add bug.txt program.txt
git commit -m "Finish original work with all files"
git verify
```
Passed.