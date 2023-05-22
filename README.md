# Scrapy

**Scrapy** is a CLI tool used to scrape data from various websites. To install the Scrapy CLI, use the command `go install github.com/Pradumnasaraf/scrapy@latest`. Go will automatically install it in your `$GOPATH/bin` directory, which should be in your `$PATH`.

Once installed, you can use the `scrapy` CLI command. To confirm installation, type `scrapy` at the command line.

> **Note** If you are getting an error like `command not found: scrapy`, then you need to add `$GOPATH/bin` to your `$PATH` environment variable. For that you can refer to [this](https://gist.github.com/Pradumnasaraf/ca6f9a0507089a4c44881446cdda4aa3).

## ⭐️ Features

- Scape Ebay products and their prices
- Scrape TechCrunch articles in different categories

## 📝 Usage

```
Usage:
  scrapy [command] [Flags]

Available Commands:
  ebay [query] Scrape Ebay products and their prices.
    ebay [query] --page [Number of pages]
  tc          Scrape TechCrunch articles in different categories.
    tc --category [Category name] 
```

eg: `scrapy ebay "iphone 12 mini" --page 5`
eg: `scrapy tc --category "startups"`

## 📜 License

This project is licensed under the Apache-2.0 license - see the [LICENSE](LICENSE) file for details.

## 🛡 Security

If you discover a security vulnerability within this project, please check the [SECURITY](SECURITY.md) for more information.
