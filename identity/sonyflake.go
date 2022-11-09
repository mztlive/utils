package identity

import (
	"fmt"
	"sync"

	"github.com/sony/sonyflake"
	"github.com/spf13/cast"
)

var sf *sonyflake.Sonyflake
var sfOnceLock sync.Once

// Next 获得下一个唯一ID
func Next() string {
	//TODO: NodeID 换成读取配置

	sfOnceLock.Do(func() {
		st := sonyflake.Settings{}
		sf = sonyflake.NewSonyflake(st)
	})

	ID, err := sf.NextID()

	if err != nil {
		panic(fmt.Errorf("generate identity fail. error: %w", err))
	}

	return cast.ToString(ID)
}
