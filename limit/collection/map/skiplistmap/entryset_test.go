package skiplistmap

import "testing"

var mes *EntrySet
var smes *EntrySet

func RestartEntrySet() {
	RestartSubMap()
	es1 := m.EntrySet()
	mes = (*es1).(*EntrySet)
	es2 := sm.EntrySet()
	smes = (*es2).(*EntrySet)
}

func TestEntrySet_Clear(t *testing.T) {

}

func TestEntrySet_Contains(t *testing.T) {

}

func TestEntrySet_Empty(t *testing.T) {

}

func TestEntrySet_Equals(t *testing.T) {

}

func TestEntrySet_GetIterator(t *testing.T) {

}

func TestEntrySet_Remove(t *testing.T) {

}

func TestEntrySet_RemoveAll(t *testing.T) {

}

func TestEntrySet_RetainAll(t *testing.T) {

}

func TestEntrySet_String(t *testing.T) {

}