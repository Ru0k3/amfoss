## Exercise 19: pick-your-features

Instructions: Features A, B, and C each live as separate commits on their own branches (feature-a, feature-b, feature-c). Bring all three onto pick-your-features as individual commits, in order, without merging the branches directly.

Ran:
```bash
git branch
git log --oneline --all --graph
```
This showed pick-your-features sitting on its own base commit, with feature-a, feature-b, and feature-c each branching off separately with their own unique commits.

Cherry-picked Feature A:
```bash
git cherry-pick feature-a
```
Output:
```
[pick-your-features 14fe1e2] Implement Feature A
```
Applied cleanly.

Cherry-picked Feature B:
```bash
git cherry-pick feature-b
```
Output:
```
Auto-merging program.txt
[pick-your-features f96ac1b] Implement Feature B
```
Also applied cleanly.

Cherry-picked Feature C:
```bash
git cherry-pick feature-c
```
Output:
```
Auto-merging program.txt
CONFLICT (content): Merge conflict in program.txt
error: could not apply 549dc97... Complete Feature C
```
This conflicted because Feature C's commit was originally written against the base version of program.txt, but by this point A and B's changes had already been layered onto the file — so Git couldn't automatically figure out how to combine C's changes with what A and B had already added.

Checked the conflict:
```bash
cat program.txt
```
Output:
```
This is complete feature B
This is base version of the program.
It has only two lines at the beginning.
<<<<<<< HEAD
This is complete feature A
=======
This is first part Feature C
This is second part of Feature C
>>>>>>> 549dc97 (Complete Feature C)
```

Resolved it by hand in nano, keeping both A's line and C's two lines, removing the conflict markers:
```
This is complete feature B
This is base version of the program.
It has only two lines at the beginning.
This is complete feature A
This is first part Feature C
This is second part of Feature C
```

Ran:
```bash
git add program.txt
git cherry-pick --continue
```
This opened an editor to confirm the commit message ("Complete Feature C") — left it unchanged, saved and exited.

Ran:
```bash
git verify
```
Passed.

---

## Why this happened

Cherry-pick takes a single commit from anywhere in the repo and replays it onto whatever branch you're currently on — it's like a targeted, one-commit version of rebase. Feature A and B applied cleanly because their changes touched separate lines in program.txt that didn't overlap with anything already there.

Feature C conflicted because its original commit assumed program.txt still looked like the base version — but by the time it got cherry-picked, program.txt already had Feature A's and Feature B's lines added on top. Git tried to apply Feature C's diff onto a file that no longer matched what Feature C expected, so it couldn't merge the change automatically and left conflict markers for a manual decision on how to combine both sets of changes.

The end result was pick-your-features ending up with three independent commits (A, B, C) stacked in order — each brought in individually via cherry-pick rather than through a merge of the three branches.