package chanlog

import (
	// "fmt"
	// "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	SetChanSize(10)
}

func TestPrint(t *testing.T) {
	Print("TestPrint")
	time.Sleep(time.Second)
}

func TestPrintf(t *testing.T) {
	Printf("%s", "TestPrintf")
}
