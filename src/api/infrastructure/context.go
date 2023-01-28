package infrastructure

type ctxKey string

func (c ctxKey) String() string {
	return "fast_truck_context_key_" + string(c)
}

const ActionKey = ctxKey("action")
