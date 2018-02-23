# _murex_ command reference

## try

> Exit block when error found. Similar in function to Bash's `set -e`


`try` forces a different execution behaviour where a failed process at the end
of a pipeline will cause the block to terminate regardless of any functions that
might follow.

    try {
        out: "Hello, World!" -> grep: "non-existent string"
        out: "This process will be ignored"
    }

A failure is determined by:

* Any process that returns a non-zero exit number
* Any process that returns more output via STDERR than it does via STDOUT

You can see which run mode your functions are executing under via the `fid-list`
command:


### See also

* trypipe
* evil
* catch
* fid-list