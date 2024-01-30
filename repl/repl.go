package repl
import(
	"bufio"
	"fmt"
	"io"
	"github.com/mrsicsasa/interpreter-in-go/token"
	"github.com/mrsicsasa/interpreter-in-go/lexer"
)
const PROMT=">>"
func Start(in io.Reader, out io.Writer){
	scanner:=bufio.NewScanner(in)
	
	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		lex := lexer.New(line) // Change variable name from l to lex
		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
	
}