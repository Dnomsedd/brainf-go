import sys

ALLOWED = '+-<>,.[]'

src = open(sys.argv[1], 'r').read()
out = open(sys.argv[2], 'w')

for c in src:
    if c in ALLOWED:
        out.write(c)
