# Go Commands Cheat Sheet

This is a handy reference for common `go` commands used during development.

---

## 🚀 Running and Building

| Command               | Description                                 |
|-----------------------|---------------------------------------------|
| `go run main.go`      | Run a Go file directly                      |
| `go build`            | Compile the code and create an executable   |
| `go install`          | Build and install to `$GOPATH/bin`          |

---

## 🧹 Formatting and Linting

| Command               | Description                                 |
|-----------------------|---------------------------------------------|
| `go fmt`              | Format your Go source code                  |
| `go vet`              | Report suspicious code (static analysis)    |

---

## 🧪 Testing

| Command               | Description                                 |
|-----------------------|---------------------------------------------|
| `go test`             | Run tests in the current directory          |
| `go test -v`          | Verbose output during tests                 |

---

## 📦 Dependency Management (Modules)

| Command                   | Description                                   |
|---------------------------|-----------------------------------------------|
| `go mod init <module>`    | Initialize a new module (creates `go.mod`)    |
| `go mod tidy`             | Add missing and remove unused modules         |
| `go list -m all`          | List all modules used in your project         |

---

## 📖 Documentation and Help

| Command                    | Description                                     |
|----------------------------|-------------------------------------------------|
| `go doc fmt.Println`       | Show documentation for a specific function      |
| `go help`                  | Display help for Go commands                    |
| `go env`                   | Print Go environment info                       |
| `go version`               | Show installed Go version                       |

---

## 🧼 Cleanup

| Command       | Description                            |
|----------------|----------------------------------------|
| `go clean`     | Remove object files and cached files   |

---

Use these commands as you work on your Go projects to speed up development and maintain code quality.
