# Task Tracker CLI

Task Tracker CLI is a simple command-line application for tracking and managing tasks. This application allows you to add, update, delete, and mark tasks with various statuses. Task data is stored in a JSON file, allowing for persistent task tracking between sessions.
Example task from https://roadmap.sh/projects/task-tracker
## Features

- Add new tasks with a description
- Update task descriptions
- Delete tasks
- Mark tasks as "in-progress" or "done"
- List all tasks or filter by status (e.g., "done", "todo", "in-progress")

## Requirements

- Go (Golang) 1.16 or higher

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/jun-fajr/TaskTrackerGo.git
   cd task-tracker-cli
   ```

2. Compile the code into an executable:

   ```bash
   go build -o task-cli main.go
   ```

3. (Optional) Move `task-cli` to a directory in your `PATH` (e.g., `/usr/local/bin`) for easy access:

   ```bash
   sudo mv task-cli /usr/local/bin/
   ```

## Usage

The following commands demonstrate how to use the `task-cli` application.

### Adding a Task

Add a new task with a description:

```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### Updating a Task

Update the description of an existing task:

```bash
task-cli update 1 "Buy groceries and cook dinner"
# Output: Task updated successfully (ID: 1)
```

### Deleting a Task

Delete a task by its ID:

```bash
task-cli delete 1
# Output: Task deleted successfully (ID: 1)
```

### Marking a Task as In Progress or Done

Mark a task as "in-progress":

```bash
task-cli mark-in-progress 1
# Output: Task marked as in progress (ID: 1)
```

Mark a task as "done":

```bash
task-cli mark-done 1
# Output: Task marked as done (ID: 1)
```

### Listing Tasks

List all tasks:

```bash
task-cli list
```

List tasks by status:

```bash
task-cli list done
task-cli list todo
task-cli list in-progress
```

## Task Properties

Each task has the following properties stored in `tasks.json`:

- **id**: Unique identifier for the task
- **description**: Short description of the task
- **status**: Status of the task (e.g., "todo", "in-progress", "done")
- **createdAt**: Date and time when the task was created
- **updatedAt**: Date and time when the task was last updated

## Error Handling

The CLI handles common errors, such as invalid task IDs, missing commands, and incorrect arguments. Error messages are displayed to guide you on correct usage.

## License

This project is licensed under the MIT License.
```

Feel free to replace `https://github.com/jun-fajr/TaskTrackerGo.git` with your actual GitHub repository URL if you plan to share the project online. Let me know if you need further customization!