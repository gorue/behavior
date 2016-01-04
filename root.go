package behavior

import "sync"

type behaviorDataTreeStack struct {
	Nodes []*behaviorDataTreeStack
	Data  BehaviorData
}

type Runner struct {
	b             Behavior
	vars          map[string]interface{}
	varlock       sync.Mutex
	dataTree      *behaviorDataTreeStack
	dataTreeStack *behaviorDataTreeStack
	stack         Behavior
}

func NewRunner(b Behavior) *Runner {
	dt := &behaviorDataTreeStack{}
	return &Runner{
		b:        b,
		dataTree: dt,
		vars:     make(map[string]interface{}),
	}
}

func (r *Runner) Step() Result {
	return r.Next(0)
}

func (r *Runner) Next(i int) (res Result) {
	preStack := r.stack
	preDataStack := r.dataTreeStack
	var b Behavior
	var bds *behaviorDataTreeStack
	if preStack == nil {
		b = r.b
		bds = r.dataTree
	} else {
		bp := preStack.(BehaviorParent)
		b = bp.Child(i)
		if preDataStack.Nodes == nil {
			preDataStack.Nodes = make([]*behaviorDataTreeStack, bp.NumChild())
		}
		bds = preDataStack.Nodes[i]
		if bds == nil {
			bds = &behaviorDataTreeStack{}
			preDataStack.Nodes[i] = bds
		}
	}
	r.stack = b
	r.dataTreeStack = bds

	if bds.Data == nil {
		bds.Data = b.Init()
	}

	res, bds.Data = b.Step(r, bds.Data)

	if res != RUNNING {
		bds.Data = nil
	}

	r.dataTreeStack = preDataStack
	r.stack = preStack
	return
}

func (r *Runner) SetVar(name string, val interface{}) {
	r.varlock.Lock()
	defer r.varlock.Unlock()
	r.vars[name] = val
}

func (r *Runner) Var(name string) interface{} {
	r.varlock.Lock()
	defer r.varlock.Unlock()
	return r.vars[name]
}
