package ghostfield

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Mukhameds/MVP_ARU_AGI/types"
	"github.com/Mukhameds/MVP_ARU_AGI/core/signalengine"
)

type Phantom struct {
	ID           string
	Topic        string
	GeneratedAt  time.Time
	TTL          int
	EmotionBias  map[string]float64
	Goal         string
	Active       bool
}

var PhantomPool []Phantom

func TickGhostField() {
	if rand.Float64() < 0.1 {
		gen := GeneratePhantom("unresolved_thought", "clarify_direction", map[string]float64{"doubt": 0.6})
		PhantomPool = append(PhantomPool, gen)
		fmt.Println("[GhostField] New phantom:", gen.ID)
	}
	processActivePhantoms()
}

func GeneratePhantom(topic string, goal string, emotion map[string]float64) Phantom {
	return Phantom{
		ID:          fmt.Sprintf("phantom_%d", rand.Intn(999999)),
		Topic:       topic,
		GeneratedAt: time.Now(),
		TTL:         15,
		EmotionBias: emotion,
		Goal:        goal,
		Active:      true,
	}
}

func processActivePhantoms() {
	var survivors []Phantom
	for _, p := range PhantomPool {
		if int(time.Since(p.GeneratedAt).Seconds()) > p.TTL {
			fmt.Println("[GhostField] Phantom expired:", p.ID)
			continue
		}

		if rand.Float64() < 0.3 {
			sig := signalengine.GenerateSignal(
				"GhostField",
				p.Goal,
				"phantom",
				0.5,
				p.EmotionBias,
			)
			signalengine.SignalInbox <- sig
			fmt.Println("[GhostField] Phantom fired:", p.ID)
		}
		survivors = append(survivors, p)
	}
	PhantomPool = survivors
}
