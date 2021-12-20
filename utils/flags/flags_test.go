package flags_test

import (
	"testing"

	"github.com/angel-one/ws-load-test/constants"
	"github.com/angel-one/ws-load-test/utils/flags"
	"github.com/stretchr/testify/assert"
)

func TestBaseConfigPath(t *testing.T) {
	assert.Equal(t, constants.BaseConfigPathDefaultValue, flags.BaseConfigPath())
}
