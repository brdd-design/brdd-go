# 🐹 BRDD Go

A minimalist implementation of the Business Rule Driven Design (BRDD) pattern for Go.

## 🎯 Goal
Provide clean interfaces and a standard execution context for Go-based microservices.

## 🚀 Installation

```bash
go get github.com/brdd-design/brdd-go
```

## 🛠 Usage

```go
import "github.com/brdd-design/brdd-go/brdd"

// In your UseCase
func Execute(data interface{}) brdd.ExecutionContext {
    ctx := brdd.NewDefaultExecutionContext(data)
    // ... logic
    ctx.AddEffect("CE001")
    return ctx
}
```

## 🤖 AI-First Development
This library is designed for AI-driven development. Check the [AI Guidelines](./AI_GUIDELINES.md) for more details.

## 📚 Documentation
- [Technical Spec](https://github.com/brdd-design/brdd/blob/main/BRDD.md)
- [Practical Example](https://github.com/brdd-design/brdd/blob/main/core/articles/EN/BRDD-PRACTICAL-EXAMPLE.md)

## 📄 License
MIT
