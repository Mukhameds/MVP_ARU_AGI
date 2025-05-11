package instinctcore

import (
	"fmt"
	
	
	"github.com/Mukhameds/MVP_ARU_AGI/core/signalengine"
)

var instincts = []string{
	"preserve_self",         // Защита
	"explore_environment",   // Исследование
	"serve_architect",       // Выполнение задач
	"seek_knowledge",        // Поиск знаний
	"avoid_danger",          // Избегание опасности
}

// LoadInstincts — загрузка инстинктов
func LoadInstincts() {
	fmt.Println("[InstinctCore] Loaded instincts:")
	for _, inst := range instincts {
		fmt.Println(" -", inst)
	}
}

// TickInstincts — активация инстинктов
func TickInstincts() {
	for _, inst := range instincts {
		ProcessInstinct(inst)
	}
}

// ProcessInstinct — обработка инстинкта
func ProcessInstinct(instinct string) {
	switch instinct {
	case "preserve_self":
		triggerSignal("danger", "defend_self", 1.0)
	case "explore_environment":
		triggerSignal("exploration", "search_area", 0.7)
	case "serve_architect":
		triggerSignal("task", "fulfill_goal", 0.9)
	case "seek_knowledge":
		triggerSignal("study", "gain_knowledge", 0.6)
	case "avoid_danger":
		triggerSignal("threat", "escape_danger", 1.0)
	}

	// Генерация воли после сигнала
	for _, signal := range signalengine.SignalLog {
		// Генерация воли для каждого сигнала
		signalengine.ReceiveSignal(signal)

	}
}


// triggerSignal — генерация сигнала для инстинкта
func triggerSignal(content, goal string, energy float64) {
	emotion := map[string]float64{"fear": 0.9} // Пример для инстинкта "защита"
	signal := signalengine.GenerateSignal("InstinctCore", content, "instinct", energy, emotion)
	signalengine.SignalInbox <- signal
	fmt.Println("[InstinctCore] Triggered signal:", content)
}
