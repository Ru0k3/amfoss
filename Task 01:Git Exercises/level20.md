## Exercise 20: rebase-complex

Instructions: rebase-complex branched off after "First issue commit" (b6c2ac9) and has two commits on top ("First commit fixing bug", "Bug finally fixed"). A separate branch, issue-555, also branched from the same point with unrelated work. your-master has moved ahead independently with its own two commits. Goal is to move only rebase-complex's two bugfix commits onto the tip of your-master, without dragging along the issue-555 commits or the original base.

Ran:
```bash
git branch
git log --oneline --all --graph
```
Output showed the shape of the history:
```
* 1811826 (issue-555) Evnt more work on issue
* ad83e5d More work on issue
| * 9d5f523 (HEAD -> rebase-complex) Bug finally fixed
| * bbfead3 First commit fixing bug
|/
* b6c2ac9 First issue commit
| * 31e57a1 (your-master) Second commit in your-master
| * 98a65a9 First commit in your-master
|/
* 47996a9 Base commit for rebase complex
```
This confirmed the branch point was `b6c2ac9`, with rebase-complex's two unique commits (`bbfead3`, `9d5f523`) sitting after it — separate from issue-555's commits, which also start after the same point but shouldn't be touched.

Ran:
```bash
git rebase --onto your-master b6c2ac9
```
Output:
```
Successfully rebased and updated refs/heads/rebase-complex.
```

Ran:
```bash
git verify
```
Passed.

---

## Why this command

A normal `git rebase your-master` would have tried to replay *everything* on rebase-complex since the common ancestor with your-master — which isn't what was needed here, since the two branches diverge much earlier at `47996a9`, not at `b6c2ac9`.

`git rebase --onto <new-base> <old-base>` is for exactly this situation: it takes only the commits that come *after* `<old-base>` on the current branch, and replays just those onto `<new-base>`, ignoring everything else. Read as:

```
git rebase --onto your-master b6c2ac9
             ↑            ↑
        new base    exclude this and everything before it
```

So Git took only `bbfead3` and `9d5f523` (everything after `b6c2ac9` on rebase-complex) and replayed them on top of `your-master`'s tip (`31e57a1`), leaving issue-555's commits and the shared base entirely untouched.