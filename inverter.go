package behavior

type inverter struct {
	Child Behavior
}

func (i *inverter) Init() BehaviorData {
	return nil
}

func (i *inverter) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	result := r.Next(i.Child)
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
		Child: Child,
	}
}
