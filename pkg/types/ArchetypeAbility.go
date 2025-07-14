package types

type ArchetypeAbility struct {
	FluffyAbility *FluffyAbility
	String        *string
}

func (x *ArchetypeAbility) UnmarshalJSON(data []byte) error {
	x.FluffyAbility = nil
	var c FluffyAbility
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyAbility = &c
	}
	return nil
}

func (x *ArchetypeAbility) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyAbility != nil, x.FluffyAbility, false, nil, false, nil, false)
}
