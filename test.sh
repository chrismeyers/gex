#/usr/bin/env bash

set -x
set +o posix

diff --color=always <(go run . input.bin) <(hexdump input.bin)
diff --color=always <(go run . -C input.bin) <(hexdump -C input.bin)
diff --color=always <(go run . -v input.bin) <(hexdump -v input.bin)
diff --color=always <(go run . -C -v input.bin) <(hexdump -Cv input.bin)
