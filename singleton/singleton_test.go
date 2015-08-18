package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	instance := singleton{}
	instance.doSomething()
}