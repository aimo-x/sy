//1541
var jle=[];// 记录元素信息
var jlm=[];// 存放要炒作的对象
var jle_i=0;
//计时器
var jsq_time=0;
var jls_score = 0;
window.jsq_f = function (){

    jsq_time=jsq_time+1;
    jsq_o.text=jsq_time;
  	score.text=jls_score;
    setTimeoutID = setTimeout("jsq_f()",1000)
}

//点击后执行函数 m 为点击的元素
window.dj_start = function(m){
  fpmp3.audio.play()
  // 透明图层 执行盖子，动画过程不让点击
  var b = scene.getObjectByName('tmtc').dom;
  b.style.zIndex=999;
  // 执行信息记录存放
  jl_e(m);
  // 相同元素
  var bd = if_eab();
  if(bd===0){
    // 复原所有状态
    recover_e();
    console.log('dj_start(m) and if_eab()=0');
  }else if(bd===1){
    // 执行当前元素动画并继续判断图案是否相同，复原状态或者得分进行下一步动画
    TweenMax.to(m, 0.15, {rotateY:Math.PI/2,onComplete:function(e){next_p(e,m)}});
    console.log('dj_start(m) and if_eab()=1');
  }else if(bd===2){
    // 执行当前元素动画
    TweenMax.to(m, 0.15, {rotateY:Math.PI/2, onComplete:function(e){next_n(e,m)}});
    console.log('dj_start(m) and if_eab()=2');
  }
}

// 翻转2_需要判断
window.next_p = function (e,m){
    // 设置元素的rotateY,以进行自然动画
    m.rotateY = -Math.PI/2;
    // 获取图片地址并且设置
    m.src = g_src(m.name);
    // 动画执行完进行判断
    TweenMax.to(m, 0.15, {rotateY:0,onComplete:if_p});
  console.log('next_p(m)');
  
}
//翻转2_直接执行无需判断
window.next_n = function (e,m){
    //设置元素的rotateY,以进行自然动画
    m.rotateY = -Math.PI/2;
    //获取图片地址并且设置
    m.src = g_src(m.name);
    //var e = scene.getObjectByName(m.name);
    //e.src = g_src(m.name);
    TweenMax.to(m, 0.15, {rotateY:0,onComplete:re_atc});
  console.log(m);
}
//翻转2动画执行完-进行判断图案是否相同
window.if_p = function (){
  if(jlm[jle[0]]==jlm[jle[1]]){// 相同
    bs_yes(jle[0],jle[1]);
    cgmp3.audio.play()
    console.log('if_p(m) 图案相同该加分',jle[0],jle[1]);
    jls_score=jls_score+1;
    score.text=jls_score;
    //大于16的视为通关
    if(jls_score>15){
      // 延迟 500 ms 动画完成执行逻辑
      setTimeout(function(){
        // 重置关卡
        gkname = "b";
        //游戏结束
        clearTimeout(setTimeoutID) // 清楚定时器
        jsmp3.audio.play()
        var s = scene.getObjectByName('time2') // 设置结束页的时间
        s.text=jsq_time;
        scene.gotoAndPause(1, 1); // 跳转到结束针
        SubmitScore() // 提交成绩
      },500);
    }

  }else{
    bs_err(jle[0],jle[1]);
    cwmp3.audio.play()
    console.log('if_p(m) 图案不相同',jle[0],jle[1]);
  }
}
// 执行透明图层盖着恢复
window.re_atc = function (){
  var b = scene.getObjectByName('tmtc').dom;
  b.style.zIndex=-1;
}
//记录元素信息 记录点击的元素
window.jl_e = function (m){
  jle[jle_i]=m.name;
  jlm[m.name]=g_src(m.name);
  jle_i=jle_i+1;
  console.log(' jl_e(m)');
}
//判断是否相同元素 
window.if_eab = function (){
  console.log(' if_eab()');
  //jle_i 不等于0
  if(jle_i>1){
    if(jle[0]==jle[1]){
      // 相同
      // 复原所有状态
      recover_e();
      return 0;
    }else{
      // 不相同
      // 执行当前元素动画
      jle_i=0;
      return 1;
    }
  }else{
    //执行当前元素动画
    return 2;
  }
}
// 相同时执行动画并且复原所有状态
function bs_yes(a,b){
  var ea = scene.getObjectByName(a);
  var eb = scene.getObjectByName(b);
  TweenMax.to(ea, 0.15, {scaleX:0,scaleY:0,x:ea.width/2+ea.left,y:ea.height/2+ea.top,onComplete:re_atc});
  TweenMax.to(eb, 0.15, {scaleX:0,scaleY:0,x:eb.width/2+eb.left,y:eb.height/2+eb.top});
  //复原所有状态
  recover_e();
  console.log(' bs_yes(a,b)',a,b);
    
}
// 不相同时执行动画并且复原所有状态
function bs_err(a,b){
  var ea = scene.getObjectByName(a);
  var eb = scene.getObjectByName(b);
  TweenMax.to(ea, 0.15, {rotateY:-Math.PI/2,onComplete:function(e){bs_last(e,ea)}});
  TweenMax.to(eb, 0.15, {rotateY:-Math.PI/2,onComplete:function(e){bs_last(e,eb)}});
  //复原所有状态
  recover_e();
	console.log(' bs_err(a,b)',a,b);
}
//错误后的动画重新盖上
function bs_last(e,g){
  var gksrc = scene.getObjectByName(gkname+"0").src
  g.src= gksrc
  //eb.src='https://www.mugeda.com/c/user/data/5748fa30a3664e1b3c0001d6/5a7ff7b5347a1955126c7a7c.png';
  g.rotateY=Math.PI/2;
  //eb.rotateY=Math.PI/2;
  console.log(g);
  //console.log(eb);
  TweenMax.to(g, 0.15, {rotateY:0,onComplete:re_atc});
  //TweenMax.to(eb, 0.15, {rotateY:0});
}
// 抓取图像地址
function g_src(name){
  console.log("g_src",name.slice(1));
  return y_s[name.slice(1)];
}
//复原所有状态
function recover_e(){
  jle=[];
  jlm=[];
  jle_i=0;
  console.log(' recover_e()');
}
//通关或者重玩重置函数
function sb_cw(){
  //scene.gotoPage(1, options);
  jls_score=0;
  score.text=jls_score;
  recover();
  console.log('开始重玩')
}
// 重置游戏关卡
function recover(){
  //获取所有牌
  var b = [];
  for(var i=0;i<16;i++){
    b[i] = scene.getObjectByName("a"+i);
    var gksrc = scene.getObjectByName(gkname+"0").src
    b[i].src =  gksrc;
    TweenMax.to(b[i], 0.15, {rotateY:o_a[i][2],scaleX:o_a[i][3],scaleY:o_a[i][4],x:o_a[i][0],y:o_a[i][1],onComplete:function(e){
      // 重置时间
      jsq_time=0;
      // 重置舞台时间
      jsq_o.text=jsq_time;
      re_atc()
    }});
  }
  	
}
//分配src
i_s=[];//图标地址数组
l_s=[];//预先加载数组
y_s=[];//映射图片数组
//加载图片函数
function load_img(){
  for(var i=1;i<7;i++){
    i_s[i] = 'https://cdns1.oovmi.com/2018/0211/'+i+'.png';
    l_s[i] = new Image();
    l_s[i].src = i_s[i];
  }
}

