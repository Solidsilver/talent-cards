package types

type PriceUnion struct {
	Enum    *PriceEnum
	Integer *int64
}

func (x *PriceUnion) UnmarshalJSON(data []byte) error {
	x.Enum = nil
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, nil, false, nil, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *PriceUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, nil, false, nil, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}
