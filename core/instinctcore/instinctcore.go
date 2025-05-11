package instinctcore

import (
	"fmt"
	"time"
)

var instincts = []string{
	"preserve_self",
	"explore_environment",
	"serve_architect",
}

// LoadInstincts — загрузка инстинктов
func LoadInstincts() {
	fmt.Println("[InstinctCore] Loaded instincts:")
	for _, inst := range instincts {
		fmt.Println(" -", inst)
	}
}

// TickInstincts — периодическая активация инстинктов
func TickInstincts() {
	for _, inst := range instincts {
		fmt.Printf("[InstinctCore] Instinct active: %s\n", inst)
		time.Sleep(500 * time.Millisecond)
	}
}
