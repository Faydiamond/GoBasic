package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ct := context.Background()
	ct = context.WithValue(ct, "my-key", "Que")
	viewContext(ct)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) //se cancela en dos segundos

	defer cancel()

	myProcess(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Se excedio el tiempo")
		fmt.Println(ctx.Err())
	}
}

func viewContext(ct context.Context) {
	fmt.Printf("my value is %s  ", ct.Value("my-key"))
}

func myProcess(ctx context.Context) {
	for {
		select { //conanales
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		default:
			fmt.Println("Procesando")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

//estudiar canales go y select
