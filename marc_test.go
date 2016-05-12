package fastprng

import (
	"log"
	"testing"
	"time"
)

func TestGenerateMarc(t *testing.T) {
	rnd := NewMARC(time.Now().UnixNano())

	log.Println("MARC RAND", rnd.Next())
}

func TestGenerateMarcBB(t *testing.T) {
	rnd := NewMARCBB(time.Now().UnixNano())

	log.Println("MARCBB RAND", rnd.Next())
}
