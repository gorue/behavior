// behavior project behavior.go
package behavior

type Result int

const (
	RUNNING Result = iota
	SUCCESS
	FAILURE
)

type Behavior interface {
	Init() BehaviorData
	Step(*Runner, BehaviorData) (Result, BehaviorData)
}

type BehaviorData interface{}
