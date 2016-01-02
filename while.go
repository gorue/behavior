package behavior

type while struct {
	Children []Behavior
}

func (w *while) Init() BehaviorData {
	return 0
}

func (w *while) Step(run *Runner, bd BehaviorData) (Result, BehaviorData) {
	d := bd.(int)
	result := run.Next(w.Children[d])
	if d == 0 {
		if result == FAILURE {
			return FAILURE, d
		}
	}
	if result == FAILURE {
		return RUNNING, 0
	}
	if result == SUCCESS {
		d++
		if d >= len(w.Children) {
			return SUCCESS, d
		}
	}
	return RUNNING, d
}

func While(Condition Behavior, Children ...Behavior) Behavior {
	return &while{
		Children: append([]Behavior{Condition}, Children...),
	}
}
