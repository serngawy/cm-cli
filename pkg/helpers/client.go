// Copyright Contributors to the Open Cluster Management project

package helpers

import "k8s.io/apimachinery/pkg/runtime/schema"

const (
	QPS   = 200
	Burst = 200
)

var (
	GvrCC                       schema.GroupVersionResource = schema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "clusterclaims"}
	GvrCP                       schema.GroupVersionResource = schema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "clusterpools"}
	GvrCD                       schema.GroupVersionResource = schema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "clusterdeployments"}
	GvrCIS                      schema.GroupVersionResource = schema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "clusterimagesets"}
	GvrPol                      schema.GroupVersionResource = schema.GroupVersionResource{Group: "policy.open-cluster-management.io", Version: "v1", Resource: "policies"}
	GvrMCH                      schema.GroupVersionResource = schema.GroupVersionResource{Group: "operator.open-cluster-management.io", Version: "v1", Resource: "multiclusterhubs"}
	GvrMCE                      schema.GroupVersionResource = schema.GroupVersionResource{Group: "multicluster.openshift.io", Version: "v1alpha1", Resource: "multiclusterengines"}
	GvrMC                       schema.GroupVersionResource = schema.GroupVersionResource{Group: "cluster.open-cluster-management.io", Version: "v1", Resource: "managedclusters"}
	GvrCSV                      schema.GroupVersionResource = schema.GroupVersionResource{Group: "operators.coreos.com", Version: "v1alpha1", Resource: "clusterserviceversions"}
	GvrPLC                      schema.GroupVersionResource = schema.GroupVersionResource{Group: "cluster.open-cluster-management.io", Version: "v1alpha1", Resource: "placements"}
	GvrMCS                      schema.GroupVersionResource = schema.GroupVersionResource{Group: "cluster.open-cluster-management.io", Version: "v1beta1", Resource: "managedclustersets"}
	GvrMCSB                     schema.GroupVersionResource = schema.GroupVersionResource{Group: "cluster.open-cluster-management.io", Version: "v1beta1", Resource: "managedclustersetbindings"}
	GvrIDPConfig                schema.GroupVersionResource = schema.GroupVersionResource{Group: "identityconfig.identitatem.io", Version: "v1alpha1", Resource: "idpconfigs"}
	GvrOpenshiftClusterVersions schema.GroupVersionResource = schema.GroupVersionResource{Group: "config.openshift.io", Version: "v1", Resource: "clusterversions"}
	GvrHC                       schema.GroupVersionResource = schema.GroupVersionResource{Group: "hypershift.openshift.io", Version: "v1alpha1", Resource: "hostedclusters"}
)
