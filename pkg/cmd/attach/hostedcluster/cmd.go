// Copyright Contributors to the Open Cluster Management project
package hostedcluster

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/stolostron/cm-cli/pkg/cmd/attach/cluster/scenario"
	genericclioptionscm "github.com/stolostron/cm-cli/pkg/genericclioptions"
	"github.com/stolostron/cm-cli/pkg/helpers"
	clusteradmhelpers "open-cluster-management.io/clusteradm/pkg/helpers"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var example = `
# Attach a cluster
%[1]s attach hc --values values.yaml
`

const (
	scenarioDirectory = "attach"
)

var valuesTemplatePath = filepath.Join(scenarioDirectory, "values-template.yaml")
var valuesDefaultPath = filepath.Join(scenarioDirectory, "values-default.yaml")

// NewCmd provides a cobra command wrapping NewCmdImportCluster
func NewCmd(cmFlags *genericclioptionscm.CMFlags, streams genericclioptions.IOStreams) *cobra.Command {
	o := newOptions(cmFlags, streams)

	cmd := &cobra.Command{
		Use:          "hostedcluster",
		Aliases:      []string{"hc", "hostedclusters", "ccs"},
		Short:        "Import a hostedcluster",
		Example:      fmt.Sprintf(example, helpers.GetExampleHeader()),
		SilenceUsage: true,
		PreRunE: func(c *cobra.Command, args []string) error {
			clusteradmhelpers.DryRunMessage(cmFlags.DryRun)
			isHypershift, err := helpers.IsHypershift(o.CMFlags)
			if err != nil {
				return err
			}
			isSupported, err := helpers.IsSupported(o.CMFlags)
			if err != nil {
				return err
			}
			if !isSupported || !isHypershift {
				return fmt.Errorf("this command '%s %s' is only available on %s or %s on %s",
					helpers.GetExampleHeader(),
					strings.Join(os.Args[1:], " "),
					helpers.RHACM,
					helpers.MCE,
					helpers.HYPERSHIFT)
			}

			return nil
			// return clusterpoolhost.BackupCurrentContexts()
		},
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.complete(c, args); err != nil {
				return err
			}
			if err := o.validate(); err != nil {
				return err
			}
			if err := o.run(); err != nil {
				return err
			}

			return nil
		},
		// PostRunE: func(cmd *cobra.Command, args []string) error {
		// 	return clusterpoolhost.RestoreCurrentContexts()
		// },
	}

	cmd.SetUsageTemplate(clusteradmhelpers.UsageTempate(cmd, scenario.GetScenarioResourcesReader(), valuesTemplatePath))
	cmd.Flags().StringVar(&o.valuesPath, "values", "", "The files containing the values")
	cmd.Flags().StringVar(&o.outputFile, "output-file", "", "The generated resources will be copied in the specified file")
	cmd.Flags().BoolVar(&o.waitAgent, "wait", false, "Wait until the klusterlet agent is installed")
	//Not implemented as it requires to import all addon packages
	// cmd.Flags().BoolVar(&o.waitAddOns, "wait-addons", false, "Wait until the klusterlet agent and the addons are is installed")
	cmd.Flags().IntVar(&o.timeout, "timeout", 180, "Timeout to get the klusterlet agent or addons ready in seconds")

	return cmd
}
