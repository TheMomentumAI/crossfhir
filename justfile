default:
  @just --list

deploy version desc:
  git tag -a v{{version}} -m "{{desc}}" && goreleaser release --clean
