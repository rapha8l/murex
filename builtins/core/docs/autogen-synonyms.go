package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`echo`:            `out`,
	`!if`:             `if`,
	`!set`:            `set`,
	`(`:               `brace-quote`,
	`!or`:             `or`,
	`!catch`:          `catch`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!global`:         `global`,
	`!event`:          `event`,
	`!and`:            `and`,
	`export`:          `export`,
	`g`:               `g`,
	`and`:             `and`,
	`global`:          `global`,
	`event`:           `event`,
	`murex-docs`:      `murex-docs`,
	`if`:              `if`,
	`swivel-datatype`: `swivel-datatype`,
	`get`:             `get`,
	`err`:             `err`,
	`>`:               `>`,
	`f`:               `f`,
	`rx`:              `rx`,
	`alter`:           `alter`,
	`append`:          `append`,
	`post`:            `post`,
	`out`:             `out`,
	`>>`:              `>>`,
	`tread`:           `tread`,
	`trypipe`:         `trypipe`,
	`swivel-table`:    `swivel-table`,
	`getfile`:         `getfile`,
	`try`:             `try`,
	`set`:             `set`,
	`tout`:            `tout`,
	`or`:              `or`,
	`ttyfd`:           `ttyfd`,
	`prepend`:         `prepend`,
	`pt`:              `pt`,
	`read`:            `read`,
	`catch`:           `catch`,
	`brace-quote`:     `brace-quote`,
}
