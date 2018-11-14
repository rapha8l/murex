package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`(`:               `brace-quote`,
	`!and`:            `and`,
	`!if`:             `if`,
	`!set`:            `set`,
	`!event`:          `event`,
	`echo`:            `out`,
	`!or`:             `or`,
	`!catch`:          `catch`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!global`:         `global`,
	`swivel-table`:    `swivel-table`,
	`get`:             `get`,
	`rx`:              `rx`,
	`if`:              `if`,
	`prepend`:         `prepend`,
	`catch`:           `catch`,
	`global`:          `global`,
	`out`:             `out`,
	`murex-docs`:      `murex-docs`,
	`getfile`:         `getfile`,
	`err`:             `err`,
	`ttyfd`:           `ttyfd`,
	`g`:               `g`,
	`append`:          `append`,
	`>`:               `>`,
	`>>`:              `>>`,
	`read`:            `read`,
	`trypipe`:         `trypipe`,
	`set`:             `set`,
	`event`:           `event`,
	`tout`:            `tout`,
	`swivel-datatype`: `swivel-datatype`,
	`post`:            `post`,
	`f`:               `f`,
	`and`:             `and`,
	`export`:          `export`,
	`alter`:           `alter`,
	`or`:              `or`,
	`tread`:           `tread`,
	`pt`:              `pt`,
	`try`:             `try`,
	`brace-quote`:     `brace-quote`,
}
