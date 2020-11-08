package skiplistmap

import "testing"

var mvs *Values
var smvs *Values

func RestartValues() {
	RestartSubMap()
	vs1 := m.Values()
	mvs = (*vs1).(*Values)
	vs2 := sm.Values()
	smvs = (*vs2).(*Values)
}

func TestValues_Clear(t *testing.T) {

}

func TestValues_Contains(t *testing.T) {

}

func TestValues_Empty(t *testing.T) {

}

func TestValues_GetIterator(t *testing.T) {

}

func TestValues_String(t *testing.T) {

}
