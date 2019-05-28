Go API Experiment
---

## Installation

### 1. Install go

Edit `~/.bash_profile`:
```sh
export GOPATH="${HOME}/.go"
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
```

Create necessary folders
```sh
source ~/.bash_profile
mkdir "${GOPATH}"
mkdir -p "${GOPATH}/src/github.com"
```

Install and re-source profile (updates $GOROOT)
```sh
brew install go
source ~/.bash_profile
```

### 2. Install dev libraries

Package Manager
```sh
brew install dep
```

Linter
```sh
go get github.com/golang/lint/golint
```

### 3. Clone project in $GOPATH

```sh
cd $GOPATH/src
git clone git@github.com:arempe93/experiment.git github.com/arempe93/experiment

cd github.com/arempe93/experiment
```

*Why do I have to clone it here?*

> When go is installed it expects everything its needs to be in `$GOPATH`. Compiled commands (like `golint`) are searched for in `$GOPATH/bin`.
> Package imports are searched for in `$GOPATH/src`, and since we are making a "package" in a sense it needs to be in `$GOPATH/src` for go to find
> the import.

*Then why not just clone it to `$GOPATH/src/experiment`?*

> Not sure, it's possible to do that, but community best practices seem to heavily prefer this folder structure. The pros are it's predictable and is
> automatically generated with `go get`. Cons are the imports are quite long: `"github.com/arempe93/experiment"` vs `"experiment"` for example

### 4. Install dependencies

```sh
dep ensure
```

### 5. Change configs

Check `./config.yml` and `./config.go` for options. Values can also be set via environment in this pattern:

```yml
database:
  username: andrew
```

becomes

```sh
export DATABASE_USERNAME="andrew"
```

Environment variables added to `./.env` are included automatically

### 6. Run

```sh
./run.sh
```

### 7. Create an Audit

```sh
http POST http://localhost:7000/api/audits action="my.action" -v
```

### 8. Query Audit

```sh
http GET http://localhost:7000/api/audits/1
```