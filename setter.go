package defaults

// Setter is an interface for setting default values
type Setter interface {
	SetDefaults()
}

func callSetter(v any) {
	if ds, ok := v.(Setter); ok {
		ds.SetDefaults()
	}
}
