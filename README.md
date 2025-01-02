# Bril: The Terminal-Centric Daily Task Manager on Steroids

Welcome to **Bril**, the task manager that's so fast, minimal, and terminal-friendly it might just make you forget about all those fancy GUI tools.
Built by **Branila Claudiu Stefan**, a high school student from ITIS P. Paleocapa in Bergamo (Italy), Bril is crafted with love (and a lot of Go) to help you manage your tasks without ever leaving your beloved terminal.
Because let’s face it: if you are reading this documentation from github, you most likely have a terminal open somewhere.

## Why Bril?
Bril is perfect for anyone who:
- Wants a no-nonsense, lightning-fast task manager.
- Prefers typing commands over clicking buttons.
- Appreciates simplicity, minimalism, and immediacy.
- Loves open-source tools and enjoys tinkering with code.

## Key Features
- **Command-line simplicity**: Manage tasks with intuitive commands. Bril works as fast as you can type.
- **Minimal dependencies**: Written entirely in Go, using only the standard library.
- **Flexible task management**: Add, list, edit, delete, restore, and even remind yourself about tasks.
- **Custom tagging**: Organize tasks with tags, deadlines, priorities and more.
- **Future-proof**: Planned features include cloud database integration and a Telegram bot interface.

## Installation

### Prerequisites

Before installing Bril, ensure you have:
- Go installed on your system (latest version)
- Git for cloning the repository
- Appropriate system permissions for installing executables

### Installation Steps

1. First, clone the repository and navigate to the project directory:
```bash
git clone https://github.com/branila/bril.git
cd bril
```

2. Build the project:
```bash
go build -o bril
```

3. Install the executable:
```bash
# On Unix-like systems (Linux/macOS)
sudo mv bril /usr/local/bin/

# On Windows (run PowerShell as Administrator)
move bril.exe C:\Windows\System32\
```

### Database Location

When you run Bril for the first time, it will automagically create its database in your system's user configuration directory:

- Linux: `~/.config/bril/bril.json`
- macOS: `~/Library/Application Support/bril/bril.json`
- Windows: `%AppData%\bril\bril.json`

### Verifying the Installation

To verify that Bril was installed correctly:

1. Open a new terminal window (this ensures your PATH is updated)
2. Run:
```bash
bril list
```
You should see an empty task list, indicating that Bril is working correctly.

### Troubleshooting

If you encounter any issues:

1. Ensure the executable is in your system PATH:
```bash
# On Unix-like systems
which bril

# On Windows
where bril
```

2. Check that you have write permissions in your user configuration directory:
```bash
# On Unix-like systems
ls -la ~/.config/bril/

# On Windows
dir %AppData%\bril
```

3. If Bril isn't creating its database, ensure you have adequate permissions in your user configuration directory.

### Uninstalling

To uninstall Bril:

1. Remove the executable:
```bash
# On Unix-like systems
sudo rm /usr/local/bin/bril

# On Windows (run PowerShell as Administrator)
del C:\Windows\System32\bril.exe
```

2. Remove the database and configuration directory:
```bash
# On Linux
rm -rf ~/.config/bril

# On macOS
rm -rf ~/Library/Application\ Support/bril

# On Windows
rmdir /s /q %AppData%\bril
```

## Platform Support
Bril is designed to work seamlessly across all major operating systems.
The application automagically adapts to your operating system, storing data in the appropriate system-specific configuration directories and using platform-specific commands where necessary.

## Usage
Bril offers a variety of commands to manage your tasks. Here’s a detailed overview with syntax and examples:

### 1. Add a Task
**Syntax:**
```bash
bril add <task name> [--priority <int> | -p <int>] [--due <date> | -d <date>] [--tag <string> | -t <string>] [--note <string> | -n <string>]
```
**Example:**
```bash
bril add "Try bril" --priority 10 --due "13/11/2025" --tag "important" --note "It will change your life"
```
Adds a task named "Buy groceries" with priority 1, a deadline of December 31, 2024, tagged as "shopping," and a note.

### 2. List Tasks
**Syntax:**
```bash
bril list [--all | -a] [--done | -d] [--tag <string> | -t <string>] [--expired | -e] [--deleted | -x] [--sort <name|priority|deadline> | -s <name|priority|deadline>]
# or
bril ls [--all | -a] [--done | -d] [--tag <string> | -t <string>] [--expired | -e] [--deleted | -x] [--sort <name|priority|deadline> | -s <name|priority|deadline>]
```
**Examples:**
```bash
bril list --all
```
Lists all tasks, including completed and deleted ones.

