package util

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type VisitorData struct {
	IP    string
	Agent string
	Count int
}

// 聚合数据函数
func AggregateByIP(visitors []*VisitorData) map[string]int {
	// 创建一个用于存储按 IP 聚合后的计数的映射
	ipCounts := make(map[string]int)

	for _, visitor := range visitors {
		if visitor.IP != "" {
			// 将相同 IP 的计数累加
			ipCounts[visitor.IP] += visitor.Count
		}
	}
	return ipCounts
}
func TopNVisitors(visitors []*VisitorData, topN int) []VisitorData {
	// 创建一个map用于根据IP聚合count
	aggregate := make(map[string]int)
	for _, visitor := range visitors {
		aggregate[visitor.IP] += visitor.Count
	}

	// 将map转换为slice以便排序
	aggregatedList := make([]VisitorData, 0, len(aggregate))
	for ip, count := range aggregate {
		aggregatedList = append(aggregatedList, VisitorData{
			IP:    ip,
			Count: count,
		})
	}

	// 按 count 降序排序，如果 count 相同，则按 IP 降序排序
	sort.Slice(aggregatedList, func(i, j int) bool {
		if aggregatedList[i].Count == aggregatedList[j].Count {
			return aggregatedList[i].IP > aggregatedList[j].IP // IP 降序
		}
		return aggregatedList[i].Count > aggregatedList[j].Count
	})
	// 如果 topN 是负数或 0，返回空列表
	if topN <= 0 {
		return []VisitorData{}
	}

	// 取前topN个IP
	if len(aggregatedList) > topN {
		aggregatedList = aggregatedList[:topN]
	}

	return aggregatedList
}

// AgentCounts 结构体定义
type AgentCounts struct {
	Bot      int
	Computer int
	Phone    int
	Tablet   int
	Unknown  int
}

// 聚合数据函数
func SumAgentCounts(visitors []*VisitorData) AgentCounts {
	agentCounts := AgentCounts{}

	for _, visitor := range visitors {
		agent := strings.ToLower(visitor.Agent)
		visited := false

		// 分类为 Bot
		if strings.Contains(agent, "bot") || strings.Contains(agent, "spider") {
			visited = true
			agentCounts.Bot += visitor.Count
		}

		// 如果已经分类为 Bot，则跳过其余分类
		if visited {
			continue
		}

		// 分类为 Tablet
		if strings.Contains(agent, "ipad") ||
			strings.Contains(agent, "android tablet") ||
			strings.Contains(agent, "kindle") {
			agentCounts.Tablet += visitor.Count
		} else if strings.Contains(agent, "iphone") ||
			strings.Contains(agent, "android") ||
			strings.Contains(agent, "windows phone") ||
			strings.Contains(agent, "blackberry") {
			agentCounts.Phone += visitor.Count
		} else if strings.Contains(agent, "windows") ||
			strings.Contains(agent, "macintosh") ||
			(strings.Contains(agent, "linux") && !strings.Contains(agent, "android")) {
			agentCounts.Computer += visitor.Count
		} else {
			agentCounts.Unknown += visitor.Count
		}
	}

	return agentCounts
}

// UserAgentData holds parsed user agent information
type UserAgentData struct {
	Platform       string
	OS             string
	Browser        string
	BrowserVersion string
	Engine         string
	EngineVersion  string
} // parseUserAgent parses the user agent string and returns UserAgentData
func parseUserAgent(agent string) *UserAgentData {
	uaData := &UserAgentData{}

	if agent == "" {
		return uaData
	}

	// Extract platform and OS
	platformRegex := regexp.MustCompile(`\(([^)]+)\)`)
	platformMatch := platformRegex.FindStringSubmatch(agent)
	fmt.Println("plat:", platformMatch)
	if len(platformMatch) > 1 {

		platformParts := strings.Split(platformMatch[1], ";")
		if len(platformParts) > 0 {
			uaData.Platform = strings.TrimSpace(platformParts[0])
		}
		if len(platformParts) > 1 {
			uaData.OS = strings.TrimSpace(platformParts[1])
		}
	}

	// Extract browser and version
	browserRegexes := []struct {
		regex *regexp.Regexp
		name  string
	}{
		{regexp.MustCompile(`(Chrome)\/([0-9.]+)`), "Chrome"},
		{regexp.MustCompile(`(Safari)\/([0-9.]+)`), "Safari"},
		{regexp.MustCompile(`(Firefox)\/([0-9.]+)`), "Firefox"},
		{regexp.MustCompile(`(Edg)\/([0-9.]+)`), "Edge"},
		{regexp.MustCompile(`(OPR)\/([0-9.]+)`), "Opera"},
		{regexp.MustCompile(`(MSIE) ([0-9.]+)`), "Internet Explorer"},
		{regexp.MustCompile(`(Trident)\/.*rv:([0-9.]+)`), "Internet Explorer"},
	}

	for _, br := range browserRegexes {
		match := br.regex.FindStringSubmatch(agent)
		if len(match) > 2 {
			uaData.Browser = br.name
			uaData.BrowserVersion = match[2]
			break
		}
	}

	// Extract rendering engine and version
	engineRegexes := []struct {
		regex *regexp.Regexp
		name  string
	}{
		{regexp.MustCompile(`(AppleWebKit)\/([0-9.]+)`), "AppleWebKit"},
		{regexp.MustCompile(`(Gecko)\/([0-9.]+)`), "Gecko"},
		{regexp.MustCompile(`(KHTML)\/([0-9.]+)`), "KHTML"},
		{regexp.MustCompile(`(Trident)\/([0-9.]+)`), "Trident"},
	}

	for _, eng := range engineRegexes {
		match := eng.regex.FindStringSubmatch(agent)
		if len(match) > 2 {
			uaData.Engine = eng.name
			uaData.EngineVersion = match[2]
			break
		}
	}

	return uaData
}
