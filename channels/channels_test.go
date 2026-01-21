package channels

import (
	"log"
	"testing"
)

func TestRun(t *testing.T) {
	s := Run()
	log.Println(s)
}
