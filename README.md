# Gommerce

## **⚠️WIP⚠️**

## Description

RESTful API for a simple e-commerce web service written in Go programming language. This is a capstone project for [Project-Based Internship at Evermos](https://www.rakamin.com/virtual-internship-experience/back-end-developer-evermos).

It uses [Fiber](https://gofiber.io) as the web framework, [GORM](https://gorm.io) as the ORM library, and [MySQL](https://www.mysql.com) as the database. It also utilizes [Zap](https://github.com/uber-go/zap) as the logger, [Viper](https://github.com/spf13/viper) as the configuration manager, [Gomock](https://github.com/uber-go/mock) as the mocking library, and [Ginkgo](https://onsi.github.io/ginkgo) as the testing framework.

## Getting Started

1. Ensure you have [Go](https://go.dev/dl/) 1.21.3 or higher and [Task](https://taskfile.dev/installation/) installed on your machine:

   ```bash
   go version && task --version
   ```

2. Install all required tools for the project:

   ```bash
   task install
   ```

3. Create a copy of the `.env.example` file and rename it to `.env`:

   ```bash
   cp .env.example .env
   ```

   Update configuration values as needed.

## Documentation

For database schema documentation, see [here](https://dbdocs.io/bagashiz/Gommerce), powered by [dbdocs.io](https://dbdocs.io).

API documentation is on progress.

## Contributing

Developers interested in contributing to Gommerce project can refer to the [CONTRIBUTING](CONTRIBUTING.md) file for detailed guidelines and instructions on how to contribute.

## License

Gommerce project is licensed under the [MIT License](LICENSE), providing an open and permissive licensing approach for further development and usage.

## Learning References

- [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) by Uncle Bob
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch) by Bxcodec
- [go-clean-template](https://github.com/evrone/go-clean-template) by Evrone
