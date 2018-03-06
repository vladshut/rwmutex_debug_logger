# Debug wrapper for sync.RWMutex.
## Why
To debug for deadlocks.

## Installation
```sh
go get github.com/vladshut/rwmutext_debug_wrapper
```

## Usage
```go
import "github.com/vladshut/rwmutext_debug_wrapper"
var rwm rwmutex_debug_wrapper.RWMutex
rwm.UID = "1"
// Use normally, it works exactly like sync.Mutex does.
rwm.Lock()

defer rwm.Unlock()
```