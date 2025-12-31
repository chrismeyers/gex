#/usr/bin/env bash

set -x
set +o posix

diff --color=always <(go run . input/sample.txt) <(hexdump input/sample.txt)
diff --color=always <(go run . -c input/sample.txt) <(hexdump -c input/sample.txt)
diff --color=always <(go run . -C input/sample.txt) <(hexdump -C input/sample.txt)
diff --color=always <(go run . -v input/sample.txt) <(hexdump -v input/sample.txt)
diff --color=always <(go run . -C -v input/sample.txt) <(hexdump -Cv input/sample.txt)

diff --color=always <(go run . input/repeats.txt) <(hexdump input/repeats.txt)
diff --color=always <(go run . -c input/repeats.txt) <(hexdump -c input/repeats.txt)
diff --color=always <(go run . -C input/repeats.txt) <(hexdump -C input/repeats.txt)
diff --color=always <(go run . -v input/repeats.txt) <(hexdump -v input/repeats.txt)
diff --color=always <(go run . -C -v input/repeats.txt) <(hexdump -Cv input/repeats.txt)

diff --color=always <(go run . input/1.webp) <(hexdump input/1.webp)
# diff --color=always <(go run . -c input/1.webp) <(hexdump -c input/1.webp)
diff --color=always <(go run . -C input/1.webp) <(hexdump -C input/1.webp)
diff --color=always <(go run . -v input/1.webp) <(hexdump -v input/1.webp)
diff --color=always <(go run . -C -v input/1.webp) <(hexdump -Cv input/1.webp)

diff --color=always <(go run . input/wikipedia.png) <(hexdump input/wikipedia.png)
# diff --color=always <(go run . -c input/wikipedia.png) <(hexdump -c input/wikipedia.png)
diff --color=always <(go run . -C input/wikipedia.png) <(hexdump -C input/wikipedia.png)
diff --color=always <(go run . -v input/wikipedia.png) <(hexdump -v input/wikipedia.png)
diff --color=always <(go run . -C -v input/wikipedia.png) <(hexdump -Cv input/wikipedia.png)

diff --color=always <(go run . input/apple.jpg) <(hexdump input/apple.jpg)
# diff --color=always <(go run . -c input/apple.jpg) <(hexdump -c input/apple.jpg)
diff --color=always <(go run . -C input/apple.jpg) <(hexdump -C input/apple.jpg)
diff --color=always <(go run . -v input/apple.jpg) <(hexdump -v input/apple.jpg)
diff --color=always <(go run . -C -v input/apple.jpg) <(hexdump -Cv input/apple.jpg)
