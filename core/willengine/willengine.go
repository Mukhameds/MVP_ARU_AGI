package willengine

import (
	"fmt"

	"github.com/Mukhameds/MVP_ARU_AGI/types"
	"github.com/Mukhameds/MVP_ARU_AGI/core/memoryengine" // 🔧 импорт
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
	
	goal := signal.Content // Цель по сигналу

	w := Will{
		ID:       "will_" + signal.ID,
		SignalID: signal.ID,
		Power:    power,
		Goal:     goal,
		Active:   true,
	}

	WillPool = append(WillPool, w)
	fmt.Printf("[WillEngine] Will generated: %s → %s (power=%.2f)\n", w.ID, w.Goal, w.Power)

	// 🔧 Усиливаем логическую связь в семантической памяти
	memoryengine.LinkQBits("thought_" + signal.ID, w.ID, 1.0)


	return w
}

// ProcessWill — выполнение воли
func ProcessWill(w Will) {
	if !w.Active {
		return
	}

	fmt.Printf("[WillEngine] Acting on will: %s → %s (power=%.2f)\n", w.ID, w.Goal, w.Power)

	switch w.Goal {
	case "defend_self":
		act("Activate defense system")
	case "search_area":
		act("Scan environment for anomalies")
	case "fulfill_goal":
		act("Execute assigned objective from Architect")
	case "gain_knowledge":
		act("Query internal memory and seek new patterns")
	case "escape_danger":
		act("Initiate retreat or avoidance maneuver")
	case "clarify_direction":
		act("Generate phantom: clarify current purpose")
	default:
		act("Reflect internally on goal: " + w.Goal)
	}

	w.Active = false
}

// act — безопасный внутренний исполняемый механизм
func act(description string) {
	fmt.Printf("[Action] ➤ %s\n", description)
}
