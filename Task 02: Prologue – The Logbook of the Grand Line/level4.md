# Level 4: The Camouflaged Blueprints of Water 7

**Objective:** Uncover the disguised blueprint file's true format, extract its contents, and recover the second Poneglyph cipher fragment.

### Step 1: Initial Reconnaissance

I switched back to the `canonical-timeline` branch and navigated to the `Water_7` directory.

**Commands:**

```bash
git checkout canonical-timeline
cd ../Water_7/
ls -la
```

**Output:**

```text
drwxr-xr-x - aizen 22 Jul 22:43  galley_la_company
```

**Reasoning:** There was only one directory. I entered it to find the blueprint files.

**Commands:**

```bash
cd galley_la_company/
ls -la
```

**Output:**

```text
.rw-r--r-- 0 aizen 22 Jul 22:43 󰡯 puffing_tom_blueprints
```

### Step 2: Analyzing the Disguised Blueprint

I attempted to read the file using `cat`, but it printed gibberish to the terminal.

**Command:**

```bash
cat puffing_tom_blueprints
```

**Output:**

```text
�T^jstep2_blueprints.tar��1o�@�3I�(
                                  ,�r+fc�N�,���DM�T	qr�k��u��D��)�"����
# ... (binary gibberish) ...
```

**Problem Faced:** The file had no file extension, and `cat` outputted unreadable binary data. The Level 4 README stated: *"He stripped the blueprint of its identity... Those who judge by appearances will see only meaningless corruption. But every object remembers what it truly is. Ask it not for its name... Ask it for its nature."* In Linux, asking a file for its "nature" means checking its file signature or magic bytes, regardless of its name or extension.

### Step 3: Revealing its True Nature

I used the `file` command to determine the actual file type.

**Command:**

```bash
file puffing_tom_blueprints
```

**Output:**

```text
puffing_tom_blueprints: gzip compressed data, was "step2_blueprints.tar", last modified: Mon Jul 20 17:02:24 2026, from Unix, original size modulo 2^32 10240
```

**Reasoning:** The `file` command reads the file's header bytes. It revealed that `puffing_tom_blueprints` was actually a **gzip compressed tar archive**.

### Step 4: Multi-Layered Extraction

Since I knew it was a tar/gzip archive, I used the `tar` command to extract it.

**Command:**

```bash
tar -xvzf puffing_tom_blueprints
```

**Output:**

```text
step1_blueprints.zip
```

**Reasoning:** It extracted a new file called `step1_blueprints.zip`. Earlier in the CTF, when I ran `git log --all --oneline` during Level 2, one of the commit messages was: *"Level 4: Implemented multi-layered archive obfuscation for Water 7 blueprints"*. This indicated there were multiple layers of compression to get through.

I checked the nature of this new file to confirm how to extract it next.

**Command:**

```bash
file step1_blueprints.zip
```

**Output:**

```text
step1_blueprints.zip: Zip archive data, made by v3.0 UNIX, extract using at least v1.0...
```

**Reasoning:** It was a standard Zip archive. I used `unzip` to extract it.

**Command:**

```bash
unzip step1_blueprints.zip
```

**Output:**

```text
Archive:  step1_blueprints.zip
   creating: blueprints_extracted/
   creating: blueprints_extracted/hull_design/
 extracting: blueprints_extracted/hull_design/frame_specs.dat  
 extracting: blueprints_extracted/secret_link.txt  
```

### Step 5: Locating the Cipher Fragment

The extraction created a `blueprints_extracted` directory with two files. I checked the nature of both to see if they were readable text.

**Commands:**

```bash
cd blueprints_extracted/
file frame_specs.dat
file secret_link.txt
```

**Output:**

```text
frame_specs.dat: ASCII text
secret_link.txt: ASCII text
```

**Reasoning:** Both were text files. I read the contents of the decoy file first, then the file that sounded like it contained the real data.

**Commands:**

```bash
cat frame_specs.dat
cat secret_link.txt
```

**Output:**

```text
DECOY_DATA_01: Structural keel blueprint verification pending...
PONEGLYPH_FRAGMENT_II="SwnbzptDiM3JSpvFiMuJ28PJzAlJ28VIzA="
```

**Reasoning:** `frame_specs.dat` was explicitly labeled as decoy data. `secret_link.txt` contained the second Poneglyph fragment, matching the format of the fragment found in Level 3.

### Flag / Reward

> `PONEGLYPH_FRAGMENT_II="SwnbzptDiM3JSpvFiMuJ28PJzAlJ28VIzA="`
