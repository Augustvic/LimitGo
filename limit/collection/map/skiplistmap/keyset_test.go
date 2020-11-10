package skiplistmap

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/linear/hashset"
	"strconv"
	"testing"
)

var mks *KeySet
var smks *KeySet

func RestartKeySet() {
	RestartSubMap()
	ks1 := m.KeySet()
	mks = (*ks1).(*KeySet)
	ks2 := sm.KeySet()
	smks = (*ks2).(*KeySet)
}

func TestKeySetAll(t *testing.T) {
	TestKeySet_Ceiling(t)
	TestKeySet_Clear(t)
	TestKeySet_Contains(t)
	TestKeySet_Empty(t)
	TestKeySet_Equals(t)
	TestKeySet_First(t)
	TestKeySet_Floor(t)
	TestKeySet_GetIterator(t)
	TestKeySet_HeadSet(t)
	TestKeySet_Higher(t)
	TestKeySet_Last(t)
	TestKeySet_Lower(t)
	TestKeySet_PollFirst(t)
	TestKeySet_PollLast(t)
	TestKeySet_Remove(t)
	TestKeySet_RemoveAll(t)
	TestKeySet_RetainAll(t)
	TestKeySet_String(t)
	TestKeySet_SubSet(t)
	TestKeySet_TailSet(t)
}

