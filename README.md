# twitter-text-parse-go

[twitter-text](https://github.com/twitter/twitter-text) Go binding for `parse` func of Rust impl

`main` branch is for `go get`. Refer to `dev` branch for development.

## Get Started

For target `x86_64-unknown-linux-gnu`, use the package in dir `pkg/gnu`:

```go
package main

import (
    "log"

    twtextparse "github.com/myl7/twitter-text-parse-go/pkg/gnu"
)

func main() {
    res, err := twtextparse.Parse("test")
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(res)
}
```

For target `x86_64-unknown-linux-musl`, use the package in dir `pkg/musl`, and all API are the same

For other targets, you may need to build by yourself.
`build/build.sh` contains the actual build commands.

## License

MIT
