package inliner
import (
    
"net/http" 
"os"
"fmt"
"bufio" 
 "golang.org/x/net/html"
"strings"
"log"

"path"
)


func RenderToHttp(mainfile string,w http.ResponseWriter, r *http.Request) {
 _, err := os.Stat("./temp")
if err != nil {
    // no such file or dir
    //return
      os.Mkdir("./temp",0777);
}
    aa,_:=loadFile(mainfile);
//fmt.Println(aa);  
doc, err := html.Parse(strings.NewReader(aa))

if err != nil {
    log.Fatal(err)
}
var f func(*html.Node)
par:=html.Node{};
var src[] *html.Node
var lin[] *html.Node
f = func(n *html.Node) {
   var bb string
   
    
  
    
    if n.Type == html.ElementNode && strings.ToLower(n.Data) == "script" {
        
       
        for _, a := range n.Attr {
            if strings.ToLower(a.Key) == "src" {
               bb,_=loadFile(path.Dir(mainfile)+"/"+a.Val);
              
              
              
                sctr:=&html.Node{Data:"script",Type:html.ElementNode,Attr:[]html.Attribute{html.Attribute{Key:"type",Val:"text/javascript"}}}
                art:=&html.Node{Data:bb,Type:html.TextNode}
                sctr.AppendChild(art);
               src=append(src,sctr)
                break
            } 
        }
    }
    if n.Type == html.ElementNode && strings.ToLower(n.Data) == "link" {
        
       
        for _, a := range n.Attr {
            if strings.ToLower(a.Key) == "href" {
               bb,_=loadFile(path.Dir(mainfile)+"/"+a.Val);
              
            }
               if strings.ToLower(a.Key) == "rel" {
                sctr:=&html.Node{Data:"style",Type:html.ElementNode}
                art:=&html.Node{Data:bb,Type:html.TextNode}
                sctr.AppendChild(art);
               lin=append(lin,sctr)
                break
            } 
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
      
        f(c)
    }
}

par=*doc;
f(doc)
var body html.Node
var  head html.Node
ar:=*par.LastChild;
if strings.Compare(ar.Data,"html")==0{
 head=*ar.FirstChild   
body=*ar.LastChild

}else{
 body=ar;   
 head=*par.FirstChild
}
//html.Render(os.Stdout,&head) 
//fmt.Println(len(src));
j:=0;
adf:=0;
for c := head.FirstChild; c != nil&&j<len(src); c = c.NextSibling {
      //html.Render(os.Stdout,c) 
      if c.Type == html.ElementNode && strings.ToLower(c.Data) == "script" {
         
              for _, a := range c.Attr {
            if strings.ToLower(a.Key) == "src" {
               //fmt.Println(c.Parent)
                body.InsertBefore(src[j],c);
               c.Parent.RemoveChild(c);
               c=src[j];
                
                 j++;
                break
            } 
        }
           
       }
}
for c := head.FirstChild; c != nil&&adf<len(lin); c = c.NextSibling {
       if c.Type == html.ElementNode && strings.ToLower(c.Data) == "link" {
          
              for _, a := range c.Attr {
            if strings.ToLower(a.Key) == "rel" && strings.ToLower(a.Val) == "stylesheet" {
                head.InsertBefore(lin[adf],c);
               c.Parent.RemoveChild(c);
               c=src[adf];
               adf++;
                break
            } 
        }
           
       }
}
    
for c := body.FirstChild; c != nil; c = c.NextSibling {
    //  html.Render(os.Stdout,c) 
       if c.Type == html.ElementNode && strings.ToLower(c.Data) == "script" {
         
              for _, a := range c.Attr {
            if strings.ToLower(a.Key) == "src" {
               //fmt.Println(c.Parent)
                body.InsertBefore(src[j],c);
               c.Parent.RemoveChild(c);
               c=src[j];
                
                 j++;
                break
            } 
        }
           
       }
    }
    


//

   fil, err := os.OpenFile("./temp/myFile."+"html", os.O_WRONLY|os.O_CREATE, 0666)
  if err != nil {
          fmt.Println(err)
        //  os.Exit(1)
  }
  defer fil.Close();
html.Render(fil,&par) 
http.ServeFile(w,r,"./temp/myFile."+"html");

err3:=os.Remove("./temp/myFile."+"html");
if(err3!=nil){
fmt.Println(err3);
}

/** **/

}
func RenderToFile(mainfile string,a *os.File) {
     
    aa,_:=loadFile(mainfile);
//fmt.Println(aa);  
doc, err := html.Parse(strings.NewReader(aa))

if err != nil {
    log.Fatal(err)
}
var f func(*html.Node)
par:=html.Node{};
var src[] *html.Node
var lin[] *html.Node
f = func(n *html.Node) {
   var bb string
   
    
  
    
    if n.Type == html.ElementNode && strings.ToLower(n.Data) == "script" {
        
       
        for _, a := range n.Attr {
            if strings.ToLower(a.Key) == "src" {
               bb,_=loadFile(path.Dir(mainfile)+"/"+a.Val);
              
              
              
                sctr:=&html.Node{Data:"script",Type:html.ElementNode,Attr:[]html.Attribute{html.Attribute{Key:"type",Val:"text/javascript"}}}
                art:=&html.Node{Data:bb,Type:html.TextNode}
                sctr.AppendChild(art);
               src=append(src,sctr)
                break
            } 
        }
    }
    if n.Type == html.ElementNode && strings.ToLower(n.Data) == "link" {
        
       
        for _, a := range n.Attr {
            if strings.ToLower(a.Key) == "href" {
               bb,_=loadFile(path.Dir(mainfile)+"/"+a.Val);
              
            }
               if strings.ToLower(a.Key) == "rel" {
                sctr:=&html.Node{Data:"style",Type:html.ElementNode}
                art:=&html.Node{Data:bb,Type:html.TextNode}
                sctr.AppendChild(art);
               lin=append(lin,sctr)
                break
            } 
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
      
        f(c)
    }
}

par=*doc;
f(doc)
var body html.Node
var  head html.Node
ar:=*par.LastChild;
if strings.Compare(ar.Data,"html")==0{
 head=*ar.FirstChild   
body=*ar.LastChild

}else{
 body=ar;   
 head=*par.FirstChild
}
//html.Render(os.Stdout,&head) 
//fmt.Println(len(src));
j:=0;
adf:=0;
for c := head.FirstChild; c != nil&&j<len(src); c = c.NextSibling {
      //html.Render(os.Stdout,c) 
      if c.Type == html.ElementNode && strings.ToLower(c.Data) == "script" {
         
              for _, a := range c.Attr {
            if strings.ToLower(a.Key) == "src" {
               //fmt.Println(c.Parent)
                body.InsertBefore(src[j],c);
               c.Parent.RemoveChild(c);
               c=src[j];
                
                 j++;
                break
            } 
        }
           
       }
}
for c := head.FirstChild; c != nil&&adf<len(lin); c = c.NextSibling {
       if c.Type == html.ElementNode && strings.ToLower(c.Data) == "link" {
          
              for _, a := range c.Attr {
            if strings.ToLower(a.Key) == "rel" && strings.ToLower(a.Val) == "stylesheet" {
                head.InsertBefore(lin[adf],c);
               c.Parent.RemoveChild(c);
               c=src[adf];
               adf++;
                break
            } 
        }
           
       }
}
    
for c := body.FirstChild; c != nil; c = c.NextSibling {
    //  html.Render(os.Stdout,c) 
       if c.Type == html.ElementNode && strings.ToLower(c.Data) == "script" {
         
              for _, a := range c.Attr {
            if strings.ToLower(a.Key) == "src" {
               //fmt.Println(c.Parent)
                body.InsertBefore(src[j],c);
               c.Parent.RemoveChild(c);
               c=src[j];
                
                 j++;
                break
            } 
        }
           
       }
    }
    


//

html.Render(a,&par) 
/** **/

}

func loadFile(filepath string)(doc string,err error) {
   fil,err:=os.Open(filepath);
  if err != nil {
          fmt.Println(err)
        //  os.Exit(1)
  }
  defer fil.Close()

  fileInfo, _ := fil.Stat()
  var size int64 = fileInfo.Size()
  bytes := make([]byte, size)

  // read file into bytes
  //data:image/jpeg;
  buffer := bufio.NewReader(fil)
  _, err = buffer.Read(bytes)

  //filetype := http.DetectContentType(bytes)  
  // fmt.Println(filetype);
   doc=string(bytes); 
   return
}