package main

import (
	"github.com/mvdan/sh/syntax"
	"os"
	"fmt"
)

// Example AST format
//0: *syntax.Stmt {
//.  .  .  .  Comments: []syntax.Comment (len = 0) {}
//.  .  .  .  Cmd: *syntax.DeclClause {
//.  .  .  .  .  Variant: *syntax.Lit {
//.  .  .  .  .  .  ValuePos: 1:1
//.  .  .  .  .  .  ValueEnd: 1:7
//.  .  .  .  .  .  Value: "export"
//.  .  .  .  .  }
//.  .  .  .  .  Opts: []*syntax.Word (len = 0) {}
//.  .  .  .  .  Assigns: []*syntax.Assign (len = 1) {
//.  .  .  .  .  .  0: *syntax.Assign {
//.  .  .  .  .  .  .  Append: false
//.  .  .  .  .  .  .  Naked: false
//.  .  .  .  .  .  .  Name: *syntax.Lit {
//.  .  .  .  .  .  .  .  ValuePos: 1:8
//.  .  .  .  .  .  .  .  ValueEnd: 1:27
//.  .  .  .  .  .  .  .  Value: "TZANALYTIC_USERNAME"
//.  .  .  .  .  .  .  }
//.  .  .  .  .  .  .  Index: nil
//.  .  .  .  .  .  .  Value: *syntax.Word {
//.  .  .  .  .  .  .  .  Parts: []syntax.WordPart (len = 1) {
//.  .  .  .  .  .  .  .  .  0: *syntax.Lit {
//.  .  .  .  .  .  .  .  .  .  ValuePos: 1:28
//.  .  .  .  .  .  .  .  .  .  ValueEnd: 1:44
//.  .  .  .  .  .  .  .  .  .  Value: "v-read-0qy0pw25x"
//.  .  .  .  .  .  .  .  .  }
//.  .  .  .  .  .  .  .  }
//.  .  .  .  .  .  .  }
//.  .  .  .  .  .  .  Array: nil
//.  .  .  .  .  .  }
//.  .  .  .  .  }
//.  .  .  .  }
//.  .  .  .  Position: 1:1
//.  .  .  .  Semicolon: 0:0
//.  .  .  .  Negated: false
//.  .  .  .  Background: false
//.  .  .  .  Coprocess: false
//.  .  .  .  Redirs: []*syntax.Redirect (len = 0) {}
//.  .  .  }

func createLit(value string) *syntax.Lit {
	return &syntax.Lit{
		Value: value,
	}
}

func createExportStatement(variableName, assignedValue string) *syntax.Stmt {
	return &syntax.Stmt{
		Cmd: &syntax.DeclClause{
			Variant: createLit("export"),
			Assigns: []*syntax.Assign{
				{
					Name:  createLit(variableName),
					Value: createWord(assignedValue),
				},
			},
		},
	}
}

func createWord(value string) *syntax.Word {
	return &syntax.Word{
		Parts: []syntax.WordPart{createLit(value)},
	}
}

type ShellFile struct {
	AST *syntax.File
	filepath string
}

func (sf *ShellFile) UpdateCredentials(usernameVariable, passwordVariable string, credentials *Credentials) {
	usernameUpdated := false
	passwordUpdated := false

	// Walk the tree to update existing username and password variables
	syntax.Walk(sf.AST, func(node syntax.Node) bool {
		switch x := node.(type) {
		case *syntax.DeclClause:
			if x.Variant.Value == "export" && len(x.Assigns) == 1 {
				if x.Assigns[0].Name.Value == usernameVariable {
					x.Assigns[0].Value = createWord(credentials.Username)
					usernameUpdated = true
				}

				if x.Assigns[0].Name.Value == passwordVariable {
					x.Assigns[0].Value = createWord(credentials.Password)
					passwordUpdated = true
				}

				return false
			}
		}
		return true
	})

	// Create additional environment variables at the end of the file if
	// they do not already exist
	if !usernameUpdated {
		sf.AST.Stmts = append(sf.AST.Stmts, createExportStatement(usernameVariable, credentials.Username))
	}

	if !passwordUpdated {
		sf.AST.Stmts = append(sf.AST.Stmts, createExportStatement(passwordVariable, credentials.Password))
	}
}

func (sf *ShellFile) WriteToDisk() error {
	file, err := os.OpenFile(sf.filepath, os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		return fmt.Errorf("could not open shell script '%s' for writing: %s", sf.filepath, err)
	}

	printer := syntax.NewPrinter()
	err = printer.Print(file, sf.AST)
	if err != nil {
		return fmt.Errorf("could not write AST to filepath '%s': %s", sf.filepath, err)
	}

	return nil
}

func NewShellFile(filepath string) (*ShellFile, error) {
	p := syntax.NewParser(syntax.KeepComments, syntax.Variant(syntax.LangBash))
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	syntaxTree, err := p.Parse(file, filepath)
	if err != nil {
		return nil, err
	}

	shellFile := &ShellFile{
		AST: syntaxTree,
		filepath: filepath,
	}

	return shellFile, nil
}
