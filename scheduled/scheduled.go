package scheduled

import (
	"go-database/routes"
	"math/rand"
	"strconv"
	"time"
)

func StartScheduled() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rand.Seed(time.Now().UnixNano())
			numeroAleatorio := rand.Intn(100)

			routes.TablesInMemory[0] = strconv.Itoa(numeroAleatorio)
		}
	}
}
