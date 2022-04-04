package template_test

import (
	"crypto/tls"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-cli/director/template"
	"github.com/cloudfoundry/bosh-cli/testutils"
)

var (
	cert        tls.Certificate
	cacertBytes []byte
	validCACert string
)
var _ = SynchronizedBeforeSuite(func() []byte {
	var err error
	cert, cacertBytes, err = testutils.Certsetup()
	validCACert = string(cacertBytes)
	Expect(err).ToNot(HaveOccurred())
	return []byte{}
}, func(in []byte) {})

func TestReg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "director/template")
}

type FakeVariables struct {
	GetFunc      func(VariableDefinition) (interface{}, bool, error)
	GetVarDef    VariableDefinition
	GetErr       error
	GetCallCount int
}

func (v *FakeVariables) Get(varDef VariableDefinition) (interface{}, bool, error) {
	v.GetCallCount++
	v.GetVarDef = varDef
	if v.GetFunc != nil {
		return v.GetFunc(varDef)
	}
	return nil, false, v.GetErr
}

func (v *FakeVariables) List() ([]VariableDefinition, error) {
	return nil, nil
}
