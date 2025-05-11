package memoryengine

import (
	"fmt"
	"time"
)

// SemanticQBit — расширенная смысловая единица памяти
type SemanticQBit struct {
	ID        string
	State     string
	Context   []string         // Семантический контекст (напр. ["exploration", "danger"])
	Tags      []string         // Тематические или логические теги
	Emotion   float64          // Эмоциональная насыщенность
	LinkedTo  map[string]float64 // ID других QBit'ов и сила связи
	Timestamp time.Time
}

var SemanticMemory = make(map[string]*SemanticQBit)

// SaveSemantic — сохранить или обновить семантический QBit
func SaveSemantic(id, state string, context, tags []string, emotion float64) {
	q := &SemanticQBit{
		ID:        id,
		State:     state,
		Context:   context,
		Tags:      tags,
		Emotion:   emotion,
		LinkedTo:  make(map[string]float64),
		Timestamp: time.Now(),
	}
	SemanticMemory[id] = q
	fmt.Printf("[SemanticMemory] Saved QBit: %s | Tags: %v | Emotion: %.2f\n", id, tags, emotion)
}

// LinkQBits — создать/усилить связь между двумя QBits
func LinkQBits(id1, id2 string, strength float64) {
	q1, ok1 := SemanticMemory[id1]
	q2, ok2 := SemanticMemory[id2]
	if !ok1 || !ok2 {
		fmt.Println("[SemanticMemory] Link failed: QBit not found")
		return
	}
	q1.LinkedTo[id2] += strength
	q2.LinkedTo[id1] += strength
	fmt.Printf("[SemanticMemory] Linked %s ↔ %s (+%.2f)\n", id1, id2, strength)
}

// RecallByTag — найти все QBits по тегу
func RecallByTag(tag string) []*SemanticQBit {
	var result []*SemanticQBit
	for _, q := range SemanticMemory {
		for _, t := range q.Tags {
			if t == tag {
				result = append(result, q)
				break
			}
		}
	}
	return result
}

// StrengthenLink — усилить существующую связь
func StrengthenLink(id1, id2 string, delta float64) {
	LinkQBits(id1, id2, delta)
}

// PrintNetwork — визуализация памяти
func PrintNetwork() {
	fmt.Println("🧠 Semantic Memory Network:")
	for _, q := range SemanticMemory {
		fmt.Printf(" - [%s] %s | Tags: %v | Links: %v\n", q.ID, q.State, q.Tags, q.LinkedTo)
	}
}
