# Language Guide: _murex_ Type System

_murex_ is a strictly typed shell with a few concessions for usability:

* Variables when outputted are automatically converted to strings

* `eval` and `let` functions evaluate the data type as well as the value.
An example of strict typing in `eval` can be seen with these 2 blocks:

```
set a=1  # define 'a' as string
let b=1  # define 'b' as number
eval a+b # returns '11' as 'a' is string so values are concatenated
```

```
let a=1  # define 'a' as number
let b=1  # define 'b' as number
eval a+b # returns '2' as both 'a' and 'b' are numbers
```

For more on the `set`, `let` and `eval` functions see [GUIDE.variables-and-evaluation.md](./GUIDE.variables-and-evaluation.md).

## Supported types

The types natively supported by this shell are:

* Generic   (defined: *)
* Null      (defined: null)
* Die       (defined: die)
* String    (defined: str)
* Boolean   (defined: bool)
* Number    (defined: num); this is the preferred type for numbers
* Integer   (defined: int)
* Float     (defined: float)
* Code Block (defined: block)
* JSON      (defined: json)
* CSV       (defined: csv)

Support for other mark ups such as XML and YAML will likely follow.
However JSON will always be a first class citizen because it is the
primary format for transmitting objects between methods (much like
Javascript's relationship with JSON).

#### Generic

This is used by methods to state they can accept any data type or output
any data type. Use of a `generic` input type can all define that a
method call also operate as a function (ie with no STDIN).

#### Null

This states that no data expected. Use `null` input to define functions
and/or `null` output to state the process doesn't write to STDOUT.

#### Die

If a `die` object is created it kills the shell.

#### Boolean

True or False. Generic input can be translated to boolean:

* 0 == False, none zero numbers == True
* "0" == False
* "null" == False
* "false" == False, "true" == True
* "no" == False, "yes" == True
* "off" == False, "on" == True
* "fail" == False, "pass" == True
* "failed" == False, "passed" == True
* "" == False, all other strings == True

Strings are not case sensitive when converted to boolean.

#### Number

This is the preferred (and default) method for storing numeric data. All
numbers are stored as a floating point value (in fact `float` and `num`
are one and the same data type internally).

#### Integer

As with normal programming languages, any number that doesn't have a
decimal point.

#### Float

A number which does have a decimal point.

This data type shouldn't ever be needed because its functionality is
duplicated by the default numeric data type, number (`num`).

#### Code Block

A sub-shell. This is used to inline code. eg for loops. Blocks are
encapsulated by curly brackets, `{}`.

#### JSON

A JSON object.

#### CSV

A comma separated list of records.