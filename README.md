# glog

This is a log package for golang.
This package is based on offical `log` pkg.

#### Usage
``` golang
import "glog"
```

#### Log Level and Log Format
This package supports three log levels:
* INFO
```bash
2016/03/25 14:42:00 [INFO] /Users/chideat/Projects/go/src/test/main.go:8 this is info
```
* WARN
```bash
2016/03/25 14:42:00 [WARN] /Users/chideat/Projects/go/src/test/main.go:8 this is warn
```
* ERROR
```bash
2016/03/25 14:42:00 [ERROR] /Users/chideat/Projects/go/src/test/main.go:8 this is error
```

Modifiy the log level at runtime
``` golang
SetLevel(INFO|WARN|ERROR)
```

#### Debug Model
In this model, all log info is output to stdout.

Examples:
``` golang
SetDebug(true)
```
