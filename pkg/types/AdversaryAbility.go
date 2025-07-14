package types

type AdversaryAbility struct {
	PurpleAbility *PurpleAbility
	String        *string
}

func (x *AdversaryAbility) UnmarshalJSON(data []byte) error {
	x.PurpleAbility = nil
	var c PurpleAbility
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleAbility = &c
	}
	return nil
}

func (x *AdversaryAbility) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleAbility != nil, x.PurpleAbility, false, nil, false, nil, false)
}
