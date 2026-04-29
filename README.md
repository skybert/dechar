# dechar

Decomposes characters and shows them in all their Unicode splendour.

# Usage
```perl
$ dechar <one or more characters>
```
or:
```text
$ echo <one or more characters> | dechar
```

Example:
```perl
$ dechar 👻f
Byte pos: 0 Code point: 128123 Name: GHOST Bytes: 4
Byte pos: 4 Code point: 102 Name: LATIN SMALL LETTER F Bytes: 1
```

Verbose mode:
```perl
$ dechar --verbose 👻f
Byte 0 has code point 128123 U+1F47B '👻'
Code point 128123 is encoded using 4 bytes
Code point 128123 is called GHOST
Byte 4 has code point 102 U+66 'f'
Code point 102 is encoded using 1 bytes
Code point 102 is called LATIN SMALL LETTER F
```

# Human made 💪

`dechar` is coded by a human.
