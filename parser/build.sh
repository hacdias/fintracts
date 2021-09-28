rm lexer.ml parser.ml parser.mli

ocamllex lexer.mll       # generates lexer.ml
ocamlyacc parser.mly     # generates parser.ml and parser.mli

# ocamlc -c parser.mli
# ocamlc -c lexer.ml
# ocamlc -c parser.ml
# ocamlc -c compiler.ml
# ocamlc -o compiler lexer.cmo parser.cmo compiler.cmo
