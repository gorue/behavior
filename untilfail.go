package behavior

type untilFail struct {
	BehaviorParent
}

func (r *untilFail) Init() BehaviorData {
	return nil
}

func (r *untilFail) Step(run *Runner, bd BehaviorData) (Result, BehaviorData) {
	result := run.Next(0)
	if result == FAILURE {
		return SUCCESS, nil
	}
	return RUNNING, nil
}

func UntilFail(Child Behavior) Behavior {
	return &untilFail{
		BehaviorParent: Children(Child),
	}
}
