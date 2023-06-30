package base_test

import (
	"testing"

	"github.com/buffreak/domain-scraper/base"
	"github.com/stretchr/testify/assert"
)

func TestGetDomainsDan(t *testing.T) {
	dan := &base.Dan{}
	domains, err := dan.GetDomains("al", "0")
	assert.Nil(t, err)
	t.Logf("%+v", domains)
	convert, _ := base.StructToMap(dan)
	t.Logf("%+v", convert)
}
