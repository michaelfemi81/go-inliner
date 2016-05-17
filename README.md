# go-inliner
Converts javascripts and css in an html file to inline providing two options 
which are:
1)Piping to Http Stream
2)Creating a Bundled inline file

### INTRODUCTION

    go get github.com/michaelfemi81/go-inliner
### API DOCUMENTATION
 
RenderToHttp(mainfile string,w http.ResponseWriter, r *http.Request)


         
RenderToFile(mainfile string,a *os.File)