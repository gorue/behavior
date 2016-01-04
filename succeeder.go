package behavior

type succeeder struct {
	BehaviorParent
}

func (i *succeeder) Init() BehaviorData {
	return nil
}

func (i *succeeder) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	result := r.Next(0)
	if result == SUCCESS || result == FAILURE {
		return SUCCESS, nil
	}
	return RUNNING, nil
}

func Succeeder(Child Behavior) Behavior {
	return &succeeder{
		BehaviorParent: Children(Child),
	}
}
