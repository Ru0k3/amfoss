## Exercise 18: commit-parts

Instructions: file.txt has multiple unrelated changes mixed together in one working copy — some belong to "Task 1" and some to "Task 2". Split them into two separate commits, one per task, without touching the other's lines.

Ran:
```bash
git diff
```
Output:
```
+I forgot to add file header.
 This is a program
+And this is new feature done in task 1.
+It lasts for many lines as task 1 was big.
 It is supposed to work.
-It works
+It works!
+This is not related, it is task 2.
 It is quite brilliant, actually.
+Task 1 is finished.
```
All the changes were in the same file, so a normal `git add file.txt` would've bundled everything into one commit — not what was needed.

Ran:
```bash
git add -p
```
This walks through the file change by change ("hunks") and asks whether to stage each one individually. The first hunk covered multiple unrelated lines together, so:
```
(1/1) Stage this hunk [y,n,q,a,d,s,e,p,P,?]? s
```
Used `s` to split it further into 4 smaller hunks, then answered per hunk based on which task each line belonged to:
```
(1/4) ... n   → "I forgot to add file header." (Task 2, skip for now)
(2/4) ... y   → Task 1 feature lines (stage)
(3/4) ... n   → "It works!" / "task 2" line (Task 2, skip)
(4/4) ... y   → "Task 1 is finished." (stage)
```

Checked what actually got staged:
```bash
git diff --cached
```
Output:
```
+And this is new feature done in task 1.
+It lasts for many lines as task 1 was big.
+Task 1 is finished.
```
Only the Task 1 lines were staged — exactly what should go into the first commit.

Committed Task 1:
```bash
git commit -m "Complete Task 1"
```

Checked what was left:
```bash
git diff
```
Output:
```
+I forgot to add file header.
-It works
+It works!
+This is not related, it is task 2.
```
Everything remaining was Task 2 content.

Staged and committed the rest:
```bash
git add file.txt
git commit -m "Complete Task 2"
git verify
```
Passed.

---

## Why this approach

The core problem was that both tasks' changes lived in the same file and even some in overlapping hunks — a plain `git add file.txt` would stage everything at once, making it impossible to commit them separately.

`git add -p` (patch mode) breaks a file's diff into hunks and lets you decide, hunk by hunk, whether it goes into the current staging area. When a hunk itself mixes lines from both tasks (like the first one here), the `s` option splits it into smaller hunks until each one is small enough to belong cleanly to just one task.

Once only the Task 1 hunks were staged, `git diff --cached` confirmed exactly what would go into the next commit — checking this before committing avoided accidentally mixing tasks together. After committing Task 1, whatever was left in `git diff` (the unstaged remainder) was, by definition, everything that hadn't been touched yet — which was all Task 2 content, so it could be added and committed as a whole without needing another patch-mode pass.