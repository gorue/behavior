package behavior

type repeater struct {
	Child    Behavior
	MaxCount int64
}

func (r *repeater) Init() BehaviorData {
	return int64(0)
}

func (r *repeater) Step(run *Runner, bd BehaviorData) (Result, BehaviorData) {
	d := bd.(int64)
	result := run.Next(r.Child)
	if result == SUCCESS || result == FAILURE {
		d++
		if d >= r.MaxCount && r.MaxCount != 0 {
			return SUCCESS, d
		}
	}
	return RUNNING, d
}

func Repeater(Child Behavior, maxCount int64) Behavior {
	return &repeater{
		Child:    Child,
		MaxCount: maxCount,
	}
}
