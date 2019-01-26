package config

// IwdParameter struct to holds static parameter of iwd
type IwdParameter struct {
	MaximumIteration int     `json:"maximum_iteration"`
	PopulationSize   int     `json:"population_size"`
	As               float64 `json:"as"`
	Bs               float64 `json:"bs"`
	Cs               float64 `json:"cs"`
	Av               float64 `json:"av"`
	Bv               float64 `json:"bv`
	Cv               float64 `json:"cv"`
	InitSoil         float64 `json:"init_soil"`
	InitIWDVel       float64 `json:"init_iwd_vel"`
	InitIWDSoil      float64 `json:"init_iwd_soil"`
	P                float64 `json:"p"`
	MinDSoil         float64 `json:"min_d_soil"`
	MaxDSoil         float64 `json:"max_d_soil"`
}
