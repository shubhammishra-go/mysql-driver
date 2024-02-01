package model

type Details struct {
	Name        string `json: "name"`
	Roll_number string `json: "rollnumber"`
	University  string `json: "university"`
	Course      string `json: "course"`
	Branch      string `json: "branch"`
	Start_Year  int64  `json: "start"`
	End_Year    int64  `json: "end"`
}
