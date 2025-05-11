package hemisphere

import (
	"fmt"
)

// Hemisphere — структура полушария
type Hemisphere struct {
	Name    string
	Active  bool
	Purpose string
}

var Hemispheres = make(map[string]*Hemisphere)

// ActivateHemispheres — активация нужных полушарий
func ActivateHemispheres(names []string) {
	for _, name := range names {
		h := &Hemisphere{
			Name:    name,
			Active:  true,
			Purpose: describe(name),
		}
		Hemispheres[name] = h
		fmt.Printf("[Hemisphere] Activated: %s → %s\n", h.Name, h.Purpose)
	}
}

// describe — назначение по имени
func describe(name string) string {
	switch name {
	case "logic":
		return "Logical thinking, analysis"
	case "emotion":
		return "Emotional response and tagging"
	case "goal":
		return "Motivation and directional thinking"
	default:
		return "Generic hemisphere"
	}
}
