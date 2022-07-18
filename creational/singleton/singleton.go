package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type singleton struct{}

var singletonInstance *singleton

func getInstance() *singleton {

	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			fmt.Println("Creating singleton instance now.")
			singletonInstance = &singleton{}
		} else {
			fmt.Println("Singleton instance already created.")
		}
	} else {
		fmt.Println("Singleton instance already created.")
	}

	return singletonInstance
}