function randomMax(max){
  return Math.round(Math.random()*max)+1
}
// 映射关系
function ys(){
  /*
  y_s[0] = scene.getObjectByName(gkname+randomMax(5)).src
  y_s[1] = scene.getObjectByName(gkname+randomMax(5)).src
  y_s[2] = scene.getObjectByName(gkname+randomMax(5)).src
  y_s[3] = scene.getObjectByName(gkname+randomMax(5)).src
  y_s[4] = scene.getObjectByName(gkname+randomMax(5)).src
  y_s[5] = scene.getObjectByName(gkname+randomMax(5)).src
  */
  y_s[0] = scene.getObjectByName(gkname+"0").src
  y_s[1] = scene.getObjectByName(gkname+"1").src
  y_s[2] = scene.getObjectByName(gkname+"2").src
  y_s[3] = scene.getObjectByName(gkname+"3").src
  y_s[4] = scene.getObjectByName(gkname+"4").src
  y_s[5] = scene.getObjectByName(gkname+"5").src
  y_s[6] = scene.getObjectByName(gkname+"6").src
  y_s[7] = scene.getObjectByName(gkname+"7").src
  
  y_s[8] = y_s[0];
  y_s[9] = y_s[1];
  y_s[10] = y_s[2];
  y_s[11] = y_s[3];
  y_s[12] = y_s[4];
  y_s[13] = y_s[5];
  y_s[14] = y_s[6];
  y_s[15] = y_s[7];

  console.log(y_s)
  /*
  y_s[0]=i_s[Math.round(Math.random()*5)+1];
  y_s[1]=i_s[Math.round(Math.random()*5)+1];
  y_s[2]=i_s[Math.round(Math.random()*5)+1];
  y_s[3]=i_s[Math.round(Math.random()*5)+1];
  y_s[4]=i_s[Math.round(Math.random()*5)+1];
  y_s[5]=i_s[Math.round(Math.random()*5)+1];
  
  y_s[6]=y_s[0];
  y_s[7]=y_s[1];
  y_s[8]=y_s[2];
  y_s[9]=y_s[3];
  y_s[10]=y_s[4];
  y_s[11]=y_s[5];
  */
  y_s.sort(randomsort);
}
//打乱数组函数
function randomsort(a, b) {
  return Math.random()>0.5 ? -1 : 1;
  //用Math.random()函数生成0~1之间的随机数与0.5比较，返回-1或1
}
