package config

var ModelReverseMap = map[string]string{}
var ModelMap = map[string]string{
	"claude-4-sonnet":         "claude2",
	"claude-4-sonnet-think":   "claude37sonnetthinking",
	"claude-4-5-sonnet":       "claude45sonnet",
	"claude-4.5-sonnet-think": "claude45sonnetthinking",
	"claude-4-6-sonnet":       "claude46sonnet",
	"claude-4.6-sonnet-think": "claude46sonnetthinking",
	"gemini-2.5-pro":          "gemini25pro",
	"gemini-3-pro":            "gemini30pro",
	"kimi-k2-thinking":        "kimik2thinking",
	"grok":                    "grok",
	"grok-4":                  "grok4",
	"grok-4-non-thinking":     "grok4nonthinking",
    "grok-4.1-reasoning":      "grok41reasoning",
	"grok-4.1-non-reasoning":  "grok41nonreasoning",
	"o3-pro":                  "o3pro",
	"o4-pro":                  "o4mini",
	"gpt-4o":                  "gpt4o",
	"gpt-4.1":                 "gpt41",
	"gpt-5.1":                 "gpt51",
	"gpt-5-think":             "gpt5_thinking",
	"claude-4-opus":           "claude40opus",
	"claude-4-opus-think":     "claude40opusthinking",
	"claude-4.1-opus":         "claude41opus",
	"claude-4.1-opus-think":   "claude41opusthinking",
	"claude-4.5-opus":         "claude45opus",
	"claude-4.5-opus-think":   "claude45opusthinking",
    "claude-4.6-opus":         "claude46opus",
	"claude-4.6-opus-think":   "claude46opusthinking",
	"alpha":                   "pplx_alpha",
	"beta":                    "pplx_beta",
	"study":                   "pplx_study",
	"r1":                      "r1",
    "claude40sonnetthinking-labs":   "claude40sonnetthinking_labs",
	"claude40opusthinking-labs":   "claude40opusthinking_labs",
}
var MaxModelMap = map[string]string{
	"o3-pro":                "o3pro",
	"claude-4.1-opus-think": "claude41opusthinking",
}

// Get returns the value for the given key from the ModelMap.
// If the key doesn't exist, it returns the provided default value.
func ModelMapGet(key string, defaultValue string) string {
	if value, exists := ModelMap[key]; exists {
		return value
	}
	return defaultValue
}

// GetReverse returns the value for the given key from the ModelReverseMap.
// If the key doesn't exist, it returns the provided default value.
func ModelReverseMapGet(key string, defaultValue string) string {
	if value, exists := ModelReverseMap[key]; exists {
		return value
	}
	return defaultValue
}

var ResponseModels []map[string]string

func init() {
	// 构建反向映射
	for k, v := range ModelMap {
		ModelReverseMap[v] = k
	}
	buildResponseModels()
}

// buildResponseModels 构建响应模型列表
func buildResponseModels() {
	ResponseModels = make([]map[string]string, 0, len(ModelMap)*2)

	for modelID := range ModelMap {
		// 如果不是最大订阅用户，跳过最大模型
		if !ConfigInstance.IsMaxSubscribe {
			if _, isMaxModel := MaxModelMap[modelID]; isMaxModel {
				continue
			}
		}

		// 添加普通模型
		ResponseModels = append(ResponseModels, map[string]string{
			"id": modelID,
		})

		// 添加搜索模型
		ResponseModels = append(ResponseModels, map[string]string{
			"id": modelID + "-search",
		})
	}
}



