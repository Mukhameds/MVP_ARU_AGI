package willengine

import (
	"fmt"
	"github.com/Mukhameds/MVP_ARU_AGI/types"
)

// Will — воля на основе мысли
type Will struct {
	ID       string
	SignalID string
	Goal     string
	Power    float64
	Active   bool
}

var WillPool []Will

// GenerateWill — на основе мысли формируем волю
func GenerateWill(signal types.Signal) Will {
	power := signal.Mass * 0.5 // Усиление на основе массы сигнала
	if tag, ok := signal.EmotionalTag["fear"]; ok {
		power += tag * 0.5 // Усиливаем волю, если есть страх
	}
	
	// Обновим цель на основе инстинкта
	goal := signal.Content // Цель по сигналу

	w := Will{
		ID:       "will_" + signal.ID,
		SignalID: signal.ID,
		Power:    power,
		Goal:     goal,
		Active:   true,
	}

	// Добавляем в пул
	WillPool = append(WillPool, w)
	fmt.Printf("[WillEngine] Will generated: %s → %s (power=%.2f)\n", w.ID, w.Goal, w.Power)

	return w
}

// ProcessWill — выполнение воли
func ProcessWill(w Will) {
	if w.Active {
		// Логика исполнения воли, например, защита, исследование
		fmt.Printf("[WillEngine] Acting on will: %s → %s\n", w.ID, w.Goal)
		// Выполняем действие (могут быть сигналы или физические действия)
	}
}
