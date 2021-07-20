package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	ListAnimalTypeForQurban = []string{
		"Goat", "Sheep", "Cow", "Camel",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Receiver(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for {
		time.Sleep(time.Second)
		log.Printf("Receiving %d %s for qurban.\n",
			rand.Intn(10)+1, ListAnimalTypeForQurban[rand.Intn(
				len(ListAnimalTypeForQurban),
			)],
		)
	}
}

func main() {
	log.Println("Happy Eid Al Adha!")

	waitGroup := sync.WaitGroup{}
	maxReceiver := 20

	for i := 0; i < maxReceiver; i++ {
		waitGroup.Add(1)
		go Receiver(&waitGroup)
	}

	waitGroup.Wait()
}
