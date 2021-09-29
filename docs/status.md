# Status

## 29 Sept 2021

I've been working with OCaml and with `ocamllex` in [../parser](../parser) (not working right now). However, there's some issues:

- Overly complicated for what we need to do.
- Not simply to add new parts to the grammar.
- Limitation: once `ocamllex` returns a token, it automatically returns to the main entry point instead of going to the caller entry point. Losing this kind of context makes it very difficult to implement our grammar. [Read more.](https://medium.com/@huund/recipes-for-ocamllex-bb4efa0afe53)

I also tried using Go with [`participle`](https://github.com/alecthomas/participle). We can define the syntax as a tag of our data structure and it automatically fills the data structure. See the file [../parser-go/parser.go](../parser-go/parser.go) to understand what I mean. It would probably be very simple to add new contracts if we used Go. This version is correctly parsing bond agreements, while the OCaml version is simply not working yet.

If we want to keep using OCaml, new alternatives must be found: I asked on a forum ([see post](https://discuss.ocaml.org/t/define-literals-on-parser-using-ocamlyacc-menhir/8541)) and I was suggested to look into scannerless parsers. Both of the mentioned libraries ([ocaml-earley](https://github.com/rlepigre/ocaml-earley) and [dypgen](http://dypgen.free.fr/)) are either old or don't have much documentation.

Another option is to implement our own parser. This parser would go word by word, ignoring whitespaces and recursively calling methods that'd parse specific bits. This would probably overcomplicate the code and I'm not fluent in OCaml. I can try, but I don't think it'd be necessarily the best idea :)
