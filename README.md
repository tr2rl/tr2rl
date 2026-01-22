<div align="center">

# 🌳 tr2rl 🗄️
### (Tree to Reality)

**The CLI utility that turns messy text, directory trees, and indented outlines into real project or directory structures.**
*Build structures from scratch. Refine messy inputs. Automate your setup.*

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/cytificlabs/tr2rl?style=flat-square)](https://goreportcard.com/report/github.com/cytificlabs/tr2rl)
[![Release](https://img.shields.io/github/v/release/cytificlabs/tr2rl?style=flat-square&color=blue)](https://github.com/cytificlabs/tr2rl/releases)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

[Features](#-features) • [Installation](#-installation) • [Examples](#-examples) • [Usage](#-usage) • [Contributing](#-contributing)

<br>

<a href="https://cytificlabs.github.io/tr2rl/">
  <img src="https://img.shields.io/badge/🌐_Website-Live_Demo-blue?style=for-the-badge&logo=github" height="30">
</a>

</div>

---

> **The Scenario:** You're sketching a project architecture in a text file, or maybe ChatGPT just generated a perfect folder structure for you. 
>
> **The Problem:** Manually creating specifically nested folders and files is tedious, error-prone, and slow.
>
> **The Solution:** Feed that text into **tr2rl**. It parses seemingly "broken" or "messy" trees and instantly builds the real file structure for you.

---

## ⚡ Why Use tr2rl? (The Problem Solved)

*   **For AI Users**: ChatGPT and Claude give you ASCII trees. You can't execute them. `tr2rl` lets you **copy-paste-build** in seconds.
*   **For System Architects**: Stop writing 20 lines of `mkdir -p` and `touch` commands. Just draw the tree.
*   **For Tutorials**: Share a lightweight text tree instead of a heavy zip file. Your users can build the repo instantly.

> *Keywords: Directory tree generator, Project scaffolding, ChatGPT to folder structure, Text to filesystem, Golang CLI tool.*

---

## 🐣 Quick Start for Humans

**"I just want to create folders from ChatGPT, Gemini, or any AI."**
1.  Copy the tree text from your AI chat.
2.  Open your terminal in the folder where you want the project.
3.  Run: `tr2rl build --clipboard`
    > *What this does: It reads the text directly from your clipboard and **instantly creates** the real files and folders on your computer.*

**"I want to see what it will do BEFORE it creates files."**
1.  Use the "Dry Run" flag:
2.  Run: `tr2rl build --clipboard --dry-run`
    > *What this does: It shows you a **preview list** of every file that WOULD be created, but it **does not touch your disk** yet. Use this to be safe.*

**"I'm following a tutorial with a big file tree."**
1.  Copy the tree text from the blog post or documentation.
2.  Run: `tr2rl build --clipboard`
    > *What this does: It turns that static text diagram into a real starter project instantly.*

**"I have this messy text from a friend."**
1.  Paste it into a file called `plan.txt`.
2.  Run: `tr2rl build plan.txt`
    > *What this does: It reads the structure from the file `plan.txt` instead of the clipboard.*

---

## ✨ Features

*   **🪄 Magic Parsing**: Smartly understands Unicode trees (`└──`), ASCII trees (`|--`), Indented lists, and even messy, mixed-format text.
*   **🛠️ Build & Populate**: Creates directories and files instantly. Can auto-fill files with language-specific boilerplate (e.g., `package main` for Go).
*   **🧹 Refine & Format**: Takes messy "napkin-sketch" text and formats it into a pristine, professional Unicode tree string.
*   **🛡️ Safety Preview**: Use `--dry-run` to see exactly what *would* happen before it touches your disk.
*   **📋 Clipboard Mode**: Build or format directly from your system clipboard. No temporary files needed.
*   **🚀 Zero Dependencies**: A single, static binary. Runs anywhere (Windows, macOS, Linux).

---

## 🚀 Examples

Here are some real-world commands to get you started.

### 1. The "Grand Finale" Build
Build a complex project structure from a file, verifying it first.
```powershell
# Preview what will be created (Dry Run)
.\tr2rl.exe build testdata/grand_finale.tree --dry-run

# Actually create the structure
.\tr2rl.exe build testdata/grand_finale.tree
```

### 2. Real-World Cookbooks
We've included production-ready architectures in the `examples/` folder.
```powershell
# Microservices with K8s & Docker
.\tr2rl.exe build examples/microservices-k8s.tree ./my-cluster --populate

# Next.js Fullstack (App Router, Tailwind, Supabase)
.\tr2rl.exe build examples/nextjs-fullstack.tree ./my-app --populate

# Python Data Science Project (Cookiecutter style)
.\tr2rl.exe build examples/data-science.tree ./analysis --populate
```

### 3. Instant Clipboard Build
Copy a tree from a chat window or website, then run:
```powershell
.\tr2rl.exe build --clipboard
```

### 3. Cleanup & Formatting
Turn a messy text file into a clean, shareable tree diagram.
```powershell
# Format a file
.\tr2rl.exe format testdata/grand_finale.tree

# Format text currently in your clipboard
.\tr2rl.exe format --clipboard
```

---

## 📦 Installation Guide

### Option 1: Download Binaries (Recommended)
Go to the [**Releases Page**](https://github.com/cytificlabs/tr2rl/releases) and download the file for your system:

| Cloud / OS | Architecture | File to Download | Notes |
|:---|:---|:---|:---|
| **Windows** | Intel/AMD (Standard) | `tr2rl_..._Windows_x86_64.zip` | Most common for PC/Laptop. |
| **Windows** | ARM64 (Snapdragon) | `tr2rl_..._Windows_arm64.zip` | For Surface Pro X, etc. |
| **macOS** | Apple Silicon (M1/M2/M3) | `tr2rl_..._Darwin_arm64.tar.gz` | Most new Macs. |
| **macOS** | Intel | `tr2rl_..._Darwin_x86_64.tar.gz` | Older Macs. |
| **Linux** | Intel/AMD | `tr2rl_..._Linux_x86_64.tar.gz` | Standard servers/desktops. |
| **Linux** | ARM64 | `tr2rl_..._Linux_arm64.tar.gz` | Raspberry Pi 4/5, AWS Graviton. |

**How to use:**
1.  **Unzip/Extract** the downloaded file.
2.  **Move** the `tr2rl` (or `tr2rl.exe`) binary to a folder in your system `PATH` (e.g., `/usr/local/bin` on Linux/Mac).
3.  **Run** `tr2rl --version` to verify.

### Option 2: Build from Source
```bash
git clone https://github.com/cytificlabs/tr2rl.git
cd tr2rl
go build -o tr2rl.exe .
```

---

## 📖 Usage Guide

### `build`
Parses input text and creates the corresponding directory structure.

**Syntax:**
```bash
tr2rl build [file] [output-dir] [flags]
```

**Flags:**
*   `--dry-run`: Enable preview mode (do not write to disk). Default: `false`.
*   `--force`: Overwrite existing files.
*   `--populate`: Auto-fill created files with boilerplate content.
*   `--clipboard`: Read input from clipboard instead of a file.

### `format`
Reads messy input and outputs a clean, canonical Unicode tree. Great for documentation.

**Syntax:**
```bash
tr2rl format [file] [flags]
```

**Flags:**
*   `--style`: Output format. Options: `unicode` (default) or `ascii`.
*   `--clipboard`: Read input from clipboard.

### `template`
View built-in project templates to quick-start your development.

```bash
# List all templates
tr2rl template list

# Use a template (pipe it to build)
tr2rl template show react-vite | tr2rl build - ./my-app --dry-run=false
```

---

## 🧩 Supported Input Formats
tr2rl is context-aware and handles "noisy" input.

| Style | Example Input |
|:---|:---|
| **Unicode Tree** | `├── src/` |
| **ASCII Tree** | `|-- src/` or `+--- src/` |
| **Indented List** | `  src` (just spaces) |
| **Path List** | `root/src/main.go` |

---

## 🤝 Contributing
Contributions are welcome!
1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes
4. Push to the Branch
5. Open a Pull Request

---

## ❤️ Support the Project

If `tr2rl` saved you time, consider supporting the development!

*   ⭐ **Star this repo**: It helps others find the tool.
*   ☕ **Buy me a coffee**: https://ko-fi.com/cosmicquark
*   🗣️ **Share it**: Tell your friends or tweet about it!

---

## 📜 License
Distributed under the MIT License. See `LICENSE` for more information.

<br>

