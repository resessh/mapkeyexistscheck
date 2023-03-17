# mapkeyexistscheck

mapkeyexistscheck is a lint program that enforces checking for the existence of a key when accessing a map.

## Install

```
go install github.com/resessh/mapkeyexistscheck/cmd/mapkeyexistscheck@latest
```

## Usage

```go
package main

func main()  {
  m := map[int]string{
    1: "1",
  }

  val := m[2]

  // ...
}
```

```
$ go vet -vettool=(which mapkeyexistscheck) ./...

main.go:8:2: map key exists check is not performed
```
