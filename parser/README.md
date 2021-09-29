# `ocamellex` parser

Problems:

- When returning token from sub-state, it always returns to the main state. This is making it very difficult to work with. See https://medium.com/@huund/recipes-for-ocamllex-bb4efa0afe53

Solution:

- Implement our own parser? Use Go? See [../parser-go]