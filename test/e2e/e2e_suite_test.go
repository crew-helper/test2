package e2e_test

import (
	"os"
	// "os/exec"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "github.com/onsi/gomega/gexec"
	// . "github.com/onsi/gomega/gbytes"

	"testing"
)

const(
	ConfigAll = "../../deploy/all-in-one.yaml" 	// basic configuration (release)
	ProjectSample = "data/atlasproject.yaml"
	ClusterSample = "data/atlascluster_basic.yaml"
	//TODO const from CI/CD matrix

)
var(
	Platform = "kind"
	K8sVersion = "1.17"
	TestName = "" //TODO not sure about it
)

func TestE2e(t *testing.T) {
	setUpMongoCLI()
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}

// not sure
func setUpMongoCLI() {
	// TODO
	Platform = os.Getenv("Platform")
	K8sVersion = os.Getenv("")"1.17"

	os.Setenv("MCLI_PUBLIC_API_KEY", "DJPQJCNN")
	os.Setenv("MCLI_PRIVATE_API_KEY", "864c88bc-531a-46af-bebd-f3d6637bab37")
	os.Setenv("MCLI_ORG_ID", "5ffdac0657666b4b84836fd4")
	os.Setenv("MCLI_OPS_MANAGER_URL", "https://cloud-qa.mongodb.com/")


	// TODO: change secrets
	os.Setenv("MCLI_PUBLIC_API_KEY", "DJPQJCNN")
	os.Setenv("MCLI_PRIVATE_API_KEY", "864c88bc-531a-46af-bebd-f3d6637bab37")
	os.Setenv("MCLI_ORG_ID", "5ffdac0657666b4b84836fd4")
	os.Setenv("MCLI_OPS_MANAGER_URL", "https://cloud-qa.mongodb.com/")
}
