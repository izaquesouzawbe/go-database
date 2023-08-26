package scheduleds

import (
	"time"
)

func ScheduledTask() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			/*rand.Seed(time.Now().UnixNano())
			numeroAleatorio := rand.Intn(100)*/

			//tablesInMemory[0] = strconv.Itoa(numeroAleatorio)
		}
	}
}
