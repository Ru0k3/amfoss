## Exercise 13: fix-old-typo

Instructions: file.txt has a typo (wordl instead of world), but another commit was made on top of it. Fix the typo by editing the older commit in history, not just the latest one. Pay attention to the commit message too.

Ran:
```bash
git log --oneline
```
Output showed the typo commit sitting 2 commits back from HEAD.

Ran:
```bash
git rebase -i HEAD~2
```
Changed `pick` to `edit` on the typo commit, left the other as `pick`. Saved and exited.

Fixed the typo using nano:
```bash
nano file.txt
```
Changed `Hello wordl` to `Hello world`, saved and exited.

Ran:
```bash
git add file.txt
git commit --amend -m "Add Hello world"
git rebase --continue
```
Output:
```
CONFLICT (content): Merge conflict in file.txt
```
This happens because the second commit was originally written on top of the old (typo) content — once the old commit changed, Git couldn't cleanly replay the second commit on top of it.

Opened file.txt with nano to resolve the conflict markers by hand, writing the final two lines (`Hello world` followed by `Further work on Hello world`), saved and exited.

Ran:
```bash
git add file.txt
git rebase --continue
```
This opened an editor to confirm the commit message for the second commit — left it unchanged, saved and quit.

Ran:
```bash
git verify
```
Passed.

---

## Why this one is hard

`git commit --amend` only touches the **latest** commit. Here the typo was 2 commits back, with another commit already sitting on top of it — so amend alone couldn't reach it. That's what interactive rebase is for: it lets you pause at an older commit, fix it, and have Git rebuild everything that came after it.

Running `git rebase -i HEAD~2` opens a todo list of the last 2 commits, each marked `pick`. Changing `pick` to `edit` on the typo commit tells Git: "replay this one, but stop right here so I can fix it."

Once Git stops, fixing the file and running `git commit --amend` doesn't literally edit the old commit — since commits are immutable, Git creates a **new** commit with the fix and treats it as the replacement.

The problem is the second commit was originally built assuming the *old* (typo) content. Once the old commit's content changes, Git tries to replay the second commit on top of the corrected one and finds a mismatch — hence the merge conflict. The conflict has to be resolved by hand in the file itself so the final content has both the fix and the second commit's actual work, then `git add` + `git rebase --continue` tells Git to keep rebuilding history from there.

`git rebase --abort` is the escape hatch — it cancels the whole rebase and puts you back exactly where you started, if you ever want to restart the exercise cleanly.