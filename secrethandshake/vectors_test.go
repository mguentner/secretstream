package secrethandshake

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/cryptix/secretstream/secrethandshake/stateless"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestVectors(t *testing.T) {
	dataf, err := os.Open("test-vectors/data.json")
	assert.Nil(t, err)

	defer dataf.Close()

	var data []map[string]interface{}
	assert.Nil(t, json.NewDecoder(dataf).Decode(&data))

	for i, v := range data {
		if i > 1 {
			return
		}
		switch v["name"] {
		case "initialize": //
			args := v["args"].([]interface{})
			assert.Len(t, args, 1, "init test %d", i)

			// parse args
			var argState stateless.JsonState
			err := mapstructure.Decode(args[0], &argState)
			assert.Nil(t, err, "init test %d", i)
			initState, err := stateless.JsonStateToOurState(argState)
			assert.Nil(t, err, "init test %d", i)

			// parse result
			var resultState stateless.JsonState
			err = mapstructure.Decode(v["result"], &resultState)
			assert.Nil(t, err, "init test %d", i)

			assert.Equal(t, resultState, *initState.ToJsonState())

		default:
			// t.Logf("unhandled case testing %d: %s", i, v["name"])
		}
	}
}