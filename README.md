# ⚙️ MWF (Windows Management Framework)

# Make. Windows. Faster. ⚡⚡

MWF is a Go-based CLI tool for safely applying Windows system configuration profiles focused on performance, startup optimization, power management, and developer-friendly system setups.

It is designed as a deterministic, safety-first execution framework for Windows configuration and tuning — **not a script collection**.

---

# 🚀 Current Version: v0.3 — Stateful System Management Foundation

MWF v0.3 introduces the foundation required for a real Windows management framework.

This version focuses on:

- Profile-based system configuration
- Typed system actions
- Risk-aware execution
- Power plan management
- Startup application management
- Windows configuration actions
- Registry support foundation
- Snapshot and restore system
- Dry-run preview mode
- Planner → Validator → Snapshot → Executor pipeline

> ⚠️ MWF is under active development and is not yet production-stable.

---

# 🎯 Goals

MWF aims to become a structured Windows configuration framework that:

- Applies system optimizations safely and predictably
- Avoids destructive or undocumented modifications
- Provides transparent execution of all changes
- Supports reproducible system setup through declarative profiles
- Allows users to preview and rollback changes

---

# ⚙️ Features

## ✅ Implemented

### Profile System

- YAML-based configuration profiles
- Default profile support
- Declarative system actions

---

### Execution Engine

MWF uses a deterministic execution pipeline:

```
Profile YAML
      ↓
Planner
      ↓
Validator
      ↓
Snapshot
      ↓
Executor
      ↓
Reporter
```

---

### Supported System Actions

Current action categories:

- Power plan management
- Startup application management
- Windows settings framework
- Registry foundation

---

### Safety Features

- Dry-run mode
- Action validation
- Risk levels
- Snapshot creation before changes
- Restore through reverse actions
- No hidden system modifications

---

# 📦 Installation

## Recommended Installation

Install MWF using the PowerShell installer:

```powershell
irm https://raw.githubusercontent.com/J4Edev/MWF/main/installers/install.ps1 | iex
```

The installer:

- Installs MWF for the current user
- Places the executable in:

```
%LOCALAPPDATA%\Programs\MWF
```

- Adds MWF to the user PATH

Restart PowerShell after installation.

Verify:

```powershell
mwf --help
```

---

## Build From Source

### Requirements

- Go 1.22+
- Windows 10/11

Clone:

```bash
git clone https://github.com/J4Edev/MWF.git
cd MWF
```

Install dependencies:

```bash
go mod tidy
```

Build:

```bash
go build -o mwf.exe ./cmd/mwf
```

---

# 🧪 Usage

## View available commands

```bash
mwf
```

---

## List profiles

```bash
mwf list
```

Example:

```
Available profiles:

- default
```

---

## Apply a profile

```bash
mwf apply default
```

---

## Dry run

Preview changes without modifying the system:

```bash
mwf apply default --dry-run
```

Dry-run will:

- Load the profile
- Validate actions
- Generate execution report
- Skip system writes

---

## Create a backup

```bash
mwf backup default
```

Creates a snapshot of the current system state before changes.

---

## Restore a snapshot

```bash
mwf restore <snapshot-id>
```

Restoration works by generating reverse actions and sending them through the normal execution pipeline.

---

# 📄 Profile System

MWF uses YAML-based profiles to define system configuration actions.

Example:

`default.yaml`

```yaml
name: default

risk:
  level: low

actions:

  - type: power_plan
    value: balanced

  - type: startup_remove
    target:
      - OneDrive
      - Teams

  - type: windows_setting
    target:
      - example_setting
```

Profiles define **what should happen**.

The engine determines **how it happens safely**.

---

# 🧠 Architecture

MWF separates system logic from execution logic.

## Execution Flow

```
Profile YAML
      ↓
Planner
      ↓
Validator
      ↓
Snapshot
      ↓
Executor
      ↓
Reporter
```

---

# Components

## Planner

Responsible for:

- Parsing profiles
- Converting YAML into actions
- Creating deterministic execution plans

---

## Validator

Responsible for:

- Checking action validity
- Checking permissions
- Evaluating risk
- Blocking unsupported operations

---

## Snapshot System

Responsible for:

- Capturing current state
- Storing previous values
- Creating rollback information

---

## Executor

Responsible for:

- Running validated actions
- Communicating with Windows system managers

---

# 🛡️ Safety Model

MWF is designed with safety as a core principle.

Features:

- ⚠️ Dry-run mode available before execution
- 🔒 No hidden system modifications
- 🧾 Actions are explicitly planned and reported
- 🔄 Changes can be restored using snapshots
- 🧠 Validation occurs before execution

---

# 📁 Project Structure

```
MWF/

├── cmd/
│   └── mwf/
│       └── main.go

├── internal/
│
│   ├── actions/
│   │   └── action.go
│
│   ├── cli/
│   │   ├── apply.go
│   │   ├── backup.go
│   │   └── restore.go
│
│   ├── engine/
│   │   ├── planner.go
│   │   ├── validator.go
│   │   ├── executor.go
│   │   ├── pipeline.go
│   │   └── reporter.go
│
│   ├── snapshot/
│   │   ├── manager.go
│   │   └── models.go
│
│   ├── system/
│   │   ├── registry/
│   │   ├── power/
│   │   ├── startup/
│   │   └── windows/
│
│   ├── profiles/
│
│   └── logging/
│
├── configs/
│   └── profiles/
│       └── default.yaml
│
├── installers/
│   └── install.ps1
│
├── go.mod
└── README.md
```

---

# 🧪 Example Workflow

### 1. View available profiles

```bash
mwf list
```

---

### 2. Preview changes

```bash
mwf apply default --dry-run
```

---

### 3. Create snapshot

```bash
mwf backup default
```

---

### 4. Apply configuration

```bash
mwf apply default
```

---

### 5. Restore if needed

```bash
mwf restore <snapshot-id>
```

---

# 🧭 Design Principles

MWF follows these principles:

- **Deterministic execution** — same input produces the same actions
- **Explicit over implicit** — nothing happens without being declared
- **Safety-first design** — validation before execution
- **Modular architecture** — system logic separated from CLI logic
- **Windows-first, extensible later**

---

# 📌 Roadmap

## v0.3 ✅ Current

- Typed action system
- Risk levels
- Registry foundation
- Snapshot system
- Restore pipeline
- Expanded system architecture
- Installer support

---

## v0.4

Planned:

- Expanded Windows settings
- Improved reporting
- More system actions
- Better profile management

---

## v0.5

Planned:

- Doctor mode
- System diagnostics
- Performance analysis

---

## v1.0

Goal:

A complete Windows configuration framework featuring:

- Stable profiles
- Safe rollback
- Extensible actions
- Production-grade reliability

---

# ⚠️ Disclaimer

MWF modifies system-level Windows settings.

Always review changes first:

```powershell
mwf apply default --dry-run
```

before applying modifications.

Use at your own risk.

---

# 🤝 Contributing

This project is currently in active development.

Guidelines:

- Follow Go idioms
- Keep system logic isolated from CLI logic
- Route all system changes through the execution pipeline
- Prefer safe, reversible changes

---

# 📜 License

Apache License 2.0