```bash
bril list --tag "work" --sort priority
```
Lists tasks tagged as "work", sorted by priority.

### 3. View a Task
**Syntax:**
```bash
bril view <task id>
```
**Example:**
```bash
bril view 13
```
Displays detailed information about the task with id 13.

### 4. Mark a Task as Done
**Syntax:**
```bash
bril do <task id>
# or
bril done <task id>
```
**Example:**
```bash
bril do 13
```
Marks the task with id 13 as completed.

### 5. Undo Task Completion
**Syntax:**
```bash
bril undo <task id>
```
**Example:**
```bash
bril undo 13
```
Reverts the task with id 13 back to an incomplete state.

### 6. Delete a Task
**Syntax:**
```bash
bril rm <task id>
# or
bril delete <task id>
```
**Example:**
```bash
bril rm 13
```
Logically deletes the task with id 13.

### 7. Restore a Deleted Task
**Syntax:**
```bash
bril restore <task id>
```
**Example:**
```bash
bril restore 13
```
Restores the task with id 13 that was previously deleted.

### 8. Set a Reminder
**Syntax:**
```bash
bril remind <task id> <time>
```
**Example:**
```bash
bril remind 13 "15:30"
```
Sets a reminder for the task with id 13 at 3:30 PM. At the moment bril only supports european time format.

**Note:** Reminders currently use `notify-send` and are not persistent after reboot. I plan to implement a more robust reminder system in future releases.

Supported durations:
- Regexp: `^(?:\d+[yMwdhms])+$`
- Example: `1d2h30m` (1 day, 2 hours, 30 minutes)

### 9. Find Tasks
**Syntax:**
```bash
bril find <query> [--all | -a] [--done | -d] [--tag <string> | -t <string>] [--expired | -e] [--deleted | -x] [--sort <name|priority|deadline> | -s <name|priority|deadline>]
```
**Example:**
```bash
bril find "rocket science" --sort deadline
```
Searches for tasks containing "rocket science" and sorts them by deadline.

### 10. Reset
**Syntax:**
```bash
bril reset
```
Completely resets the database, removing all tasks.

### 11. Edit a Task
**Syntax:**
```bash
bril edit <task id> [--name <string> | -n <string>] [--state <done|deleted|undone|active>] [--tag <string> | -t <string>] [--priority <int> | -p <int>] [--note <string> | -m <string>] [--deadline <date> | -n <date>]
```
**Example:**
```bash
bril edit 13 --name "Win a nobel prize" --priority 10 --deadline "next year"
```
Edits the task with id 123, changing its name, priority, and deadline.

### 12. Manage Tags
**Syntax:**
```bash
bril tag <command> [args]
```
Supported tag commands:
- **Add a Tag:**
  ```bash
  bril tag add <tag> [--priority <int> | -p <int>]
  ```
  **Example:**
  ```bash
  bril tag add "urgent" --priority 10
  ```
  The priority of a tag acts as the default priority for tasks assigned that tag, provided no other priority is explicitly set.

- **Delete a Tag:**
  ```bash
  bril tag delete <tag name>
  # or
  bril tag rm <tag name>
  ```
  **Example:**
  ```bash
  bril tag delete "shopping"
  ```

- **Edit a Tag:**
  ```bash
  bril tag edit <tag name> [--name <string> | -n <string>] [--priority <int> | -p <int>]
  ```
  **Example:**
  ```bash
  bril tag edit "work" --name "office" --priority 2
  ```

- **List Tags:**
  ```bash
  bril tag list
  # or
  bril tag ls
  ```
  **Example:**
  ```bash
  bril tag list
  ```

## Contributing
Bril is open source, and contributions are welcome! Here’s how you can help:
1. Fork the repository on GitHub: [github.com/branila/bril](https://github.com/branila/bril).
2. Clone your fork and create a new branch for your changes.
3. Submit a pull request with a clear description of your changes.

Whether it’s fixing bugs, adding features, or improving documentation, your help is greatly appreciated.

## Roadmap
Exciting features planned for future releases:
- **Cloud Database Integration**: Sync your tasks across devices.
- **Telegram Bot Interface**: Manage tasks directly from your favorite messaging app.
- **Persistent Reminders**: Reliable notifications powered by `cron`.

Stay tuned for updates, and don’t forget to star the project on GitHub if you find it useful.

Now go forth and be productive :)
