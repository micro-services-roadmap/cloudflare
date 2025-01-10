package worker

import (
	"fmt"
	"testing"
)

func TestNextWorkerID(t *testing.T) {
	for i := 0; i < 0; i++ {
		id, err := NextWorkerID()
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Println(id)
		}
	}
}