func TestKeySet_Ceiling(t *testing.T) {
	RestartKeySet()
	var temp1 collection.Object = Teacher{1, "t1", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{8, "t8", 0}
	var temp4 collection.Object = Teacher{10, "t10", 0}
	mk1 := mks.Ceiling(&temp1)
	if *mk1 != t2 {
		t.Error("Ceiling operation fail!")
	}
	mk2 := mks.Ceiling(&temp2)
	if *mk2 != t6 {
		t.Error("Ceiling operation fail!")
	}
	mk3 := mks.Ceiling(&temp3)
	if *mk3 != t8 {
		t.Error("Ceiling operation fail!")
	}
	mk4 := mks.Ceiling(&temp4)
	if mk4 != nil {
		t.Error("Ceiling operation fail!")
	}
	smk1 := smks.Ceiling(&temp1)
	if *smk1 != t4 {
		t.Error("Ceiling operation fail!")
	}
	smk2 := smks.Ceiling(&temp2)
	if *smk2 != t6 {
		t.Error("Ceiling operation fail!")
	}
	smk3 := mks.Ceiling(&temp3)
	if *smk3 != t8 {
		t.Error("Ceiling operation fail!")
	}
	smk4 := mks.Ceiling(&temp4)
	if smk4 != nil {
		t.Error("Ceiling operation fail!")
	}
}

func TestKeySet_Clear(t *testing.T) {
	RestartKeySet()
	if mks.Size() == 0 {
		t.Error("Start operation fail!")
	}
	mks.Clear()
	if mks.Size() != 0 || m.Size() != 0 {
		t.Error("Clear operation fail!")
	}
	RestartKeySet()
	if smks.Size() == 0 {
		t.Error("Start operation fail!")
	}
	smks.Clear()
	if smks.Size() != 0 || sm.Size() != 0 || m.Size() != 1 {
		t.Error("Clear operation fail!")
	}
}

func TestKeySet_Contains(t *testing.T) {
	RestartKeySet()
	var temp1 collection.Object = Teacher{2, "t2", 0}
	var temp2 collection.Object = Teacher{4, "t4", 0}
	var temp3 collection.Object = Teacher{5, "t5", 0}
	if !mks.Contains(&temp1) {
		t.Error("Contains operation fail!")
	}
	if !mks.Contains(&temp2) {
		t.Error("Contains operation fail!")
	}
	if mks.Contains(&temp3) {
		t.Error("Contains operation fail!")
	}
	if mks.Contains(nil) {
		t.Error("Contains operation fail!")
	}
	if smks.Contains(&temp1) {
		t.Error("Contains operation fail!")
	}
	if !smks.Contains(&temp2) {
		t.Error("Contains operation fail!")
	}
	if smks.Contains(&temp3) {
		t.Error("Contains operation fail!")
	}
	if smks.Contains(nil) {
		t.Error("Contains operation fail!")
	}
}

func TestKeySet_Empty(t *testing.T) {
	RestartKeySet()
	if mks.Empty() {
		t.Error("Start operation fail!")
	}
	mks.Clear()
	if !mks.Empty() {
		t.Error("Empty operation fail!")
	}
	RestartKeySet()
	if smks.Empty() {
		t.Error("Start operation fail!")
	}
	smks.Clear()
	if !smks.Empty() || mks.Empty() {
		t.Error("Empty operation fail!")
	}
}

func TestKeySet_Equals(t *testing.T) {
	RestartKeySet()
	var tt2 collection.Object
	var tt4 collection.Object
	var tt6 collection.Object
	var tt8 collection.Object
	var tks collection.Set
	var tsks collection.Set
	tt2 = Teacher{2, "t2", 0}
	tt4 = Teacher{4, "t4", 0}
	tt6 = Teacher{6, "t6", 0}
	tt8 = Teacher{8, "t8", 0}
	tks = hashset.New()
	tsks = hashset.New()
	tks.Add(&tt2)
	tks.Add(&tt4)
	tks.Add(&tt6)
	tks.Add(&tt8)
	tsks.Add(&tt4)
	tsks.Add(&tt6)
	tsks.Add(&tt8)
	if !mks.Equals(&tks) || !smks.Equals(&tsks) {
		t.Error("Equals operation fail!")
	}
}

func TestKeySet_First(t *testing.T) {
	RestartKeySet()
	k1 := mks.First()
	if *k1 != t2 {
		t.Error("First operation fail!")
	}
	k2 := smks.First()
	if *k2 != t4 {
		t.Error("First operation fail!")
	}
}

func TestKeySet_Floor(t *testing.T) {
	RestartKeySet()
	var temp1 collection.Object = Teacher{3, "t3", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{8, "t8", 0}
	var temp4 collection.Object = Teacher{11, "t10", 0}
	mk1 := mks.Floor(&temp1)
	if *mk1 != t2 {
		t.Error("Floor operation fail!")
	}
	mk2 := mks.Floor(&temp2)
	if *mk2 != t4 {
		t.Error("Floor operation fail!")
	}
	mk3 := mks.Floor(&temp3)
	if *mk3 != t8 {
		t.Error("Floor operation fail!")
	}
	mk4 := mks.Floor(&temp4)
	if *mk4 != t8 {
		t.Error("Floor operation fail!")
	}
	smk1 := smks.Floor(&temp1)
	if smk1 != nil {
		t.Error("Floor operation fail!")
	}
	smk2 := smks.Floor(&temp2)
	if *smk2 != t4 {
		t.Error("Floor operation fail!")
	}
	smk3 := mks.Floor(&temp3)
	if *smk3 != t8 {
		t.Error("Floor operation fail!")
	}
	smk4 := mks.Floor(&temp4)
	if *smk4 != t8 {
		t.Error("Floor operation fail!")
	}
}

func TestKeySet_GetIterator(t *testing.T) {
	RestartKeySet()
	it := mks.GetIterator()
	index := 2
	s := ""
	for i := 0; it.HashNext(); i++ {
		if i == index {
			it.Remove()
		} else {
			key := it.Next()
			teacher := (*key).(Teacher)
			k := strconv.Itoa(teacher.Id)
			s += k
		}
	}
	if mks.Size() != 3 || m.Size() != 3 || s != "2468" {
		t.Error("GetIterator operation fail!")
	}
}

func TestKeySet_HeadSet(t *testing.T) {
	RestartKeySet()
	tt1 := mks.HeadSet(&t6, false)
	if (*tt1).Size() != 2 || *(*tt1).First() != t2 ||  *(*tt1).Last() != t4 {
		t.Error("HeadSet operation fail!")
	}
	tt2 := smks.HeadSet(&t6, true)
	if (*tt2).Size() != 2 || *(*tt2).First() != t4 ||  *(*tt2).Last() != t6 {
		t.Error("HeadSet operation fail!")
	}
}

func TestKeySet_Higher(t *testing.T) {
	RestartKeySet()
	var temp1 collection.Object = Teacher{1, "t1", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{6, "t8", 0}
	var temp4 collection.Object = Teacher{10, "t10", 0}
	mk1 := mks.Higher(&temp1)
	if *mk1 != t2 {
		t.Error("Higher operation fail!")
	}
	mk2 := mks.Higher(&temp2)
	if *mk2 != t6 {
		t.Error("Higher operation fail!")
	}
	mk3 := mks.Higher(&temp3)
	if *mk3 != t8 {
		t.Error("Higher operation fail!")
	}
	mk4 := mks.Higher(&temp4)
	if mk4 != nil {
		t.Error("Higher operation fail!")
	}
	smk1 := smks.Higher(&temp1)
	if *smk1 != t4 {
		t.Error("Higher operation fail!")
	}
	smk2 := smks.Higher(&temp2)
	if *smk2 != t6 {
		t.Error("Higher operation fail!")
	}
	smk3 := mks.Higher(&temp3)
	if *smk3 != t8 {
		t.Error("Higher operation fail!")
	}
	smk4 := mks.Higher(&temp4)
	if smk4 != nil {
		t.Error("Higher operation fail!")
	}
}

func TestKeySet_Last(t *testing.T) {
	RestartKeySet()
	k1 := mks.Last()
	if *k1 != t8 {
		t.Error("Last operation fail!")
	}
	k2 := smks.Last()
	if *k2 != t8 {
		t.Error("Last operation fail!")
	}
}

func TestKeySet_Lower(t *testing.T) {
	RestartKeySet()
	var temp1 collection.Object = Teacher{3, "t1", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{8, "t8", 0}
	var temp4 collection.Object = Teacher{11, "t10", 0}
	mk1 := mks.Lower(&temp1)
	if *mk1 != t2 {
		t.Error("Lower operation fail!")
	}
	mk2 := mks.Lower(&temp2)
	if *mk2 != t4 {
		t.Error("Lower operation fail!")
	}
	mk3 := mks.Lower(&temp3)
	if *mk3 != t6 {
		t.Error("Lower operation fail!")
	}
	mk4 := mks.Lower(&temp4)
	if *mk4 != t8 {
		t.Error("Lower operation fail!")
	}
	smk1 := smks.Lower(&temp1)
	if smk1 != nil {
		t.Error("Lower operation fail!")
	}
	smk2 := smks.Lower(&temp2)
	if *smk2 != t4 {
		t.Error("Lower operation fail!")
	}
	smk3 := mks.Lower(&temp3)
	if *smk3 != t6 {
		t.Error("Lower operation fail!")
	}
	smk4 := mks.Lower(&temp4)
	if *smk4 != t8 {
		t.Error("Lower operation fail!")
	}
}

func TestKeySet_PollFirst(t *testing.T) {
	RestartKeySet()
	k1 := mks.PollFirst()
	if mks.Size() != 3 || *k1 != t2 || m.Size() != 3 {
		t.Error("PollFirst operation fail!")
	}
	k2 := smks.PollFirst()
	if smks.Size() != 2 || *k2 != t4 || m.Size() != 2 || sm.Size() != 2 {
		t.Error("PollFirst operation fail!")
	}
}

func TestKeySet_PollLast(t *testing.T) {
	RestartKeySet()
	k1 := mks.PollLast()
	if mks.Size() != 3 || *k1 != t8 || m.Size() != 3 {
		t.Error("PollLast operation fail!")
	}
	k2 := smks.PollLast()
	if smks.Size() != 1 || *k2 != t6 || m.Size() != 2 || sm.Size() != 1 {
		t.Error("PollLast operation fail!")
	}
}

func TestKeySet_Remove(t *testing.T) {
	RestartKeySet()
	mks.Remove(&t2)
	mks.Remove(&t4)
	k1 := mks.First()
	if mks.Size() != 2 || *k1 != t6 || m.Size() != 2 {
		t.Error("Remove operation fail!")
	}
	smks.Remove(&t8)
	k2 := smks.First()
	if smks.Size() != 1 || *k2 != t6 || sm.Size() != 1 || m.Size() != 1 {
		t.Error("Remove operation fail!")
	}
}

func TestKeySet_RemoveAll(t *testing.T) {
	RestartKeySet()
	var s1 collection.Linear = hashset.New()
	s1.Add(&t2)
	s1.Add(&t4)
	mks.RemoveAll(&s1)
	if mks.Size() != 2 || mks.Contains(&t2) || mks.Contains(&t4) {
		t.Error("RemoveAll operation fail!")
	}
	s1.Add(&t8)
	smks.RemoveAll(&s1)
	if smks.Size() != 1 || smks.Contains(&t2) || smks.Contains(&t4) || smks.Contains(&t8) {
		t.Error("RemoveAll operation fail!")
	}
}

func TestKeySet_RetainAll(t *testing.T) {
	RestartKeySet()
	var s1 collection.Linear = hashset.New()
	s1.Add(&t2)
	s1.Add(&t4)
	mks.RetainAll(&s1)
	if mks.Size() != 2 || !mks.Contains(&t2) || !mks.Contains(&t4) || mks.Contains(&t6) {
		t.Error("RetainAll operation fail!")
	}
	if smks.Size() != 1 || smks.Contains(&t2) || !smks.Contains(&t4) || smks.Contains(&t8) {
		t.Error("RetainAll operation fail!")
	}
}

func TestKeySet_String(t *testing.T) {
	RestartKeySet()
	if mks.String() != "{{\"Id\":2,\"Name\":\"t2\",\"Sex\":0},{\"Id\":4,\"Name\":\"t4\",\"Sex\":0},{\"Id\":6,\"Name\":\"t6\",\"Sex\":0},{\"Id\":8,\"Name\":\"t8\",\"Sex\":0}}" {
		t.Error("String operation fail!")
	}
	if smks.String() != "{{\"Id\":4,\"Name\":\"t4\",\"Sex\":0},{\"Id\":6,\"Name\":\"t6\",\"Sex\":0},{\"Id\":8,\"Name\":\"t8\",\"Sex\":0}}" {
		t.Error("String operation fail!")
	}
}

func TestKeySet_SubSet(t *testing.T) {
	RestartKeySet()
	tt1 := mks.SubSet(&t4, true, &t6, false)
	if (*tt1).Size() != 1 || *(*tt1).First() != t4 ||  *(*tt1).Last() != t4 {
		t.Error("SubSet operation fail!")
	}
	tt2 := smks.SubSet(nil, false, &t6, true)
	if (*tt2).Size() != 2 || *(*tt2).First() != t4 ||  *(*tt2).Last() != t6 {
		t.Error("SubSet operation fail!")
	}
	tt3 := smks.SubSet(&t2, false, &t6, true)
	if tt3 != nil {
		t.Error("SubSet operation fail!")
	}
}

func TestKeySet_TailSet(t *testing.T) {
	RestartKeySet()
	tt1 := mks.TailSet(&t4, true)
	if (*tt1).Size() != 3 || *(*tt1).First() != t4 ||  *(*tt1).Last() != t8 {
		t.Error("SubSet operation fail!")
	}
	tt2 := smks.TailSet(nil, false)
	if (*tt2).Size() != 3 || *(*tt2).First() != t4 ||  *(*tt2).Last() != t8 {
		t.Error("SubSet operation fail!")
	}
}