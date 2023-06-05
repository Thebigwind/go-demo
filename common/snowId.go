package common

import (
	"github.com/bwmarrin/snowflake"
)

func GetSeqMaker() (int64, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		println(err.Error())
		return 0, err
	}
	return int64(node.Generate()), nil
}
