package lib

type Args struct {
	Show   *ShowArgs   `arg:"subcommand"`
	Config *ConfigArgs `arg:"subcommand"`
}
