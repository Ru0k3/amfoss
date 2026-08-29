# Level 3: The Wax Labyrinth of Little Garden

**Objective:** Find the genuine Baroque Works report hidden among hundreds of replicas by identifying the transformed Executive Transmission Code, and recover the first cipher fragment.

### Step 1: Initial Reconnaissance and Branch Investigation

I navigated to the `Wax_Jungle` directory while on the `canonical-timeline` branch.

**Commands:**

```bash
cd ../
cd Wax_Jungle/
ls -la
```

**Output:**

```text
.rw-r--r-- 0 aizen 22 Jul 22:54  .gitkeep
```

**Reasoning:** The directory was completely empty except for a `.gitkeep` file. In Level 2, we discovered that challenges were hidden in alternate Git branches. To find the correct timeline for Level 3, I listed all available branches in the repository.

**Command:**

```bash
git branch -a
```

**Output:**

```text
* canonical-timeline
  remotes/origin/HEAD -> origin/canonical-timeline
  remotes/origin/alternate_timeline
  remotes/origin/canonical-timeline
  remotes/origin/little_garden
  remotes/origin/whiskey_peak_investigation
```

**Reasoning:** There were several branches. We already used `whiskey_peak_investigation` for Level 2. The `alternate_timeline` branch was a trap/destroyed timeline (evidenced by its commit message "Vaults REMOVED"). The only remaining logical branch for Level 3 was `little_garden`. I switched to it to populate the `Wax_Jungle` directory.

**Command:**

```bash
git checkout little_garden
```

### Step 2: Assessing the Labyrinth

I listed the directory contents recursively to understand the structure of the files.

**Command:**

```bash
tree
```

**Output:**

```text
.
├── report_001.log
├── report_002.log
...
└── sector_gamma
    └── swamp
        └── south
            ├── report_337.log
            ...
61 directories, 489 files
```

**Reasoning:** The output revealed 489 files. Almost all of them strictly followed the naming convention `report_XXX.log`. Manually inspecting 489 files is inefficient. I needed a way to identify the genuine report among the fakes.

### Step 3: Locating the Anomalous File

While reviewing the `tree` output, I noticed a single file that broke the `report_XXX.log` naming convention. It was located inside `sector_beta/outpost/camp/watchtower/storage/archive/` and was named `agent_manifest.log`. I used the `cat` command to read its contents directly.

**Command:**

```bash
cat sector_beta/outpost/camp/watchtower/storage/archive/agent_manifest.log
```

**Output:**

```text
SECURITY LOG ACCESS // LEVEL 3 CLEARANCE REQUIRED
-------------------------------------------------
STATUS: METALLIC WAX SUIT ACTIVE

SECURITY_TAG:
QkFST1FVRV9ESUFMe1NQTElUX1RJTUVMSU5FX01JU0RJUkVDVElPTn0K

-------------------------------------------------

BAROQUE WORKS EXECUTIVE REPORT

PONEGLYPH_FRAGMENT_I = "KjY2MjF4bW0lKzYqNyBsIS0vbTAtJTcnL"

-------------------------------------------------
```

### Step 4: Verifying the Broadcast Identifier (The Security Tag)

The file contained a `SECURITY_TAG` with a Base64 encoded string. The Level 3 README stated: *"Executive Codes are never transmitted in their original form. Before being sent across the Den Den Mushi network, every transmission identifier is converted into its broadcast representation."* In CTF challenges, "broadcast representation" or "transformed identifier" almost always refers to Base64 encoding. I decoded the string to see if it matched our Level 2 flag.

**Command:**

```bash
echo "QkFST1FVRV9ESUFMe1NQTElUX1RJTUVMSU9JUkVDVElPTn0K" | base64 -d
```

**Output:**

```text
BAROQUE_DIAL{SPLIT_TIMELINE_MISDIRECTION}
```

**Reasoning:** The decoded string was an exact match for the Executive Transmission Code (Level 2 flag). This proved that `agent_manifest.log` was the genuine report. The fakes would not have contained this valid, transformed identifier.

### Step 5: Extracting the Cipher Fragment (The Reward)

**How did I conclude the reward was the Poneglyph fragment?**

The connection is explicitly stated in the Level 3 README: *"Somewhere inside this labyrinth, one intercepted report still carries that transformed identifier. Find the genuine report. Within it lies the first cipher fragment needed to uncover the road ahead."*

Because I verified that `agent_manifest.log` was the genuine report (by confirming its Security Tag matched the Executive Code), I knew the text inside it was the promised reward. The file contained the variable `PONEGLYPH_FRAGMENT_I` and its corresponding value. This string is not a standard flag format like `FLAG{...}`, but rather a piece of data required to solve Level 5, acting as a progress item for the later stages of the CTF.

### Flag / Reward

> `PONEGLYPH_FRAGMENT_I = "KjY2MjF4bW0lKzYqNyBsIS0vbTAtJTcnL"`
