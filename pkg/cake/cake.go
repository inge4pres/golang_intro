package cake

import "fmt"

func PrepareCake(orders <-chan (string), deliveries chan (int), delivery int) {
	for order := range orders {
		fmt.Printf("Order %s prepared\n", order)
		delivery++
		deliveries <- delivery
		fmt.Printf("Delivery %d scheduled\n", delivery)
	}
}

func DeliverCake(addrs <-chan (int)) {
	for addr := range addrs {
		fmt.Printf("Delivery %d done!\n", addr)
	}
}
