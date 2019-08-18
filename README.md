# canon-go
## What is this?
- Simple load test tool.
- Can write scenario (HTTP requests like user action) in YAML format.
- Can get report summarizing HTTP responses as JSON format.

## Install
### install Go
- Requires Go 1.12 or later.

### install canon-go
- Clone the source from github
```
git clone https://github.com/HiroyukiNakatsuma/canon-go.git
```

- Or go get
```
go get -u github.com/HiroyukiNakatsuma/canon-go
```

## Usage
### Define test scenario
- Define test scenario in YAML format
```
timeout: 30
actions:   # Choose "request" or "sleep"
  - request:
      method: GET   # Choose "GET" or "POST" or "PUT" or "PATCH" or "DELETE"
      url: http://localhost:80/?greet="Hello World!"
  - sleep: 10   # Define sleep time seconds in Integer
  - request:
      method: POST
      url: http://localhost:80/
      body: '{"greet":"Hello World!"}'
      ###
      # Define Multiple contents header like this:
      #   Accept: application/json,text/csv,application/xml
      #
      # not like this:
      #   Accept: application/json
      #   Accept: text/csv
      #   Accept: application/xml
      ###
      headers:
        Content-Type: application/json
        Authorization: Bearer token
```

### Execute test
```
go run cmd/canon-go/main.go --input /your_input_filepath.yml --output /your_output_filepath.json
```

### Confirm test report
-  Confirm `/your_output_filepath.json` file.
    - You can see summaries average response time by request. 

## Notice
If you use this tool and have trouble, the author do not take any responsibility.

## License
This tool is based on ["gotling"](https://github.com/eriklupander/gotling).
The license is as below and follow this:

```
The MIT License (MIT)

Copyright (c) 2015 ErikL

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
