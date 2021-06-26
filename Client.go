package main

type Configurations struct {
	RunType byte   `json:"run_type"` // 0 Morning, 1 Afternoon, 2 Night, 3 Random, 4 Free
	Pram    bool   `json:"pram"`     // Pram or Not
	Ubers   string `json:"ubers"`    // If 4 is enabled
	Request string `json:"request"`
}

type Client struct {
	Id     int            `json:"id"`
	World  *world         `json:"world"`
	Config Configurations `json:"config"`
}
