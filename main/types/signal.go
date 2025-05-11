package types

import (
	"time"
)

// Signal — базовая единица сигнального мышления ARU
type Signal struct {
	ID            string
	Type          string            // "external", "emotion", "memory", "phantom", "instinct", "will"
	Content       string            // Текст, команда, гипотеза и т.д.
	Energy        float64           // Энергия сигнала
	Mass          float64           // Масса (влияет на приоритет мышления)
	EmotionalTag  map[string]float64 // {"fear": 0.9, "curiosity": 0.5}
	Origin        string            // Источник: User, Reflex, Memory и др.
	Dimensions    map[string]float64 // Пространственные и ассоциативные координаты
	Timestamp     time.Time         // Время генерации сигнала
}

// Emotion — внутренняя эмоциональная реакция на сигнал
type Emotion struct {
	Type  string
	Power float64
}
