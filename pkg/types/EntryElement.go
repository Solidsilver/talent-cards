package types

type EntryElement struct {
	ChapterClass *ChapterClass
	String       *string
}

func (x *EntryElement) UnmarshalJSON(data []byte) error {
	x.ChapterClass = nil
	var c ChapterClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.ChapterClass = &c
	}
	return nil
}

func (x *EntryElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.ChapterClass != nil, x.ChapterClass, false, nil, false, nil, false)
}
