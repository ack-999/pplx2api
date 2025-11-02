package config

var ModelReverseMap = map[string]string{}
var ModelMap = map[string]string{
	"claude-2": "claude2",
	"claude-4.5-sonnet": "claude37sonnet",
	"claude-3.7-sonnet-think": "claude37sonnetthinking",
	"claude-4-5-sonnet":       "claude45sonnet",
	"claude-4.5-sonnet-think": "claude45sonnetthinking",
	"gemini-2.5-pro-06-05":    "gemini2flash",
	"grok4":   "grok4",
	"grok4-non-thinking":   "grok4nonthinking",
	"gpt-4o":   "gpt4o",
	"gpt-41":   "gpt41",
	"o4-mini":   "o4mini",
	"o3-pro":                  "o3pro",
	"gpt-5":                   "gpt5",
	"gpt-5-think":             "gpt5_thinking",
	"gpt5-pro":   "gpt5_pro",
	"claude-3-opus":   "claude3opus",
	"claude-4.0-opus":   "claude40opus",
	"claude-4.0-opus-think":   "claude40opusthinking",
	"claude-4.1-opus":   "claude41opus",
	"claude-4.1-opus-think":   "claude41opusthinking",
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



