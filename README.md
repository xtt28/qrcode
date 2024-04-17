# xtt28/qrcode

A REST API for generating QR codes.

## Description

Built with the [go-qrcode](github.com/skip2/go-qrcode) library, this project
allows for generation of QR codes from URLs.

## Setup

1. `git clone https://github.com/xtt28/qrcode.git`
2. `cd qrcode`
3. `go run .`
4. The application will be accessible at <http://localhost:8080>

## Environment Variables

- **PORT** - The port to run the application on. Default is 8080.

## API Reference

### Generate a QR code

```http
  GET /
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `url` | `string` | **Required**. The URL to make the QR code resolve to. |
| `size` | `int` | The size in pixels of the QR code. |

## License

MIT