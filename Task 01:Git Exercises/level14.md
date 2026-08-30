## Exercise 14: commit-lost

Instructions: a commit was accidentally lost (overwritten). Recover it and get commit-lost branch pointing back to the original commit.

Ran:
```bash
git log --oneline
```
Output:
```
bb18a33 (HEAD -> commit-lost) Accidental change
10a1c3c (origin/commit-lost) Ex: commit-lost
8fbeef9 (tag: exercise-base) Exercise base
d0f3fe9 Initial commit
```
"Accidental change" was visible — not the original commit.

Ran:
```bash
git reflog
```
Top of the output:
```
bb18a33 HEAD@{0}: commit (amend): Accidental change
bfc3bf3 HEAD@{1}: commit: Very imporant piece of work
```
This showed that `bfc3bf3` ("Very imporant piece of work") was the original commit before it got amended into `bb18a33`.

Ran:
```bash
git reset --hard bfc3bf3
git log --oneline
```
Output:
```
bfc3bf3 (HEAD -> commit-lost) Very imporant piece of work
10a1c3c (origin/commit-lost) Ex: commit-lost
8fbeef9 (tag: exercise-base) Exercise base
d0f3fe9 Initial commit
```
commit-lost now points to the original commit instead of the amended one.

Ran:
```bash
git verify
```
Passed.

---

## Why this works

Normal `git log` only shows the commits currently reachable from your branch's HEAD. Once a commit gets amended, the old version isn't reachable anymore through `log` — it looks gone.

But Git doesn't actually delete it right away. Every time HEAD moves (a commit, an amend, a reset, a checkout), Git records that move in the **reflog** — a local, personal log of where your HEAD has pointed over time. The old commit object still exists in Git's database; it's just not part of any branch's visible history anymore.

So `git reflog` lets you look back and find the commit hash from before the amend happened. Once you have that hash (`bfc3bf3`), `git reset --hard <hash>` moves the branch pointer straight back onto it, making it the tip of `commit-lost` again — recovering the "lost" commit.

The reflog is local-only and temporary (it expires after a while), which is why this trick works for recovering your own recent mistakes but isn't a permanent history log like `git log`.