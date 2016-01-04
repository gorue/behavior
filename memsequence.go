package behavior

type memSequence struct {
	BehaviorParent
}

func (s *memSequence) Init() BehaviorData {
	return 0
}

func (s *memSequence) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	d := bd.(int)
	for {
		result := r.Next(d)
		switch result {
		case SUCCESS:
			d++
			if d >= s.NumChild() {
				return SUCCESS, d
			}
		case FAILURE:
			return FAILURE, d
		case RUNNING:
			return RUNNING, d
		}
	}
	return FAILURE, d
}

func MemSequence(Steps ...Behavior) Behavior {
	return &memSequence{
		BehaviorParent: Children(Steps...),
	}
}
