package function

import (
	"encoding/json"
	"os"

	"github.com/ermes-labs/api-go/infrastructure"
)

func init() {
	var node infrastructure.Node

	// Get env variable
	env, ok := os.LookupEnv("ermes-node")

	if !ok {
		panic("ermes-node env variable is not set")
	}

	err := json.Unmarshal([]byte(env), &node)

	if err != nil {
		panic(err)
	}

}
