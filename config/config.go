package config

// Config for the solver, to be added : iwd-sa constants
type Config struct {
	DeliveryDate      string       `json:"delivery_date"`
	MaxDriverCapacity int          `json:"max_driver_capacity"`
	MaxDriverDistance float64      `json:"max_driver_distance"`
	MaxTreeEntry      int          `json:"max_tree_entry"`
	IwdParameter      IwdParameter `json:"iwd_parameter"`
}

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
	InitVel          float64 `json:"init_vel"`
}
