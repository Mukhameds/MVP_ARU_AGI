package emotionengine

import (
	"fmt"

	"github.com/Mukhameds/MVP_ARU_AGI/types"
)

// Вес эмоций — влияет на приоритет сигнала
var EmotionWeights = map[string]float64{
	"fear":          1.4,
	"inspiration":   1.3,
	"anger":         1.2,
	"humility":      0.9,
	"curiosity":     1.1,
	"responsibility": 1.25,
}

// ComputePriority — вычисление приоритета на основе эмоции
func ComputePriority(signal types.Signal) float64 {
	base := signal.Mass * signal.Energy
	modifier := 1.0

	for emo, power := range signal.EmotionalTag {
		weight, exists := EmotionWeights[emo]
		if exists {
			modifier += weight * power
		}
	}

	// Усиление эмоции в вычислениях
	if modifier > 2.0 {
		modifier = 2.0
	}

	priority := base * modifier
	return priority
}



// GenerateEmotion — формирует доминирующую эмоцию
func GenerateEmotion(signal types.Signal) types.Emotion {
	instinctInfluence := extractInstinctComponent(signal)
	memoryInfluence := extractMemoryComponent(signal)
	conflictInfluence := detectConflict(signal)

	total := instinctInfluence + memoryInfluence + conflictInfluence
	dominant := selectDominantEmotion(signal)

	em := types.Emotion{
		Type:  dominant,
		Power: normalize(total),
	}
	fmt.Println("[EmotionEngine] Generated emotion:", em.Type, em.Power)
	return em
}

// ─── Вспомогательные ───

func extractInstinctComponent(signal types.Signal) float64 {
	if signal.Origin == "InstinctCore" {
		return 0.5
	}
	return 0.0
}

func extractMemoryComponent(signal types.Signal) float64 {
	for key := range signal.EmotionalTag {
		if key == "nostalgia" || key == "regret" {
			return 0.3
		}
	}
	return 0.0
}

func detectConflict(signal types.Signal) float64 {
	if signal.Type == "error" || signal.Type == "paradox" {
		return 0.4
	}
	return 0.0
}

func selectDominantEmotion(signal types.Signal) string {
	maxPower := 0.0
	dominant := "neutral"
	for emo, power := range signal.EmotionalTag {
		if power > maxPower {
			maxPower = power
			dominant = emo
		}
	}
	return dominant
}

func normalize(x float64) float64 {
	if x > 1.0 {
		return 1.0
	}
	return x
}


