package thoughtengine

import (
	"fmt"
	"time"
	"github.com/Mukhameds/MVP_ARU_AGI/types"
	"github.com/Mukhameds/MVP_ARU_AGI/core/flowengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/emotionengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/willengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/memoryengine"
)

type Thought struct {
	ID       string
	Signal   types.Signal
	Priority float64
	Form     string
	Thread   flowengine.ThoughtThread
}

var ThoughtPool []Thought
var lastEmotion types.Emotion

// ReceiveSignal — приём сигнала и формирование мысли
func ReceiveSignal(signal types.Signal) Thought {
	priority := emotionengine.ComputePriority(signal)
	emotion := emotionengine.GenerateEmotion(signal)

	th := Thought{
		ID:       "thought_" + signal.ID,
		Signal:   signal,
		Priority: priority,
		Form:     DetectForm(signal),
	}

	// создаём поток и сохраняем в структуру
	thread := flowengine.CreateThread(signal)
	th.Thread = thread

	// сохраняем в Pool
	ThoughtPool = append(ThoughtPool, th)
	lastEmotion = emotion

	fmt.Printf("[ThoughtEngine] New thought: %s (form=%s, priority=%.2f)\n", th.ID, th.Form, th.Priority)

// Если мысль с таким ID уже в памяти — усиливаем связь и не дублируем
if existing, ok := memoryengine.SemanticMemory[th.ID]; ok {
	memoryengine.StrengthenLink(th.ID, th.ID, 0.2) // саморефлексия
	existing.Emotion = (existing.Emotion + emotion.Power) / 2
	existing.Timestamp = time.Now()

	fmt.Printf("[ThoughtEngine] 🔁 Thought already exists: %s — strengthened.\n", th.ID)
	return th // ⛔ не продолжаем — воля уже могла быть вызвана
}


	// 🧠 Сохраняем мысль в семантическую память
	memoryengine.SaveSemantic(
		th.ID,
		signal.Content,
		[]string{signal.Type, th.Form},
		[]string{"thought"},
		emotion.Power,
	)

	// ✅ Теперь, когда мысль уже записана — формируем волю
	if priority > 0.5 {
		flowengine.Schedule(signal)
		willengine.GenerateWill(signal)
	}

	return th
}


// LastEmotion — возвращает последнюю эмоцию
func LastEmotion() types.Emotion {
	return lastEmotion
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
