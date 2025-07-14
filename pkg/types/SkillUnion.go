package types

type SkillUnion struct {
	SkillSkill *SkillSkill
	String     *string
}

func (x *SkillUnion) UnmarshalJSON(data []byte) error {
	x.SkillSkill = nil
	var c SkillSkill
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.SkillSkill = &c
	}
	return nil
}

func (x *SkillUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.SkillSkill != nil, x.SkillSkill, false, nil, false, nil, false)
}
