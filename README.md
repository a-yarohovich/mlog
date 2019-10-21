# mlog
Micro logger for go 

# Example

```go
package main

import (
	"io"
	"os"

	"github.com/a-yarohovich/mlog"
)

func main() {
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	multi := io.MultiWriter(file, os.Stdout)
	mlog.SetFlags(mlog.Ldate | mlog.Lmicroseconds | mlog.Lshortfile | mlog.Lcolor)
	mlog.SetOutput(multi)

	mlog.Error("This is an error")
	mlog.Debug("Debug")
	mlog.Info("Info")
	mlog.Warning("Warning")
	//mlog.Fatal("Fatal")
}
```
