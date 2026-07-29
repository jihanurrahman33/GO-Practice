# Go Practice

A collection of small Go example projects and exercises used for learning
Go language concepts and patterns (modules, interfaces, error handling,
HTTP servers, slices, maps, pointers, structs, and more).

## Repository layout

Top-level folders are small, self-contained practice projects. Examples:

- bank-account-demo
- contact_management
- dependency_injection_practice
- error_handling
- inventory-demo
- maps_practice
- net_http_practice
- payment-interface-demo
- pointer-demo
- sales_order
- slice-exercises
- slot-machine-demo
- struct-exercises
- student-grade-demo

Each project generally includes a `main.go` (entrypoint) and a Go module
(`go.mod`) when needed.

## Prerequisites

- Go 1.20+ installed (modules-enabled workflow). Verify with:

```bash
go version
```

## Running a project

From the repository root you can run any example directly with `go run`.

Examples:

```bash
go run ./bank-account-demo
go run ./slot-machine-demo
go run ./net_http_practice
```

Or run a specific file:

```bash
go run ./slice-exercises/main.go
```

To build a binary:

```bash
go build -o bin/bank-account-demo ./bank-account-demo
```

## Tests

Some folders may include tests. Run `go test` in the project directory:

```bash
cd slice-exercises
go test ./...
```

## Contributing

- Add new practice folders at the repo root.
- Include a short `README.md` inside larger exercises explaining how to
	run and what the exercise covers.