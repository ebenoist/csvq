CSVQ
---

A simple CSV tool

```BASH
NAME:
   csvq - A Simple CSV Tool

USAGE:
   cat my.csv | csvq

VERSION:
   0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --col, -c    Specify the columns to print by name or number
   --help, -h   show help
   --version, -v  print the version
```

## Examples
```BASH
$ cat routes.csv | csvq -c 1,2
route_short_name route_long_name
0                South Broadway
0L               South Broadway Limited
1                1st Avenue
10               East 12th Avenue
100              Kipling Street
```

```BASH
$ cat routes.csv | csvq -c route_short_name,route_long_name
route_short_name route_long_name
0                South Broadway
0L               South Broadway Limited
1                1st Avenue
10               East 12th Avenue
100              Kipling Street
```
