package types

type HardPoints struct {
	Double  *float64
	Integer *int64
}

func (x *HardPoints) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, &x.Double, nil, nil, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *HardPoints) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, x.Double, nil, nil, false, nil, false, nil, false, nil, false, nil, false)
}
