package e2e_test

import (
	"fmt"
	"time"

	"os"
	"math/rand"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "github.com/pborman/uuid"

	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	// "github.com/onsi/ginkgo/config"
	cli "github.com/mongodb/mongodb-atlas-kubernetes/test/e2e/utils"
	// "go.mongodb.org/atlas/mongodbatlas"
)


var _ = Describe("Deploy simple cluster", func() {

	It("Release sample all-in-one.yaml should work", func() {
		By("Prepare namespaces")
		// namespaceUserResources := uuid.NewRandom().String() //TODO for another tests
		namespaceUserResources := "testdata"
		// namespaceUserResources := "mongodb-atlas-kubernetes-system"
		// namespaceOperator := "mongodb-atlas-kubernetes-system"
		session := cli.Execute("kubectl", "create", "namespace", namespaceUserResources)
		Expect(session).ShouldNot(Say("created"))
		userProjectConfig := cli.LoadUserProjectConfig("data/atlasproject.yaml")
		userClusterConfig := cli.LoadUserClusterConfig("data/atlascluster_basic.yaml")

		By("Check Kubernetes version\n")
		session = cli.Execute("kubectl", "version")
		Eventually(session).Should(Say(K8sVersion))

		By("Apply All-in-one configuration\n in ")
		session = cli.Execute("kubectl", "apply", "-f", ConfigAll)
		Eventually(session.Wait()).Should(Say("customresourcedefinition.apiextensions.k8s.io/atlasclusters.atlas.mongodb.com"))

		By("Create secret")
		session = cli.Execute("kubectl", "create", "secret", "generic", "my-atlas-key",
			"--from-literal=orgId=" + os.Getenv("MCLI_ORG_ID"),
			"--from-literal=publicApiKey=" + os.Getenv("MCLI_PUBLIC_API_KEY"),
			"--from-literal=privateApiKey=" + os.Getenv("MCLI_PRIVATE_API_KEY"),
			"-n", namespaceUserResources,
		)
		// Eventually(session).Should(Say("my-atlas-key created"))
		// Eventually(session).Should(Say("my-atlas-key"))

		By("Create Sample Project\n")
		session = cli.Execute("kubectl", "apply", "-f", ProjectSample, "-n", namespaceUserResources)
		// Eventually(session).Should(Say("my-project created"))
		Eventually(session).Should(Say("atlasproject.atlas.mongodb.com/my-project"))

		By("Sample Cluster\n")
		session = cli.Execute("kubectl", "apply", "-f", ClusterSample, "-n", namespaceUserResources)
		// session = cli.Execute("kubectl", "apply", "-f", ClusterSample)
		// Eventually(session).Should(Say("atlascluster-sample created"))
		Eventually(session).Should(Say("atlascluster-sample"))

		By("Wait creating and check that it was created")
		session = cli.Execute("mongocli", "--version")
		Eventually(session).Should(gexec.Exit(0)) //TODO exit status
		Expect(os.Getenv("MCLI_OPS_MANAGER_URL")).Should(Equal("https://cloud-qa.mongodb.com/")) //TODO remove

		Eventually(cli.GetProjectID(userProjectConfig.Spec.Name)).ShouldNot(BeNil())
		projectID := cli.GetProjectID(userProjectConfig.Spec.Name)
		GinkgoWriter.Write([]byte("projectID = " + projectID))

		Eventually(
			cli.GetClusterStatus(projectID, userClusterConfig.Spec.Name),
			"35m", "1m",
		).Should(Equal("IDLE"))

		// cli.WaitCluster(
		// 	projectID,
		// 	userClusterConfig.Spec.Name,
		// 	"IDLE",
		// ) //TODO UPDATING?

		By("check cluster Attribute") //TODO ...
		cluster := cli.GetClustersInfo(projectID, userClusterConfig.Spec.Name)
		Expect(
			cluster.ProviderSettings.InstanceSizeName,
		).Should(Equal(userClusterConfig.Spec.ProviderSettings.InstanceSizeName))
		Expect(
			cluster.ProviderSettings.ProviderName,
		).Should(Equal(userClusterConfig.Spec.ProviderSettings.ProviderName))
		Expect(
			cluster.ProviderSettings.RegionName,
		).Should(Equal(userClusterConfig.Spec.ProviderSettings.RegionName))

		By("Update cluster\n")
		session = cli.Execute("kubectl", "apply", "-f", "data/updated_atlascluster_basic.yaml", "-n", namespaceUserResources) //TODO param
		// Eventually(session).Should(Say("atlascluster-sample configured"))
		Eventually(session).Should(Say("atlascluster-sample"))

		By("Wait creation")
		userClusterConfig = cli.LoadUserClusterConfig("data/updated_atlascluster_basic.yaml")
		Expect(projectID).ShouldNot(BeNil())
		Eventually(
			cli.GetClusterStatus(projectID, userClusterConfig.Spec.Name),
			"35m", "1m",
		).Should(Equal("IDLE"))

		uCluster := cli.GetClustersInfo(projectID, userClusterConfig.Spec.Name)
		Expect(
			uCluster.ProviderSettings.InstanceSizeName,
		).Should(Equal(
			userClusterConfig.Spec.ProviderSettings.InstanceSizeName,
		))

		By("Delete cluster")
		session = cli.Execute("kubectl", "delete", "-f", "data/updated_atlascluster_basic.yaml", "-n", namespaceUserResources)
		Eventually(session.Wait("7m")).Should(gexec.Exit(0))
		Eventually(
			cli.IsClusterExist(projectID, userClusterConfig.Spec.Name),
			"10m", "1m",
		).Should(BeFalse())

		// By("Delete project") //TODO
	})
})

func t44() func() int {
	return func() int {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(5)
		fmt.Println(i)
		return i
	}

}
