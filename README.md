# Url Shortener

My goal is to study Golang, and for that I am developing a simple application that is a Shortener Url.

To save the generated urls I use a slice as if it were a memory database, very simple.

There is no validation for a new generated url, it just generates a response for you :)

## How to test

Clone this repository:

```bash
git clone git@github.com:aristidesneto/shortener-url.git && cd shortener-url
```

Export host variable:

```bash
export BASE_URL=http://localhost:8080
```

Run `main.go`

```bash
go run main.go
```

Execute the cURL command:

```bash
# Creating shortened urls
curl -X POST --location 'http://localhost:8080/short-url' \
    --header 'Content-Type: application/json' \
    --data '{
        "url": "https://github.com/aristidesneto/shortener-url"
    }'
```

The response will be a json with a new shortened url. Copy and paste into your browser.

```bash
# Example response
{"url":"http://localhost:8080/ddKc0"}
```
