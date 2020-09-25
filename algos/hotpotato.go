package algos

import (
	"log"

	"github.com/arunsivasankaran/go-algorithms-and-datastructures/datastructures"
)

// Hotpotato takes a list of names and a number and tells us who survives a game of hotpotato
func Hotpotato(names []string, num int) string {

	var newQueue datastructures.Queue

	for _, value := range names {
		newQueue.Enqueue(value)
	}

	for newQueue.Size() > 1 {

		for i := 0; i < num; i++ {
			name, _ := newQueue.Dequeue()
			newQueue.Enqueue(name)
		}

		_, err := newQueue.Dequeue() // dead

		if err != nil {
			log.Fatalf("Hotpotato: %v\n", err)
		}
	}

	survivor, _ := newQueue.Dequeue()

	return survivor
}
