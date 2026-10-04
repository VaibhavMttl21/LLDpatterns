// not thread safe
// Goroutine 1                 Goroutine 2
//      │                           │
//      │ GetLogger()               │ GetLogger()
//      ↓                           ↓
// instance == nil?             instance == nil?
//      │                           │
//     YES                         YES
//      │                           │
//      ↓                           ↓
// newLogger()                  newLogger()
//      │                           │
//      ↓                           ↓
//  Logger #1                    Logger #2

// package logger

// import "fmt"

// type logger struct{}  // private

// var loggerInstance *logger // private loggerInstance stores a pointer to a Logger object.

// func newLogger() *logger { // private
//     return &logger{}
// }

// func GetLogger() *logger { // public
//     if loggerInstance == nil {
//         loggerInstance = newLogger()
//     }

//     return loggerInstance
// }

// func (l *logger) Log(msg string) {
//     fmt.Println(msg)
// }

// package logger

// import (
// 	"fmt"
// 	"time"
// )

// type Logger struct{}

// var loggerInstance *Logger

// func newLogger() *Logger {
// 	fmt.Println("Creating new Logger instance...")

// 	time.Sleep(100 * time.Millisecond)

// 	return &Logger{}
// }

// func GetLogger() *Logger {
// 	if loggerInstance == nil {
// 		loggerInstance = newLogger()
// 	}

// 	return loggerInstance
// }

// func (l *Logger) Log(msg string) {
// 	fmt.Println(msg)
// }


// ======================== Mutext added so to save it from multthreading that distrbs singleton ========

// Goroutine 1                 Goroutine 2
//      │                           │
//      ↓                           ↓
// GetLogger()                 GetLogger()
//      │                           │
//      ↓                           ↓
//  mutex.Lock()               mutex.Lock()
//      │                           │
//      ↓                           │
//    LOCKED                       WAIT
//      │                           │
//      ↓                           │
// instance == nil                 │
//      │                           │
//     YES                          │
//      │                           │
//      ↓                           │
// create Logger                   │
//      │                           │
//      ↓                           │
// instance = Logger #1            │
//      │                           │
//      ↓                           │
// mutex.Unlock()                 │
//                                  ↓
//                             LOCK ACQUIRED
//                                  │
//                                  ↓
//                          instance == nil?
//                                  │
//                                 NO
//                                  │
//                                  ↓
//                          return Logger #1


package logger

import (
	"fmt"
	"sync"
)

type Logger struct{}

var loggerInstance *Logger
var mutex sync.Mutex

func newLogger() *Logger {
	fmt.Println("Creating new Logger instance...")
	return &Logger{}
}
// heir the thing is that we don't need lock after the first initalisation and it is a expensive method also
// so we have to make it once so in go we have sync.once shich take care of all the cases 
func GetLogger() *Logger {

	mutex.Lock()
	defer mutex.Unlock()

	if loggerInstance == nil {
		loggerInstance = newLogger()
	}

	return loggerInstance
}

func (l *Logger) Log(msg string) {
	fmt.Println(msg)
}

// package logger

// import (
// 	"fmt"
// 	"sync"
// )

// type Logger struct{}

// var (
// 	loggerInstance *Logger
// 	once           sync.Once
// )

// func newLogger() *Logger {
// 	fmt.Println("Creating new Logger instance...")
// 	return &Logger{}
// }

// func GetLogger() *Logger {

// 	once.Do(func() {
// 		loggerInstance = newLogger()
// 	})

// 	return loggerInstance
// }

// func (l *Logger) Log(msg string) {
// 	fmt.Println(msg)
// }