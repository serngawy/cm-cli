// Copyright Contributors to the Open Cluster Management project
package clusterclaim

import (
	genericclioptionscm "github.com/stolostron/cm-cli/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type Options struct {
	//CMFlags: The generic optiosn from the cm cli-runtime.
	CMFlags *genericclioptionscm.CMFlags
	//The list of cluster claim name to create (comma-separated)
	ClusterClaims   string
	ClusterPoolHost string
	SkipSchedule    bool
	//The file to output the resources will be sent to the file.
	outputFile string
}

func newOptions(cmFlags *genericclioptionscm.CMFlags, streams genericclioptions.IOStreams) *Options {
	return &Options{
		CMFlags: cmFlags,
	}
}
