package main

import (
    "singletonpattern/logger"
    "sync"
)

func User1Logs(wg *sync.WaitGroup) {
    defer wg.Done()

    logger1 := logger.GetLogger()
    logger1.Log("This message is from user 1")
}

func User2Logs(wg *sync.WaitGroup) {
    defer wg.Done()

    logger2 := logger.GetLogger()
    logger2.Log("This message is from user 2")
}
