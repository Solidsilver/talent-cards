package types

type Variant string

const (
	Error   Variant = "error"
	Info    Variant = "info"
	Success Variant = "success"
	Warn    Variant = "warn"
)
