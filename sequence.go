package behavior

type sequence struct {
	Steps []Behavior
}

type sequenceState struct {
	step int
}

func (s *sequence) Init() BehaviorData {
	bd := &sequenceState{}
	return bd
}

func (s *sequence) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	d := bd.(*sequenceState)
	result := r.Next(s.Steps[d.step])
	if result.Complete {
		d.step++
		if d.step >= len(s.Steps) {
			return Result{Complete: true}, d
		}
	}
	if result.Failed {
		return Result{Failed: true}, d
	}
	return Result{}, d
}

func Sequence(Steps ...Behavior) Behavior {
	return &sequence{
		Steps: Steps,
	}
}
