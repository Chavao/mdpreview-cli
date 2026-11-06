#!/bin/sh
set -eu

tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
make_cmd=$(command -v make)
mkdir -p "$tmp/tools" "$tmp/bin" "$tmp/gopath/bin" "$tmp/other"
for cmd in ln rm chmod cat readlink; do
  ln -s "$(command -v "$cmd")" "$tmp/tools/$cmd"
done
cat > "$tmp/tools/go" <<'EOF'
#!/bin/sh
case "$1 $2" in
  'env GOBIN') printf '%s\n' "$TEST_GOBIN" ;;
  'env GOPATH') printf '%s\n' "$TEST_GOPATH" ;;
  install*) : > "${TEST_GOBIN:-$TEST_GOPATH/bin}/mdpreview-cli" ;;
  *) exit 1 ;;
esac
EOF
chmod +x "$tmp/tools/go"

export TEST_GOBIN="$tmp/bin" TEST_GOPATH="$tmp/gopath" PATH="$tmp/tools"
"$make_cmd" --no-print-directory install
test "$(readlink "$tmp/bin/mdc")" = mdpreview-cli

rm "$tmp/bin/mdc"
printf '#!/bin/sh\n' > "$tmp/other/mdc"
chmod +x "$tmp/other/mdc"
PATH="$tmp/other:$PATH" "$make_cmd" --no-print-directory install
test ! -e "$tmp/bin/mdc"

printf 'preserve' > "$tmp/bin/mdc"
"$make_cmd" --no-print-directory install
test "$(cat "$tmp/bin/mdc")" = preserve

TEST_GOBIN= "$make_cmd" --no-print-directory install
test "$(readlink "$tmp/gopath/bin/mdc")" = mdpreview-cli
