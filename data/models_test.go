package data_test

import (
	. "github.com/cloudfoundry-community/stannis/data"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Data", func() {
	var (
		db DeploymentsPerBOSH
	)
	BeforeEach(func() {
		db = NewDeploymentsPerBOSH()
		db.FixtureBosh("../upload/fixtures/bosh-lite.json")
		db.FixtureDeployment("../upload/fixtures/deployment-bosh-lite-cf1.json")
		db.FixtureDeployment("../upload/fixtures/deployment-bosh-lite-cf2.json")

		db.FixtureBosh("../upload/fixtures/bosh-vsphere-sandbox.json")
		db.FixtureDeployment("../upload/fixtures/deployment-vsphere-sandbox-cf.json")

		db.FixtureBosh("../upload/fixtures/bosh-aws-production.json")
		db.FixtureDeployment("../upload/fixtures/deployment-aws-production-cf.json")
	})

	It("finds releases", func() {
		Expect(db.ReleaseNames()).To(Equal([]string{"cf", "cf-haproxy", "concourse", "garden-linux"}))
	})
})
