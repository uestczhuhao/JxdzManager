function view(){
            a=document.getElementById("gra1").style.display;
            if(a=="none"){
                document.getElementById("gra1").style.display="block";
                document.getElementsByClassName("view")[0].style.backgroundImage="url(static/img/close.png)";
            }else{
                document.getElementById("gra1").style.display="none";
                document.getElementsByClassName("view")[0].style.backgroundImage="url(static/img/open.png)";
            }
        }
