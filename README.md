# details

[![CI](https://github.com/thekuwayama/details/workflows/CI/badge.svg)](https://github.com/thekuwayama/details/actions?workflow=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/thekuwayama/details)](https://goreportcard.com/report/github.com/thekuwayama/details)
[![MIT licensed](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://raw.githubusercontent.com/thekuwayama/details/master/LICENSE.txt)

`details` is the command that adds details tag header and footer.


## Usage

```bash
$ details -help
Usage of details:
  -file string
    	input file path
  -summary string
    	summary tag
```

````bash
$ echo "This is a sample string." | details
<details>

```
This is a sample string.
```

</details>
````

````bash
$ echo "This is a sample string." | details -summary placeholder
<details><summary>placeholder</summary>

```
This is a sample string.
```

</details>
````


## License

The CLI is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
