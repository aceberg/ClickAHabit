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
	BtnWidth string
}

// Plan - check from plan.yaml
type Plan struct {
	ID      int    `yaml:"-"`
	Group   string `yaml:"group"`
	Name    string `yaml:"name"`
	Color   string `yaml:"color,omitempty"`
	NoColor bool   `yaml:"nocolor,omitempty"`
	Icon    string `yaml:"icon,omitempty"`
	Place   int    `yaml:"place,omitempty"`
	Link    string `yaml:"link,omitempty"`
	Pause   bool   `yaml:"pause,omitempty"`
	Weekly  bool   `yaml:"weekly,omitempty"`
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
	Link  string `db:"LINK"`
}

// Stat - statistic
type Stat struct {
	Name   string
	Group  string
	DTotal int
	CTotal int
	Checks []Check
}

// GuiData - web gui data
type GuiData struct {
	Config  Conf
	Themes  []string
	Version string
	OnePlan Plan
	Plans   []Plan
	Checks  []Check
	Stats   map[string]Stat
}
