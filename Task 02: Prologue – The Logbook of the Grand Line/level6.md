# Level 6: The Great Merge War at Laugh Tale

**Objective:** Clone the final repository, resolve a Git merge conflict to stabilize the timeline, reconstruct the Pirate King's Password, and execute the victory script to claim the One Piece.

### Step 1: Cloning the Final Repository

I obtained the URL from the Level 5 deciphering step and cloned the repository to my local machine.

**Command:**

```bash
git clone https://github.com/rogueone-x/Laugh-Tale-Merge-War
cd Laugh-Tale-Merge-War
```

**Reasoning:** I needed to pull the final challenge files to my local environment to interact with them.

### Step 2: Initial Reconnaissance

I listed the files and read the README and the victory script to understand the win condition.

**Commands:**

```bash
ls -la
cat README.md
cat victory.sh
```

**Output (victory.sh excerpt):**

```bash
if git diff --quiet && git diff --cached --quiet; then
    :
else
    echo "History is still unstable."
    echo "Complete the merge first."
    exit 1
fi

echo -n "Enter the Pirate King's Password: "
read -r PASS
EXPECTED_HASH="2abfc485e42e701824a6340b3b12e54f0dfad6647d56fb095b50bd4d6384700e"
INPUT_HASH=$(printf "%s" "$PASS" | sha256sum | awk '{print $1}')

if [[ "$INPUT_HASH" == "$EXPECTED_HASH" ]]; then
    # ... prints the final flag ...
```

**Reasoning:** The script checked for a clean working tree (`git diff --quiet`), meaning no merge conflicts could be active. It then asked for a password and checked its SHA-256 hash against a hardcoded expected hash. The README stated: *"Only by reconciling both histories can the true inscription be restored."* This meant I needed to trigger a merge conflict.

### Step 3: Investigating Timelines and Initiating the Merge

I checked the available branches to find the conflicting histories.

**Command:**

```bash
git branch -a
```

**Output:**

```text
* ancient_history
  remotes/origin/HEAD -> origin/ancient_history
  remotes/origin/ancient_history
  remotes/origin/pirate_king_path
```

**Reasoning:** There were two branches: `ancient_history` (current) and `pirate_king_path`. I initiated a merge to collide the timelines.

**Command:**

```bash
git merge origin/pirate_king_path
```

**Output:**

```text
Auto-merging treasure/key_part_1.txt
CONFLICT (content): Merge conflict in treasure/key_part_1.txt
Auto-merging treasure/key_part_2.txt
CONFLICT (content): Merge conflict in treasure/key_part_2.txt
Automatic merge failed; fix conflicts and then commit the result.
```

### Step 4: Analyzing the Fractured Records

I read the conflicted files to see both sides of the history.

**Commands:**

```bash
cat treasure/key_part_1.txt
cat treasure/key_part_2.txt
```

**Output (key_part_1.txt):**

```text
PONEGLYPH FRAGMENT α

Recovered Inscription:

<<<<<<< HEAD
Line
=======
TheGrand
>>>>>>> origin/pirate_king_path
```

**Output (key_part_2.txt):**

```text
PONEGLYPH FRAGMENT β

Recovered Inscription:

<<<<<<< HEAD
bers
=======
Remem
>>>>>>> origin/pirate_king_path
```

**Reasoning:** Neither timeline held the complete truth.

* Part 1: `ancient_history` had "Line", `pirate_king_path` had "TheGrand".
* Part 2: `ancient_history` had "bers", `pirate_king_path` had "Remem".

By combining the missing pieces from both timelines, the true inscription was restored:

* Part 1: `TheGrand` + `Line` = `TheGrandLine`
* Part 2: `Remem` + `bers` = `Remembers`
* **Pirate King's Password:** `TheGrandLineRemembers`

### Step 5: Resolving the Conflict

I created a new branch and opened each conflicted file using `nano`.

**Command:**

```bash
git checkout -b pirate_king_resolution
nano treasure/key_part_1.txt
```

I changed the content from:

```text
PONEGLYPH FRAGMENT α

Recovered Inscription:

<<<<<<< HEAD
Line
=======
TheGrand
>>>>>>> origin/pirate_king_path
```

to:

```text
PONEGLYPH FRAGMENT α

Recovered Inscription:

TheGrandLine
```

I then opened the second file:

```bash
nano treasure/key_part_2.txt
```

I changed the content from:

```text
PONEGLYPH FRAGMENT β

Recovered Inscription:

<<<<<<< HEAD
bers
=======
Remem
>>>>>>> origin/pirate_king_path
```

to:

```text
PONEGLYPH FRAGMENT β

Recovered Inscription:

Remembers
```

I saved both files and then staged and committed the resolved files.

**Commands:**

```bash
git add treasure/key_part_1.txt treasure/key_part_2.txt
git commit -m "Resolved merge conflict: Reconciled the timelines"
```

### Step 6: Claiming the One Piece

I ran the victory script and entered the reconstructed password.

**Command:**

```bash
./victory.sh
```

**Output:**

```text
==============================
 Verifying Timeline Integrity 
==============================

Enter the Pirate King's Password: TheGrandLineRemembers
Timeline Integrity ............. OK
Merge Conflict ................. Resolved
Repository ..................... Restored
History ........................ Preserved

====================================================

        THE ONE PIECE HAS BEEN FOUND

====================================================

Congratulations, Captain.
...
FLAG{The_Grand_Line_Remembers_Your_Commit}
...
```

### Flag

> `FLAG{The_Grand_Line_Remembers_Your_Commit}`
