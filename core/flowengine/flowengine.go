package flowengine

import (
	"fmt"
	"time"

	"github.com/Mukhameds/MVP_ARU_AGI/types"
)

type ThoughtThread struct {
	ID        string
	Signal    types.Signal
	StartTime time.Time
	Active    bool
}

type FlowEngine struct {
	ActiveThreads map[string]*ThoughtThread
	PriorityQueue []types.Signal
	MaxConcurrent int
}

var engine = FlowEngine{
	ActiveThreads: make(map[string]*ThoughtThread),
	PriorityQueue: []types.Signal{},
	MaxConcurrent: 1,
}

// Инициализация потока мышления
func InitFlowEngine(maxThreads int) {
	engine.MaxConcurrent = maxThreads
	fmt.Println("[FlowEngine] Initialized with MaxConcurrent =", maxThreads)
}

// Планирование сигнала
func Schedule(th types.Signal) {
	if len(engine.ActiveThreads) < engine.MaxConcurrent {
		thread := &ThoughtThread{
			ID:        th.ID,
			Signal:    th,
			StartTime: time.Now(),
			Active:    true,
		}
		engine.ActiveThreads[th.ID] = thread
		go runThread(thread)
	} else {
		engine.PriorityQueue = append(engine.PriorityQueue, th)
		fmt.Println("[FlowEngine] Queued:", th.Content)
	}
}

// Исполнение потока мышления
func runThread(t *ThoughtThread) {
	fmt.Printf("[FlowEngine] Thinking on signal: %s (%s)\n", t.Signal.ID, t.Signal.Content)
	time.Sleep(2 * time.Second)

	t.Active = false
	delete(engine.ActiveThreads, t.ID)
	fmt.Println("[FlowEngine] Thread completed:", t.Signal.ID)

	checkQueue()
}

func checkQueue() {
	if len(engine.PriorityQueue) > 0 {
		next := engine.PriorityQueue[0]
		engine.PriorityQueue = engine.PriorityQueue[1:]
		Schedule(next)
	}
}

// CreateThread — создаёт объект потока, но не запускает его
func CreateThread(signal types.Signal) ThoughtThread {
	return ThoughtThread{
		ID:        signal.ID,
		Signal:    signal,
		StartTime: time.Now(),
		Active:    false, // ещё не запущен
	}
}
