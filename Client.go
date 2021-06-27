package main

type (
	UberFormat struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	RequestFormat struct {
		Xi int `json:"xi"`
		Yi int `json:"yi"`
		Xf int `json:"xf"`
		Yf int `json:"yf"`
		T  int `json:"t"`
	}

	Configurations struct {
		RunType byte            `json:"run_type"` // 0 Morning, 1 Afternoon, 2 Night, 3 Random, 4 Free
		Pram    bool            `json:"pram"`     // Pram or Not
		Ubers   []UberFormat    `json:"ubers"`    // If 4 is enabled
		Request []RequestFormat `json:"request"`
	}

	Client struct {
		Id     int            `json:"id"`
		World  *world         `json:"world"`
		Config Configurations `json:"config"`
	}
)
