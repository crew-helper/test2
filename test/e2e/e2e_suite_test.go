package e2e_test

import (
	"os"
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const(
	//TODO data provider?
	ConfigAll = "../../deploy/all-in-one.yaml" 	// basic configuration (release)
	ProjectSample = "data/atlasproject.yaml"
	ClusterSample = "data/atlascluster_basic.yaml"
)
var(
	//default
	Platform = "kind"
	K8sVersion = "v1.17.17"
)

func TestE2e(t *testing.T) {
	setUpMongoCLI()
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}

//setUpMongoCLI initial setup
func setUpMongoCLI() {
	Platform = os.Getenv("K8s_PLATFORM")
	K8sVersion = os.Getenv("K8s_VERSION")
}
