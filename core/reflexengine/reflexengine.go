package reflexengine

import (
	"fmt"
)

type Reflex struct {
	ID             string
	TriggerPattern string
	Reaction       string
	Threshold      float64
	Source         string
	Modifiable     bool
}

var ReflexStorage []Reflex

// LoadDefaultReflexes — загрузка базовых рефлексов
func LoadDefaultReflexes() {
	ReflexStorage = []Reflex{
		{
			ID:             "protect_core",
			TriggerPattern: "danger",
			Reaction:       "activate_defense",
			Threshold:      0.6,
			Source:         "Instinct",
			Modifiable:     false,
		},
		{
			ID:             "logic_error",
			TriggerPattern: "contradiction",
			Reaction:       "start_self_diagnosis",
			Threshold:      0.7,
			Source:         "Adapted",
			Modifiable:     true,
		},
	}
	fmt.Println("[ReflexEngine] Default reflexes loaded.")
}

// TickReflexes — проверка (заглушка)
func TickReflexes() {
	fmt.Println("[ReflexEngine] Reflex system tick.")
}
