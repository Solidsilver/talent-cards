package types

type Description struct {
	AnythingArray []interface{}
	String        *string
}

func (x *Description) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.AnythingArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Description) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.AnythingArray != nil, x.AnythingArray, false, nil, false, nil, false, nil, false)
}
