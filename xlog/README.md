
## Features
- 按日期自动拆分日志记录
- 日志级别：DEBUG,INFO,WARN,ERROR




## Usage
```
package main

import "github.com/xiuos/mozi/xlog"

func main() {
	logger := xlog.NewLogger(".", "test_", xlog.Ldate|xlog.Ltime|xlog.Lmicroseconds|xlog.Llongfile, xlog.LevelInfo)
	logger.Debug("test debug.")
	logger.Info("test info.")
	logger.Warn("test warn.")
	logger.Error("test error.")
}

```

## Documentation

http://godoc.org/github.com/xiuos/xlog
