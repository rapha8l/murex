package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!set`:            `set`,
	`echo`:            `out`,
	`!and`:            `and`,
	`!if`:             `if`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!global`:         `global`,
	`(`:               `brace-quote`,
	`!or`:             `or`,
	`!catch`:          `catch`,
	`!event`:          `event`,
	`murex-docs`:      `murex-docs`,
	`err`:             `err`,
	`out`:             `out`,
	`set`:             `set`,
	`alter`:           `alter`,
	`append`:          `append`,
	`prepend`:         `prepend`,
	`read`:            `read`,
	`and`:             `and`,
	`global`:          `global`,
	`>`:               `>`,
	`pt`:              `pt`,
	`f`:               `f`,
	`event`:           `event`,
	`rx`:              `rx`,
	`try`:             `try`,
	`trypipe`:         `trypipe`,
	`swivel-datatype`: `swivel-datatype`,
	`get`:             `get`,
	`>>`:              `>>`,
	`tread`:           `tread`,
	`or`:              `or`,
	`if`:              `if`,
	`swivel-table`:    `swivel-table`,
	`g`:               `g`,
	`export`:          `export`,
	`getfile`:         `getfile`,
	`post`:            `post`,
	`brace-quote`:     `brace-quote`,
	`tout`:            `tout`,
	`ttyfd`:           `ttyfd`,
	`catch`:           `catch`,
}
