
//var jq = jQuery.noConflict(true);
function genSsqBalls(){
	return tryGenBalls(6,33).toString()+" +"+tryGenBalls(1,16).toString();
}

function genDltBalls(){
	return tryGenBalls(5,35).toString()+" +"+tryGenBalls(1,12).toString()+" "+tryGenBalls(1,12).toString();
}

//len 生成长度
//max 最大数
function tryGenBalls(len,max){
  var balls =new Array();
  while(balls.length<len){
	   var vl=rnd(1,max);
	   var index=balls.indexOf(vl);
	   if(index < 0){
		    balls.push(vl)
	   }
	  /// console.log("vl="+vl +" index="+index);	   
  }
  balls.sort(function(a,b){return a-b});
  return balls;  
}

//生成随机数
function rnd(n, m){
  var random = Math.floor(Math.random()*(m-n+1)+n);
  return random;
}


$(document).ready(function(){
  $("button").click(function(){
    var red= tryGenBalls(6,33);
    var blue= tryGenBalls(1,16);
    var ballsHtml = ""
    for (i = 0; i< red.length; i++){
      ballsHtml = ballsHtml + "<li>" + red[i] + "</li>"
    } 
    for (i = 0; i< blue.length; i++){
      ballsHtml = ballsHtml + "<li class=\"lQiu\">" + blue[i] + "</li>"
    }
    $("#kjh").html(ballsHtml)
  });
});

// $(".btn  btn-info").click(function(){//点击事件  

//   //console.log($(this).html());//打印日志
//   //console.log($(this).attr("value"));//打印日志
//   var balls=genSsqBalls();
//    //取当前选中 li 的显示值 赋到input标签 选择器id="orderBy"
//    //$("#orderBy").html($(this).html());

//    //取当前选中 li 的value值 
//    var orderColumn = $(this).attr("value");
//    //赋到input标签 选择器id="orderColumn"
//    //$("#orderColumn").val($(this).attr("value"));

// });
