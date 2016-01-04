package behavior

type sequence struct {
	BehaviorParent
}

func (s *sequence) Init() BehaviorData {
	return nil
}

func (s *sequence) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	d := 0
	for {
		result := r.Next(d)
		switch result {
		case SUCCESS:
			d++
			if d >= s.NumChild() {
				return SUCCESS, nil
			}
		case FAILURE:
			return FAILURE, nil
		case RUNNING:
			return RUNNING, nil
		}
	}
	return FAILURE, nil
}

func Sequence(Steps ...Behavior) Behavior {
	return &sequence{
		BehaviorParent: Children(Steps...),
	}
}
