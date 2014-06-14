package safemap

import "testing"

const (
	K         = "huang"
	V         = "anxin"
	TESTVALUE = "foo"
)

func TestInsert(t *testing.T) {
	smap := New()
	smap.Insert(K, V)
	inserted := smap.Close()
	if inserted[K] != V {
		t.Errorf(`Unexpected value %q, expected "Yo bob"`, inserted[K])
	}
}
func TestLen(t *testing.T) {
	smap := New()
	if smap.Len() != 0 {
		t.Errorf(`Unexpected length %d, expected 0`, smap.Len())
	}
	smap.Insert(K, V)
	if smap.Len() != 1 {
		t.Errorf(`Unexpected length %d, expected 1`, smap.Len())
	}
}
func updater(originValue interface{}, found bool) interface{} {
	if found {
		return originValue.(string) + TESTVALUE
	}
	return originValue
}
func TestUpdate(t *testing.T) {

	smap := New()
	smap.Insert(K, V)
	smap.Update(K, updater)
	updated := smap.Close()
	if updated[K] != (V + TESTVALUE) {
		t.Errorf(`Unexpected value %s, expected %s`, updated[K].(string), (V + TESTVALUE))
	}
}

func TestFind(t *testing.T) {

	smap := New()
	smap.Insert(K, V)
	v, found := smap.Find(K)
	if !found {
		t.Errorf(`Unexpected !`)
	}
	if v.(string) != V {
		t.Errorf(`Unexpected value %s, expected %s`, v.(string), V)
	}

}
