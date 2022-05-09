package command

import (
	"strings"

	flag "github.com/spf13/pflag"
	"github.com/zirain/ubrain/pkg/generic"
	"github.com/zirain/ubrain/pkg/plugin/karmada"
)

type KarmadaInstallCommand struct {
	*generic.Options
}

func (c *KarmadaInstallCommand) Run(args []string) int {
	cmdFlags := c.karmadaFlagSet(c.Options)
	cmdFlags.Usage = func() {
		c.Errorf(c.Help())
		c.Errorf(cmdFlags.FlagUsages())
	}
	if err := cmdFlags.Parse(args); err != nil {
		c.Errorf("Error parsing command-line flags: %s\n", err.Error())
		return 1
	}

	plugin, err := karmada.NewKarmadaPlugin(c.Options)
	if err != nil {
		c.Errorf("karmada init error: %v", err)
		return 1
	}

	c.Infof("start install karmada: %+v ", c.Options)
	if err := plugin.Execute(args, nil); err != nil {
		c.Infof("istio execute error: %v", err)
		return 1
	}

	return 0
}

func (c *KarmadaInstallCommand) Help() string {
	helpText := `
Usage: ubrain install karmada [Options]
`
	return strings.TrimSpace(helpText)
}

func (c *KarmadaInstallCommand) Synopsis() string {
	return "Install karmada component"
}

func (c *KarmadaInstallCommand) karmadaFlagSet(s *generic.Options) *flag.FlagSet {
	f := c.FlagSet("karmada")
	s.AddFlags(f)
	return f
}