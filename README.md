# bootstrap-cli
A scaffold CLI for Golang projects that helps you create new projects **without worrying about installing dependencies, setting up linters, or boilerplate configuration**.  
Just use the CLI to generate **pre-configured templates** and start building instantly.

---

## Installation

Make sure you have [Go](https://golang.org/doc/install) installed (Go 1.20+ recommended).  
Then install the CLI:

```bash
go install github.com/upsaurav12/bootstrap@latest

```

## Usage
```bash
bootstrap new <your_project_name --type=rest --router=gin --port=8080
```

- ``--type`` : what kind of project ?.
- ``--router`` : framework (i.e gin or chi ).
- ``--port`` : what should be your port where your server would be running.

```base
cd <your_project_name> && go mod tidy
```
this will give take you in your project directory and install the dependencies, here we go your project is now ready.


## TODO's
- [ ] Add more project types like microservices, cli-app, rest, etc.
- [ ] Add more frameworks in REST like echo, fiber
- [ ] Add flags for automating setup of linters, CI/CD pipelines, tests, etc.
- [ ] Add `add` command for adding features like auth, db, services, etc.
- [ ] Add `update` command for automating available dependency update.




