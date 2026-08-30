## Exercise 23: find-bug (final exercise)

Instructions: somewhere across ~300 commits, a bug was introduced into an encoded file (home-screen-text.txt) — the word "jackass" should no longer be appearing, but does in some past commits. Find the exact commit that introduced it.

Ran:
```bash
git log --oneline --all --decorate
```
Output showed ~300 commits, all named "Change text on home screen #N" — no way to tell which one introduced the bug just by reading messages.

Checked the current state:
```bash
openssl enc -base64 -A -d < home-screen-text.txt | grep jackass
```
Output showed "jackass" present in the decoded text — confirming the bug is there right now.

Started a bisect:
```bash
git bisect start
git bisect bad
git bisect good 1.0
```
This marked the current commit as "bad" (bug present) and an older tag (`1.0`) as "good" (bug absent), giving Git a range to binary-search through. Git checked out a commit roughly in the middle:
```
Bisecting: 149 revisions left to test after this (roughly 7 steps)
[b39e252] Change text on home screen #150
```

Repeated the same two-step check at each point Git jumped to — decode and grep, then tell Git the result:
```bash
openssl enc -base64 -A -d < home-screen-text.txt | grep jackass
```
- If "jackass" showed up → `git bisect bad`
- If nothing showed up → `git bisect good`

Went through several rounds (#150 → #75 → #112 → #93 → #84 → #79 → #77 → #78), each time Git narrowing the range and jumping to a new midpoint automatically. On the final round:
```bash
git bisect bad
```
Output:
```
8d1ce3f is the first 'bad' commit
commit 8d1ce3f
    Change text on home screen #78
```
Git had narrowed it down to the exact commit that introduced the word.

Pushed that commit for verification:
```bash
git push origin 8d1ce3f:find-bug
```
Output included:
```
remote: Exercise: find-bug
remote: Status: PASSED
remote: Congratulations! You have done all exercises!
```
The push itself was then rejected by the server's hook afterward (`hook declined`), but that's expected — the exercise server confirms the result during the push attempt itself, and rejects the actual push since it's just a verification mechanism, not a real update.

---

## Why bisect works

`git bisect` automates the manual process of narrowing down a bug by binary search instead of checking every single commit one by one. Once you give it one "good" commit (bug absent) and one "bad" commit (bug present), it checks out the commit exactly halfway between them and asks you to test it. Based on whether you say `good` or `bad`, it eliminates half of the remaining range each time — so instead of possibly checking ~300 commits one at a time, it took roughly 7-8 checks (log₂ of the range) to land on the exact commit that introduced the bug.

The only manual part is the test itself — here, decoding the file and checking for the string "jackass" — everything else (checking out commits, narrowing the range, reporting the final answer) is handled by `git bisect`.