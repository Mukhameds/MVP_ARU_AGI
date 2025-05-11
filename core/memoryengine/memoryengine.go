package memoryengine

import (
	"fmt"
)

// MemoryEngine — простая инициализация памяти
var Memory = make(map[string]string)

// InitMemory — запуск хранилища памяти
func InitMemory() {
	fmt.Println("[MemoryEngine] Memory initialized")
}

// Save — сохранить элемент в память
func Save(key, value string) {
	Memory[key] = value
	fmt.Printf("[MemoryEngine] Saved: %s → %s\n", key, value)
}

// Recall — извлечь элемент из памяти
func Recall(key string) string {
	if val, ok := Memory[key]; ok {
		fmt.Printf("[MemoryEngine] Recalled: %s → %s\n", key, val)
		return val
	}
	fmt.Printf("[MemoryEngine] Key not found: %s\n", key)
	return ""
}
