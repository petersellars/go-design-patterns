package singleton

import "fmt"

type singleton struct {

}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func (instance *singleton) doSomething() {
	fmt.Print("Doing Something\n")
}