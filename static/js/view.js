function view(){
            a=document.getElementById("gra").style.display;
            if(a=="none"){
                document.getElementById("gra").style.display="block";
                document.getElementsByClassName("view")[0].style.backgroundImage="url(img/close.png)";
            }else{
                document.getElementById("gra").style.display="none";
                document.getElementsByClassName("view")[0].style.backgroundImage="url(img/open.png)";
            }
        }