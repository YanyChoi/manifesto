package flag

import "flag"

type Flags struct {
	Template string
	Output   string
	Version  string
}

func NewFlags() *Flags {
	flags := &Flags{}

	flag.StringVar(&flags.Template, "template", "template.yaml", "path to the template file")
	flag.StringVar(&flags.Output, "output", "deployment.yaml", "path to the output file")
	flag.StringVar(&flags.Version, "version", "v1", "version of the template")
	flag.Parse()
	return flags
}
