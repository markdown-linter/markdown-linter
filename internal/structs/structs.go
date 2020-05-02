package structs

type PluginInfo struct {
	Name             string
	Contributors     []string
	ErrorDescription string
}

type Offence struct {
	Line        int
	Description string
}

type LinterResult struct {
	FileName         string
	Line             int
	Plugin           string
	ErrorDescription string
}
