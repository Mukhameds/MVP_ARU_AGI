package thoughtengine

import (
	"fmt"

	"github.com/Mukhameds/MVP_ARU_AGI/types"
	"github.com/Mukhameds/MVP_ARU_AGI/core/flowengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/emotionengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/willengine"
)

type Thought struct {
	ID       string
	Signal   types.Signal
	Priority float64
	Form     string
}

var ThoughtPool []Thought

// ReceiveSignal — приём сигнала и формирование мысли
func ReceiveSignal(signal types.Signal) {
	priority := emotionengine.ComputePriority(signal)
	th := Thought{
		ID:       "thought_" + signal.ID,
		Signal:   signal,
		Priority: priority,
		Form:     DetectForm(signal),
	}
	ThoughtPool = append(ThoughtPool, th)

	fmt.Printf("[ThoughtEngine] New thought: %s (form=%s, priority=%.2f)\n", th.ID, th.Form, th.Priority)

	if priority > 0.5 {
		flowengine.Schedule(signal)
		willengine.GenerateWill(signal)
	}
}

// DetectForm — упрощённая классификация типа мысли
func DetectForm(signal types.Signal) string {
	switch signal.Type {
	case "goal":
		return "deductive"
	case "error", "paradox":
		return "critical"
	case "emotion":
		return "introspective"
	case "phantom":
		return "abductive"
	default:
		return "generic"
	}
}
