package behavior

type succeeder struct {
	Child Behavior
}

func (i *succeeder) Init() BehaviorData {
	return nil
}

func (i *succeeder) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	result := r.Next(i.Child)
	if result.Complete || result.Failed {
		return Result{Complete: true}, nil
	}
	return Result{}, nil
}

func Succeeder(Child Behavior) Behavior {
	return &succeeder{
		Child: Child,
	}
}
