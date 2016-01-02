package behavior

type untilFail struct {
	Child Behavior
}

func (r *untilFail) Init() BehaviorData {
	return nil
}

func (r *untilFail) Step(run *Runner, bd BehaviorData) (Result, BehaviorData) {
	result := run.Next(r.Child)
	if result == FAILURE {
		return SUCCESS, nil
	}
	return RUNNING, nil
}

func UntilFail(Child Behavior) Behavior {
	return &untilFail{
		Child: Child,
	}
}
