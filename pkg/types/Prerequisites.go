package types

type Prerequisites struct {
	Bool        *bool
	StringArray []string
}

func (x *Prerequisites) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Prerequisites) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}
