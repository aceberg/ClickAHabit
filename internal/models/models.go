package models

// Conf - web gui config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Color    string
	DBPath   string
	DirPath  string
	ConfPath string
	NodePath string
}

// Item - element of Plan
type Item struct {
	Name  string `yaml:"name"`
	Color string `yaml:"color"`
}

// Plan - check from plan.yaml
type Plan struct {
	Group string `yaml:"group"`
	Items []Item `yaml:"items"`
}

// Check - check for DB
type Check struct {
	ID    int    `db:"ID"`
	Date  string `db:"DATE"`
	Name  string `db:"NAME"`
	Group string `db:"GR"`
	Color string `db:"COLOR"`
	Count int    `db:"COUNT"`
}

// GuiData - web gui data
type GuiData struct {
	Config  Conf
	Themes  []string
	Version string
	Checks  []Check
}
