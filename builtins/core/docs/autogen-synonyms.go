package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!export`:         `export`,
	`unset`:           `export`,
	`!set`:            `set`,
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!and`:            `and`,
	`!or`:             `or`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`!global`:         `global`,
	`and`:             `and`,
	`swivel-table`:    `swivel-table`,
	`murex-docs`:      `murex-docs`,
	`getfile`:         `getfile`,
	`f`:               `f`,
	`g`:               `g`,
	`ttyfd`:           `ttyfd`,
	`catch`:           `catch`,
	`err`:             `err`,
	`out`:             `out`,
	`read`:            `read`,
	`or`:              `or`,
	`set`:             `set`,
	`export`:          `export`,
	`event`:           `event`,
	`alter`:           `alter`,
	`>`:               `>`,
	`tread`:           `tread`,
	`if`:              `if`,
	`prepend`:         `prepend`,
	`post`:            `post`,
	`pt`:              `pt`,
	`try`:             `try`,
	`get`:             `get`,
	`brace-quote`:     `brace-quote`,
	`>>`:              `>>`,
	`rx`:              `rx`,
	`trypipe`:         `trypipe`,
	`append`:          `append`,
	`swivel-datatype`: `swivel-datatype`,
	`tout`:            `tout`,
	`global`:          `global`,
}
