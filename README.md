# ⚙️ MWF (Windows Management Framework)
Make. Windows. Faster.⚡⚡

MWF is a Go-based CLI tool for safely applying Windows system configuration profiles focused on performance, startup optimization, power management, and developer-friendly system setups.

It is designed as a deterministic, safety-first execution framework for Windows tuning — not a script collection.

---

## 🚀 Current Version: v0.2 (Early System Control Layer)

This version focuses on:

- Profile-based system configuration (single default profile)
- Power plan switching
- Startup application management
- Basic Windows settings adjustments
- Safe execution pipeline (Planner → Validator → Executor)
- Dry-run preview mode

> ⚠️ MWF is under active development and is not yet production-stable.

---

## 🎯 Goals

MWF aims to become a structured Windows configuration framework that:

- Applies system optimizations safely and predictably
- Avoids destructive or undocumented system modifications
- Provides transparent execution of all changes
- Supports reproducible system setup via declarative profiles

---

## ⚙️ Features (v0.2 Scope)

### ✅ Implemented
- CLI-based profile execution
- Default system profile (`default`)
- Power plan switching
- Startup application management
- Basic Windows configuration adjustments
- Dry-run mode (preview changes without applying)
- Deterministic execution pipeline:
  - Planner → Validator → Executor

### 🚧 In Progress / Not Yet Included
- Registry modifications
- System diagnostics ("doctor mode")
- Multiple built-in profiles
- Rollback / snapshot system
- Plugin system

---

## 📦 Installation

### Prerequisites
- Go 1.22+
- Windows 10/11
- Administrator privileges (required for system changes)

### Build from source

```bash
git clone https://github.com/yourname/mwf.git
cd mwf
go mod tidy
go build -o mwf.exe ./cmd/mwf
```

## 🧪 Usage

View available commands
`mwf`
List profiles
`mwf list`
Apply default profile
`mwf apply default`
Dry run (preview changes)
`mwf apply default --dry-run`

## 📄 Profile System

MWF uses YAML-based profiles to define system configuration actions.

In v0.2, only the `default` profile is included.

Example: `default.yaml`
```name: default
description: Baseline system configuration

actions:
  - type: power_plan
    value: balanced

  - type: startup_cleanup
    target:
      - OneDrive
      - Teams

  - type: windows_setting
    target: visual_effects
    value: default
```

## 🧠 Architecture

MWF uses a deterministic execution pipeline to ensure safe and predictable system changes.

Execution Flow 
```
Profile YAML
    ↓
Planner (converts profile → actions)
    ↓
Validator (checks safety & correctness)
    ↓
Executor (applies system changes)
```

## Components
### Planner
Parses YAML profiles
Converts configuration into structured system actions
Ensures consistent internal representation
### Validator
Ensures actions are safe to execute
Checks required permissions
Blocks unsupported or invalid operations
### Executor
Applies validated actions to the system
Handles:
Power plan changes
Startup modifications
Windows configuration updates

## 🛡️ Safety Model

MWF is designed with safety as a core principle:

⚠️ **Dry-run mode available for all operations**
🔒 No hidden or implicit system changes
🧾 Every action is explicitly logged before execution
🧠 Validator prevents unsafe operations from executing

## 📁 Project Structure
```
mwf/
│
├── cmd/mwf/              # CLI entry point
├── internal/
│   ├── cli/              # Command definitions
│   ├── engine/           # Planner / Validator / Executor
│   ├── system/           # Windows system operations
│   │   ├── power/        # Power plan management
│   │   ├── startup/      # Startup app management
│   │
│   ├── profiles/         # YAML profile loader
│   ├── safety/           # Dry-run & validation logic
│   ├── logging/          # Structured logging
│   └── types/            # Shared data models
│
├── configs/
│   └── profiles/
│       └── default.yaml
│
├── go.mod
└── README.md
```

## 🧪 Example Workflow
1. Check available profiles
mwf list
2. Preview system changes
mwf apply default --dry-run
3. Apply configuration
mwf apply default

## 🧭 Design Principles

MWF follows these core principles:

Deterministic execution — same input always produces same actions
Explicit over implicit — nothing happens without being declared
Safety-first design — validation before execution
Modular architecture — system logic separated from CLI layer
Windows-first, extensible later

## 📌 Roadmap
**v0.2 (Current)**
CLI framework
Default profile support
Power plan management
Startup management
Basic Windows settings
Dry-run execution pipeline

**v0.3**
Registry support
Snapshot/rollback system
Expanded system actions

**v0.4**
System diagnostics (doctor mode)
Performance profiling tools

**v1.0**
Stable release
Full Windows configuration framework
Production-grade safety and logging system

## ⚠️ Disclaimer

MWF modifies system-level Windows settings.

Use at your own risk. Always run:

mwf apply default --dry-run

before applying changes.

## 🤝 Contributing

This project is currently in early development.

Guidelines:

Follow Go idioms and `internal/` architecture
Keep system logic isolated from CLI layer
All system changes must go through the engine pipeline

## 📜 License

Apache License 2.0
