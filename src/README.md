
# Go REST API Setup Guide

This guide walks you through creating a modular Go REST API with a backend package and a separate entry point.

## Step 1: Create Backend Folder

### Guidance
The backend folder will contain your HTTP server logic. This module is designed to be reusable and can be imported by other packages.

**What you'll do:**
- Create a new directory called `backend` to house your server code
- Initialize a Go module for the backend package
- This creates a `go.mod` file that manages dependencies for the backend

```bash
mkdir backend
cd backend
go mod init example.com/backend
```

### backend.go Implementation

The `backend.go` file contains the HTTP server setup and route handlers:

```go
package backend

import (
  "fmt"
  "log"
  "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World!")
}

func Run(addr string) {
  http.HandleFunc("/", helloWorld)
  fmt.Println("Starting server on: ", addr)
  log.Fatal(http.ListenAndServe(addr, nil))
}
```

**Key Components:**
- **helloWorld handler** - Responds with "Hello, World!" to HTTP requests
- **Run function** - Initializes the HTTP server on a specified address (must be exported with capital R)
- **HandleFunc** - Routes the root path "/" to the helloWorld handler
- **ListenAndServe** - Starts the server and listens for incoming connections

**Important:** The `Run` function must start with a capital letter to be exported and accessible from other packages.

---

## Step 2: Create Practice Source Directory

### Guidance
The practice_src (or practice source) directory is your main application entry point. It will import and use the backend module you created in Step 1.

**What you'll do:**
- Create a `practice_src` directory for your main application
- Initialize a separate Go module for this entry point
- This module will depend on the backend module

```bash
cd ..
mkdir practice_src
cd practice_src
go mod init example.com/practice_src
```

### practice_src.go Implementation

The main entry point that imports and uses the backend module:

```go
package main

import (
  "example.com/backend"
)

func main() {
  backend.Run(":8080")
}
```

**Key Components:**
- **Package declaration** - `package main` marks this as an executable program (required for `go run`)
- **Import statement** - Imports the backend module by its module name
- **Main function** - Entry point that starts the server on port 8080
- **backend.Run()** - Calls the exported Run function from the backend package

**Important:** Only one `package main` per application is allowed, and it must have a `main()` function.

---

## Step 3: Link Modules and Run

### Guidance
Now you need to tell the Go module system where to find the backend module. Since it's in a local directory (not published to a registry), you'll use a `replace` directive to map the module name to the local path.

**What you'll do:**
- Edit the go.mod file to replace the backend module import path with a local directory path
- Run `go mod tidy` to clean up and verify dependencies
- Run the application with `go run .`

**Why replace?** The `replace` directive tells Go: "When you see `example.com/backend`, use the code from `../backend` instead of looking for it online."

```bash
go mod edit -replace example.com/backend=../backend
go mod tidy
go run .
```

### go.mod Configuration

After running the `go mod edit -replace` command, your `go.mod` file will look like:

```go
module example.com/practice_src

go 1.18

replace example.com/backend => ../backend

require example.com/backend v0.0.0-00010101000000-000000000000
```

**Explanation:**
- **replace directive** - Maps the backend module to the local `../backend` directory
  - Format: `replace <module-path> => <local-path>`
  - This is local development only; don't commit replace directives for published modules
- **require statement** - Declares the dependency on the backend module
  - The version `v0.0.0-00010101000000-000000000000` is a placeholder for local modules

---

## Step 4: Test API (Different Terminal)

### Guidance
Once your server is running (from Step 3), open a new terminal to test the API without stopping the server.

**What you'll do:**
- Use `curl` to make an HTTP GET request to `localhost:8080`
- Verify that the server responds with "Hello, World!"
- This confirms your server setup is working correctly

```bash
curl localhost:8080
```

**Expected Output:**
```
Hello, World!
```

---

## Summary

You've successfully built a modular Go REST API structure:
1. **Backend module** - Contains reusable HTTP server code
2. **Practice_src application** - Main entry point that uses the backend
3. **Local module linking** - Uses `replace` directive for local development
4. **Running server** - Accessible at `http://localhost:8080`

This architecture allows you to maintain separation of concerns and reuse the backend module in multiple applications.

```bash
curl localhost:8080
```

**Expected Output:**

>Hello, World!

