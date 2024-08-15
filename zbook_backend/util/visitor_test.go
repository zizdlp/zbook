package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTopNVisitors(t *testing.T) {
	visitors := []*VisitorData{
		{IP: "192.168.1.1", Agent: "Mozilla", Count: 10},
		{IP: "192.168.1.2", Agent: "Chrome", Count: 15},
		{IP: "192.168.1.1", Agent: "Safari", Count: 20},
		{IP: "192.168.1.3", Agent: "Mozilla", Count: 25},
		{IP: "192.168.1.2", Agent: "Edge", Count: 5},
		{IP: "192.168.1.4", Agent: "Opera", Count: 30},
	}

	topN := 2
	expectedTopVisitors := []VisitorData{
		{IP: "192.168.1.4", Count: 30},
		{IP: "192.168.1.1", Count: 30},
	}

	topVisitors := TopNVisitors(visitors, topN)
	require.Equal(t, expectedTopVisitors, topVisitors, "Top N visitors do not match expected output")

	// 测试当N大于输入长度时
	topN = 10
	expectedTopVisitors = []VisitorData{
		{IP: "192.168.1.4", Count: 30},
		{IP: "192.168.1.1", Count: 30},
		{IP: "192.168.1.3", Count: 25},
		{IP: "192.168.1.2", Count: 20},
	}
	topVisitors = TopNVisitors(visitors, topN)
	require.Equal(t, expectedTopVisitors, topVisitors, "Top visitors when N > len(visitors) do not match expected output")

	// 测试当N为0时
	topN = 0
	expectedTopVisitors = []VisitorData{}
	topVisitors = TopNVisitors(visitors, topN)
	require.Equal(t, expectedTopVisitors, topVisitors, "Top visitors when N=0 should be an empty list")
}
func TestParseUserAgent(t *testing.T) {

	tests := []struct {
		agent    string
		expected *UserAgentData
	}{
		{
			agent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
			expected: &UserAgentData{
				Platform:       "Windows NT 10.0",
				OS:             "Win64",
				Browser:        "Chrome",
				BrowserVersion: "91.0.4472.124",
				Engine:         "AppleWebKit",
				EngineVersion:  "537.36",
			},
		},
		{
			agent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) Gecko/20100101 Firefox/89.0",
			expected: &UserAgentData{
				Platform:       "Macintosh",
				OS:             "Intel Mac OS X 10_15_7",
				Browser:        "Firefox",
				BrowserVersion: "89.0",
				Engine:         "Gecko",
				EngineVersion:  "20100101",
			},
		},
		{
			agent:    "", // 测试空字符串
			expected: &UserAgentData{},
		},
		{
			agent: "Mozilla/5.0 (Linux; Android 10; SM-G973U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.109 Mobile Safari/537.36",
			expected: &UserAgentData{
				Platform:       "Linux",
				OS:             "Android 10",
				Browser:        "Chrome",
				BrowserVersion: "85.0.4183.109",
				Engine:         "AppleWebKit",
				EngineVersion:  "537.36",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.agent, func(t *testing.T) {
			actual := parseUserAgent(tt.agent)
			require.Equal(t, tt.expected, actual, "User agent parsing did not match expected output")
		})
	}
}

// 测试 SumAgentCounts 函数
func TestSumAgentCounts(t *testing.T) {
	tests := []struct {
		name     string
		visitors []*VisitorData
		expected AgentCounts
	}{
		{
			name: "Test with different agents",
			visitors: []*VisitorData{
				{Agent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64)", Count: 10},
				{Agent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X)", Count: 5},
				{Agent: "Mozilla/5.0 (Linux; Android 10; Pixel 4 XL Build/QD1A.190821.016)", Count: 7},
				{Agent: "Mozilla/5.0 (iPad; CPU OS 14_0 like Mac OS X)", Count: 3},
				{Agent: "Googlebot/2.1 (+http://www.google.com/bot.html)", Count: 2},
				{Agent: "UnknownAgent/1.0", Count: 1},
			},
			expected: AgentCounts{
				Bot:      2,
				Computer: 10,
				Phone:    12,
				Tablet:   3,
				Unknown:  1,
			},
		},
		{
			name:     "Test with empty visitors",
			visitors: []*VisitorData{},
			expected: AgentCounts{},
		},
		{
			name: "Test with only bots",
			visitors: []*VisitorData{
				{Agent: "Bingbot/2.0 (+http://www.bing.com/bingbot.htm)", Count: 5},
				{Agent: "DuckDuckBot/1.0; (+http://duckduckgo.com)", Count: 3},
			},
			expected: AgentCounts{
				Bot: 8,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumAgentCounts(tt.visitors)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
