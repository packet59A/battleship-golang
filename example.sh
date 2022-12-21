export GITHUB_TOKEN="X"

goreleaser release --rm-dist

read -rsp $'Press enter to continue...\n'