# Go Bootstrapper 🐹⚡

[![Go Version](https://img.shields.io/badge/Go-1.22-blue)](https://go.dev/)
[![Build Status](https://github.com/upsaurav12/bootstrapper/actions/workflows/go.yml/badge.svg)](https://github.com/upsaurav12/bootstrapper/actions)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

A **CLI tool** to scaffold Go projects quickly — think of it like [Vite](https://vitejs.dev/), but for **Golang developers**.  
It saves you from boilerplate, folder structure confusion, and manual dependency setup.

---

## ✨ Features
- 🚀 Create new Go projects in seconds
- 📦 Preconfigured templates (REST API with Gin, Chi, etc.)
- 🗂 Standardized folder structure (`cmd/`, `internal/`, `pkg/`)
- 🔌 Incrementally add dependencies (DB, gRPC, logging, etc.)
- 🧪 Built-in **Makefile** for build, test, and lint
- ⚙️ GitHub Actions workflow for CI/CD
- 🛠 Extensible via templates

---

## 📂 Example Project Structure

When you run `bootstrap new myapp --type=rest --router=gin --port=9000`,  
you’ll get something like this:

