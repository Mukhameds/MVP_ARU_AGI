package main

import (
	"fmt"
	"time"

	"github.com/Mukhameds/MVP_ARU_AGI/core/flowengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/signalengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/memoryengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/instinctcore"
	"github.com/Mukhameds/MVP_ARU_AGI/core/reflexengine"
	"github.com/Mukhameds/MVP_ARU_AGI/core/ghostfield"
	"github.com/Mukhameds/MVP_ARU_AGI/core/hemisphere"
)

func main() {
	fmt.Println("🔵 ARU-MVP initializing...")

	memoryengine.InitMemory()
	flowengine.InitFlowEngine(1)
	hemisphere.ActivateHemispheres([]string{"logic", "emotion", "goal"})
	reflexengine.LoadDefaultReflexes()
	instinctcore.LoadInstincts()
	signalengine.InitSignalEngine()

	go loopGhostField()
	go loopInstincts()

	fmt.Println("✅ ARU-MVP ready.")
	select {}
}

func loopGhostField() {
	for {
		ghostfield.TickGhostField()
		time.Sleep(5 * time.Second)
	}
}

func loopInstincts() {
	for {
		instinctcore.TickInstincts()
		time.Sleep(3 * time.Second)
	}
}
