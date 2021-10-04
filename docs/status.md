# Status

## 4th October 2021

I've been trying to use OCaml with `ocamllex` to write the parser (see [../parser](../parser) - not currently working). However, I did not manage to make it work successfully. It seems to be:

- Overcomplicated as we need a token per each word. We cannot match thing in the parser with string literals.
- Not simple to add new grammar parts.
- Very limited support for [regular expressions](https://ocaml.org/manual/lexyacc.html#ss:ocamllex-regexp).
- Once `ocamllex` returns a token, it automatically returns to the main entry point instead of the caller entry point. Losing this context makes it more difficult to implement our grammar specifically. [Read more.](https://medium.com/@huund/recipes-for-ocamllex-bb4efa0afe53)

I tried using Go with [`participle`](https://github.com/alecthomas/participle). We can define the syntax as a tag of our data structure and it automatically fills the data structure. See the files [parser.go](../parser-go/parser.go),  [types.go](../parser-go/types.go) and  [agreements.go](../parser-go/agreements.go) to understand what I mean. It would probably be very simple to add new contracts if we used Go. Advantages:

- It is a recursive descent parser with backtracking.
- It works and I understand it, while the OCaml version is simply not working yet.
- Parses >= 2 parties, >= 2 signatures, >= 1 different agreements within the same contract.
- It is simple to add new contract types. Easy syntax.
- You annotate the parsing on the data structure, making it clear where the fields go.

### Ways to Go

#### Keep OCaml

1. Use another parser. I asked on a forum [see post](https://discuss.ocaml.org/t/define-literals-on-parser-using-ocamlyacc-menhir/8541)), but both of the mentioned libraries ([ocaml-earley](https://github.com/rlepigre/ocaml-earley) and [dypgen](http://dypgen.free.fr/)) are either old or don't have much documentation. This would require investigating more the field. I could not find many options.
2. Write our own parser. Probably over complicating things.
3. Keep using `ocamllex`. In this case, we would have to add a token per each **word** (the, a, contract, this, by, ...) and make it case-insensitive (i.e., not differentiate "and" and "And", as both would yield the token 'AND'). That would be a lot of tokens and we'd need to be careful not to reach the maximum size of the generated automata. We could also try to generate the vocabulary from a text template. Even though this would make it easier for us to add new contracts, it won't help with the underlying issues.

#### Use Go

- It is much simpler.
- Very easy to add new contract types.
- Easy language to learn.
- I already implemented it for the bond purchase agreement.

### Complexity Example

I made an example to compare the parsing of:

> Signed by \<Name\> (, \<Name\>)* on the <Date>.

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
