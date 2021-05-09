# gghtml - (Go GET HTML) a threaded html source downloader.

> Speed up the GET request and HTML retrieval.

I created gghtml to speed up the downloading of a webpage html source code which was previously done in python. Below is an example usage from my python project.

## Usage

```py
import subprocess

# a: string = list of urls comma separated
a = "<url1>,<url2>"

return_json = subprocess.run(
  ["go", "<gghtml_DirPath>", a],
  capture_output=True
)
```

### Return

```go
// The return object is a string dumped to 
// stdout in the formatted object below
// Articles (list)
//  |-> webpage (items)

{"Articles": [
  {"InputOrder": int, "html": string}]
}
```

#### TODO

- [ ] **comparison Go vs Python**
- [ ] Shared Binary