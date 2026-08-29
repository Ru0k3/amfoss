# Level 1: Awakening at Loguetown Reef

## Objective

Find the genuine Devil Fruit among four sectors of replica files and execute it using the `eat.sh` script to obtain the first flag.

---

## Step 1: Initial Reconnaissance

I started by listing the files in the current directory and reading the `README.md` for instructions.

### Command

```bash
ls
```

### Output

```text
GrandLine
README.md
```

### Command

```bash
cat README.md
```

### Output

```text
# 🏴‍☠️ One Piece Terminal Adventure
...
# LEVEL 1 — AWAKENING AT LOGUETOWN REEF
...
Only by eating the genuine fruit will its true power awaken. (Use ./eat.sh <file_name>)
```

The instructions indicated I needed to find the real fruit and feed it to `./eat.sh`.

---

## Step 2: Analyzing the Execution Script

I navigated into the `Loguetown_Reef` directory and inspected the `eat.sh` script to understand how it validates the fruit.

### Commands

```bash
cd GrandLine/Loguetown_Reef/
cat eat.sh
```

### Output

```bash
#!/bin/bash

if [[ $# -ne 1 ]]; then
    echo "Usage: ./eat.sh <fruit>"
    exit 1
fi

FRUIT="$1"

if [[ ! -f "$FRUIT" ]]; then
    echo "That fruit does not exist."
    exit 1
fi

if [[ -x "$FRUIT" ]]; then
    echo
    echo "------------------------------------------------------------"
    # ... decryption logic ...
    echo "U2FsdGVkX1+uCmgR0ns+u4FKrvHfxhi3bbfdWOB3EpZWwcw4BFHKchR+6/rcU3Xs5EpcM88dLRI49CSbRZK5KQ==" |
    openssl enc -aes-256-cbc -a -d -pbkdf2 \
        -k "GrandLineHistory"
    # ...
else
    echo
    echo "*** CRUNCH! ***"
    echo
    echo "It's just another Marine replica."
    echo "Nothing happens."
fi
```

### Reasoning

The script checks if the provided file argument has the executable permission bit set:

```bash
if [[ -x "$FRUIT" ]]
```

If it does, it runs an OpenSSL decryption command to reveal a flag. If not, it prints a failure message.

Therefore, the real fruit must be an executable file.

---

## Step 3: Searching for the Executable File

There are four sectors (`sector_A`, `sector_B`, `sector_C`, `sector_D`) each containing multiple text files. Instead of manually checking them, I used the `find` command to locate any file with the executable bit set.

### Command

```bash
find . -type f -executable
```

### Output

```text
./sector_C/devil_fruit_6.txt
```

### Reasoning

The search pinpointed the exact location of the genuine file in `sector_C`.

---

## Step 4: Troubleshooting Execution

I initially attempted to run the script from inside `sector_D` without specifying the correct path to the file in `sector_C`.

### Command (from `sector_D`)

```bash
../eat.sh devil_fruit_6.txt
```

### Output

```text
*** CRUNCH! ***

It's just another Marine replica.
Nothing happens.
```

### Problem Faced

Because I was in `sector_D`, the script tried to read `sector_D/devil_fruit_6.txt`, which was a replica.

I needed to target the specific file found by the `find` command.

---

## Step 5: Successful Execution

I moved back to the `Loguetown_Reef` root directory and executed the script with the absolute relative path to the executable file in `sector_C`.

### Command

```bash
./eat.sh ./sector_C/devil_fruit_6.txt
```

### Output

```text
------------------------------------------------------------

*** CRUNCH! ***

The fruit tastes absolutely terrible...

Reality begins to fracture.

Forgotten histories rush into your mind.

You have awakened the legendary...

          Gito Gito no Mi

AWAKENING_SIGNATURE:

ONE_PIECE{GITO_GITO_NO_AWAKENING}
------------------------------------------------------------
```

---

## Flag

> `ONE_PIECE{GITO_GITO_NO_AWAKENING}`
