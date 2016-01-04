package behavior

type memSelector struct {
	BehaviorParent
}

func (s *memSelector) Init() BehaviorData {
	return 0
}

func (s *memSelector) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	d := bd.(int)
	for {
		result := r.Next(d)
		switch result {
		case SUCCESS:
			return SUCCESS, d
		case FAILURE:
			d++
			if d >= s.NumChild() {
				return FAILURE, d
			}
		case RUNNING:
			return RUNNING, d
		}
	}
	return FAILURE, d
}

func MemSelector(Choices ...Behavior) Behavior {
	return &memSelector{
		BehaviorParent: Children(Choices...),
	}
}
