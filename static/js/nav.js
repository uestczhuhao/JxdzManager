function changeView(){
		var s=document.getElementById("searchText").style.display;
		if(s=="none"){
			document.getElementById("searchText").style.display="block";
		}else{
			document.getElementById("searchText").style.display="none";
			
		}
	}
	
	function drop(a,b){
		document.getElementById(a).style.height=b+"px";
	}