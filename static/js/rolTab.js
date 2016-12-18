	var num=0;
	var b=new Array();
	function del(){
		for(i=0;i<8;i++){
		a=document.getElementById("li"+i).style.top;
		b[i]=getNum(a);
		}
	}
	setTimeout("del()",1);

	function rolTab(){
		num=num+1;
		if(num<76){
			for(i=0;i<8;i++){
				b[i]=b[i]-1;
				if(b[i]==-75){
					p=(i+1)%8;
					document.getElementById("li"+i).innerHTML=document.getElementById("li"+p).innerHTML;
					document.getElementById("li"+i).style.top="525px";
					b[i]=525;
				}
				document.getElementById("li"+i).style.top=b[i]+"px";		
			}
		}else{
			if(num==150){
				num=0;
			}
		}
	}
	function getNum(text){
		var value = text.replace(/[^0-9]/ig,""); 
		return value;
	}
	setInterval("rolTab()",10);