# go-game-management-system
A go project example to try to manage a repository of games (just info) 

## Steps to build the project in windows:
1. go to directory cmd/main and run `go build`
2. if there is any error related to the package **github.com/mattn/go-sqlite3** and some variable called **CGO_ENABLED** you need to install gcc in windows to do that follow the next [link](https://jmeubank.github.io/tdm-gcc/)
3. Then, run the command `go env -w CGO_ENABLED=1` to enable CGO.
4. After that, run the command `go build -a -ldflags "-linkmode external -extldflags '-static' -s -w"` again and it should work. In case it doesn't work, try to reboot the cmd or terminal you are using.
5. Finally, you can run the command `go run main.go` to run the project.