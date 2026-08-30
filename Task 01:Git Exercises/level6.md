## Exercise 6: merge-conflict

Instructions: merge another-piece-of-work into merge-conflict branch. Both branches changed the same part of the same file, so this will cause a conflict that needs to be resolved by hand.

Ran:
```bash
git merge another-piece-of-work
```
Conflict triggered as expected.

Ran:
```bash
git status
```
Showed equation.txt as the conflicted file.

Opened equation.txt:
```
<<<<<<< HEAD
2 + ? = 5
=======
? + 3 = 5
>>>>>>> another-piece-of-work
```
Removed the conflict markers and combined both sides into one correct equation:
```
2 + 3 = 5
```

Ran:
```bash
git add equation.txt
git commit -m "Resolve merge conflict"
git verify
```
Passed.