package behavior

import "sync"

type Runner struct {
	b          Behavior
	bd         BehaviorData
	vars       map[string]interface{}
	varlock    sync.Mutex
	stack      []BehaviorData
	stackDepth int
}

func NewRunner(b Behavior) *Runner {
	return &Runner{
		b:    b,
		bd:   b.Init(),
		vars: make(map[string]interface{}),
	}
}

func (r *Runner) Step() Result {
	return r.Next(r.b)
}

func (r *Runner) Next(b Behavior) (res Result) {
	if r.stackDepth >= len(r.stack) {
		r.stack = append(r.stack, b.Init())
	}
	bd := r.stack[r.stackDepth]
	r.stackDepth++
	res, bd = b.Step(r, bd)
	r.stackDepth--
	if res.Failed || res.Complete {
		r.stack = r.stack[:r.stackDepth]
	} else {
		r.stack[r.stackDepth] = bd
	}
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
