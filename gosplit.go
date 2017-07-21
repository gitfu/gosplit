

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

//var comments []string
//var code []string
var bufcomments bytes.Buffer
var bufcode bytes.Buffer

var theHead = `<!DOCTYPE html>
<html>
  <head>
    <meta content="text/html; charset=utf-8" http-equiv="content-type">
    <title>fu</title>
    <style type="text/css">
    body{ font-size:1em;}
  .commentbox{ left:2%; width:30%;}

    .codebox {right:2%; width:60%;}
    pre {position:absolute;
	top: 2%;
	bottom:2%;
	height: 80%;
	overflow-x: hidden;
	background:white;
	whitespace:prewrap;
	padding:2em;
	font-size: 1.3em;
	font-family: "Bitstream Vera";
	font-stretch: extra-expanded;
	text-shadow: 1px 2px 1px #808280;
}
    </style>
  </head>
  <body>`

var theTail = `
  </body>
  <script>
  var codebox=document.getElementsByClassName("codebox")[0]
  var commentbox=document.getElementsByClassName("commentbox")[0]
  function syncscroll(evt){
	codebox.scrollTop=evt.target.scrollTop
	commentbox.scrollTop=evt.target.scrollTop
}
codebox.onscroll=syncscroll
commentbox.onscroll=syncscroll
</script>	
</html>
`

func codeTag(instring string) string {
	return fmt.Sprintf("<code> %s</code><br>", instring)
}

func mkP(aline string) (string, string) {
	r := strings.NewReplacer("{", "<span style='color:lime'>{</span>", "}", "<span style='color:lime'>}</span>",
		" func ", "<span style='color:blue'> func </span>", " var ", "<span style='color:red'> var </span>")
	data = r.Replace(aline)
	if strings.Contains(data,'//'){
		splitLine := string.Split(data,'//')
		comment := codeTag(splitline[1])
		code := codeTag(splitline[0])
	} else {
		comment := codeTag(" ")
		code := codeTag(data)
	}	
	return comment,code
}

func main() {
	file, err := os.Open("./manifesto2.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bufcomments.WriteString("<pre class='commentbox'>")
	bufcode.WriteString("<pre class='codebox'>")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l, r := mkP(scanner.Text())
		bufcomments.WriteString(l)
		bufcode.WriteString(r)
	}
	bufcomments.WriteString(" </pre>")
	bufcode.WriteString("</pre>")
	fmt.Println(theHead)
	fmt.Println(bufcomments.String())
	fmt.Println(bufcode.String())
	fmt.Println(theTail)

}
