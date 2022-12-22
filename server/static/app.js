const origin = window.location.origin;

document.write("<h1>Hello world</h1>");
document.write(`<p>API server is running on ${origin}</p>`);

function callopenFile(){
    //title string,save bool,extention []string,initPath string
    result=openfile("open file",false,[],".").then(result=>{
        alert(result.File)
    })
}
function openDir(){
    result=opendir("open dir",".").then(result=>{
        alert(result.File)
    })
}