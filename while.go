package behavior

type while struct {
	BehaviorParent
}

func (w *while) Init() BehaviorData {
	return 0
}

func (w *while) Step(run *Runner, bd BehaviorData) (Result, BehaviorData) {
	d := bd.(int)
	result := run.Next(d)
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
		if d >= w.NumChild() {
			return SUCCESS, d
		}
	}
	return RUNNING, d
}

func While(Condition Behavior, Child ...Behavior) Behavior {
	return &while{
		BehaviorParent: Children(append([]Behavior{Condition}, Child...)...),
	}
}
