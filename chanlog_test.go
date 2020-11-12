package main

import (
	// "fmt"
	// "github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	Print("HelloWorld")
}

func TestPrintf(t *testing.T) {
	Printf("%s", "HelloWorld")
}
