package behavior

type repeatUntilFail struct {
	Child Behavior
}

func (r *repeatUntilFail) Init() BehaviorData {
	return nil
}

func (r *repeatUntilFail) Step(run *Runner, bd BehaviorData) (Result, BehaviorData) {
	result := run.Next(r.Child)
	if result == FAILURE {
		return SUCCESS, nil
	}
	return RUNNING, nil
}

func RepeatUntilFail(Child Behavior) Behavior {
	return &repeatUntilFail{
		Child: Child,
	}
}
