# Example Go web app with YAML config

This is full project structure example of article **[Let's write config for your Golang web app on right way — YAML 👌](https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp)**.

Published on [Dev.to](https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp) @ 13 Feb 2020.

## Requirements

- Go `1.11+`
- Go Modules

## Usage

- Clone this repository and go to folder:

```console
git clone https://github.com/koddr/example-go-config-yaml.git
cd example-go-config-yaml
```

- Run server (by one of commands below):

```console
# Build server, run tests, collect coverage info and run server:
make

# Only server with default config and without build and coverage info:
make run

# Use your own config file:
go run ./... -config my-config.yml
```

- For show info about coverage (will open in browser):

```console
make show_coverage
```

- That's all! 🎉

## Author

- [Vic Shóstak](https://github.com/koddr) (aka Koddr).

## Code reviewer

- [Jordan Gregory](https://github.com/j4ng5y) (aka j4ng5y).

## Article assistance

If you want to say «thank you»:

1. Twit about article [on your Twitter](https://twitter.com/intent/tweet?text=Let%27s%20write%20config%20for%20your%20Golang%20web%20app%20on%20right%20way%20%E2%80%94%20YAML%20%F0%9F%91%8C%20https%3A%2F%2Fdev.to%2Fkoddr%2Flet-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp).
2. Add a GitHub Star and make Fork to this repository.
3. Donate some money to project author via PayPal: [@paypal.me/koddr](https://paypal.me/koddr?locale.x=en_EN) or [LiberaPay](https://liberapay.com/koddr/donate).
4. Join DigitalOcean at my [referral link](https://m.do.co/c/b41859fa9b6e) (your profit is **\$100** and we get \$25).

Thanks for your support! 😘

## License

MIT
