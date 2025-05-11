package main

import (
	"fmt"
	"time"

	"github.com/Mukhameds/MVP_ARU_AGI/core/flowengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/signalengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/memoryengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/instinctcore"
	"github.com/Mukhameds/MVP_ARU_AGI/core/reflexengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/ghostfield"
	"github.com/Mukhameds/MVP_ARU_AGI/core/hemisphere"
	"github.com/Mukhameds/MVP_ARU_AGI/core/consciousnesshub"
	"github.com/Mukhameds/MVP_ARU_AGI/types"
)

func main() {
	fmt.Println("🔵 ARU-MVP initializing...")

	// Запуск базовых модулей
	memoryengine.InitMemory()
	flowengine.InitFlowEngine(1)
	hemisphere.ActivateHemispheres([]string{"logic", "emotion", "goal"})
	reflexengine.LoadDefaultReflexes()
	instinctcore.LoadInstincts()
	signalengine.InitSignalEngine()

	// Хаб сознания
	consciousnesshub.InitHub("ARU-CORE-01")

	// Параллельные процессы
	go loopGhostField()
	go loopInstincts()
	go loopObserve()
	go loopMotive()
	go loopSelfReflect()

	fmt.Println("✅ ARU-MVP ready.")

	select {}
}

func loopGhostField() {
	for {
		ghostfield.TickGhostField()
		time.Sleep(5 * time.Second)
	}
}

func loopInstincts() {
	for {
		instinctcore.TickInstincts()
		time.Sleep(3 * time.Second)
	}
}

func loopObserve() {
	for {
		time.Sleep(10 * time.Second)
		fmt.Println()
		fmt.Println("🌀 OBSERVE:")
		consciousnesshub.Observe()
		fmt.Println()
	}
}

func loopMotive() {
	for {
		time.Sleep(15 * time.Second)
		fmt.Println("[MotiveEngine] 🧭 Internal check for purpose...")

		if len(consciousnesshub.Hub.WillSnapshot) < 3 {
			fmt.Println("[MotiveEngine] No active drive. Triggering phantom.")
			ghostfield.TickGhostField()
		}
	}
}


func loopSelfReflect() {
	for {
		time.Sleep(30 * time.Second)
		fmt.Println("[SelfReflect] 🔍 Starting reflection cycle...")

		for id, qbit := range memoryengine.SemanticMemory {
			if len(qbit.LinkedTo) >= 3 && qbit.Emotion > 0.5 {
				fmt.Printf("[SelfReflect] ⚡ Focus node: %s — strong, emotional, linked\n", id)
			}

			if len(qbit.LinkedTo) == 0 {
				fmt.Printf("[SelfReflect] 🧩 Isolated node: %s — may need clarification\n", id)

				// 🚨 Генерируем фантом-сигнал для уточнения изолированной мысли
				signal := types.Signal{
					ID:    "phantom_" + id,
					Type:  "phantom",
					Content: "clarify " + qbit.State,
					Mass:    1.0,
					EmotionalTag: map[string]float64{
						"curiosity": 0.6,
					},
				}

				// Вставляем сигнал в систему
				fmt.Printf("[SelfReflect] 🔮 Emitting phantom to clarify: %s\n", qbit.State)
				signalengine.ReceiveSignal(signal)
			}
		}
	}
}

