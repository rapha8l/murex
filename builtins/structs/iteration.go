package structs

import (
	"bytes"
	"errors"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"strings"
)

func init() {
	proc.GoFunctions["for"] = proc.GoFunction{Func: cmdFor, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["foreach"] = proc.GoFunction{Func: cmdForEach, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["while"] = proc.GoFunction{Func: cmdWhile, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["!while"] = proc.GoFunction{Func: cmdWhile, TypeIn: types.Null, TypeOut: types.Generic}
}

func cmdFor(p *proc.Process) (err error) {
	cblock, err := p.Parameters.Block(0)
	if err != nil {
		return err
	}

	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	parameters := strings.Split(string(cblock), ";")
	if len(parameters) != 3 {
		return errors.New("Invalid syntax. Must be { variable; conditional; incremental }")
	}

	variable := "let " + parameters[0]
	conditional := "eval " + parameters[1]
	incremental := "let " + parameters[2]

	_, err = lang.ProcessNewBlock([]rune(variable), nil, nil, p.Stderr, types.Null)
	if err != nil {
		return err
	}

	for {
		stdout := streams.NewStdin()
		i, err := lang.ProcessNewBlock([]rune(conditional), nil, stdout, p.Stderr, types.Null)
		stdout.Close()
		if err != nil {
			return err
		}

		if !types.IsTrue(stdout.ReadAll(), i) {
			return nil
		}

		lang.ProcessNewBlock(block, nil, p.Stdout, p.Stderr, types.Null)

		_, err = lang.ProcessNewBlock([]rune(incremental), nil, nil, p.Stderr, types.Null)
		if err != nil {
			return err
		}
	}

	return nil
}

func cmdForEach(p *proc.Process) (err error) {
	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	varName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	p.Stdin.ReadLineFunc(func(b []byte) {
		b = bytes.TrimSpace(b)
		if len(b) == 0 {
			return
		}

		proc.GlobalVars.Set(varName, string(b), p.Previous.ReturnType)

		stdin := streams.NewStdin()
		stdin.Writeln(b)
		stdin.Close()

		lang.ProcessNewBlock(block, stdin, p.Stdout, p.Stderr, p.Previous.Name)
	})

	return nil
}

func cmdWhile(p *proc.Process) error {
	switch p.Parameters.Len() {
	case 1:
		// Condition is taken from the while loop.
		block, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		for {
			stdout := streams.NewStdin()
			i, err := lang.ProcessNewBlock(block, nil, stdout, p.Stderr, types.Null)
			stdout.Close()
			if err != nil {
				return err
			}
			b := stdout.ReadAll()

			_, err = p.Stdout.Write(b)
			if err != nil {
				return err
			}

			conditional := types.IsTrue(b, i)

			if (!p.IsNot && !conditional) ||
				(p.IsNot && conditional) {
				return nil
			}

		}

	case 2:
		// Condition is first parameter, while loop is second.
		ifBlock, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		whileBlock, err := p.Parameters.Block(1)
		if err != nil {
			return err
		}

		for {
			stdout := streams.NewStdin()
			i, err := lang.ProcessNewBlock(ifBlock, nil, stdout, nil, types.Null)
			stdout.Close()
			if err != nil {
				return err
			}
			b := stdout.ReadAll()
			conditional := types.IsTrue(b, i)

			if (!p.IsNot && !conditional) ||
				(p.IsNot && conditional) {
				return nil
			}

			lang.ProcessNewBlock(whileBlock, nil, p.Stdout, p.Stderr, types.Null)
		}

	default:
		// Error
		return errors.New("Invalid number of parameters. Please read usage notes.")
	}

	return errors.New("cmdWhile(p *proc.Process) unexpected escaped a switch with default case.")
}