// old code bad singleton
// package main

// import "fmt"

// type Logger struct{}

// var instance *Logger

// func GetLogger() *Logger {
//     if instance == nil {
//         instance = &Logger{}
//     }

//     return instance
// }

// func (l *Logger) Log(message string) {
//     fmt.Println(message)
// }

// func main() {
//     logger1 := GetLogger()
//     logger2 := GetLogger()

//     logger1.Log("Hello")
//     logger2.Log("World")

//	    fmt.Println(logger1 == logger2)
//	}

// this implementation is also not thread safe
package main

import (
	// "sync"
"fmt"
"singletonpattern/logger"
)
// it have the multi thread error
func main() {

    logger1 := logger.GetLogger()

    logger1.Log("This message is from user 1")

    logger2 := logger.GetLogger()

    logger2.Log("This message is from user 2")

    fmt.Println(logger1 == logger2)
	
}
// this the multi thread error 
// func main(){
// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	go User1Logs(&wg)
// 	go User2Logs(&wg)

// 	wg.Wait()
// }
        //             logger package
        //         ┌────────────────────┐
        //         │                    │
        //         │  logger            │
        //         │                    │
        //         │  loggerInstance    │
        //         │       │            │
        //         │       ▼            │
        //         │   ┌────────┐       │
        //         │   │ Logger │       │
        //         │   └────────┘       │
        //         │                    │
        //         │  GetLogger()       │
        //         │       │            │
        //         └───────┼────────────┘
        //                 │
        //       ┌─────────┴─────────┐
        //       ↓                   ↓
        //   logger1              logger2
        //       │                   │
        //       └────────┬──────────┘
        //                ↓
        //          SAME OBJECT