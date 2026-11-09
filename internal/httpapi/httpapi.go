package httpapi

// 稳定错误码枚举。新增错误场景必须复用或在此登记，禁止散落字符串。
const (
	CodeOK                    = "OK"
	CodeInvalidArgument       = "INVALID_ARGUMENT"       // 400：请求参数不合法
	CodeUnauthenticated       = "UNAUTHENTICATED"        // 401：缺少或错误的 API Key
	CodeNotFound              = "NOT_FOUND"              // 404：资源或路由不存在
	CodeFeatureDisabled       = "FEATURE_DISABLED"       // 404：能力未启用（配置关闭）
	CodeRateLimited           = "RATE_LIMITED"           // 429：触发限流
	CodeDependencyUnavailable = "DEPENDENCY_UNAVAILABLE" // 503：Redis/ES/MySQL/观测数据源不可用
	CodeUpstreamError         = "UPSTREAM_ERROR"         // 502：LLM 等上游模型服务失败
	CodeInternal              = "INTERNAL"               // 500：未分类内部错误
)

// Response 是统一响应封装。
type Response struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
