package docs

// Digest stores a 1 line summary of each builtins
var Digest map[string]string = map[string]string{
	`err`:             `Print a line to the STDERR`,
	`out`:             `'echo' a string to the STDOUT with a trailing new line character`,
	`>`:               `Writes STDIN to disk - overwriting contents if file already exists`,
	`>>`:              `Writes STDIN to disk - appending contents if file already exists`,
	`post`:            `HTTP POST request with a JSON-parsable return`,
	`brace-quote`:     `Write a string to the STDOUT without new line`,
	`ttyfd`:           `Returns the TTY device of the parent.`,
	`global`:          `Define a global variable and set it's value`,
	`alter`:           `Change a value within a structured data-type and pass that change along the pipeline without altering the original source input`,
	`murex-docs`:      `Displays the man pages for _murex_ builtins`,
	`rx`:              `Regexp pattern matching for file system objects (eg '.*\.txt')`,
	`tread`:           `'read' a line of input from the user and store as a user defined typed variable`,
	`try`:             `Handles errors inside a block of code`,
	`trypipe`:         `Checks state of each function in a pipeline and exits block on error`,
	`set`:             `Define a local variable and set it's value`,
	`event`:           `Event driven programming for shell scripts`,
	`append`:          `Add data to the end of an array`,
	`and`:             `Returns 'true' or 'false' depending on whether multiple conditions are met`,
	`if`:              `Conditional statement to execute different blocks of code depending on the result of the condition`,
	`pt`:              `Pipe telemetry. Writes data-types and bytes written`,
	`getfile`:         `Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.`,
	`prepend`:         `Add data to the start of an array`,
	`swivel-table`:    `Rotates a table by 90 degrees`,
	`read`:            `'read' a line of input from the user and store as a variable`,
	`swivel-datatype`: `Converts tabulated data into a map of values for serialised data-types such as JSON and YAML`,
	`or`:              `Returns 'true' or 'false' depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.`,
	`export`:          `Define an environmental variable and set it's value`,
	`tout`:            `Print a string to the STDOUT and set it's data-type`,
	`f`:               `Lists objects (eg files) in the current working directory`,
	`g`:               `Glob pattern matching for file system objects (eg *.txt)`,
	`catch`:           `Handles the exception code raised by 'try' or 'trypipe'`,
	`get`:             `Makes a standard HTTP request and returns the result as a JSON object`,
}
