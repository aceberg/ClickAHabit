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

// Plan - check from plan.yaml
type Plan struct {
	Group string `yaml:"group"`
	Name  string `yaml:"name"`
	Color string `yaml:"color"`
	Icon  string `yaml:"icon"`
	Place int    `yaml:"place"`
}

// Check - check for DB
type Check struct {
	ID    int    `db:"ID"`
	Date  string `db:"DATE"`
	Name  string `db:"NAME"`
	Group string `db:"GR"`
	Color string `db:"COLOR"`
	Icon  string `db:"ICON"`
	Place int    `db:"PLACE"`
	Count int    `db:"COUNT"`
}

// GuiData - web gui data
type GuiData struct {
	Config  Conf
	Themes  []string
	Version string
	Checks  []Check
}
