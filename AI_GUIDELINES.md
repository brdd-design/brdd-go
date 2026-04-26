# 🤖 AI Guidelines for brdd-go

## 🏗 Core Components to Implement

### 1. `ExecutionContext` Struct
```go
type ExecutionContext struct {
    Data    interface{} `json:"data"`
    Errors  []BRDDError `json:"errors"`
    Setters []string    `json:"setters"`
    Effects []string    `json:"effects"`
    Status  int         `json:"status"`
}
```

### 2. Implementation Rules for AI Agents
- **No Globals:** Use dependency injection for specialized services.
- **Error Handling:** Don't return `error` in UseCases; instead, return the `ExecutionContext` with the error list populated.
- **Context Package:** Always accept `context.Context` as the first argument in service methods for timeout/cancellation.
