package e2e_test

import (
	// "fmt"
	// "context"
	"fmt"
	// "reflect"
	// "os"
	"os/exec"

	// "runtime"
	// "time"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// . "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	// "github.com/onsi/ginkgo/config"
	// corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/errors"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/tools/clientcmd"
)


var _ = Describe("Deploy simple cluster", func() {

	It("Should deploy and update simple cluster", func() {
		// namespace := "mongodb-atlas-kubernetes-system"//TODO autogen

		// By("Check Kubernetes version\n")
		// session := execute("kubectl", "version")
		// version := genKubeVersion(k8sVersion)
		// Eventually(session).Should(Say(version))

		// By("Apply All-in-one configuration\n") //TODO why? move to beforeEach? do we need to do that?
		// session = execute("kubectl", "apply", "-f", ConfigAll, "-n", namespace)
		// Eventually(session).Should(Say("customresourcedefinition.apiextensions.k8s.io/atlasclusters.atlas.mongodb.com"))

		// By("Create secret")
		// session = execute("kubectl", "create", "secret", "generic", "my-atlas-key",
		// 	"--from-literal=orgId=" + os.Getenv("MCLI_ORG_ID"),
		// 	"--from-literal=publicApiKey=" + os.Getenv("MCLI_PUBLIC_API_KEY"),
		// 	"--from-literal=privateApiKey=" + os.Getenv("MCLI_PRIVATE_API_KEY"),
		// 	"-n", namespace)
		// Eventually(session).Should(Say("my-atlas-key created"))
		// // Eventually(session).Should(Say("my-atlas-key"))

		// By("Sample Project\n")
		// session = execute("kubectl", "apply", "-f", ProjectSample, "-n", namespace)
		// // Eventually(session).Should(Say("my-project created"))
		// Eventually(session).Should(Say("atlasproject.atlas.mongodb.com/my-project"))

		// By("Sample Cluster\n")
		// session = execute("kubectl", "apply", "-f", ClusterSample, "-n", namespace)
		// // Eventually(session).Should(Say("atlascluster-sample created"))
		// Eventually(session).Should(Say("atlascluster-sample"))

		// By("Wait creating and check that it was created")
		// session := execute("mongocli", "--version")
		// Expect(session).ShouldNot(BeNil())
		// Expect(os.Getenv("MCLI_OPS_MANAGER_URL")).Should(Equal("https://cloud-qa.mongodb.com/")) //TODO remove

		getProjectID := func() []byte{
			session := execute("mongocli", "iam", "projects", "list", "-o", "go-template=\"{{ range .Results }}{{ if eq .Name \"Test Atlas Operator Project\"}}{{ .ID }}{{end}}{{end}}\"")
			projectID := session.Out.Contents()
			return projectID
		}

		Eventually(getProjectID()).ShouldNot(BeNil())
		// projectID = getProjectID()
		// Eventually(session).Should(Say("60082b3e31fbe32df4d4470d")) //TODO param

		By("check cluster name")
		getClusterName := func() string {
			session := execute("mongocli", "atlas", "clusters", "list", "--projectId", "60082b3e31fbe32df4d4470d", "-o", "go-template=\"{{ range . }}{{ .Name }} {{ end }}\"")
			name := strings.TrimRight(string(session.Out.Contents()), " ")
			GinkgoWriter([]byte(name))
			return name
		}
		Eventually(getClusterName()).Should(BeEquivalentTo("cross44 "))

		// By("check provider settings")
		// 	name: "cluster44"
		//   instanceSizeName: M10 providerSettings.instanceSizeName
		//   providerName: AWS providerSettings.providerName
		//   regionName: US_EAST_1 providerSettings.regionName


		// projectID := "5fff1401869e8f54b7d5fd6c"
		// command = exec.Command("mongocli", "iam", "projects", "ls")
		// session, _ = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		// GinkgoWriter.Write(session.Buffer().Contents())
		// Eventually(session).Should(Say("Test Atlas Operator Project")) //TODO param
		// // Eventually(session)


		// By("Update cluster\n")
		// session = execute("kubectl", "apply", "-f", "data/updated_atlascluster_basic.yaml", "-n", namespace) //TODO param
		// Eventually(session).Should(Say("customresourcedefinition.apiextensions.k8s.io/atlasclusters.atlas.mongodb.com"))

		// By("Check updated field")
		// // session = execute("kubectl", )

		// By("Delete")
		// session = execute("kubectl", "delete", "projects", "list", "-o", "go-template=\"{{ range .Results }}{{ if eq .Name \"Test Atlas Operator Project\"}}{{ .ID }}{{end}}{{end}}\"")
		// projectID := session.Out.Contents()
		// Eventually(session).Should(Say("60082b3e31fbe32df4d4470d")) //TODO param
	})
})

func genKubeVersion(fullVersion string) string {
	version := strings.Split(fullVersion, ".")
	return fmt.Sprintf("Major:\"%s\", Minor:\"%s\"", version[0], version[1])
}

func execute(command string, args ...string) *gexec.Session {
	cmd := exec.Command(command, args...)
	session, _ := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	return session
}
