package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//Contexto em Branco
	ctx := context.Background()

	//com o WithCancel eu posso cancelar esse contexto sempre que chamar a função cancel()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel Booking Cancelled. Timeout Reached")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel Booked")
	}
}
