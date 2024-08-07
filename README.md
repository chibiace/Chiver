# Chiver

VERSION file incrementor written in Go, should work on versions with the format **x.x.x**

## Usage:

```console
❯./chiver --help
Usage of chiver:
  -M    Major release x.0.0
  -m    Minor release 1.x.0
  -r    Revision 1.0.x

❯ ./chiver
Version: 419.5.7 -> 419.5.8

❯ ./chiver -r
Version: 419.5.8 -> 419.5.9

❯ ./chiver -m
Version: 419.5.9 -> 419.6.9

❯ ./chiver -M
Version: 419.6.9 -> 420.6.9
```
