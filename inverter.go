package behavior

type inverter struct {
	BehaviorParent
}

func (i *inverter) Init() BehaviorData {
	return nil
}

func (i *inverter) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	result := r.Next(0)
	if result == SUCCESS {
		return FAILURE, nil
	}
	if result == FAILURE {
		return SUCCESS, nil
	}
	return RUNNING, nil
}

func Inverter(Child Behavior) Behavior {
	return &inverter{
		BehaviorParent: Children(Child),
	}
}
