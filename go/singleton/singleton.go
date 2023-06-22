package singleton

import (
	"fmt"
	"sync"
	"time"
)

// singleton is nil at the beginning of application
var singleton *Logger

// once can be used to prevent multiple initialization of singleton
var once sync.Once

// Logger is the example, we want to use it as singleton for whatever reason
type Logger struct {
}

// NewLogger initializes singleton variable only once! and returns it
func NewLogger() *Logger {
	// once.Do accept a function as a parameter
	// this function is called only once, the first time NewLogger is called
	once.Do(func() {
		singleton = &Logger{}
	})
	return singleton
}

func (l *Logger) Info(msg string) {
	fmt.Printf("%s - INFO: %s\n", time.Now().Format(time.RFC822), msg)
}
