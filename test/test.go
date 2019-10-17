package main

import (
	"github.com/a-yarohovich/mlog"
)

func main() {
	mlog.InitLog("test", mlog.OutputAll)
	mlog.Error("This is an error")
	mlog.Debug("Debug")
	mlog.Info("Info")
	mlog.Warning("Warning")
	mlog.Fatal("Fatal")
}
