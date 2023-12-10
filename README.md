# MailHedgehog logger

Log helper for MailHedgehog package. Can be reused in any other app/package.

## Usage

Full usage expects creation custom labeled logger for each package and then execute:

```go
package main

import (
    "github.com/mailhedgehog/logger"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
    if configuredLogger == nil {
        configuredLogger = logger.CreateLogger("server.api")
    }
    return configuredLogger
}

func main() {
    logManager().Emergency("I see error!")
}
```

Direct quick usage:

```go
package main

import (
    "github.com/mailhedgehog/logger"
)

func main() {
    logger.CreateLogger("server.api").Emergency("I see error!")
}
```

```go
package main

import (
    "github.com/mailhedgehog/logger"
)

func main() {
    logger.Emergency("I see error!")
}
```

## Development

```shell
go mod tidy
go mod verify
go mod vendor
go test --cover
```


## Credits

- [![Think Studio](https://yaroslawww.github.io/images/sponsors/packages/logo-think-studio.png)](https://think.studio/)
