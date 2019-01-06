package docs

func init() {

	Definition["post"] = "# _murex_ Language Guide\n\n## Command Reference: `post`\n\n> HTTP POST request with a JSON-parsable return\n\n### Description\n\nFetches a page from a URL via HTTP/S POST request.\n\n### Usage\n\n    post url -> <stdout>\n    \n    <stdin> -> post url content-type -> <stdout>\n\n### Examples\n\n    » post google.com -> [ Status ] \n    {\n        \"Code\": 405,\n        \"Message\": \"Method Not Allowed\"\n    }\n\n### Detail\n\n#### JSON return\n\n`POST` returns a JSON object with the following fields:\n\n    {\n        \"Status\": {\n            \"Code\": integer,\n            \"Message\": string,\n        },\n        \"Headers\": {\n            string [\n                string...\n            ]\n        },\n        \"Body\": string\n    }\n    \nThe concept behind this is it provides and easier path for scripting eg pulling\nspecific fields via the index, `[`, function.\n\n#### `post` as a method\n\nRunning `post` as a method will transmit the contents of STDIN as part of the\nbody of the HTTP POST request. When run as a method you have to include a second\nparameter specifying the Content-Type MIME.\n\n#### Configurable options\n\n`post` has a number of behavioral options which can be configured via _murex_'s\nstandard `config` tool:\n\n    config: -> [ http ]\n    \nTo change a default, for example the user agent string:\n\n    config: set http user-agent \"bob\"\n    post: google.com\n    \nThis enables sane, repeatable and readable defaults. Read the documents on\n`config` for more details about it's usage and the rational behind the command.\n\n### See Also\n\n* [`get`](../docs/commands/get.md):\n  Makes a standard HTTP request and returns the result as a JSON object\n* [`getfile`](../docs/commands/getfile.md):\n  Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.\n* [config](../docs/commands/commands/config.md):\n  \n* [square-bracket-open](../docs/commands/commands/square-bracket-open.md):\n  "

}