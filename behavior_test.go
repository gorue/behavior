package behavior

import (
	"fmt"
	"testing"
)

type PrintVar struct {
	Name string
}

func (pb *PrintVar) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	if bd.(bool) == false {
		fmt.Println(r.Var(pb.Name))
		return RUNNING, true
	}
	return SUCCESS, true
}

func (pb *PrintVar) Init() BehaviorData {
	return false
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
		&SetVar{Name: "test2", Val: "this2"},
		MemSequence(
			&PrintVar{Name: "test2"},
			&SetVar{Name: "test", Val: "that"},
			Repeater(&PrintVar{Name: "test"}, 5),
		),
	)

	runner := NewRunner(simpleTree)

	for l1 := 0; l1 < 30; l1++ {
		res := runner.Step()
		if res == SUCCESS {
			break
		}
	}
}

/*
func BenchmarkBehaviourTree(b *testing.B) {

	simpleTree := Repeater(MemSequence(
		&SetVar{Name: "test", Val: "this"},
		&PrintVar{Name: "test"},
		Sequence(
			&SetVar{Name: "test", Val: "that"},
			Repeater(&PrintVar{Name: "test"}, 5),
		),
	), 0)
	runner := NewRunner(simpleTree)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runner.Step()
	}
}
*/
