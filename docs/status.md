# Status

## 29 Sept 2021

I've been working with OCaml and with `ocamllex` in [../parser](../parser) (not working right now). However, there's some issues:

- Overly complicated for what we need to do.
- Not simply to add new parts to the grammar.
- Limitation: once `ocamllex` returns a token, it automatically returns to the main entry point instead of going to the caller entry point. Losing this kind of context makes it very difficult to implement our grammar. [Read more.](https://medium.com/@huund/recipes-for-ocamllex-bb4efa0afe53)

I also tried using Go with [`participle`](https://github.com/alecthomas/participle). We can define the syntax as a tag of our data structure and it automatically fills the data structure. See the files [parser.go](../parser-go/parser.go),  [types.go](../parser-go/types.go) and  [agreements.go](../parser-go/agreements.go) to understand what I mean. It would probably be very simple to add new contracts if we used Go. Advantages:

- It is a recursive descent parser with backtracking.
- It works and I understand it, while the OCaml version is simply not working yet.
- Parses >= 2 parties, >= 2 signatures, >= 1 different agreements within the same contract.
- It is simple to add new contract types. Easy syntax.
- You annotate the parsing on the data structure, making it clear where the fields go.

If we want to keep using OCaml, new alternatives must be found: I asked on a forum ([see post](https://discuss.ocaml.org/t/define-literals-on-parser-using-ocamlyacc-menhir/8541)) and I was suggested to look into scannerless parsers. Both of the mentioned libraries ([ocaml-earley](https://github.com/rlepigre/ocaml-earley) and [dypgen](http://dypgen.free.fr/)) are either old or don't have much documentation.

Another option is to implement our own parser. This parser would go word by word, ignoring whitespaces and recursively calling methods that'd parse specific bits. This would probably overcomplicate the code and I'm not fluent in OCaml. I can try, but I don't think it'd be necessarily the best idea :)

### Complexity Example

I made an example to compare the parsing of:

> Signed by <Name> (, <Name>)* on the <Date>.

In both OCaml and Go using the libraries I found.

#### Go

main.go

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.Rule{
	{"Integer", `[-+]?(\d*\.)?\d+`, nil},
	{"Ident", `[a-zA-Z_]\w*`, nil},
	{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},
	{"eol", `[\n\r]+`, nil},
	{"whitespace", `[ \t]+`, nil},
})

var parser = participle.MustBuild(&Signature{},
	participle.Lexer(basicLexer),
	participle.UseLookahead(20),
)

func Parse(signature []byte) (*Signature, error) {
	ast := &Signature{}
	err := parser.ParseBytes("", signature, ast)
	return ast, err
}

type Date struct {
	Day   int    `parser:"@Integer" json:"day"`
	Month string `parser:"('th' | 'rd' | 'st') 'of' @Ident" json:"month"`
	Year  int    `parser:"@Integer" json:"year"`
}

type Signature struct {
	Parties []string `parser:"'Signed' 'by'  @Ident (',' @Ident)* 'and' @Ident" json:"parties"`
	Date    Date     `parser:"'on' 'the' @@ '.'" json:"date"`
}

func main() {
	file, err := ioutil.ReadFile("../signature.txt")
	if err != nil {
		log.Fatal(err)
	}

	signature, err := Parse(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bytes, err := json.MarshalIndent(signature, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(bytes))
}
```

### OCaml

fintracts.ml

```ocaml
open Types

module JSON = Yojson.Safe

let main =
  try
    let lexbuf = Lexing.from_channel stdin in
      let result = Parser.main Lexer.token lexbuf in
        let yo = signature_to_yojson result in
          print_endline (JSON.to_string yo);
  with Lexer.Eof ->
    exit 0
```

types.ml

```ocaml
type date = {
  day: int;
  month: string;
  year: int
} [@@deriving yojson]

type signature = {
  parties: string list;
  date: date
} [@@deriving yojson]
```

lexer.mll (we have to define tokens for every single word basically)

```
{
  open Parser
  exception Eof
}

let ws = [' ' '\t' '\n' '\r']*
let word = ['a'-'z''A'-'Z']+
let int = ['0'-'9']+

rule token = parse
  | ws                          { token lexbuf }
  | "Signed" ws "by"            { SIGNED_BY }
  | "on" ws "the"               { ON_THE }
  | "and"                       { AND }
  | "st"                        { DATE_SEP }
  | "th"                        { DATE_SEP }
  | "rd"                        { DATE_SEP }
  | "of"                        { OF }
  | ','                         { COMMA }
  | '.'                         { DOT }
  | word as w                   { WORD(w) }
  | int as i                    { INT(int_of_string i) }
  | eof                         { raise Eof }
```

parser.mly

```
%{
  open Types
%}

%token <string> WORD
%token <int> INT
%token SIGNED_BY
%token AND
%token ON_THE
%token DOT
%token COMMA
%token DATE_SEP
%token OF

%start main
%type <signature> main
%type <date> date
%type <string list> parties

%%

main
  : SIGNED_BY parties ON_THE date DOT { { parties = $2; date = $4 } }
;

parties
  : WORD parties                    { $1 :: $2  }
  | COMMA WORD parties              { $2 :: $3 }
  | AND WORD                        { [$2] }
;

date
  : INT DATE_SEP OF WORD INT { { day = $1; month = $4; year = $5 }}
;
```

Adding more things imply adding a lot new tokens and make the files barely readable. Besides, the ocamllex has a limit of tokens due to the automata reaching the maximum number of states. The alternative is to build a hash table but that also requires to separate words instead of constructions: for example "Signed by" (SIGNED_BY) would need to be separated in "Signed" (SIGNED) and "by" (BY). We'd need to almost manually define all the vocabulary. We can also generate the vocabulary from text templates: we give a text template and generate a table with all the words. It feels a bit painful, but possible.
