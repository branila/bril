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
Currently, the best way to install Bril is to clone the repository, compile the source code, and place the executable in a directory accessible from your terminal. Here’s how:

```bash
# Clone the repository
git clone https://github.com/branila/bril.git
cd bril

# Build the project
go build -o bril

# Move the executable to a directory in your PATH
sudo mv bril /usr/local/share/

# Ensure the directory is in your PATH (if not already)
echo 'export PATH=$PATH:/usr/local/share' >> ~/.bashrc
source ~/.bashrc
```
When you run Bril for the first time, a JSON file will be created in the same directory as the executable.
This file serves as the database for your tasks.

## Platform Support
Bril currently works flawlessly on all Linux distributions I have tested.
With a few minor adjustments, it also works on Windows and macOS.
In the future, I plan to ensure full support for all major operating systems, making Bril even more versatile and accessible.

## Usage
Bril offers a variety of commands to manage your tasks. Here’s a detailed overview with syntax and examples:

### 1. Add a Task
**Syntax:**
```bash
bril add <task name> [--priority <int>] [--due <date>] [--tag <string>] [--note <string>]
```
**Example:**
```bash
bril add "Buy groceries" --priority 1 --due "31/12/2024" --tag "shopping" --note "Remember to buy milk"
```
Adds a task named "Buy groceries" with priority 1, a deadline of December 31, 2024, tagged as "shopping," and a note.

### 2. List Tasks
**Syntax:**
```bash
bril list [--all] [--done] [--tag <string>] [--expired] [--deleted] [--sort <name|priority|deadline>]
```
**Examples:**
```bash
bril list --all
```
Lists all tasks, including completed and deleted ones.

```bash
bril list --tag "work" --sort priority
```
Lists tasks tagged as "work," sorted by priority.

### 3. View a Task
**Syntax:**
```bash
bril view <task id>
```
**Example:**
```bash
bril view 123
```
Displays detailed information about the task with id 123.

### 4. Mark a Task as Done
**Syntax:**
```bash
bril do <task id> | bril done <task id>
```
**Example:**
```bash
bril do 123
```
Marks the task with id 123 as completed.

### 5. Undo Task Completion
**Syntax:**
```bash
bril undo <task id>
```
**Example:**
```bash
bril undo 123
```
Reverts the task with id 123 back to an incomplete state.

### 6. Delete a Task
**Syntax:**
```bash
bril rm <task id> | bril delete <task id>
```
**Example:**
```bash
bril rm 123
```
Logically deletes the task with id 123.

### 7. Restore a Deleted Task
**Syntax:**
```bash
bril restore <task id>
```
**Example:**
```bash
bril restore 123
```
Restores the task with id 123 that was previously deleted.

### 8. Set a Reminder
**Syntax:**
```bash
bril remind <task id> <time>
```
**Example:**
```bash
bril remind 123 "15:30"
```
Sets a reminder for the task with id 123 at 3:30 PM.

**Note:** Reminders currently use `notify-send` and are not persistent after reboot. I plan to implement a more robust reminder system in future releases.

Supported time formats:
- `DD/MM/YYYY-HH:MM` (e.g., `31/12/2024-15:30`)
- `HH:MM-DD/MM/YYYY` (e.g., `15:30-31/12/2024`)
- `DD/MM/YYYY` (e.g., `31/12/2024`)
- `DD-MM-YYYY` (e.g., `31-12-2024`)
- `DD/MM` (e.g., `31/12`)
- `HH:MM` (e.g., `15:30`)
- `HH.MM` (e.g., `15.30`)

Supported durations:
- Regex: `^(?:\d+[yMwdhms])+$`
- Example: `1d2h30m` (1 day, 2 hours, 30 minutes)

### 9. Find Tasks
**Syntax:**
```bash
bril find <query> [--all] [--done] [--tag <string>] [--expired] [--deleted] [--sort <name|priority|deadline>]
```
**Example:**
```bash
bril find "groceries" --sort deadline
```
Searches for tasks containing the word "groceries" and sorts them by deadline.

### 10. Reset
**Syntax:**
```bash
bril reset
```
Completely resets the database, removing all tasks.

### 11. Edit a Task
**Syntax:**
```bash
bril edit <task id> [--name <string>] [--state <string>] [--tag <string>] [--priority <int>] [--note <string>] [--deadline <date>]
```
**Example:**
```bash
bril edit 123 --name "Buy vegetables" --priority 2 --deadline "01/01/2025"
```
Edits the task with id 123, changing its name, priority, and deadline.

### 12. Manage Tags
**Syntax:**
```bash
bril tag <command> [args]
```
Supported commands:
- **Add a Tag:**
  ```bash
  bril tag add <tag> [--priority <int>]
  ```
  **Example:**
  ```bash
  bril tag add "urgent" --priority 1
  ```
  The priority of a tag acts as the default priority for tasks assigned that tag, provided no other priority is explicitly set.

- **Delete a Tag:**
  ```bash
  bril tag delete <tag name>
  ```
  **Example:**
  ```bash
  bril tag delete "shopping"
  ```

- **Edit a Tag:**
  ```bash
  bril tag edit <tag name> [--name <string>] [--priority <int>]
  ```
  **Example:**
  ```bash
  bril tag edit "work" --name "office" --priority 2
  ```

- **List Tags:**
  ```bash
  bril tag list
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