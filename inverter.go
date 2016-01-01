package behavior

type inverter struct {
	Child Behavior
}

func (i *inverter) Init() BehaviorData {
	return nil
}

func (i *inverter) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	result := r.Next(i.Child)
	if result.Complete {
		return Result{Failed: true}, nil
	}
	if result.Failed {
		return Result{Complete: true}, nil
	}
	return Result{}, nil
}

func Inverter(Child Behavior) Behavior {
	return &inverter{
		Child: Child,
	}
}
