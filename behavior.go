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

type BehaviorParent interface {
	Child(int) Behavior
	NumChild() int
}

type BehaviorData interface{}

type children struct {
	nodes []Behavior
}

func (c *children) Child(i int) Behavior {
	return c.nodes[i]
}

func (c *children) NumChild() int {
	return len(c.nodes)
}

func Children(child ...Behavior) *children {
	return &children{
		nodes: child,
	}
}
