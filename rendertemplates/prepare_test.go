package rendertemplates_test

import (
	"github.com/cloudfoundry-community/stannis/config"
	"github.com/cloudfoundry-community/stannis/data"
	. "github.com/cloudfoundry-community/stannis/rendertemplates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Prepare data for templates", func() {
	var (
		pipelineConfig     *config.PipelinesConfig
		expectedRenderData RenderData
		db                 data.DeploymentsPerBOSH
		renderdata         *RenderData
	)
	BeforeEach(func() {
		expectedRenderData = *TestScenarioData()
		db = data.NewDeploymentsPerBOSH()

		db.FixtureBosh("../upload/fixtures/bosh-lite.json")
		db.FixtureDeployment("../upload/fixtures/deployment-bosh-lite-cf1.json")
		db.FixtureDeployment("../upload/fixtures/deployment-bosh-lite-cf2.json")

		db.FixtureBosh("../upload/fixtures/bosh-vsphere-sandbox.json")
		db.FixtureDeployment("../upload/fixtures/deployment-vsphere-sandbox-cf.json")

		db.FixtureBosh("../upload/fixtures/bosh-aws-production.json")
		db.FixtureDeployment("../upload/fixtures/deployment-aws-production-cf.json")

		var err error
		pipelineConfig, err = config.LoadConfigFromYAMLFile("../config/webserver.config.example.yml")
		Expect(err).NotTo(HaveOccurred())

		renderdata = PrepareRenderData(pipelineConfig, db, "")
	})

	It("has tiers", func() {
		Expect(len(renderdata.Tiers)).To(Equal(len(expectedRenderData.Tiers)))
	})

	It("has slots in tier", func() {
		for tierIndex := range renderdata.Tiers {
			renderTier := renderdata.Tiers[tierIndex]
			Expect(renderTier).ToNot(BeNil())
			expectedTier := *expectedRenderData.Tiers[tierIndex]
			Expect(expectedTier).ToNot(BeNil())

			Expect(renderTier.Name).To(Equal(expectedTier.Name))
			Expect(len(renderTier.Slots)).To(Equal(len(expectedTier.Slots)))

			for slotIndex := range renderTier.Slots {
				renderSlot := renderTier.Slots[slotIndex]
				Expect(renderSlot).ToNot(BeNil())
				expectedSlot := expectedTier.Slots[slotIndex]
				Expect(expectedSlot).ToNot(BeNil())

				// fmt.Println(renderTier.Name, slotIndex)
				Expect(len(renderSlot.Deployments)).To(Equal(len(expectedSlot.Deployments)))

				for deploymentIndex := range renderSlot.Deployments {
					renderDeployment := renderSlot.Deployments[deploymentIndex]
					Expect(renderDeployment).ToNot(BeNil())
					expectedDeployment := expectedSlot.Deployments[deploymentIndex]
					Expect(expectedDeployment).ToNot(BeNil())

					Expect(renderDeployment.Name).To(Equal(expectedDeployment.Name))

					Expect(len(renderDeployment.Releases)).To(Equal(len(expectedDeployment.Releases)))
					Expect(len(renderDeployment.Stemcells)).To(Equal(len(expectedDeployment.Stemcells)))
				}
			}
		}
	})
})
