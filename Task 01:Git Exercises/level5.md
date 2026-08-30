## Exercise 5: chase-branch

Instructions: chase-branch is behind escaped by two commits, straight line, no divergence. Goal is to make chase-branch point to the same commit as escaped.

Note: this isn't the classic "two features merged together" case. chase-branch hadn't done any work of its own — it was just idle while escaped moved ahead. So this merge just catches chase-branch up, no real combining of histories needed.

Ran:
```bash
git start chase-branch
git branch
```
Output:
```
* chase-branch
  commit-one-file
  commit-one-file-staged
  escaped
  ignore-them
  master
  split-commit
```
Confirmed I was already on chase-branch.

Ran:
```bash
git merge escaped
git verify
```
Passed. Since chase-branch was a direct ancestor of escaped, Git fast-forwarded the pointer instead of creating a new merge commit.