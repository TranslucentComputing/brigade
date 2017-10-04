package commands

import (
	"github.com/spf13/cobra"
	"os"
)

const mainUsage = `Interact with the Acid cluster service.

Acid is a tool for scripting cluster workflows.

The most common use for thi tool is to send an Acid JavaScript file to the
cluster for execution. This is done with the 'acid run' command.

	$ acid run -f my.js my-project

But the 'acid' command can also be used for learning about projects and
capabilities of a cluster.
`

var (
	globalNamespace  = ""
	globalVerbose    = false
	globalKubeConfig = ""
)

func init() {
	f := Root.PersistentFlags()
	f.StringVarP(&globalNamespace, "namespace", "n", "default", "The Kubernetes namespace for Acid")
	f.StringVar(&globalKubeConfig, "kubeconfig", "", "The path to a KUBECONFIG file, overrides $KUBECONFIG.")
	f.BoolVarP(&globalVerbose, "verbose", "v", false, "Turn on verbose output")
}

var Root = &cobra.Command{
	Use:   "acid",
	Short: "The Acid client",
	Long:  mainUsage,
}

func kubeConfigPath() string {
	if globalKubeConfig != "" {
		return globalKubeConfig
	}
	if v, ok := os.LookupEnv(kubeConfig); ok {
		return v
	}
	return os.ExpandEnv("$HOME/.kube/config")
}
