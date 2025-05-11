package ghostfield

import (
	"fmt"
	"math/rand"
	"time"

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
	Life float64 // от 1.0 до 0.0

}

var PhantomPool []Phantom

func TickGhostField() {
	// Частота создания фантомов
	if rand.Float64() < 0.1 {
		gen := GeneratePhantom("unresolved_thought", "clarify_direction", map[string]float64{"doubt": 0.6})
		PhantomPool = append(PhantomPool, gen)
		fmt.Println("[GhostField] New phantom:", gen.ID)
	}

	// Объединение схожих фантомов
	MergePhantoms()

	// Обработка оставшихся фантомов
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
		Life: 1.0,

	}
}

func processActivePhantoms() {
	var survivors []Phantom
	for _, p := range PhantomPool {
		// Угасание жизни при каждом тике
		p.Life -= 0.03
		if p.Life <= 0 {
			fmt.Println("[GhostField] Phantom decayed:", p.ID)
			continue
		}

		// Если TTL вышел — это дополнительный выход
		if int(time.Since(p.GeneratedAt).Seconds()) > p.TTL {
			fmt.Println("[GhostField] Phantom expired:", p.ID)
			continue
		}

		// Вероятность активации
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

		// Сохраняем живого фантома
		survivors = append(survivors, p)
	}
	PhantomPool = survivors
}

// Объединение фантомов с похожими темами и эмоциями
func MergePhantoms() {
	var merged []Phantom

	for i := 0; i < len(PhantomPool); i++ {
		p1 := PhantomPool[i]
		mergedFlag := false

		for j := i + 1; j < len(PhantomPool); j++ {
			p2 := PhantomPool[j]

			// Объединяем, если тема и эмоции схожи
			if p1.Topic == p2.Topic && SimilarEmotion(p1.EmotionBias, p2.EmotionBias) {
				// Создаём новый фантом, слияние тем и эмоций
				mergedPhantom := Phantom{
					ID:          fmt.Sprintf("merged_%s_%s", p1.ID, p2.ID),
					Topic:       p1.Topic,
					GeneratedAt: time.Now(),
					TTL:         max(p1.TTL, p2.TTL),
					EmotionBias: mergeEmotions(p1.EmotionBias, p2.EmotionBias),
					Goal:        p1.Goal, // можно комбинировать цели позже
					Active:      true,
					Life:        1.0,
				}
				merged = append(merged, mergedPhantom)
				PhantomPool = append(PhantomPool[:i], PhantomPool[i+1:]...) // Удаляем старые фантомы
				PhantomPool = append(PhantomPool[:j-1], PhantomPool[j:]...) // Удаляем старые фантомы
				mergedFlag = true
				break
			}
		}

		if !mergedFlag {
			merged = append(merged, p1)
		}
	}

	PhantomPool = merged
}

// Функция для проверки схожести эмоций
func SimilarEmotion(emotion1, emotion2 map[string]float64) bool {
	for key, value := range emotion1 {
		if emotion2[key] != value {
			return false
		}
	}
	return true
}

// Функция для объединения эмоций
func mergeEmotions(emotion1, emotion2 map[string]float64) map[string]float64 {
	merged := make(map[string]float64)
	for key, value := range emotion1 {
		merged[key] = value
	}
	for key, value := range emotion2 {
		if _, exists := merged[key]; exists {
			merged[key] += value
		} else {
			merged[key] = value
		}
	}
	return merged
}

// Функция для нахождения максимального TTL
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
