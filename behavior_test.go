package behavior

import (
	"fmt"
	"testing"
)

type PrintVar struct {
	Name string
}

func (pb *PrintVar) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	fmt.Println(r.Var(pb.Name))
	return Result{Complete: true}, nil
}

func (pb *PrintVar) Init() BehaviorData {
	return nil
}

type SetVar struct {
	Name string
	Val  interface{}
}

func (pb *SetVar) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	r.SetVar(pb.Name, pb.Val)
	return Result{Complete: true}, nil
}

func (pb *SetVar) Init() BehaviorData {
	return nil
}

func TestBehaviourTree(t *testing.T) {

	simpleTree := Sequence(
		&SetVar{Name: "test", Val: "this"},
		&PrintVar{Name: "test"},
		Sequence(
			&SetVar{Name: "test", Val: "that"},
			&PrintVar{Name: "test"},
		),
	)

	runner := NewRunner(simpleTree)

	res := runner.Step()
	for !res.Complete {
		res = runner.Step()
	}
}
