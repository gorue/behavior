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
	return SUCCESS, nil
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
	return SUCCESS, nil
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
			Repeater(&PrintVar{Name: "test"}, 5),
		),
	)

	runner := NewRunner(simpleTree)

	res := runner.Step()
	for res == RUNNING {
		res = runner.Step()
	}
}
