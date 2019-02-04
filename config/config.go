package config

// Config for the solver, to be added : iwd-sa constants
type Config struct {
	DataSize          string       `json:"data_size"`
	MaxDriverCapacity int          `json:"max_driver_capacity"`
	MaxDriverDistance float64      `json:"max_driver_distance"`
	MaxTreeEntry      int          `json:"max_tree_entry"`
	DriverSpeed       float64      `json:"driver_speed"`
	DriverRate        int          `json:"driver_rate"`
	NeighbourCount    int          `json:"neighbour_count"`
	IwdParameter      IwdParameter `json:"iwd_parameter"`
	Weight            Weight       `json:"weight"`
	SaParam           SaParameter  `json:"sa_parameter"`
	Tolerance         Tolerance    `json:"tolerance"`
	ArchiveSize       int          `json:"archive_size"`
}
