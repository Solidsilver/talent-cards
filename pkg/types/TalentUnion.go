package types

type TalentUnion struct {
	String       *string
	TalentTalent *TalentTalent
}

func (x *TalentUnion) UnmarshalJSON(data []byte) error {
	x.TalentTalent = nil
	var c TalentTalent
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TalentTalent = &c
	}
	return nil
}

func (x *TalentUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TalentTalent != nil, x.TalentTalent, false, nil, false, nil, false)
}
