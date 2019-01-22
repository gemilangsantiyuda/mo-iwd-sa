package config

// Config for the solver, to be added : iwd-sa constants
type Config struct {
	DeliveryDate      string       `json:"delivery_date"`
	MaxDriverCapacity int          `json:"max_driver_capacity"`
	MaxDriverDistance float64      `json:"max_driver_distance"`
	MaxTreeEntry      int          `json:"max_tree_entry"`
	IwdParameter      IwdParameter `json:"iwd_parameter"`
	Weight            Weight       `json:"weight"	`
}
