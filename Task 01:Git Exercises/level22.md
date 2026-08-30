## Exercise 22: find-swearwords

Instructions: the word "shit" appears in a few commits across words.txt and list.txt. Go back and replace it with "flower" inside those original commits, not with a later fix-up commit.

Ran:
```bash
git log --oneline -S shit -- words.txt list.txt
```
Output:
```
619acec Add word #94
a2a9aec Add word #46
d2c15a2 Add word #23
```
`-S shit` finds commits that actually introduced or removed the word "shit" in those files, so this pinpointed exactly the three commits that needed fixing.

My initial approach: instead of going back into those three specific commits, I did a full rebase and replaced "shit" with a new word — but typed "boom!!!" instead of "flower" — then just amended the last commit to fix the text afterward.

Ran:
```bash
git verify
```
Output:
```
Status: FAILED
You mistakenly replaced "shit" word with "boom!!!", not with a "flower".
```
This is where I realized the wrong word had gone in. Tried fixing it by just editing the files again and amending the final commit, but that wasn't enough — the exercise checks that the *original* three commits contain "flower", not that some later commit patches it. A single amend at the end doesn't reach back into commits further down the history.

Restarted cleanly:
```bash
git start find-swearwords
git log --oneline -S shit -- words.txt list.txt
```
Same three commits found again.

Ran:
```bash
git rebase -i --root
```
Marked all three as `edit`, leaving everything else as `pick`. Git stopped at the first one:
```
Stopped at d2c15a2...  # Add word #23
```

For each of the three stops, repeated the same steps:
```bash
grep -n shit words.txt list.txt
nano <filename>       # replaced shit -> flower
git add <filename>
git commit --amend --no-edit
git rebase --continue
```
This located exactly which file had the word in that commit, fixed it, folded the fix into that same commit, and moved on to the next stop — for word #23 (list.txt), word #46 (words.txt), and word #94 (words.txt).

After the third fix:
```
Successfully rebased and updated refs/heads/find-swearwords.
```

Ran:
```bash
git verify
```
Passed.

---

## Why the first attempt failed

The exercise doesn't just check the current state of the files — it checks that each of the three original commits, individually, introduced "flower" instead of "shit". Fixing the word only in a later commit (even the very last one) leaves the original three commits unchanged in history; anyone looking back at commit #23, #46, or #94 specifically would still see "shit" sitting there. `git verify` caught that mismatch immediately, since it inspects the actual historical commits, not just the final file contents.

`git rebase -i --root` was needed instead of just `HEAD~N` because the three target commits weren't necessarily the last N commits — `--root` opens the entire history from the very first commit, letting `edit` be placed exactly on the three that mattered regardless of where they sit, while everything else stays a `pick` and passes through untouched. `git log -S shit` was what made finding those three commits possible in the first place, since it searches for commits that changed the *occurrence* of a string, not just commits that touch a given file.