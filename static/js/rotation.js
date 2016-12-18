function getStyle(obj,attr){
		if(window.getComputedStyle){
			return getComputedStyle(obj,null)[attr];
		}else{
			return obj.currentStyle[attr];
		}
	}
	function indicator(){
		var name="ro";
		var id;
		for(i=1;i<6;i++){
			id=name+i;
			if(getStyle(document.getElementById(id),"opacity")==1){
				document.getElementsByClassName("indicator")[i-1].style.background="url(/static/img/btn-active.png)";
			}else{
				document.getElementsByClassName("indicator")[i-1].style.background="url(/static/img/btn-normal.png)";
			}
		}
	}
	var q=setInterval("indicator()",100);
	function changeImg(c){
		var n;
		for(i=1;i<6;i++){
			n=(i+c-1)%5+1;
			document.getElementById("ro"+n).style.animationName="rota"+i;
		}
	}
