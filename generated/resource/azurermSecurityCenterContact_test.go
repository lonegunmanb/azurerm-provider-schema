package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzurermSecurityCenterContactSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzurermSecurityCenterContactSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
