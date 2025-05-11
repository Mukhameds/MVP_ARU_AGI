package signalengine

import (
	"fmt"
	"time"

	"github.com/Mukhameds/MVP_ARU_AGI/core/thoughtengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/consciousnesshub"
	"github.com/Mukhameds/MVP_ARU_AGI/types"
)

var SignalLog []types.Signal
var SignalInbox = make(chan types.Signal, 100)

// InitSignalEngine — запуск сигнального движка и процессора
func InitSignalEngine() {
	fmt.Println("[SignalEngine] Signal engine initialized")
	go processInbox()
}

// GenerateID — генератор ID
func GenerateID() string {
	return fmt.Sprintf("sig_%d", time.Now().UnixNano())
}

// CalculateMass — масса сигнала = энергия × (1 + сумма эмоций)
func CalculateMass(energy float64, emotions map[string]float64) float64 {
	sum := 0.0
	for _, v := range emotions {
		sum += v
	}
	return energy * (1.0 + sum)
}

// InitializeDimensions — базовые координаты
func InitializeDimensions() map[string]float64 {
	return map[string]float64{
		"time": float64(time.Now().Unix()),
	}
}

// GenerateSignal — создание нового сигнала
func GenerateSignal(origin, content, signalType string, energy float64, emotion map[string]float64) types.Signal {
	s := types.Signal{
		ID:           GenerateID(),
		Type:         signalType,
		Content:      content,
		Energy:       energy,
		Mass:         CalculateMass(energy, emotion),
		EmotionalTag: emotion,
		Origin:       origin,
		Dimensions:   InitializeDimensions(),
		Timestamp:    time.Now(),
	}
	return s
}

// LogSignal — запись сигнала
func LogSignal(s types.Signal) {
	SignalLog = append(SignalLog, s)
	fmt.Printf("[SignalEngine] New signal: [%s] %s (mass=%.2f)\n", s.Type, s.Content, s.Mass)
}

// processInbox — обработка сигналов из канала SignalInbox
func processInbox() {
	for s := range SignalInbox {
		LogSignal(s)
		thought := thoughtengine.ReceiveSignal(s)
		emotion := thoughtengine.LastEmotion()
		consciousnesshub.UpdateSnapshot(s, emotion, thought.Thread)
	}
}
