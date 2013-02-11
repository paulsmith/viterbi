package viterbi

import "fmt"

type State string
type Observation string

type TransKey struct {
	From, To State
}

type EmitKey struct {
	State State
	Obs   Observation
}

type Viterbi struct {
	States  []State
	Obs     []Observation
	StartPr map[State]float64
	TransPr map[TransKey]float64
	EmitPr  map[EmitKey]float64
}

type ViterbiPath struct {
	Pr   float64
	Path []State
}

func (v Viterbi) FindPath() ViterbiPath {
	var (
		V             []map[State]float64
		path, newpath map[State][]State
		y, state      State
		maxPr, pr     float64
	)

	V = append(V, make(map[State]float64))
	path = make(map[State][]State)

	for _, y = range v.States {
		V[0][y] = v.StartPr[y] * v.EmitPr[EmitKey{y, v.Obs[0]}]
		path[y] = append(path[y], y)
	}

	for t := 1; t < len(v.Obs); t++ {
		V = append(V, make(map[State]float64))
		newpath = make(map[State][]State)

		for _, y = range v.States {
			maxPr = 0.0
			for _, y0 := range v.States {
				pr = V[t-1][y0] * v.TransPr[TransKey{y0, y}] * v.EmitPr[EmitKey{y, v.Obs[t]}]
				if pr > maxPr {
					maxPr = pr
					state = y0
				}
			}
			V[t][y] = maxPr
			newpath[y] = append(path[state], y)
		}

		path = newpath
	}

	maxPr = 0.0
	for _, y = range v.States {
		pr = V[len(v.Obs)-1][y]
		if pr > maxPr {
			maxPr = pr
			state = y
		}
	}
	return ViterbiPath{maxPr, path[state]}
}

func printDebugTable(V []map[State]float64) {
	fmt.Printf("    ")
	for i := 0; i < len(V); i++ {
		fmt.Printf("%7d ", i)
	}
	fmt.Println("")

	for y, _ := range V[0] {
		fmt.Printf("%.5s:  ", y)
		for t := 0; t < len(V); t++ {
			fmt.Printf("%.5f ", V[t][y])
		}
		fmt.Println("")
	}
}
