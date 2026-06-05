# fn-cleaner

A simple Go executable that renames files in a specified directory by converting their names to lowercase and replacing spaces with dashes.

## How It Works

`fn-cleaner` scans every file in the given directory (skipping subdirectories) and applies two transformations:

1. **Lowercase** — all characters in the filename are converted to lowercase.
2. **Space → Dash** — all space characters are replaced with dashes (`-`).

Files whose names already match the cleaned format are left untouched. The tool reports every rename it performs.

## Prerequisites

- **[Go](https://go.dev/)** must be installed and available on your `PATH`.

## Installation

As root, install the binary system-wide to `/usr/local/bin`:

```bash
./install.sh
```

This will build the Go source and place the compiled binary in `/usr/local/bin/fn-cleaner`.

## Uninstallation

As root, remove the installed binary:

```bash
./uninstall.sh
```

## Usage

Run `fn-cleaner` followed by the path to the directory you want to clean:

```bash
fn-cleaner /path/to/directory
```

### Examples

```bash
# Clean files in the current directory
fn-cleaner .

# Clean files in a specific folder
fn-cleaner ~/Downloads
```

### Output

For every file that is renamed, you will see a line like:

```
Renamed: 'My Document.TXT' -> 'my-document.txt'
```

### Error Handling

- If no directory argument is provided, the tool prints a usage message and exits with code 1.
- If the given path does not exist or is not a directory, an error is printed to stderr and the program exits.
- If a rename fails (e.g., permission denied), the error is printed but processing continues for remaining files.

