package consciousnesshub

import (
	"fmt"

	"github.com/Mukhameds/MVP_ARU_AGI/core/willengine"  // Импортируем willengine
	"github.com/Mukhameds/MVP_ARU_AGI/core/flowengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/memoryengine"
	"github.com/Mukhameds/MVP_ARU_AGI/types"
)

type ConsciousnessHub struct {
	ActiveContext   []flowengine.ThoughtThread
	MemoryOverlay   []memoryengine.QBit
	SignalSnapshot  []types.Signal
	WillSnapshot    []willengine.Will // Добавляем волю в контекст
	EmotionalState  types.Emotion
	CoreIdentity    string
}


var Hub ConsciousnessHub

// InitHub — инициализация хаба сознания
func InitHub(coreID string) {
	Hub = ConsciousnessHub{
		CoreIdentity: coreID,
	}
	fmt.Println("[ConsciousnessHub] Consciousness initialized with ID:", coreID)
}

// UpdateSnapshot — обновление состояния сознания с волей
func UpdateSnapshot(signal types.Signal, emotion types.Emotion, context flowengine.ThoughtThread) {
	Hub.SignalSnapshot = append(Hub.SignalSnapshot, signal)
	Hub.EmotionalState = emotion
	Hub.ActiveContext = append(Hub.ActiveContext, context)

	// Добавляем волю в контекст
	will := willengine.GenerateWill(signal)
	Hub.WillSnapshot = append(Hub.WillSnapshot, will)

	fmt.Printf("[ConsciousnessHub] Context updated: %s | Emotion: %s (%.2f) | Will: %s\n", 
		signal.Content, emotion.Type, emotion.Power, will.Goal)
}


// Observe — просмотр текущего состояния сознания
func Observe() {
	fmt.Println("🧠 Consciousness Snapshot:")
	fmt.Println("  🔑 ID:", Hub.CoreIdentity)
	fmt.Println("  🧠 Last Thought:", lastThought())
	fmt.Println("  ❤️ Emotion:", Hub.EmotionalState.Type, Hub.EmotionalState.Power)
	fmt.Println("  🧠 Thoughts in Context:", len(Hub.ActiveContext))
	fmt.Println("  🔁 Signals Seen:", len(Hub.SignalSnapshot))
}

func lastThought() string {
	if len(Hub.SignalSnapshot) == 0 {
		return "<none>"
	}
	return Hub.SignalSnapshot[len(Hub.SignalSnapshot)-1].Content
}
