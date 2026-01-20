<div align="center">

# 📂 tr2rl
### (Trees to Reality)

**The missing bridge between AI prompts and your filesystem.**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/tr2rl/tr2rl?style=flat-square)](https://goreportcard.com/report/github.com/tr2rl/tr2rl)
[![Release](https://img.shields.io/github/v/release/tr2rl/tr2rl?style=flat-square&color=blue)](https://github.com/tr2rl/tr2rl/releases)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

[Features](#-features) • [Installation](#-installation) • [Usage](#-usage) • [Comparison](#-comparison) • [Contributing](#-contributing)

</div>

---

> **The Problem:** You ask ChatGPT for a project architecture. It gives you a beautiful ASCII tree. You stare at it, sigh, and start manually running `mkdir` and `touch` for 10 minutes.
>
> **The Solution:** Copy the tree. Run `tr2rl`. Done in 1 second.

![Demo GIF](https://via.placeholder.com/800x400?text=Place+Your+Demo+GIF+Here)
*(Example: Copying a tree from ChatGPT and running `tr2rl build --clipboard`)*

---

## ✨ Features

**tr2rl** is a robust, single-binary CLI tool designed to be **"unbreakable"**.

* **🧠 Magic Parsing**
    Smartly understands Unicode trees (`└──`), ASCII trees (`|--`), indented lists, and path lists. It handles mixed indentation and "broken" text gracefully.

* **🛡️ Safety First**
    Defaults to `--dry-run`. You will always see a preview of what files will be created before any changes are written to disk.

* **📋 Clipboard Mode**
    No intermediate files needed. `tr2rl build --clipboard` reads directly from your system's copy buffer.

* **📝 Auto-Populator**
    Don't just create empty files. Use `--populate` to auto-fill files with boilerplate content (e.g., adds `package main` to `.go` files, HTML skeletons to `.html`).

* **🏗️ Template Registry**
    Spin up standard environments instantly with built-in blueprints for React, Go, Python, and more.

* **🚀 Zero Dependencies**
    Written in Go. Distributed as a single, static binary for Windows, Mac, and Linux. No Node_modules or Python venv required.

---

## 📦 Installation

### Option 1: Pre-compiled Binaries (Recommended)
Download the latest binary for your OS from the [**Releases Page**](https://github.com/tr2rl/tr2rl/releases).

| OS | Installation |
|:---|:---|
| **Linux** | `wget [link-to-binary] && chmod +x tr2rl && sudo mv tr2rl /usr/local/bin/` |
| **macOS** | Download binary, `chmod +x tr2rl`, and move to `/usr/local/bin/` |
| **Windows** | Download `.exe` and add to your system `PATH` |

### Option 2: Build from Source (Go)
Requires Go 1.20+ installed.

```bash
git clone [https://github.com/tr2rl/tr2rl.git](https://github.com/tr2rl/tr2rl.git)
cd tr2rl
go build -o tr2rl main.go


🚀 Usage
1. The "Speed Run" (Clipboard)
Copy a directory tree text from ChatGPT, Claude, or DeepSeek, then run:
# Preview the structure (Safe Mode)
tr2rl build --clipboard

# Create files and fill them with boilerplate
tr2rl build --clipboard --populate --dry-run=false


2. From a Text File
If you have saved your structure to spec.txt:
tr2rl build spec.txt ./output_directory


3. Use Built-in Templates
Don't have a tree? Use one of ours.
# List available templates
tr2rl template list

# Generate a React/Vite project
tr2rl template show react-vite | tr2rl build - ./my-app --dry-run=false


4. Format & Clean Trees
Turn a messy, hand-typed list into a clean, professional directory tree string (great for documentation).
tr2rl format messy_list.txt

🧩 Supported Formats
tr2rl is context-aware. It automatically detects and parses these styles:

Style,Example Input
Unicode Tree,├── src/└── main.go
ASCII Tree,`


🆚 Comparison
Feature	📂 tr2rl	🐢 Shell Scripts	🌲 Other Tools
Messy Input	✅ Magic Parser handles anything	❌ Fails on 1 wrong space	❌ Strict JSON required
Safety	✅ Dry-Run by default	❌ Destructive immediately	⚠️ Varies
Content	✅ Auto-populates boilerplate	❌ Creates empty files	❌ Directories only
Portability	✅ Single Binary (No deps)	⚠️ Requires Bash environment	❌ Requires Node/Python

🤝 Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

Fork the Project

Create your Feature Branch (git checkout -b feature/AmazingFeature)

Commit your Changes (git commit -m 'Add some AmazingFeature')

Push to the Branch (git push origin feature/AmazingFeature)

Open a Pull Request

📜 License
Distributed under the MIT License. See LICENSE for more information.





