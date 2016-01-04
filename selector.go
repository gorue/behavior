package behavior

type selector struct {
	BehaviorParent
}

func (s *selector) Init() BehaviorData {
	return nil
}

func (s *selector) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	d := 0
	for {
		result := r.Next(d)
		switch result {
		case SUCCESS:
			return SUCCESS, nil
		case FAILURE:
			d++
			if d >= s.NumChild() {
				return FAILURE, nil
			}
		case RUNNING:
			return RUNNING, nil
		}
	}
	return FAILURE, nil
}

func Selector(Choices ...Behavior) Behavior {
	return &selector{
		BehaviorParent: Children(Choices...),
	}
}
