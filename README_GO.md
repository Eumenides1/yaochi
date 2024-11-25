
# 瑶池 (Yaochi) - Go 语言版本

**瑶池** 是一个高性能、轻量级、可扩展的键值数据库，灵感来源于 Redis。Go 语言版本以简单高效为目标，提供快速开发和性能优化的能力，同时保留可扩展性和易维护性。

---

## **功能特色**

- 🚀 **高性能**：基于 Go 的原生并发支持，优化网络与 IO 性能。
- 🔧 **高度可配置**：通过配置文件轻松管理服务器配置。
- 🌈 **模块化设计**：核心架构支持快速扩展和定制化。
- 📜 **丰富的日志功能**：支持多级日志输出，方便调试与监控。
- 🛠️ **稳健架构**：可靠的资源管理与优雅停机能力。

---

## **快速开始**

### 环境要求

- Go 1.19 或更高版本

### 安装步骤

1. 克隆代码仓库：
   ```bash
   git clone https://github.com/your-username/yaochi-go.git
   cd yaochi-go
   ```

2. 构建项目：
   ```bash
   go build -o yaochi
   ```

3. 配置服务：
   修改 `yaochi.conf` 文件以适应您的环境：
   ```conf
   bind 0.0.0.0
   port 6379
   appendOnly yes
   appendFilename appendonly.aof
   maxClients 100
   requirePass password123
   databases 16
   ```

4. 启动服务：
   ```bash
   ./yaochi
   ```

---

## **使用方法**

### 连接到服务器

使用 `telnet` 或类似工具连接到服务器：
```bash
telnet 127.0.0.1 6379
```

发送一条消息并接收回显：
```plaintext
Hello, Yaochi!
Hello, Yaochi!
```

### 日志

日志默认存储在项目根目录，支持按日期和大小滚动的日志文件，方便管理。

---

## **未来规划**

- 🔄 **持久化**：实现 RDB 和 AOF 的持久化存储功能。
- 📡 **复制**：支持主从复制模式。
- 🕸️ **集群模式**：引入分布式集群，支持横向扩展。
- 📈 **性能指标**：集成监控和分析工具。
- 🧩 **插件支持**：启用基于插件的自定义命令扩展。

---

## **贡献指南**

欢迎为瑶池 Go 版本做出贡献！以下是参与的步骤：
1. Fork 本仓库。
2. 创建功能分支：
   ```bash
   git checkout -b feature-name
   ```
3. 提交代码：
   ```bash
   git commit -m "Add a new feature"
   ```
4. 推送到远程分支：
   ```bash
   git push origin feature-name
   ```
5. 提交 Pull Request。

---

## **许可证**

本项目基于 **MIT 许可证**，详情请参阅 [LICENSE](LICENSE) 文件。

---

## **致谢**

- 灵感来源于 [Redis](https://redis.io/) 及其极简的设计。
- 项目名称取自中国神话中的 **瑶池**，象征速度、和谐与完美。

---

## **联系信息**

如果您有任何问题或反馈，请通过以下方式联系：
- **邮箱**: your-email@example.com
- **GitHub Issues**: [提交问题](https://github.com/your-username/yaochi-go/issues)

让 Yaochi 为您的数据之旅注入速度与平衡！ 🌊✨

---

**[English Version](README)**
