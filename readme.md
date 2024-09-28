# Logger
Logging with this package can use the following functions
```go
// init : true / false
log.SetIsPrintCaller(false)

// logging
log.Info("this is INFO")

log.Debug("this is DEBUG")

log.Error("this is ERROR")

log.Panic("this is PANIC")

log.Fatal("this is FATAL")
```

## example result
the example if `SetIsPrintCaller(false)`

![no-use-print-info.png](img/no-use-print-info.png)

the example if `SetIsPrintCaller(true)`

![use-print-info.png](img/use-print-info.png)