# Task Tracker

This is my solution to the [Task Tracker](https://roadmap.sh/projects/task-tracker) project on roadmap.sh using Go without any third-party packages.

## Requirements

- [Go](https://go.dev/dl/) 1.27.1 or newer
- No third-party packages. Everything uses the Go standard library.

## Build

```sh
cd task-tracker
go build -o task-tracker .
```

This creates a `task-tracker` binary in the current directory.

To run the tests:

```sh
go test ./...
```

## Usage

```sh
./task-tracker <command> [arguments]
```

Tasks are stored in `tasks.json` in the directory you run the command from. The file is created automatically the first time you add a task.

| Command | What it does |
|---|---|
| `add <description>` | Add a new task with status `todo` |
| `update <id> <description>` | Change a task's description |
| `delete <id>` | Remove a task |
| `mark-in-progress <id>` | Set a task's status to `in-progress` |
| `mark-done <id>` | Set a task's status to `done` |
| `list` | Show all tasks |
| `list <status>` | Show only tasks with that status (`todo`, `in-progress`, `done`) |

### Example

```sh
$ ./task-tracker add "Buy groceries"
Task added successfully (ID: 1)

$ ./task-tracker add "Write README"
Task added successfully (ID: 2)

$ ./task-tracker mark-in-progress 1
Task marked as in-progress (ID: 1)

$ ./task-tracker mark-done 2
Task marked as done (ID: 2)

$ ./task-tracker update 1 "Buy groceries and cook dinner"
Task updated successfully (ID: 1)

$ ./task-tracker list
 ID STATUS      DESCRIPTION
  1 in-progress Buy groceries and cook dinner
  2 done        Write README

$ ./task-tracker list todo
No tasks with status "todo".

$ ./task-tracker delete 9
task not found for id = 9
```

Errors are printed to stderr, and the program exits with status code `1`.

Each task in `tasks.json` looks like this:

```json
{
  "id": 1,
  "description": "Buy groceries and cook dinner",
  "status": "in-progress",
  "createdAt": "2026-10-20T23:48:23.410112+07:00",
  "updatedAt": "2026-10-20T23:48:23.430588+07:00"
}
```

## How I approached it

This section describes how I thought through the problem. If you're building your own version, in any language, it might be a useful reference.

### 1. Start with the problem

Before thinking about commands or files, I tried to understand what I was building. I started with three questions:

- **What are the things?** A *task*. It has a description and a status, which is one of `todo`, `in-progress`, or `done`. It's also useful to know when a task was created and when it was last changed.
- **What can people do with them?** Add a task, see their tasks (all of them or by status), change a task's description or status, and remove a task.
- **What are the rules?** A task must have a description; blank text doesn't count. A task can only be in one of the three statuses. You can only change or remove a task that actually exists.

### 2. Split the work by responsibility

Looking at those answers, I noticed the project involves several different kinds of questions. I gave each one its own file:

| Question | File |
|---|---|
| **What is a single task, and what rules does it follow?** For example, a description can't be empty, and there are only three valid statuses. | `task.go` |
| **What can be done with a group of tasks?** Find one by its ID, give a new task the next ID, remove one, keep only the ones with a certain status. | `tasks.go` |
| **How are tasks remembered?** A CLI program starts, runs one command, and exits, and everything in memory is lost when it exits. So the tasks have to be saved somewhere and loaded again on the next run. | `storage.go` |
| **How does the user talk to the program?** Reading the command and its arguments, and printing a result or an error. | `main.go` |

Each test file (`task_test.go`, `tasks_test.go`, `storage_test.go`, `main_test.go`) sits next to the file it tests.

I kept these separate because each one changes for a different reason:

- Switching from a JSON file to a database shouldn't change the rules about what makes a valid task.
- Replacing the command line with a web page or a GUI shouldn't change how tasks are saved.
- Adding a new rule, such as a maximum description length, shouldn't require changes to the storage or the CLI.

With each part handling only its own job, I could change one part without touching the others, and read each part on its own.

### 3. Build features in the order they depend on each other

I didn't build the commands in the order the project lists them. Instead, I asked: *what does this command need in order to exist and be checked?*

1. **`add` first.** Every other command works on tasks that already exist, so without a way to create a task, there's nothing to list, update, or delete.
2. **`list` next.** It's how I see the data. Without it, it's hard to check whether the other commands did what I expected, so having it early made everything after it easier to verify.
3. **Then the commands that change an existing task: `delete`, `update`, `mark-in-progress`, `mark-done`.** These share the same basic shape: *point at a task by its ID, then do something to it*. So they share the same questions: what if the ID isn't a number, and what if no task has that ID? Within this group I went from simplest to most involved:
   - `delete` only needs an ID.
   - `update` needs an ID **and** a new description, and that description has to follow the same rules as in `add`.
   - `mark-in-progress` and `mark-done` are variations of "update", except they change the status instead of the description.

For each command, I looked at it from the user's side:

- **What does the user type?** What arguments are expected, and what if there are too few or too many?
- **What can go wrong?** Empty text, an ID that isn't a number, an ID that doesn't exist, a status that doesn't exist.
- **What does the user see when it goes wrong?** A message that says what was wrong, such as `task not found for id = 9`.
- **Is the data still safe if the command fails?** A failed command should leave the saved tasks exactly as they were.

### Summary

In short, my process was:

1. **Understand** the things, the actions, and the rules.
2. **Split** the work into parts that each have one reason to change.
3. **Build** features in the order they depend on each other, and for each one think about what the user types, what can go wrong, and what they see.
