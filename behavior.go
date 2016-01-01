// behavior project behavior.go
package behavior

type Result struct {
	Complete bool
	Failed   bool
}

type Behavior interface {
	Init() BehaviorData
	Step(*Runner, BehaviorData) (Result, BehaviorData)
}

type BehaviorData interface{}
