# appleopensource

An [opensource.apple.com][opensource.apple.com] resource management package, and command line tool written in Go.

[![CircleCI][circleci-badge]][circleci] [![godoc.org][godoc-badge]][godoc] [![codecov.io][codecov-badge]][codecov] [![GA][ga-badge]][ga]

---

## Install CLI

```sh
go get -u -v github.com/zchee/appleopensource/cmd/aos
```

## Package API

[godoc.org/github.com/zchee/appleopensource][godoc]

## Contribute

~~See [CONTRIBUTING.md](CONTRIBUTING.md)~~. Not yet.

## License

The appleopensource is the released under the [BSD 3-Clause License][license-bsd-3-clause].

[Apple Open Source][opensource.apple.com] resources are under the [Apple Public Source License Version 2.0][license-apsl].


<!-- links -->
[opensource.apple.com]: https://opensource.apple.com
[license-bsd-3-clause]: https://opensource.org/licenses/BSD-3-Clause
[license-apsl]: http://www.opensource.apple.com/apsl/

<!-- badge links -->
[circleci]: https://circleci.com/gh/zchee/workflows/appleopensource
[godoc]: https://godoc.org/github.com/zchee/appleopensource
[codecov]: https://codecov.io/gh/zchee/appleopensource
[release]: https://github.com/zchee/appleopensource/releases
[ga]: https://github.com/zchee/appleopensource

[circleci-badge]: https://img.shields.io/circleci/token/c2558f666e80869d759f7ff9a5574a3ca101259a/project/github/zchee/appleopensource/master.svg?style=for-the-badge&label=darwin&logo=circleci
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=for-the-badge&label=GODOC.ORG&logoWidth=25&logo=data%3Aimage%2Fsvg%2Bxml%3Bcharset%3Dutf-8%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI0MCIgaGVpZ2h0PSI0MCIgdmlld0JveD0iODUgNTUgMTIwIDEyMCI+PHBhdGggZmlsbD0iIzJEQkNBRiIgZD0iTTQwLjIgMTAxLjFjLS40IDAtLjUtLjItLjMtLjVsMi4xLTIuN2MuMi0uMy43LS41IDEuMS0uNWgzNS43Yy40IDAgLjUuMy4zLjZsLTEuNyAyLjZjLS4yLjMtLjcuNi0xIC42bC0zNi4yLS4xek0yNS4xIDExMC4zYy0uNCAwLS41LS4yLS4zLS41bDIuMS0yLjdjLjItLjMuNy0uNSAxLjEtLjVoNDUuNmMuNCAwIC42LjMuNS42bC0uOCAyLjRjLS4xLjQtLjUuNi0uOS42bC00Ny4zLjF6TTQ5LjMgMTE5LjVjLS40IDAtLjUtLjMtLjMtLjZsMS40LTIuNWMuMi0uMy42LS42IDEtLjZoMjBjLjQgMCAuNi4zLjYuN2wtLjIgMi40YzAgLjQtLjQuNy0uNy43bC0yMS44LS4xek0xNTMuMSA5OS4zYy02LjMgMS42LTEwLjYgMi44LTE2LjggNC40LTEuNS40LTEuNi41LTIuOS0xLTEuNS0xLjctMi42LTIuOC00LjctMy44LTYuMy0zLjEtMTIuNC0yLjItMTguMSAxLjUtNi44IDQuNC0xMC4zIDEwLjktMTAuMiAxOSAuMSA4IDUuNiAxNC42IDEzLjUgMTUuNyA2LjguOSAxMi41LTEuNSAxNy02LjYuOS0xLjEgMS43LTIuMyAyLjctMy43aC0xOS4zYy0yLjEgMC0yLjYtMS4zLTEuOS0zIDEuMy0zLjEgMy43LTguMyA1LjEtMTAuOS4zLS42IDEtMS42IDIuNS0xLjZoMzYuNGMtLjIgMi43LS4yIDUuNC0uNiA4LjEtMS4xIDcuMi0zLjggMTMuOC04LjIgMTkuNi03LjIgOS41LTE2LjYgMTUuNC0yOC41IDE3LTkuOCAxLjMtMTguOS0uNi0yNi45LTYuNi03LjQtNS42LTExLjYtMTMtMTIuNy0yMi4yLTEuMy0xMC45IDEuOS0yMC43IDguNS0yOS4zIDcuMS05LjMgMTYuNS0xNS4yIDI4LTE3LjMgOS40LTEuNyAxOC40LS42IDI2LjUgNC45IDUuMyAzLjUgOS4xIDguMyAxMS42IDE0LjEuNi45LjIgMS40LTEgMS43eiIvPjxwYXRoIGZpbGw9IiMyREJDQUYiIGQ9Ik0xODYuMiAxNTQuNmMtOS4xLS4yLTE3LjQtMi44LTI0LjQtOC44LTUuOS01LjEtOS42LTExLjYtMTAuOC0xOS4zLTEuOC0xMS4zIDEuMy0yMS4zIDguMS0zMC4yIDcuMy05LjYgMTYuMS0xNC42IDI4LTE2LjcgMTAuMi0xLjggMTkuOC0uOCAyOC41IDUuMSA3LjkgNS40IDEyLjggMTIuNyAxNC4xIDIyLjMgMS43IDEzLjUtMi4yIDI0LjUtMTEuNSAzMy45LTYuNiA2LjctMTQuNyAxMC45LTI0IDEyLjgtMi43LjUtNS40LjYtOCAuOXptMjMuOC00MC40Yy0uMS0xLjMtLjEtMi4zLS4zLTMuMy0xLjgtOS45LTEwLjktMTUuNS0yMC40LTEzLjMtOS4zIDIuMS0xNS4zIDgtMTcuNSAxNy40LTEuOCA3LjggMiAxNS43IDkuMiAxOC45IDUuNSAyLjQgMTEgMi4xIDE2LjMtLjYgNy45LTQuMSAxMi4yLTEwLjUgMTIuNy0xOS4xeiIvPjwvc3ZnPg==
[codecov-badge]: https://codecov.io/gh/zchee/appleopensource/branch/master/graph/badge.svg?token=aVLPx1OMvZ
<!-- [codecov-badge]: https://img.shields.io/codecov/c/github/zchee/appleopensource/master.svg?style=for-the-badge&logo=codecov&token=aVLPx1OMvZ -->
[release-badge]: https://img.shields.io/github/release/zchee/appleopensource.svg?style=for-the-badge&logo=github
[ga-badge]: https://gh-ga-beacon.appspot.com/UA-89201129-1/zchee/appleopensource?useReferer&pixel
