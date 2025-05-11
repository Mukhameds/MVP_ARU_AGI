package willengine

import (
	"fmt"

	"github.com/Mukhameds/MVP_ARU_AGI/types"
)

type Will struct {
	ID       string
	SignalID string
	Power    float64
	Goal     string
	Active   bool
}

var WillPool []Will

// GenerateWill — на основе сигнала создаёт волю
func GenerateWill(signal types.Signal) Will {
	power := signal.Mass
	if tag, ok := signal.EmotionalTag["responsibility"]; ok {
		power += tag * 0.5
	}
	w := Will{
		ID:       "will_" + signal.ID,
		SignalID: signal.ID,
		Power:    power,
		Goal:     signal.Content,
		Active:   true,
	}
	WillPool = append(WillPool, w)
	fmt.Printf("[WillEngine] Will generated: %s → %s (power=%.2f)\n", w.ID, w.Goal, w.Power)
	return w
}

// TickWillEngine — проверка и запуск активных воль
func TickWillEngine() {
	for _, w := range WillPool {
		if w.Active && w.Power > 0.5 {
			fmt.Printf("[WillEngine] Executing will: %s → %s\n", w.ID, w.Goal)
			w.Active = false
		}
	}
}
