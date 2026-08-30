## Exercise 16: too-many-commits

Instructions: file.txt was committed in two separate small commits ("Add file.txt" and "Crap, I have forgotten to add this line."). Squash them into a single commit.

My reasoning before starting: usual approach for fixing commits is `git commit --amend`, but that only touches the latest commit — it can't combine two commits into one. Then thought about `git rebase` in general, but plain rebase doesn't do the combining work either, it just replays commits. The actual tool for merging multiple commits into one is `squash`, used inside interactive rebase. So went with `git rebase -i`.

Ran:
```bash
git start too-many-commits
git rebase -i HEAD~2
```
Changed the second commit's `pick` to `squash`, saved and exited. Nano then opened the commit message editor showing both old messages, before I made any edit:
```
GNU nano 9.2     /home/aizen/Projects/exercises/.git/COMMIT_EDITMSG

# This is a combination of 2 commits.
# This is the 1st commit message:
Add file.txt
# This is the commit message #2:
Crap, I have forgotten to add this line.
# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
#
# Date:      Sun Aug 30 11:38:34 2026 +0530
#
# interactive rebase in progress; onto e3ae7c7
# Last commands done (2 commands done):
#    pick df2937f # Add file.txt
#    squash f609762 # Crap, I have forgotten to add this line.
# No commands remaining.
# You are currently rebasing branch 'too-many-commits' on 'e3ae7c7'.
#
# Changes to be committed:
#       new file:   file.txt
#
```

My initial approach: I edited this screen and replaced the two original message lines with a single new combined message instead of keeping the first commit's original one:
```
Add file.txt with missing line

# This is a combination of 2 commits.
# This is the 1st commit message:

# This is the commit message #2:

# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
#
# interactive rebase in progress; onto e3ae7c7
# Last commands done (2 commands done):
#    pick df2937f # Add file.txt
#    squash f609762 # Crap, I have forgotten to add this line.
# No commands remaining.
# You are currently rebasing branch 'too-many-commits' on 'e3ae7c7'.
#
# Changes to be committed:
#       new file:   file.txt
```
Saved and exited nano — the rebase finished cleanly:
```
[detached HEAD fcda938] Add file.txt with missing line
Successfully rebased and updated refs/heads/too-many-commits.
```

Ran:
```bash
git verify
```
Output:
```
Status: FAILED
You should leave commit message as it was in the first commit.
```
This is how I found out my approach was wrong — the squash itself worked fine, but the exercise specifically wanted the final commit to keep the **first** commit's original message ("Add file.txt"), not a new custom one.

Fixed it:
```bash
git commit --amend -m "Add file.txt"
git verify
```
Passed.

---

## Why this happened

Squashing two commits into one still leaves you free to write whatever message you want for the merged commit — Git doesn't enforce keeping either original message, it just combines the changes. I assumed a more descriptive message was the "correct" choice, but this exercise's validator specifically checks that the message matches the original first commit exactly.

Since the squash had already succeeded and only the message was wrong, there was no need to redo the rebase — `git commit --amend -m "..."` was enough to fix just the message on the already-squashed commit.