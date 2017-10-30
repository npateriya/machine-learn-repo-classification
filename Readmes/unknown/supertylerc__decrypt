# decrypt [![Build Status](https://travis-ci.org/supertylerc/decrypt.svg?branch=master)](https://travis-ci.org/supertylerc/decrypt) [![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/supertylerc/decrypt/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/supertylerc/decrypt/?branch=master)

`decrypt` is a [Ruby][1] script for converting a given hashed password
to its plain-text equivalent.

## Version

The current version of `decrypt` is `0.0.1`.

## Installation

> Please read the installation details _completely_ before actually
> installing.

[Clone the repository][2].  Open a terminal and run the following
commands:

```bash
cd /path/to/repo
bundle install
rake install
```

> This installs `decrypt` to `$HOME/bin`.  If that directory doesn't
> exist, it creates it for you.  If it's not in [your path][3], you
> should add it.

If you want `decrypt` to be installed somewhere other than `$HOME/bin`,
you need to specify the installation directory like this:

```bash
rake install[/home/tyler/cows]
```

The rake task installs `decrypt` by creating a [symlink][4].  This
_will_ overwrite any existing file in the target directory called
`decrypt`!

Once you've done the above, simply invoke the script by typing:

```
decrypt -t [type]
```

[type] should be either `ios` (for Cisco Type 7 passwords) or `junos` 
(for `$9$` passwords).

## Examples

```bash
╭─tyler at deathstar in ~/bin using ‹ruby-2.1.1› 14-04-07 - 23:34:01
╰─○ ./decrypt -t ios
Paste hash, including leading characters, below.
> 02320C5E280918


Original hash: 320C5E280918
Plain-Text   : TheCow
╭─tyler at us160536 in ~/bin using ‹ruby-2.1.1› 14-04-07 - 23:34:03
╰─○ ./decrypt -t junos
Paste hash, including leading characters, below.
> $9$EDXheMdVYJZjqmpBEcMWdVwgJD.mT3/tevdsgJHkn/C


Original hash: $9$EDXheMdVYJZjqmpBEcMWdVwgJD.mT3/tevdsgJHkn/C
Plain-Text   : TheCowGoesOink
╭─tyler at deathstar in ~/bin using ‹ruby-2.1.1› 14-04-07 - 23:34:52
╰─○
```

## Author

The script is written by Tyler Christiansen.  You can contact him <a
href="mailto:tyler@oss-stack.io?GitHub - decrypt">via e-mail</a>,
[twitter][5], or check out [his blog][6].

## License

`decrypt` is released under [BSD 2-Clause License][8]. See the [LICENSE][10] file for more details.

[1]: https://www.ruby-lang.org/en/ "Ruby Language Home"
[2]: http://git-scm.com/book/en/Git-Basics-Getting-a-Git-Repository#Cloning-an-Existing-Repository "Clone an Existing Repository"
[3]: http://www.tech-recipes.com/rx/2621/os_x_change_path_environment_variable/ "Modify OS X Path"
[4]: http://gigaom.com/2011/04/27/how-to-create-and-use-symlinks-on-a-mac/ "Symlinking for Mac Users"
[5]: https://twitter.com/oss_stack "Tyler Christiansen's Twitter"
[6]: http://oss-stack.io/ "The Operations Supporting Systems Stack"
[8]: http://opensource.org/licenses/BSD-2-Clause "BSD 2-Clause Definition"
[10]: LICENSE "BSD 2-Clause License"
