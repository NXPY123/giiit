# Giiit - Simple Version Control System in Go

Giiit is a lightweight and simple version control system written in the Go programming language. It provides basic version control functionality, allowing you to initialize a repository, commit changes, view commit logs, and tag specific commits.

## Features

### 1. Initialize a Repository

To start using Giiit, initialize a new repository using the following command:

```bash
giiit init
```

This command sets up the necessary directory structure and configuration for version control.

### 2. Commit Changes

Commit your changes to the repository to create a snapshot of your project at a specific point in time. Use the following command:

```bash
giiit commit [project directory] "Your commit message here"
```

This command saves the current state of your files and creates a new commit with the provided commit message.

### 3. View Commit Logs

To view the commit history and log, use the following command:

```bash
giiit log [project directory]
```

This command displays a list of all commits, showing commit hashes, author details, commit messages, and timestamps.

### 4. Tagging Commits

Tagging allows you to mark specific commits for easy reference. To tag a commit, use the following command:

```bash
giiit tag [project directory] [tag name] [commit hash]
```

This example creates an annotated tag with the name "v1.0" and an associated message.

### 5. Creating Branches

To create a new branch, use the following command:

```bash
giiit branch [project directory] [branch name]
```

## Getting Started

1. Clone the Giiit repository:

```bash
git clone https://github.com/NXPY123/giiit.git
cd giiit
```

2. Build the Giiit executable:

```bash
go build -o giiit ./giiit/
```

3. Add the Giiit executable to your system's PATH or use it directly from the project directory.

