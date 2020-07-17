# issue-tracker
Simple Github issue tracker


## Flags
```
-h or --help: Help command to see available commands.
-a:   sort issues by ascending order. (default true)
-d:   Disable to display username. (default false)
-n:   Number of issues to fetch (default 10 / max: 30 for now)
-r:   Github repo name as account-name/repo-name (ex. 'golang/go')
-s:   To see open issues, type -s; otherwise, do not type -s
```

## Usage

Run from main.go;
```
go run main.go -r=golang/go -n=5
```
or, from main binary

```
go build main.go
./main -r=golang/go -n=5
```

## Acknowledgments
This repo is solution for ex 4.11 in [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440) book. The solution does not assert to be exact solution for the problem.
