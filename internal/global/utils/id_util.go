package utils

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

var (
	node   *snowflake.Node
	idOnce sync.Once
)

func getNode() *snowflake.Node {
	idOnce.Do(func() {
		var err error
		node, err = snowflake.NewNode(1)
		if err != nil {
			panic(err)
		}
	})
	return node
}

func GetSnowflakeID() uint64 {
	return uint64(getNode().Generate().Int64())
}
