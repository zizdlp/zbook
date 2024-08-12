package db

import (
	"context"
	"fmt"
	"net/netip"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetGeoInfo(t *testing.T) {
	if testing.Short() {
		fmt.Println("***** TestGetGeoInfo is ignored *****")
		t.Skip()
	}
	ipStr := "116.204.73.62"
	ip, err := netip.ParseAddr(ipStr)
	require.NoError(t, err)

	// 记录开始时间
	start := time.Now()

	for i := 0; i < 1000; i++ {
		_, err := testStore.GetGeoInfo(context.Background(), ip)
		require.NoError(t, err)
	}

	// 记录结束时间并计算总耗时
	elapsed := time.Since(start)
	fmt.Printf("1000 次查询耗时: %s\n", elapsed)
}
func TestGetGeoInfoBatch(t *testing.T) {
	if testing.Short() {
		fmt.Println("***** TestGetGeoInfo is ignored *****")
		t.Skip()
	}
	ipStr := "116.204.73.62"
	ip, err := netip.ParseAddr(ipStr)
	require.NoError(t, err)

	// 创建一个包含 100 个相同 IP 地址的切片
	ips := make([]netip.Addr, 1000)
	for i := 0; i < 1000; i++ {
		ips[i] = ip
	}

	// 记录开始时间
	start := time.Now()

	// 执行批量查询
	_, err = testStore.GetGeoInfoBatch(context.Background(), ips)
	require.NoError(t, err)

	// 记录结束时间并计算总耗时
	elapsed := time.Since(start)
	fmt.Printf("1000 次查询耗时: %s\n", elapsed)
}
