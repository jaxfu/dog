# dog

Go rewrite of the GNU coreutil [cat](https://www.gnu.org/software/coreutils/manual/html_node/cat-invocation.html)

## Todo

- select start and end line

- cli flag parsing
  - required:
    - filename|'-f'|string
  - optional:
    - linenums|'-l'|bool
    - range|'-r'|texas range e.g. 3..7, 3.., ..7
