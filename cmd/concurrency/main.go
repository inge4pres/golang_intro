package main

import (
	"time"

	"git.bravofly.com/fgualazzi/golang_intro.git/pkg/cake"
)

func main() {
	clients := []string{"lastminute.com", "rumbo.es", "volagratis.it", "weg.de"}
	orders := make(chan (string), 0)
	deliveries := make(chan (int), 0)
	var delivery int
	go cake.PrepareCake(orders, deliveries, delivery)
	go cake.DeliverCake(deliveries)
	for _, client := range clients {
		orders <- client
	}
	time.Sleep(100 * time.Millisecond)
	close(orders)
	close(deliveries)
}
