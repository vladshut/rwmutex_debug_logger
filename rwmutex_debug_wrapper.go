package rwmutex_debug_wrapper

import (
	"fmt"
	"sync"
	"runtime/debug"
	"strconv"
	"io"
	"os"
)

var Opts = struct {
	Log io.Writer
}{
	Log: os.Stdout,
}

type RWMutex struct {
	sync.RWMutex
	isLocked bool
	rLocks int
	UID string
}

func (rwm *RWMutex) Lock() {
	if rwm.isLocked {
		rwm.logDebugMessage("RWM already Locked.")
	}

	if rwm.rLocks > 0 {
		rwm.logDebugMessage("RWM already RLocked.")
	}

	rwm.RWMutex.Lock()
	rwm.isLocked = true
	rwm.logDebugMessage("RWM Locked.")
}

func (rwm *RWMutex) Unlock() {
	rwm.RWMutex.Unlock()
	rwm.isLocked = false
	rwm.logDebugMessage("RWM Unlocked.")
}

func (rwm *RWMutex) RLock() {
	rwm.RWMutex.RLock()
	rwm.rLocks += 1
	rwm.logDebugMessage("RWM RLocked.")
}

func (rwm *RWMutex) RUnlock() {
	rwm.RWMutex.RUnlock()
	rwm.rLocks -= 1
	rwm.logDebugMessage("RWM RUnlocked.")
}

func (rwm *RWMutex) createLogMessage(message string) string {

	if rwm.UID != "" {
		message = message + " UID: " + rwm.UID + ","
	}

	message += " Status:"

	if rwm.isLocked {
		message += " Locked"
	} else {
		message += " Unlocked"
	}

	message += ", Readers count: " + strconv.Itoa(rwm.rLocks)

	message += ", StackTrace: \n" + string(debug.Stack())

	return message
}

func (rwm *RWMutex) logDebugMessage(message string) {
	message = rwm.createLogMessage(message)
	fmt.Fprintln(Opts.Log, message)
	fmt.Fprintln(Opts.Log)
}
