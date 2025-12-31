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

diff --color=always <(go run . input/1.webp) <(hexdump input/1.webp)
diff --color=always <(go run . -C input/1.webp) <(hexdump -C input/1.webp)
diff --color=always <(go run . -v input/1.webp) <(hexdump -v input/1.webp)
diff --color=always <(go run . -C -v input/1.webp) <(hexdump -Cv input/1.webp)

diff --color=always <(go run . input/wikipedia.png) <(hexdump input/wikipedia.png)
diff --color=always <(go run . -C input/wikipedia.png) <(hexdump -C input/wikipedia.png)
diff --color=always <(go run . -v input/wikipedia.png) <(hexdump -v input/wikipedia.png)
diff --color=always <(go run . -C -v input/wikipedia.png) <(hexdump -Cv input/wikipedia.png)

diff --color=always <(go run . input/apple.jpg) <(hexdump input/apple.jpg)
diff --color=always <(go run . -C input/apple.jpg) <(hexdump -C input/apple.jpg)
diff --color=always <(go run . -v input/apple.jpg) <(hexdump -v input/apple.jpg)
diff --color=always <(go run . -C -v input/apple.jpg) <(hexdump -Cv input/apple.jpg)
