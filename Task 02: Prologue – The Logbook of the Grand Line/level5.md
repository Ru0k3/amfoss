# Level 5: The Buster Call Timeline Recovery

**Objective:** Time-travel backward using Git to find the untouched records at Enies Lobby, combine the two recovered Poneglyph fragments, and decipher them to reveal the final destination.

### Step 1: Locating the Peaceful Moment in History

I navigated back to the `GrandLine` root directory. The Level 5 README stated: *"Walk backward through time. Find the last peaceful moment before destruction."* This hinted that the files were deleted in a past Git commit, and I needed to find the commit *before* the deletion.

**Command:**

```bash id="a1k7wq"
git log --all --oneline
```

**Output:**

```text id="x4v9rm"
c0e19e2 (HEAD -> canonical-timeline, origin/canonical-timeline, origin/HEAD) Edited README.md
74692fb README edited
ee6f464 (origin/little_garden, little_garden) Level 3: Implemented
bc5aff3 (origin/whiskey_peak_investigation, whiskey_peak_investigation) Level 2: Implemented
c337460 (origin/alternate_timeline) Vaults REMOVED, Evidences ERASED
aa616ca Level 4: Implemented multi-layered archive obfuscation for Water 7 blueprints
a802662 Level 3 init
0c60b00 Level 2: Arrived at peaceful Whiskey Peak
23b4e67 Vaults REMOVED, Evidences ERASED
02e3e60 Level 1 - Loguetown Reef
d4e7bf5 Level 5 : Vault Sealed
082970d Add README
```

**Reasoning:** I spotted two commits related to Level 5: `d4e7bf5 Level 5 : Vault Sealed` and `23b4e67 Vaults REMOVED, Evidences ERASED`. The "peaceful moment before destruction" was the commit right before the vaults were removed. Therefore, I needed to time-travel to commit `d4e7bf5`.

**Command:**

```bash id="w4rj8e"
git checkout d4e7bf5
```

**Output:**

```text id="m3n6pa"
Note: switching to 'd4e7bf5'.
You are in 'detached HEAD' state...
```

### Step 2: Exploring Enies Lobby

I listed the directory contents to see what the untouched timeline held.

**Commands:**

```bash id="h7q2cs"
ls
cd Enies_Lobby/
tree
```

**Output:**

```text id="r8t5yv"
.
├── vault_1
│   └── decode.sh
├── vault_2
│   └── decode.sh
├── vault_3
│   └── decode.sh
├── vault_4
│   └── decode.sh
└── vault_5
    └── decode.sh

6 directories, 5 files
```

**Reasoning:** There were 5 vaults. I checked a few of them using `cat vault_1/decode.sh` etc., but they all just printed "Intruder Alert!!!". They were decoys designed to waste time. I needed to search for hidden files.

### Step 3: Locating the Secure Vault

I used `ls -a` to list all files and directories, including hidden ones.

**Command:**

```bash id="q2m8fz"
ls -a
```

**Output:**

```text id="v6j3kp"
.
..
.cp9_secure_vault
vault_1
vault_2
vault_3
vault_4
vault_5
```

**Reasoning:** The `ls -a` command revealed a hidden directory called `.cp9_secure_vault` containing the secure vault. I entered the directory and found the Python script named `poneglyph.py`.

**Commands:**

```bash id="z9c4nx"
cd .cp9_secure_vault/
ls
```

**Output:**

```text id="p7w2lm"
poneglyph.py
```

I then read the script to understand the deciphering mechanism.

**Command:**

```bash id="c5h8vd"
cat poneglyph.py
```

**Output:**

```python id="e3k9rs"
import base64

ENCODED = input("Enter code : ")
KEY = 0x42

decoded = base64.b64decode(ENCODED)
flag = bytes(b ^ KEY for b in decoded).decode()

print("Prize : ")
print(flag)
```

**Reasoning:** The script takes a Base64 encoded string, decodes it, and then XORs every byte with the hex key `0x42`. The Level 5 README stated: *"The surviving records speak of a Poneglyph inscription shattered into two fragments. Neither fragment carries meaning alone. Restore the inscription before attempting to decipher it."* This meant I needed to combine Fragment I (from Level 3) and Fragment II (from Level 4) into a single string.

### Step 4: Restoring and Deciphering the Inscription

I ran the Python script. When prompted for the code, I pasted Fragment I and Fragment II together without spaces.

**Command:**

```bash id="u6k1zp"
python3 .cp9_secure_vault/poneglyph.py
```

**Input:**

```text id="n4c7hx"
Enter code : KjY2MjF4bW0lKzYqNyBsIS0vbTAtJTcnLSwnbzptDiM3JSpvFiMuJ28PJzAlJ28VIzA=
```

**Output:**

```text id="s8v2qd"
Prize : 
https://github.com/rogueone-x/Laugh-Tale-Merge-War
```

**Reasoning:** The script successfully decoded the Base64 string and XORed the bytes. The resulting plaintext was a URL to a GitHub repository, which served as the final destination for Level 6.

### Flag / Reward

> `https://github.com/rogueone-x/Laugh-Tale-Merge-War`
