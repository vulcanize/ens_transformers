package integration_tests

import (
	"github.com/vulcanize/ens_transformers/test_config"
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

func TestIntegrationTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IntegrationTests Suite")
}

var _ = BeforeSuite(func() {
	testConfig := test_config.InfuraClient
	ipc = testConfig.IPCPath
	log.SetOutput(ioutil.Discard)
})
