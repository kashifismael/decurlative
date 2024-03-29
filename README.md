# Decurlative

A Declarative JSON abstraction of the CURL CLI tool, written in Go

## Example

Provide a JSON input via standard input
```json
{
    "url": "https://jsonplaceholder.typicode.com/users",
    "method": "GET",
    "queryParams": {
        "language": [
            "en"
        ]
    }
}
```
(Save above json into a file, e.g. `http.json`)

```shell
cat http.json | ./decurlative
```