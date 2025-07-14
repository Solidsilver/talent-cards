package types

type QualityUnion struct {
	QualityQuality *QualityQuality
	String         *string
}

func (x *QualityUnion) UnmarshalJSON(data []byte) error {
	x.QualityQuality = nil
	var c QualityQuality
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.QualityQuality = &c
	}
	return nil
}

func (x *QualityUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.QualityQuality != nil, x.QualityQuality, false, nil, false, nil, false)
}
