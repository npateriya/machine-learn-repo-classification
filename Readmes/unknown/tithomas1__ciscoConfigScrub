# __Replace patterns in a Cisco configuration file__

Simple proof-of-concept for finding and replacing specified text patterns in a Cisco
configuration file. The patterns and associated replacement text are loaded from a
definitions file.

## Basic usage

```commandline
patternReplace.py -s SOURCE [-o OUTPUT] [-c CONFIG] [-v]
```

The OUTPUT defaults to `stdout`. The CONFIG defaults to `config.txt`. The `-v` flag will
trigger verbose output to `stderr`.

## Restrictions

The regex expressions in the config file should not overlap. They're matched in order,
but every pattern is tried against every line. There's no guarantee what the output would
be if overlapping patterns were matched.
