package db

import (
	"context"
	"fmt"
	"net/netip"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetGeoInfo(t *testing.T) {
	if testing.Short() {
		fmt.Println("***** TestGetGeoInfo is ignored *****")
		t.Skip()
	}

	// 测试 IP
	ipStr := "116.204.73.62"

	// 将字符串 IP 转换为 netip.Addr 类型
	ip, err := netip.ParseAddr(ipStr)
	require.NoError(t, err)

	// 从数据库中查询对应的地理信息
	_, err = testStore.GetGeoInfo(context.Background(), ip)
	require.NoError(t, err)

}
