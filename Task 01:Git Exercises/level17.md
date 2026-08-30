## Exercise 17: executable

Instructions: script.sh exists in the repo but Git doesn't record it as executable. Make Git track the executable bit so anyone who checks out the repo can run `./script.sh` directly.

Ran:
```bash
git start executable
git log --oneline
ls
```
Output:
```
38c5319 (HEAD -> executable) Create script.sh
dec563d (origin/executable) Ex: executable
8fbeef9 (tag: exercise-base) Exercise base
d0f3fe9 Initial commit

README.md  script.sh  start.sh
```

Chose `git update-index --chmod=+x script.sh` instead of the regular `chmod +x script.sh`. Plain `chmod` only changes the file's permission bit on disk — it doesn't tell Git anything. Git tracks its own copy of the permission mode (`100644` for normal, `100755` for executable) separately in the index, so the change needs to go through Git directly for it to be recorded and eventually committed.

Ran:
```bash
git update-index --chmod=+x script.sh
git commit -m "Make script executable"
git verify
```
Passed.

---

## Why this happened

`git update-index` modifies the staging area, the same as `git add` does for file content — it prepares a change but doesn't finalize it into history. Just like editing a file and running `git add` isn't enough on its own, changing a permission bit with `update-index` still needs a `git commit` to actually become part of the repo's history that `git verify` checks against.