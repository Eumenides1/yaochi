
# Yaochi - Go Version

**Yaochi** is a high-performance, lightweight, and extensible key-value database inspired by Redis. The Go version focuses on simplicity and efficiency, providing fast development and performance optimization while retaining extensibility and maintainability.

---

## **Features**

- 🚀 **High Performance**: Optimized network and IO performance leveraging Go's native concurrency.
- 🔧 **Highly Configurable**: Manage server configurations effortlessly through a config file.
- 🌈 **Modular Design**: Core architecture supports rapid extension and customization.
- 📜 **Rich Logging**: Multi-level logging for debugging and monitoring.
- 🛠️ **Robust Architecture**: Reliable resource management and graceful shutdown capabilities.

---

## **Getting Started**

### Prerequisites

- Go 1.19 or later

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/yaochi-go.git
   cd yaochi-go
   ```

2. Build the project:
   ```bash
   go build -o yaochi
   ```

3. Configure the server:
   Edit the `yaochi.conf` file to suit your environment:
   ```conf
   bind 0.0.0.0
   port 6379
   appendOnly yes
   appendFilename appendonly.aof
   maxClients 100
   requirePass password123
   databases 16
   ```

4. Run the server:
   ```bash
   ./yaochi
   ```

---

## **Usage**

### Connecting to the Server

Use `telnet` or similar tools to interact with the server:

```bash
telnet 127.0.0.1 6379
```

Send a message and get an echo response:
```plaintext
Hello, Yaochi!
Hello, Yaochi!
```

### Logs

Logs are stored in the project root by default, supporting rolling logs based on size and date for easy management.

---

## **Roadmap**

- 🔄 **Persistence**: Implement RDB and AOF for durable storage.
- 📡 **Replication**: Add support for master-slave replication.
- 🕸️ **Cluster Mode**: Introduce distributed clustering for horizontal scaling.
- 📈 **Metrics**: Integrate monitoring and analytics tools.
- 🧩 **Plugins**: Enable plugin-based extensions for custom commands.

---

## **Contributing**

We welcome contributions to the Go version of Yaochi! Here's how you can help:
1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add a new feature"
   ```
4. Push to your branch:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request.

---

## **License**

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

---

## **Acknowledgments**

- Inspired by [Redis](https://redis.io/) and its incredible simplicity.
- Named after the **Jade Pool (瑶池)** in Chinese mythology, representing purity, balance, and perfection.

---

## **Contact**

For questions or feedback, feel free to reach out via:
- **Email**: your-email@example.com
- **GitHub Issues**: [Submit an issue](https://github.com/your-username/yaochi-go/issues)

Let Yaochi bring balance and speed to your data journey! 🌊✨
