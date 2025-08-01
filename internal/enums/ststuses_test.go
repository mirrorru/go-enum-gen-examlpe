package enums_test

import (
	"testing"

	"go-enum-gen-example/internal/enums"

	"github.com/stretchr/testify/assert"
)

func TestClientStatus_HumanString(t *testing.T) {
	t.Parallel()
	hsm := make(map[string]enums.ClientStatus) // human strings map
	for _, val := range enums.ClientStatusValues() {
		hs := val.HumanString()
		ct, ok := hsm[hs]
		assert.False(t, ok, "Для элемента %v HumanString как у элемента %v", val, ct)
		hsm[hs] = val
	}
	assert.Len(t, hsm, len(enums.ClientStatusValues())) // Контрольная проверка, что все элементы прошли

	notExistCT := enums.ClientStatus(250)
	assert.Empty(t, notExistCT.HumanString())
}
