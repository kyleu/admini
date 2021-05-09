package schematypes

import (
	"encoding/json"
	"fmt"

	"github.com/kyleu/admini/app/util"
)

type Wrapped struct {
	K string `json:"k"`
	T Type   `json:"t,omitempty"`
}

var _ Type = (*Wrapped)(nil)

func Wrap(t Type) *Wrapped {
	w, ok := t.(*Wrapped)
	if ok {
		return w
	}
	return &Wrapped{K: t.Key(), T: t}
}

func (x *Wrapped) Key() string {
	return x.K
}

func (x *Wrapped) Sortable() bool {
	return x.T.Sortable()
}

func (x *Wrapped) String() string {
	return x.T.String()
}

func (x *Wrapped) From(v interface{}) interface{} {
		return x.T.From(v)
}

type wrappedUnmarshal struct {
	K string          `json:"k"`
	T json.RawMessage `json:"t,omitempty"`
}

func (x *Wrapped) MarshalJSON() ([]byte, error) {
	b := util.ToJSONBytes(x.T, true)
	if len(b) == 2 {
		return json.Marshal(x.K)
	}
	return json.Marshal(wrappedUnmarshal{K: x.K, T: b})
}

// nolint
func (x *Wrapped) UnmarshalJSON(data []byte) error {
	var wu wrappedUnmarshal
	err := json.Unmarshal(data, &wu)
	if err != nil {
		str := ""
		newErr := json.Unmarshal(data, &str)
		if newErr != nil {
			return err
		}
		wu = wrappedUnmarshal{K: str, T: []byte("{}")}
	}
	var t Type
	switch wu.K {
	case KeyBit:
		tgt := &Bit{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyBool:
		tgt := &Bool{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyByte:
		tgt := &Byte{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyChar:
		tgt := &Char{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyDate:
		tgt := &Date{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyEnumValue:
		tgt := &EnumValue{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyError:
		tgt := &Error{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyFloat:
		tgt := &Float{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyInt:
		tgt := &Int{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyJSON:
		tgt := &JSON{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyList:
		tgt := &List{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyMap:
		tgt := &Map{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyMethod:
		tgt := &Method{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyNil:
		tgt := &Nil{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyOption:
		tgt := &Option{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyRange:
		tgt := &Range{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyReference:
		tgt := &Reference{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeySet:
		tgt := &Set{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyString:
		tgt := &String{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyTime:
		tgt := &Time{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyTimestamp:
		tgt := &Timestamp{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyTimestampZoned:
		tgt := &TimestampZoned{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyUnknown:
		tgt := &Unknown{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyUUID:
		tgt := &UUID{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	case KeyXML:
		tgt := &XML{}
		err = json.Unmarshal(wu.T, &tgt)
		t = tgt
	default:
		t = &Unknown{X: "unmarshal:" + wu.K}
	}
	if err != nil {
		return fmt.Errorf("unable to unmarshal wrapped field of type [%v]: %w", wu.K, err)
	}
	if t == nil {
		return fmt.Errorf("nil type returned from unmarshal")
	}
	x.K = wu.K
	x.T = t
	return nil
}
