# js-lexer

```sh
$ cat main.go
package main

import (
        "fmt"

        "github.com/morinokami/js-lexer/lexer"
)

func main() {
        input := `
function map(f, a) {
  let result = []; // Create a new Array
  let i; // Declare variable
  for (i = 0; i != a.length; i++)
    result[i] = f(a[i]);
  return result;
}
`

        l := lexer.New(input)

        tok := l.NextToken()
        for {
                fmt.Printf("%+v\n", tok)
                tok = l.NextToken()
                if tok.Literal == "" {
                        break
                }
        }
}

$ go run main.go
{Type:{Label:function} Literal:function Loc:{Start:{Line:1 Column:0} End:{Line:1 Column:8}}}                                                                                                                                                                                
{Type:{Label:identifier} Literal:map Loc:{Start:{Line:1 Column:9} End:{Line:1 Column:12}}}                                                                                                                                                                                  
{Type:{Label:(} Literal:( Loc:{Start:{Line:1 Column:12} End:{Line:1 Column:13}}}                                                                                                                                                                                            
{Type:{Label:identifier} Literal:f Loc:{Start:{Line:1 Column:13} End:{Line:1 Column:14}}}                                                                                                                                                                                   
{Type:{Label:,} Literal:, Loc:{Start:{Line:1 Column:14} End:{Line:1 Column:15}}}                                                                                                                                                                                            
{Type:{Label:identifier} Literal:a Loc:{Start:{Line:1 Column:16} End:{Line:1 Column:17}}}                                                                                                                                                                                   
{Type:{Label:)} Literal:) Loc:{Start:{Line:1 Column:17} End:{Line:1 Column:18}}}                                                                                                                                                                                            
{Type:{Label:{} Literal:{ Loc:{Start:{Line:1 Column:19} End:{Line:1 Column:20}}}                                                                                                                                                                                            
{Type:{Label:identifier} Literal:let Loc:{Start:{Line:2 Column:2} End:{Line:2 Column:5}}}                                                                                                                                                                                   
{Type:{Label:identifier} Literal:result Loc:{Start:{Line:2 Column:6} End:{Line:2 Column:12}}}                                                                                                                                                                               
{Type:{Label:=} Literal:= Loc:{Start:{Line:2 Column:13} End:{Line:2 Column:14}}}                                                                                                                                                                                            
{Type:{Label:[} Literal:[ Loc:{Start:{Line:2 Column:15} End:{Line:2 Column:16}}}                                                                                                                                                                                            
{Type:{Label:]} Literal:] Loc:{Start:{Line:2 Column:16} End:{Line:2 Column:17}}}                                                                                                                                                                                            
{Type:{Label:;} Literal:; Loc:{Start:{Line:2 Column:17} End:{Line:2 Column:18}}}                                                                                                                                                                                            
{Type:{Label:identifier} Literal:let Loc:{Start:{Line:3 Column:2} End:{Line:3 Column:5}}}                                                                                                                                                                                   
{Type:{Label:identifier} Literal:i Loc:{Start:{Line:3 Column:6} End:{Line:3 Column:7}}}                                                                                                                                                                                     
{Type:{Label:;} Literal:; Loc:{Start:{Line:3 Column:7} End:{Line:3 Column:8}}}                                                                                                                                                                                              
{Type:{Label:for} Literal:for Loc:{Start:{Line:4 Column:2} End:{Line:4 Column:5}}}                                                                                                                                                                                          
{Type:{Label:(} Literal:( Loc:{Start:{Line:4 Column:6} End:{Line:4 Column:7}}}                                                                                                                                                                                              
{Type:{Label:identifier} Literal:i Loc:{Start:{Line:4 Column:7} End:{Line:4 Column:8}}}                                                                                                                                                                                     
{Type:{Label:=} Literal:= Loc:{Start:{Line:4 Column:9} End:{Line:4 Column:10}}}                                                                                                                                                                                             
{Type:{Label:numeric} Literal:0 Loc:{Start:{Line:4 Column:11} End:{Line:4 Column:12}}}                                                                                                                                                                                      
{Type:{Label:;} Literal:; Loc:{Start:{Line:4 Column:12} End:{Line:4 Column:13}}}                                                                                                                                                                                            
{Type:{Label:identifier} Literal:i Loc:{Start:{Line:4 Column:14} End:{Line:4 Column:15}}}
{Type:{Label:!=} Literal:!= Loc:{Start:{Line:4 Column:16} End:{Line:4 Column:18}}}
{Type:{Label:identifier} Literal:a Loc:{Start:{Line:4 Column:19} End:{Line:4 Column:20}}}
{Type:{Label:.} Literal:. Loc:{Start:{Line:4 Column:20} End:{Line:4 Column:21}}}
{Type:{Label:identifier} Literal:length Loc:{Start:{Line:4 Column:21} End:{Line:4 Column:27}}}
{Type:{Label:;} Literal:; Loc:{Start:{Line:4 Column:27} End:{Line:4 Column:28}}}
{Type:{Label:identifier} Literal:i Loc:{Start:{Line:4 Column:29} End:{Line:4 Column:30}}}
{Type:{Label:++} Literal:++ Loc:{Start:{Line:4 Column:30} End:{Line:4 Column:32}}}
{Type:{Label:)} Literal:) Loc:{Start:{Line:4 Column:32} End:{Line:4 Column:33}}}
{Type:{Label:identifier} Literal:result Loc:{Start:{Line:5 Column:4} End:{Line:5 Column:10}}}
{Type:{Label:[} Literal:[ Loc:{Start:{Line:5 Column:10} End:{Line:5 Column:11}}}
{Type:{Label:identifier} Literal:i Loc:{Start:{Line:5 Column:11} End:{Line:5 Column:12}}}
{Type:{Label:]} Literal:] Loc:{Start:{Line:5 Column:12} End:{Line:5 Column:13}}}
{Type:{Label:=} Literal:= Loc:{Start:{Line:5 Column:14} End:{Line:5 Column:15}}}
{Type:{Label:identifier} Literal:f Loc:{Start:{Line:5 Column:16} End:{Line:5 Column:17}}}
{Type:{Label:(} Literal:( Loc:{Start:{Line:5 Column:17} End:{Line:5 Column:18}}}
{Type:{Label:identifier} Literal:a Loc:{Start:{Line:5 Column:18} End:{Line:5 Column:19}}}
{Type:{Label:[} Literal:[ Loc:{Start:{Line:5 Column:19} End:{Line:5 Column:20}}}
{Type:{Label:identifier} Literal:i Loc:{Start:{Line:5 Column:20} End:{Line:5 Column:21}}}
{Type:{Label:]} Literal:] Loc:{Start:{Line:5 Column:21} End:{Line:5 Column:22}}}
{Type:{Label:)} Literal:) Loc:{Start:{Line:5 Column:22} End:{Line:5 Column:23}}}
{Type:{Label:;} Literal:; Loc:{Start:{Line:5 Column:23} End:{Line:5 Column:24}}}
{Type:{Label:return} Literal:return Loc:{Start:{Line:6 Column:2} End:{Line:6 Column:8}}}
{Type:{Label:identifier} Literal:result Loc:{Start:{Line:6 Column:9} End:{Line:6 Column:15}}}
{Type:{Label:;} Literal:; Loc:{Start:{Line:6 Column:15} End:{Line:6 Column:16}}}
{Type:{Label:}} Literal:} Loc:{Start:{Line:7 Column:0} End:{Line:7 Column:1}}}
```
