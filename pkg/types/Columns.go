package types

type Columns struct {
	ColumnsColumnArray []ColumnsColumn
	Integer            *int64
}

func (x *Columns) UnmarshalJSON(data []byte) error {
	x.ColumnsColumnArray = nil
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, nil, true, &x.ColumnsColumnArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Columns) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, nil, x.ColumnsColumnArray != nil, x.ColumnsColumnArray, false, nil, false, nil, false, nil, false)
}
