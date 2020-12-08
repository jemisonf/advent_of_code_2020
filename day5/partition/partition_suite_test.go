package partition_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPartition(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Partition Suite")
}
