package viterbi

import "testing"

func IsViterbiEqual(a, b ViterbiPath) bool {
	if a.Pr != b.Pr || len(a.Path) != len(b.Path) {
		return false
	}
	for i, _ := range a.Path {
		if a.Path[i] != b.Path[i] {
			return false
		}
	}
	return true
}

func TestFindPath(t *testing.T) {
	obs := []Observation{"normal", "cold", "dizzy"}
	v := Viterbi{
		States:  []State{"Healthy", "Fever"},
		Obs:     obs,
		StartPr: map[State]float64{"Healthy": 0.6, "Fever": 0.4},
		TransPr: map[TransKey]float64{
			TransKey{"Healthy", "Healthy"}: 0.7,
			TransKey{"Healthy", "Fever"}:   0.3,
			TransKey{"Fever", "Healthy"}:   0.4,
			TransKey{"Fever", "Fever"}:     0.6,
		},
		EmitPr: map[EmitKey]float64{
			EmitKey{"Healthy", "normal"}: 0.5,
			EmitKey{"Healthy", "cold"}:   0.4,
			EmitKey{"Healthy", "dizzy"}:  0.1,
			EmitKey{"Fever", "normal"}:   0.1,
			EmitKey{"Fever", "cold"}:     0.3,
			EmitKey{"Fever", "dizzy"}:    0.6,
		},
	}
	expected := ViterbiPath{0.01512, []State{"Healthy", "Healthy", "Fever"}}
	if actual := v.FindPath(); !IsViterbiEqual(actual, expected) {
		t.Errorf("FindPath(): %v, got %v", expected, actual)
	}
}
