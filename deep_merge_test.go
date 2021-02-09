package yaml

import (
	"testing"
)

// TestDeepMerge tests that we can successfully merge to pointers of structs and maps (which required a patch that Lantern made in decode.go)
func TestDeepMerge(t *testing.T) {
	type Child struct {
		A int
		B int
	}

	type Parent struct {
		C map[int]*Child
		M map[int]*map[string]interface{}
		V int
	}

	data := []byte(`
v: 5
c:
  1:
    b: 6
m:
  2:
    z: 7
`)

	emptyParent := &Parent{}
	err := Unmarshal(data, emptyParent)
	if err != nil {
		t.Fatal(err)
	}
	if emptyParent.V != 5 {
		t.Fatal("Wrong V")
	}
	child1 := emptyParent.C[1]
	if child1.B != 6 {
		t.Fatal("Wrong child1.b")
	}
	child2 := *(emptyParent.M[2])
	if child2["z"] != 7 {
		t.Fatal("Wrong child2.z")
	}

	filledParent := &Parent{
		C: map[int]*Child{
			1: &Child{
				A: 1,
			},
		},
		M: map[int]*map[string]interface{}{
			2: &map[string]interface{}{
				"y": 2,
			},
		},
	}
	err = Unmarshal(data, filledParent)
	if err != nil {
		t.Fatal(err)
	}
	if filledParent.V != 5 {
		t.Fatal("Wrong V")
	}
	child1 = filledParent.C[1]
	if child1.B != 6 {
		t.Fatal("Wrong child1.b")
	}
	if child1.A != 1 {
		t.Fatal("Wrong child1.a")
	}
	child2 = *(filledParent.M[2])
	if child2["z"] != 7 {
		t.Fatal("Wrong child2.z")
	}
	if child2["y"] != 2 {
		t.Fatal("Wrong child2.y")
	}
}
