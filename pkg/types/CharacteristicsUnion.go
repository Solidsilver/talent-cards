package types

type CharacteristicsUnion struct {
	CharacteristicElementArray     []CharacteristicElement
	CharacteristicsCharacteristics *CharacteristicsCharacteristics
}

func (x *CharacteristicsUnion) UnmarshalJSON(data []byte) error {
	x.CharacteristicElementArray = nil
	x.CharacteristicsCharacteristics = nil
	var c CharacteristicsCharacteristics
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.CharacteristicElementArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CharacteristicsCharacteristics = &c
	}
	return nil
}

func (x *CharacteristicsUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.CharacteristicElementArray != nil, x.CharacteristicElementArray, x.CharacteristicsCharacteristics != nil, x.CharacteristicsCharacteristics, false, nil, false, nil, false)
}
