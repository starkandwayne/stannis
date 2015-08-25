package agent_test

import (
	. "github.com/cloudfoundry-community/stannis/agent"
	"github.com/cloudfoundry-community/stannis/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Prepare data for templates", func() {
	var (
		agentConfig *config.AgentConfig
		subject     Agent
	)
	BeforeEach(func() {
		var err error
		agentConfig, err = config.LoadAgentConfigFromYAMLFile("../config/agent.config.example.yml")
		Expect(err).NotTo(HaveOccurred())

		subject = Agent{
			Config: agentConfig,
		}
	})

	It("has max bulk upload", func() {
		Expect(agentConfig.MaxBulkUploadSize).To(Equal(5))
	})
})