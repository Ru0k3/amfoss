## Exercise 21: invalid-order

Instructions: two commits exist but in the wrong order. Swap them.

Ran:
```bash
git start invalid-order
git log -2
```
Output:
```
commit 4efb50b (HEAD -> invalid-order)
    This should be the first commit

commit 730db9d
    This should be the second commit
```
`git log` shows newest first, so the commit meant to be "first" was actually sitting on top (newest), and the one meant to be "second" was underneath (older) — backwards.

Ran:
```bash
git rebase -i HEAD~2
```
The todo list (oldest to newest, opposite order from `git log`) showed:
```
pick 730db9d This should be the second commit
pick 4efb50b This should be the first commit
```
Swapped the two lines (kept both as `pick`, just reordered them):
```
pick 4efb50b This should be the first commit
pick 730db9d This should be the second commit
```
Saved and exited.

Ran:
```bash
git verify
```
Passed.

---

## Why this works

`git rebase -i` doesn't just let you edit or squash commits — the order of lines in the todo list *is* the order Git replays them in. Since the lines are executed top to bottom, swapping two `pick` lines swaps the order the commits get rebuilt in, without needing to change their content or messages at all.

The one thing to keep straight: `git log` prints newest → oldest (top to bottom), but the interactive rebase todo list is oldest → newest. So the commit that looked "on top" in `git log` actually appears at the *bottom* of the rebase list — easy to get turned around if you don't account for that flip.