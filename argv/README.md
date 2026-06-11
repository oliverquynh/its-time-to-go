# Argument Vector

In a Go CLI program, argument vector is represent by `os.Args`.

```go
// E.g. Running "go run main.go one two three"
// The argv is ["/path/to/build/main", "one", "two", "three"]

for _, arg := range os.Args {
    fmt.Println(arg)
}
```

To parse options, we can use `flag`.

```go
// E.g. Running "go run main.go --host=127.0.0.1 --port=3000"
// Host: 127.0.0.1 - Port: 3000

host := flag.String("host", "0.0.0.0", "Host to listen on")
port := flag.Int("port", 8080, "Port to listen on")

flag.Parse()

fmt.Println("Host: " + *host)
fmt.Println("Port: " + strconv.Itoa(*port))
```

`flag` stops parsing options when it faces the first argument.

```bash
# This will work
go run main.go --host=127.0.0.1 --port=3000

# This will NOT (E.g.imagine start is a command name)
go run main.go start --host=127.0.0.1 --port=3000
```
