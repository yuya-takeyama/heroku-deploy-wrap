# heroku-deploy-wrap

Wrap git push heroku master and exit 1 when deploy failed.

## Usage

```
$ heroku-deploy-wrap -- git push heroku master
```

## How it works

`heroku-deploy-wrap` executes passed command and checks the output.

* When the output contains either of below, it exists with `1`.
  * `remote: Verifying deploy... done.`
  * `Everything up-to-date`

Or if the command exited with non-zero, then `heroku-deploy-wrap` exits with same exit-code.

## Installation

```
$ go get github.com/yuya-takeyama/heroku-deploy-wrap
```

## Author

Yuya Takeyama

## License

The MIT License
