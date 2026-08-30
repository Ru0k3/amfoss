## Exercise 8: change-branch-history

Instructions: change-branch-history has a commit built on an older base, while hot-bugfix has a bugfix on a separate branch. Goal is to move (rebase) the change-branch-history commit so it sits on top of hot-bugfix instead of its original base.

Ran:
```bash
git branch
```
Confirmed I was already on change-branch-history.

Ran:
```bash
git rebase hot-bugfix
git verify
```
Passed. Git lifted the change-branch-history commit off its old base, moved the branch pointer to hot-bugfix, then replayed the commit on top of it.

## Rebase, more generally
Rebase shows up beyond this exercise too — keeping a feature branch up to date with main, squashing messy commits before a PR, fixing a typo buried a few commits back (interactive rebase), reordering commits, splitting a big commit into smaller ones, or dropping a commit entirely. The rule: only rebase local/unshared history — once commits are pushed and others have pulled them, rebasing rewrites hashes and breaks things for everyone else, so merge instead.