# GoLang Install

First, you'll want to export paths in your `.zshrc` OR `.bash_profile` files.

```
export GOPATH=$HOME/go
export GOPATH="${HOME}/go
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
```

Then install `go`, with Homebrew.

`brew install go`

You might also want to install some other handy `go` dev tools.

```
go get golang.org/x/tools/cmd/godoc
go get github.com/golang/lint/golint
```

# MongoDB Install

```
brew tap mongodb/brew
brew install mongodb-community@4.2
```

# MongoDB Start

`brew services start mongodb/brew/mongodb-community`

OR

`mongod --config /usr/local/etc/mongod.con`

# Finally, Run the project

`go run main.go`

then navigate to `http://localhost:8080` in your favorite browser. (IE excluded)
