# Level 2: The Two Faces of Whiskey Peak

**Objective:** Uncover the hidden history of Whiskey Peak and recover the Executive Transmission Code (Baroque Works flag).

---

## Step 1: Initial Reconnaissance

After completing Level 1, I navigated back to the `GrandLine` directory and entered `Whiskey_Peak`.

### Commands

```bash
cd ../
cd Whiskey_Peak/
ls -la
```

### Output

```text
.rw-r--r-- 66 aizen 22 Jul 22:16  feast_manifest.txt
```

### Reasoning

The `ls -la` command showed only one file. The README hinted at a **"hidden history concealed beside the one everyone sees"** and mentioned **"parallel timelines."**

This strongly suggested the use of Git branches to hide files.

---

## Step 2: Investigating Git Branches

I checked the available Git branches to look for an alternate timeline where the hidden history might still exist.

### Command

```bash
git branch -a
```

### Output

```text
* canonical-timeline
  remotes/origin/HEAD -> origin/canonical-timeline
  remotes/origin/alternate_timeline
  remotes/origin/canonical-timeline
  remotes/origin/little_garden
  remotes/origin/whiskey_peak_investigation
```

### Reasoning

There was a branch named `whiskey_peak_investigation`. Based on the README's clues about a parallel timeline, I decided to switch to this branch to see if the hidden files would appear.

---

## Step 3: Switching Timelines and Discovering the Vault

I switched to the investigation branch and listed the directory contents again.

### Commands

```bash
git checkout whiskey_peak_investigation
ls -la
```

### Output

```text
drwxr-xr-x  - aizen 22 Jul 22:49  .baroque_works_cache
.rw-r--r-- 80 aizen 22 Jul 22:49  feast_manifest.txt
```

### Reasoning

Switching branches revealed a hidden directory called `.baroque_works_cache`. I navigated into it to see what was inside.

### Commands

```bash
cd .baroque_works_cache/
ls
```

### Output

```text
unlock_vault.sh
```

---

## Step 4: Analyzing the Unlock Script

I read the `unlock_vault.sh` script to understand how to access the transmission.

### Command

```bash
cat unlock_vault.sh
```

### Output

```bash
#!/bin/bash

TARGET_HASH="b23dad4218492d3bcefc3cbc47b1c367d2a16852921708fee07b5b6e98068fe9"
ENCRYPTED_FLAG="U2FsdGVkX18eGXT7fCm/5zmZmejGVicPYQQLji9cigHrIyxzalWleyVW+k3X6rBlS3baMgfv0DVe24ILF5v+rw==" 

# Compute the hash of the student's input
INPUT_HASH=$(echo -n "$AWAKENING_SIGNATURE" | sha256sum | awk '{print $1}')

if [ "$INPUT_HASH" == "$TARGET_HASH" ]; then
    echo "[SIGNATURE MATCH] Devil Fruit aura detected. Bypassing proxy firewall..."
    # ... generates two log files and injects the decrypted flag into line 42 of bounty_hunter_feed.log ...
    # ... uses $AWAKENING_SIGNATURE as the OpenSSL password ...
else
    echo "[ACCESS DENIED] Environmental Scan Failed. System user unauthorized."
fi
```

### Reasoning

The script checks for an environment variable named `AWAKENING_SIGNATURE`.

It calculates the SHA-256 hash of the variable:

```bash
INPUT_HASH=$(echo -n "$AWAKENING_SIGNATURE" | sha256sum | awk '{print $1}')
```

and compares it against the `TARGET_HASH`.

If the hashes match, the script decrypts the flag using the value of `AWAKENING_SIGNATURE` as the OpenSSL password.

The README for Level 2 stated:

> "The communications vault recognizes the aura of the awakened Gito Gito no Mi."

This meant I needed to set my Level 1 flag as an environment variable.

---

## Step 5: Setting the Environment Variable and Running the Script

I exported the Level 1 flag as the required environment variable and executed the script.

### Commands

```bash
export AWAKENING_SIGNATURE="ONE_PIECE{GITO_GITO_NO_AWAKENING}"
./unlock_vault.sh
```

### Output

```text
[SIGNATURE MATCH] Devil Fruit aura detected. Bypassing proxy firewall...
[SUCCESS] Decrypting Baroque transmission streams...
Files dropped: 'marine_intercept.log' and 'bounty_hunter_feed.log'. Run diff to compare.
```

### Reasoning

The script successfully matched the hash and generated two log files. It mentioned that the flag was injected into one of the files.

Since the files contain 100 lines of identical text, comparing them manually would be tedious.

---

## Step 6: Extracting the Hidden Flag

I used the `diff` command to instantly highlight the differences between the two generated log files.

### Command

```bash
diff marine_intercept.log bounty_hunter_feed.log
```

### Output

```text
42c42
< LOG_STREAM_ENTRY_SECURE_NODE_042_VALID
---
> BAROQUE_DIAL{SPLIT_TIMELINE_MISDIRECTION}
```

### Reasoning

The `diff` output showed that line 42 differed between the files.

The second file, `bounty_hunter_feed.log`, contained the decrypted Executive Transmission Code.

---

## Flag

> `BAROQUE_DIAL{SPLIT_TIMELINE_MISDIRECTION}`
