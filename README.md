# gex - (G)o h(ex) dump

A simple [hex dump](https://en.wikipedia.org/wiki/Hex_dump) utility.
The behavior of this program mimics `hexdump`:

```
$ go run . -C -v input.bin
00000000  30 31 32 33 34 35 36 37  38 39 41 42 43 44 45 46  |0123456789ABCDEF|
00000010  0a 2f 2a 20 2a 2a 2a 2a  2a 2a 2a 2a 2a 2a 2a 2a  |./* ************|
00000020  2a 2a 2a 2a 2a 2a 2a 2a  2a 2a 2a 2a 2a 2a 2a 2a  |****************|
00000030  2a 2a 2a 2a 2a 2a 2a 2a  2a 2a 2a 2a 2a 2a 2a 2a  |****************|
00000040  2a 2a 20 2a 2f 0a 09 54  61 62 6c 65 20 77 69 74  |** */..Table wit|
00000050  68 20 54 41 42 73 20 28  30 39 29 0a 09 31 09 32  |h TABs (09)..1.2|
00000060  09 33 0a 09 33 2e 31 34  09 36 2e 32 38 09 39 2e  |.3..3.14.6.28.9.|
00000070  34 32 0a                                          |42.|
00000073
```
