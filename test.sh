#/usr/bin/env bash

set -x
set +o posix

diff --color=always <(go run . input/wikipedia.txt) <(hexdump input/wikipedia.txt)
diff --color=always <(go run . -C input/wikipedia.txt) <(hexdump -C input/wikipedia.txt)
diff --color=always <(go run . -v input/wikipedia.txt) <(hexdump -v input/wikipedia.txt)
diff --color=always <(go run . -C -v input/wikipedia.txt) <(hexdump -Cv input/wikipedia.txt)

diff --color=always <(go run . input/repeats.txt) <(hexdump input/repeats.txt)
diff --color=always <(go run . -C input/repeats.txt) <(hexdump -C input/repeats.txt)
diff --color=always <(go run . -v input/repeats.txt) <(hexdump -v input/repeats.txt)
diff --color=always <(go run . -C -v input/repeats.txt) <(hexdump -Cv input/repeats.txt)
