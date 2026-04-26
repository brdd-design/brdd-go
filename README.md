# 🐹 BRDD Go

A minimalist implementation of the Business Rule Driven Design (BRDD) pattern for Go.

## 🎯 Goal
Provide clean interfaces and a standard execution context for Go-based microservices.

## 🚀 Quick Start (Initialization)

1. **Initialize the Module:**
   ```bash
   mkdir brdd-go && cd brdd-go
   go mod init github.com/your-org/brdd-go
   ```

2. **Create the Core Types:**
   Implement `ExecutionContext` and `ValidationContext` structs in a `repositories` package.

3. **Standardize the Response:**
   Use tags like `json:"data"` to match the [Unified Response Pattern](https://github.com/brdd-design/brdd/tree/main/articles/EN/BRDD-STORY.md).

## 🚀 Publication & Metadata

To publish this library as a **Go Module**, ensure:
- **Module Name:** `github.com/brdd-design/brdd-go`
- **Tagging:** Use Semantic Versioning tags (ex: `v0.1.0`).
- **License:** MIT

**Step-by-step to Publish:**
1. `go mod tidy`.
2. `git tag v0.1.0`.
3. `git push origin v0.1.0`.

## 📚 Documentation
- [AI Guidelines](./AI_GUIDELINES.md)
- [Technical Spec](https://github.com/brdd-design/brdd/tree/main/articles/EN/BRDD-STORY.md)
- [Publishing Checklist](../../../PUBLISHING_CHECKLIST.md)
