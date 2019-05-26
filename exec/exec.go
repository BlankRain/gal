package exec

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BlankRain/gal/evaluator"
	"github.com/BlankRain/gal/lexer"
	"github.com/BlankRain/gal/object"
	"github.com/BlankRain/gal/parser"
)

func Run(fileName string) {
	env := object.NewEnvironment()
	f, e := os.Open(fileName)
	if e != nil {
		fmt.Printf("error %v \n", e)
		return
	}
	b, e := ioutil.ReadAll(f)
	if e != nil {
		fmt.Printf("error %v", e)
		return
	}
	l := lexer.New(string(b))
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			fmt.Printf("%v", msg)
		}
	}
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}

}
