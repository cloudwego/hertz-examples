# hlog
You can learn about how to use hertz hlog

1. By default, the logger implemented by hertz is used by default.
The default logger output can also be redirected using the `SetOutput` interface, and subsequent middleware and other parts of the framework can use the global methods in hlog to output logs.
2. Hertz provides `SetLogger` interface to allow injection of your own logger.

## Interface Definition

In Hertz, the interfaces Logger, CtxLogger, FormatLogger are defined in pkg/common/hlog, and these interfaces are used to output logs in different ways, and a Control interface is defined to control the logger. If you’d like to inject your own logger implementation, you must implement all the above interfaces (i.e. FullLogger). Hertz already provides a default implementation of FullLogger.

```go
// FullLogger is the combination of Logger, FormatLogger, CtxLogger and Control.
type FullLogger interface {
   Logger
   FormatLogger
   CtxLogger
   Control
}
```