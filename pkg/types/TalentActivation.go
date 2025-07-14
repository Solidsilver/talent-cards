package types

type TalentActivation string

const (
	ActiveAction              TalentActivation = "active (action)"
	ActiveIncidental          TalentActivation = "active (incidental)"
	ActiveIncidentalOutOfTurn TalentActivation = "active (incidental, out of turn)"
	ActiveManeuver            TalentActivation = "active (maneuver)"
	FluffyPassive             TalentActivation = "passive"
)
