## Exercise 1: master

Instructions: repo should have exactly 1 commit.

Ran `git verify`:
```
Exercise: master
Status: FAILED
Expected number of commits: 1. Received 2.
```
Failed because I had an extra commit already sitting on master.

Tried:
```bash
git reset --hard HEAD~1
```

Ran `git verify` again — passed.

Alternative (if things are more messed up than one extra commit), full reset:
```bash
git start master
```