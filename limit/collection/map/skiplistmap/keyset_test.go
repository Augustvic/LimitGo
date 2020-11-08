package skiplistmap

import "testing"

var mks *KeySet
var smks *KeySet

func RestartKeySet() {
	RestartSubMap()
	ks1 := m.KeySet()
	mks = (*ks1).(*KeySet)
	ks2 := sm.KeySet()
	smks = (*ks2).(*KeySet)
}

func TestKeySet_Ceiling(t *testing.T) {

}

func TestKeySet_Clear(t *testing.T) {

}

func TestKeySet_Contains(t *testing.T) {

}

func TestKeySet_Empty(t *testing.T) {

}

func TestKeySet_Equals(t *testing.T) {

}

func TestKeySet_First(t *testing.T) {

}

func TestKeySet_Floor(t *testing.T) {

}

func TestKeySet_GetIterator(t *testing.T) {

}

func TestKeySet_HeadSet(t *testing.T) {

}

func TestKeySet_Higher(t *testing.T) {

}

func TestKeySet_Last(t *testing.T) {

}

func TestKeySet_Lower(t *testing.T) {

}

func TestKeySet_PollFirst(t *testing.T) {

}

func TestKeySet_PollLast(t *testing.T) {

}

func TestKeySet_Remove(t *testing.T) {

}

func TestKeySet_RemoveAll(t *testing.T) {

}

func TestKeySet_RetainAll(t *testing.T) {

}

func TestKeySet_String(t *testing.T) {

}

func TestKeySet_SubSet(t *testing.T) {

}

func TestKeySet_TailSet(t *testing.T) {

}