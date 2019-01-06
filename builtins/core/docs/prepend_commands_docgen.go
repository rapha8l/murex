package docs

func init() {

	Definition["prepend"] = "# _murex_ Language Guide\n\n## Command Reference: `prepend` \n\n> Add data to the start of an array\n\n### Description\n\n`prepend` a data to the start of an array.\n\n### Usage\n\n    <stdin> -> prepend: value -> <stdout>\n\n### Examples\n\n    » a: [January..December] -> prepend: 'New Year'\n    New Year\n    January\n    February\n    March\n    April\n    May\n    June\n    July\n    August\n    September\n    October\n    November\n    December\n\n### See Also\n\n* [`append`](../docs/commands/append.md):\n  Add data to the end of an array\n* [a](../docs/commands/commands/a.md):\n  \n* [cast](../docs/commands/commands/cast.md):\n  \n* [square-bracket-open](../docs/commands/commands/square-bracket-open.md):\n  \n* [update](../docs/commands/commands/update.md):\n  "

}