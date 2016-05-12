package fastprng

import (
	"log"
	"testing"
	"time"
)

func TestGenerateMad0(t *testing.T) {
	rnd := NewMaD0(time.Now().UnixNano())

	log.Println("MAD0 RAND", rnd.Next())
	// t.Fail()
}
