package types

type St struct {
	Integer      *int64
	IntegerArray []int64
}

func (x *St) UnmarshalJSON(data []byte) error {
	x.IntegerArray = nil
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, nil, true, &x.IntegerArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *St) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, nil, x.IntegerArray != nil, x.IntegerArray, false, nil, false, nil, false, nil, false)
}
