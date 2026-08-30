## Exercise 12: forge-date

Instructions: work was actually finished a week ago but committed just now. Change the date of the last commit to make it look like it was committed in 1987.

Ran:
```bash
git start forge-date
git commit --amend --date="1987-01-01 12:00:00" --no-edit
git verify
```
Passed. `--date` overwrites the author date on the commit, and `--no-edit` reuses the existing commit message so no editor pops up.