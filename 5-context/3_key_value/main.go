package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senhaToken")
	bookHotel(ctx, "Valoir")

}

// Por convenção o parâmetro de contexto sempre será o primeiro, antes dos demais...
func bookHotel(ctx context.Context, nomeHotel string) {
	token := ctx.Value("token")
	fmt.Println(token)
}
