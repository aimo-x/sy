// 通用登陆代码开始
const domain = "https://iuu.pub"
const interfaceAddr = "https://iuu.pub/v2/api/fpgame/"
const domain_test = "127.0.0.1:9999"
const uuid = "779bcec2-587f-5cfa-8207-fc0289b750cc"
const CampaignID = 1 // 活动1
// const CampaignID = 2 // 活动2

// 通用登陆代码开始
const search = (variable) => {
    var arrStr = window.location.search.substring(1).split("&");
    for (var i = 0; i < arrStr.length; i++) {
        var temp = arrStr[i].split("=");
        if (temp[0] === variable){
            return decodeURIComponent(temp[1]);
        }  
    }
    return false;
};
window.setCookie = function (name,value,Days){
    // var Days = 30;
    var exp = new Date();
    exp.setTime(exp.getTime() + Days*24*60*60*1000);
    document.cookie = name + "="+ escape (value) + ";expires=" + exp.toGMTString();
};
window.getCookie = function (name){
    var arrStr = document.cookie.split("; ");
    for (var i = 0; i < arrStr.length; i++) {
    var temp = arrStr[i].split("=");
    if (temp[0] === name){
        if(decodeURIComponent(temp[1]) == ""){
        return false;
        }else{
        return decodeURIComponent(temp[1]);
        }
        
    }    
}
return false;
};
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOiJ3eDZlNTg1ZjFiODM5ZjI1Y2QiLCJleHAiOjE1OTcwODE5MjYsIm5iZiI6MTU4NDEyMTkyNiwibmlja19uYW1lIjoiNlptRzVMcVIiLCJvcGVuX2lkIjoib2RWWXkxWnUxRENtZkFxamREY19KbVJtdHVpSSIsInVuaW9uX2lkIjoiIn0.IRrIhxz2-IH1zSEZ8haY2AqetprorH_JDWbdXyzDL-w
Authorization_token = getCookie("authorization_token_fpgame_2")
if(search("test") == "true"){
    Authorization_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjaWQiOiIxIiwiZXhwIjoxNzE0ODgyNzAxLCJuYmYiOjE1ODUyODI3MDEsIm9pZCI6Im9kVll5MVp1MURDbWZBcWpkRGNfSm1SbXR1aUkiLCJ1aWQiOiIzIn0.BaxBzo-NhveX3K3gRPZd682Sw3gME_4BC52ZMAMK-VE"
    // Authorization_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBpZCI6Ind4NmU1ODVmMWI4MzlmMjVjZCIsImV4cCI6MTcxMzcyMzA5OSwibmJmIjoxNTg0MTIzMDk5LCJuaWNrbmFtZSI6IumZhuS6kSIsIm9wZW5pZCI6Im9kVll5MVp1MURDbWZBcWpkRGNfSm1SbXR1aUkiLCJ1bmlvbmlkIjoiIn0.-krpfBpdJhja139EzYwRTuc90AyLL7pbic2KdjJLw1s"
}
wechat_oauth = () => {
    // var reuri = window.location.origin+window.location.pathname;
    var reuri = window.location.href;
    fetch(domain+"/v2/api/wxmp/oauthurl?state="+encodeURIComponent(reuri)+"&CampaignID="+CampaignID)
    .then((res)=>res.json())
    .then((res)=>{ console.log("authorization_token_fpgame_2",res.data)
        if(res.code == "success"){
            window.location.href = res.data;  
        }else{
            // Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
            swal("授权失败", res.msg, "error")
        }
    })
    .catch((err)=>{
        swal("授权失败", err, "error")
    });
}
init_oauth = () => {
if(!Authorization_token){
    if(search("code")){ // 登录成功 设置token
    fetch(domain+"/v2/api/wxmp/authorization_token?code="+search("code")+"&CampaignID="+CampaignID)
    .then((res)=>res.json())
    .then((res)=>{ console.log(res.data)
        if(res.code == "success"){
            var token = res.data
            setCookie("authorization_token_fpgame_2",token,1/48);
            Authorization_token = token; // 赋值登录信息 
        }else{
            wechat_oauth();
        }
    })
    .catch((err)=>{
        console.log(err)
        swal("授权失败", err, "error")
    });
    
    }else{
        wechat_oauth();
    }
}
};
function getIsWxClient () {
    var ua = navigator.userAgent.toLowerCase();
    if (ua.match(/MicroMessenger/i) == "micromessenger") {
        return true;
    }
    return false;
};
if(search("test") == "true"){

}else{
    init_oauth();
}

// 通用登陆代码结束
var cfdjn343jn45n34 = "tqexAUG4Tb3EnfzCadhCwwIyvViy0YPh"
// 签名
function randomString(len) {
 　　len = len || 32;
 　　var $chars = 'ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678';    /****默认去掉了容易混淆的字符oOLl,9gq,Vv,Uu,I1****/
 　　var maxPos = $chars.length;
 　　var pwd = '';
 　　for (i = 0; i < len; i++) {
 　　　　pwd += $chars.charAt(Math.floor(Math.random() * maxPos));
 　　}
  　　return pwd;
 }


// 签名
// 游戏代码
var jle=[];// 记录元素信息
var jlm=[];// 存放要炒作的对象
var jle_i=0;
//计时器
var jsq_time=0;
window.jls_score = 0;
// 计时器开始
window.jsq_f = function (){
  jsq_time= Math.round((jsq_time+0.01)*100)/100
  jsq_o.text= "用时：" + jsq_time;
  setTimeoutID = setTimeout("jsq_f()",10)
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
  console.log("if_p()",jlm[jle[0]], jlm[jle[1]])
  if(jlm[jle[0]]==jlm[jle[1]]){// 相同
    bs_yes(jle[0],jle[1]);
    cgmp3.audio.play()
    // console.log('if_p(m) 图案相同该加分',jle[0],jle[1]);
    jls_score=jls_score+1;
    //大于15的视为通关
    if(jls_score>7){
      // 延迟 500 ms 动画完成执行逻辑
      setTimeout(function(){
        // 重置关卡
        gkname = "b";
        //游戏结束
        clearTimeout(setTimeoutID) // 清楚定时器
        jsmp3.audio.play()
        var s = scene.getObjectByName('time2') // 设置结束页的时间
        s.text=jsq_time;
        // scene.gotoAndPause(1, 1); // 跳转到结束针
        // SubmitScore() // 提交成绩
        putAchievement() // 提交成绩
      },500);
    }

  }else{
    bs_err(jle[0],jle[1]);
    cwmp3.audio.play()
    // console.log('if_p(m) 图案不相同',jle[0],jle[1]);
  }
}
// 执行透明图层盖着恢复
window.re_atc = function (){
  var b = scene.getObjectByName('tmtc').dom;
  b.style.zIndex=-1;
}
// 记录元素信息 记录点击的元素
window.jl_e = function (m){
  jle[jle_i]=m.name;
  jlm[m.name]=g_src(m.name);
  jle_i=jle_i+1;
  console.log(' jl_e(m)',m.name);
}
// 判断是否相同元素 
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
  tempFP = ""
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
  // var gksrc = scene.getObjectByName("a0").src
  g.src= gksrc
  //eb.src='https://www.mugeda.com/c/user/data/5748fa30a3664e1b3c0001d6/5a7ff7b5347a1955126c7a7c.png';
  g.rotateY=Math.PI/2;
  //eb.rotateY=Math.PI/2;
  console.log(g);
  //console.log(eb);
  TweenMax.to(g, 0.15, {rotateY:0,onComplete:re_atc});
  //TweenMax.to(eb, 0.15, {rotateY:0});
}
// 抓取翻牌图像地址
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
  //score.text=jls_score;
  recover();
  console.log('开始重玩')
}
// 重置游戏关卡
function recover(){
  //获取所有牌
  var b = [];
  for(var i=0;i<16;i++){
    b[i] = scene.getObjectByName("a"+i);
    // var gksrc = scene.getObjectByName("a0").src
    b[i].src = gksrc
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
  y_s[4] = scene.getObjectByName(gkname+"0").src
  y_s[5] = scene.getObjectByName(gkname+"1").src
  y_s[6] = scene.getObjectByName(gkname+"2").src
  y_s[7] = scene.getObjectByName(gkname+"3").src
  
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
 
var kkey =  cfdjn343jn45n34

window.GameCount =  0// 当日可用游戏次数
window.ShareCount = 0 // 当日可用分享次数
// 获取当日游戏可用信息
getNowDayUserGameInfo = () => {
  layer.open({
    type: 2
    ,shadeClose: false
    ,content: '获取游戏信息中...'
  });
  fetch(interfaceAddr +"getNowDayUserGameInfo"+"?CampaignID="+CampaignID,{
    headers: {
      'Authorization': Authorization_token,
    },
    method: 'GET', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, cors, *same-origin
    redirect: 'follow', // manual, *follow, error
    referrer: 'no-referrer', // *client, no-referrer
  })
  .then((res)=>res.json())
  .then((res)=>{
    
    if(res.code == "success"){
      /*
      CampaignID uint
      UserID     uint
      Count      int     // 今日可玩游戏次数
      Share      int     // 当日可用分享次数
      Time       float64 // 游戏完成时间
      */
      GameCount = res.data.FpGame.Count
      ShareCount =  res.data.FpGame.Share
      nickname = scene.getObjectByName("微信昵称")
      nickname.text = Base64.decode(res.data.User.NickName)
      wxtx = scene.getObjectByName("头像")
      wxtx.src = "https" + res.data.User.HeadImg.substring(4)
      if(GameCount<1){
        // Mugine.Utils.Toast.info(res.msg, {type:'今日游戏次数已用完,可通过分享游戏获得游戏次数'});
        layer.open({content:"今日游戏次数已用完，分享可以获得更多次游戏机会！",btn: ['好的']})
        // 无法开始游戏
        // scene.getObjectByName("开始游戏").scene.gotoAndPause(1)
      }else{
        // scene.getObjectByName("开始游戏").scene.gotoAndPause(0)
      }
      
    }else{
      Mugine.Utils.Toast.info(res.msg, {type:'info'});
      console.log(res)
    }
    console.log(res)
    // getQrCode()
    layer.closeAll()
  })
  .catch((err)=>{
    console.log(err)
    layer.closeAll()
    Mugine.Utils.Toast.info(err, {type:'info'});
  });
}
// getSelf
getSelf = () => {
  layer.open({
    type: 2
    ,shadeClose: false
    ,content: '等等返回数据...'
  });
  fetch(interfaceAddr +"getSelf"+"?CampaignID="+CampaignID,{
    headers: {
      'Authorization': Authorization_token,
    },
    method: 'GET', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, cors, *same-origin
    redirect: 'follow', // manual, *follow, error
    referrer: 'no-referrer', // *client, no-referrer
  })
  .then((res)=>res.json())
  .then((res)=>{
    
    if(res.code == "success"){
      // GameCount = res.data.FpGame.Count
      // ShareCount =  res.data.FpGame.Share
      // mingci 超越百分比
      mingci = scene.getObjectByName('mingci')
      BFB = (res.data.MC/res.data.N) *100
      mingci.text =  BFB.toFixed(1)+ "%"// Math.round(res.data.bfb/100)*10000 + "%"
      scene.gotoAndPause(1, 1); // 跳转到结束针
      // 超越了多少用户 Math.round(res.data.bfb/100)*100
    }else{
      Mugine.Utils.Toast.info(res.msg, {type:'info'});
      console.log(res)
    }
    layer.closeAll()
  })
  .catch((err)=>{
    layer.closeAll()
    Mugine.Utils.Toast.info(err, {type:'info'});
  });
}

// 提交成绩
const putAchievement = () => {
  layer.open({
    type: 2
    ,shadeClose: false
    ,content: '正在提交成绩...'
  });
  var body = {
    time: jsq_time, 
  }
  var noncestr = randomString(12)
  var timestamp = (new Date().getTime()/1000).toFixed()
  var dataStr =  JSON.stringify(body) // "{"+jsq_time+"}"
  console.log(JSON.stringify(body))
  var signStr = "data="+dataStr+"&noncestr="+noncestr+kkey+"&timestamp="+timestamp
  var sign = md5(signStr)
  fetch(interfaceAddr +"putAchievement"+"?CampaignID="+CampaignID+"&noncestr="+noncestr+"&timestamp="+timestamp+"&sign="+sign,{
    headers: {
      'Authorization': Authorization_token,
    },
    method: 'PUT', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, cors, *same-origin
    redirect: 'follow', // manual, *follow, error
    referrer: 'no-referrer', // *client, no-referrer
    body: JSON.stringify(body)
  })
  .then((res)=>res.json())
  .then((res)=>{
    
    if(res.code == "success"){
      
      mingci = scene.getObjectByName('mingci');
      BFB = (res.data.MC/res.data.N) *100;
      mingci.text =  BFB.toFixed(1)+ "%";// Math.round(res.data.bfb/100)*10000 + "%"
      scene.gotoAndPause(1, 1); // 跳转到结束针
      GameCount = res.data.FpGame.Count
    }else{
      Mugine.Utils.Toast.info(res.msg, {type:'info'});
      console.log(res)
    }
    // getSelf()
    layer.closeAll()
  })
  .catch((err)=>{
    layer.closeAll()
    Mugine.Utils.Toast.info(err, {type:'info'});
  });
}

testTJCJ = () => {
  putAchievement()
}
// getShareGameCount 分享增加游戏
getShareGameCount = () => {
  layer.open({
    type: 2
    ,shadeClose: false
    ,content: '正在查询服务器中...'
  });
  fetch(interfaceAddr +"getShareGameCount"+"?CampaignID="+CampaignID,{
    headers: {
      'Authorization': Authorization_token,
    },
    method: 'PUT', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, cors, *same-origin
    redirect: 'follow', // manual, *follow, error
    referrer: 'no-referrer', // *client, no-referrer
  })
  .then((res)=>res.json())
  .then((res)=>{
    
    if(res.code == "success"){
      GameCount = res.data.Count
      ShareCount =  res.data.Share
    }else{
      Mugine.Utils.Toast.info(res.msg, {type:'info'});
      console.log(res)
    }
    layer.closeAll()
  })
  .catch((err)=>{
    layer.closeAll()
    Mugine.Utils.Toast.info(err, {type:'info'});
  });
}
// 初始话排名相关
var phimg = [], phname = [], phtime = []
// getRank 总排行
getRank = () => {
  // 当前游戏排行
  for(var i = 0; i < 20; i++){
    var n = i+20
    phimg[i+20] = scene.getObjectByName('txa#'+n)
    // phimg[i+20].src = txaSrc
    phimg[i+20].alpha = 1
    phname[i+20] = scene.getObjectByName('xm#'+n)
    phname[i+20].alpha = 1
    // phname[i+20].text = xmText
    phtime[i+20] = scene.getObjectByName('ys#'+n)
    phtime[i+20].alpha = 1
    // phtime[i+20].text = ysText
  }

  layer.open({
    type: 2
    ,shadeClose: false
    ,content: '请稍后...'
  });
  fetch(interfaceAddr+"getRank"+"?CampaignID="+CampaignID)
  .then((res)=>res.json())
  .then((res)=>{
    layer.closeAll()
    if(res.code == "success"){
      if(res.data == "" || res.data == null || res.data == undefined || res.data.length == 0){
        scene.gotoAndPause(1, 2);
        layer.closeAll()
        return
      }
      /*
      CampaignID uint
      UserID     uint
      Time       float64
      */
      if(res.data.length<20){
        var tmh = 20-res.data.length
        for(var i = 0; i < res.data.length; i++){
          var n = i+20
          phimg[i+20] = scene.getObjectByName('txa#'+n)
          phimg[i+20].src = "https" + res.data[i].HeadImg.substring(4)
          phimg[i+20].alpha = 1
          phname[i+20] = scene.getObjectByName('xm#'+n)
          phname[i+20].alpha = 1
          phname[i+20].text = Base64.decode(res.data[i].NickName)
          phtime[i+20] = scene.getObjectByName('ys#'+n)
          phtime[i+20].alpha = 1
          phtime[i+20].text = res.data[i].Time
        }
        for(var x = res.data.length; x < tmh;x ++){
          var n = x+20
          phimg[i+20] = scene.getObjectByName('txa#'+n)
          // phimg[i+20].alpha =  0
          phname[i+20] = scene.getObjectByName('xm#'+n)
          // phname[i+20].alpha = 0
          phtime[i+20] = scene.getObjectByName('ys#'+n)
          //phtime[i+20].alpha = 0
        }
      }else{
        for(var i = 0; i < 20; i++){
          var n = i+20
          phimg[i+20] = scene.getObjectByName('txa#'+n)
          phimg[i+20].src = "https" + res.data[i].HeadImg.substring(4)
          phimg[i+20].alpha = 1
          phname[i+20] = scene.getObjectByName('xm#'+n)
          phname[i+20].alpha = 1
          phname[i+20].text = Base64.decode(res.data[i].NickName)
          phtime[i+20] = scene.getObjectByName('ys#'+n)
          phtime[i+20].alpha = 1
          phtime[i+20].text = res.data[i].Time
        }
      }
      scene.gotoAndPause(1, 2);
    }else{
      Mugine.Utils.Toast.info(res.msg, {type:'info'});
    }
  })
  .catch((err)=>{
    layer.closeAll()
    Mugine.Utils.Toast.info(err, {type:'info'});
  });
}
// 获取当日排行
getNowDayRank = () => {
  // 当前游戏排行
  for(var i = 0; i < 20; i++){
    var n = i
    phimg[i] = scene.getObjectByName('txa#'+n)
    // phimg[i].src = txaSrc
    phimg[i].alpha = 1
    phname[i] = scene.getObjectByName('xm#'+n)
    phname[i].alpha = 1
    // phname[i].text = xmText
    phtime[i] = scene.getObjectByName('ys#'+n)
    phtime[i].alpha = 1
    // phtime[i].text = ysText
  }

  layer.open({
    type: 2
    ,shadeClose: false
    ,content: '请稍后...'
  });
  fetch(interfaceAddr+"getNowDayRank"+"?CampaignID="+CampaignID)
  .then((res)=>res.json())
  .then((res)=>{
    layer.closeAll()
    if(res.code == "success"){
      if(res.data == "" || res.data == null || res.data == undefined || res.data.length == 0){
        scene.gotoAndPause(0, 2);
        layer.closeAll()
        return
      }
      /*
      CampaignID uint
      UserID     uint
      Time       float64
      */
      if(res.data.length<20){
        var tmh = 20-res.data.length
        for(var i = 0; i < res.data.length; i++){
          var n = i
          phimg[i] = scene.getObjectByName('txa#'+n)
          phimg[i].src = "https" + res.data[i].HeadImg.substring(4)
          phimg[i].alpha = 1
          phname[i] = scene.getObjectByName('xm#'+n)
          phname[i].alpha = 1
          phname[i].text = Base64.decode(res.data[i].NickName)
          phtime[i] = scene.getObjectByName('ys#'+n)
          phtime[i].alpha = 1
          phtime[i].text = res.data[i].Time
        }
        for(var x = res.data.length; x < tmh;x ++){
          var n = x
          phimg[i] = scene.getObjectByName('txa#'+n)
          // phimg[i].alpha =  0
          phname[i] = scene.getObjectByName('xm#'+n)
          // phname[i].alpha = 0
          phtime[i] = scene.getObjectByName('ys#'+n)
          //phtime[i].alpha = 0
        }
      }else{
        for(var i = 0; i < 20; i++){
          var n = i
          phimg[i] = scene.getObjectByName('txa#'+n)
          phimg[i].src = "https" + res.data[i].HeadImg.substring(4)
          phimg[i].alpha = 1
          phname[i] = scene.getObjectByName('xm#'+n)
          phname[i].alpha = 1
          phname[i].text = Base64.decode(res.data[i].NickName)
          phtime[i] = scene.getObjectByName('ys#'+n)
          phtime[i].alpha = 1
          phtime[i].text = res.data[i].Time
        }
      }
      // 获取成功
      scene.gotoAndPause(0, 2);
    }else{
      Mugine.Utils.Toast.info(res.msg, {type:'info'});
    }
  })
  .catch((err)=>{
    layer.closeAll()
    Mugine.Utils.Toast.info(err, {type:'info'});
  });
}
// 生成二维码
getQrCode = () => {
  layer.open({
    type: 2
    ,shadeClose: false
    ,content: '正在查询服务器中...'
  });
  fetch(interfaceAddr +"getQrCode"+"?CampaignID="+CampaignID,{
    headers: {
      'Authorization': Authorization_token,
    },
    method: 'GET', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, cors, *same-origin
    redirect: 'follow', // manual, *follow, error
    referrer: 'no-referrer', // *client, no-referrer
  })
  .then((res)=>res.json())
  .then((res)=>{
    if(res.code == "success"){
      qr = scene.getObjectByName('二维码')
      qr.src = "data:image/png;base64,"+res.data
      scene.gotoAndPause(2, 0)
    }else{
      Mugine.Utils.Toast.info(res.msg, {type:'info'});
      
    }
    console.log(res)
    layer.closeAll()
  })
  .catch((err)=>{
    console.log(err)
    layer.closeAll()
    Mugine.Utils.Toast.info(err, {type:'info'});
  });
}
// https://iuu.pub/api/qr?url=asas
// goGame 开始游戏按钮
goGame = () => {
  // 需要判断是否可以参与游戏
  // 参数初始化
  jle=[];// 记录元素信息
  jlm=[];// 存放要炒作的对象
  jle_i=0;
  //计时器
  jsq_time=0;
window.jls_score = 0;
  jsq_f()
  if(GameCount<1){
    layer.open({content:"今日游戏次数已用完，分享可以获得更多次游戏机会！",btn: ['好的']})
    return
  }else{
    sb_cw()
    scene.nextPage()
  }
}



window.gkArr = ["b","b","b","b","b","b"]
window.gkid = null;
window.gkname = "b";

mugeda.addEventListener("renderready", function(){
  if(search('console')=="open"){
    var vConsole = new VConsole();
    console.log('Hello world');
  }
  
  // 当动画准备完成，开始播放前的那一刻引发回调。
  scene = mugeda.scene
  // 分享获取游戏次数
  defineWechatParameters({
    "success_share_callback": function(){
      getShareGameCount()
    },
  });
  
  getNowDayUserGameInfo()
  // 默认的盖子SRC
  gksrc = scene.getObjectByName("a0").src;
  //加载图片函数
  // load_img();
  //执行映射关系
  ys();
  a = [];
  // 存放原始数据
  o_a=[];
  // 初始化关卡
  for(var i=0;i<16;i++){
    // var gksrc = scene.getObjectByName("a0").src
    a[i] = scene.getObjectByName("a"+i);
    a[i].src = gksrc
    o_a[i] = [a[i].x, a[i].y, a[i].rotateY, a[i].scaleX,a[i].scaleY];
  }

  jsq_o = scene.getObjectByName("time");
  // score = scene.getObjectByName("score");
  
  cgmp3 = scene.getObjectByName("cgmp3");
  cwmp3 = scene.getObjectByName("cwmp3");
  jsmp3 = scene.getObjectByName("jsmp3");
  fpmp3 = scene.getObjectByName("fpmp3");

});

window.start_game=function(){
    // 计时器
    jsq_f();
};
// 重玩一次
window.re_start=function(){
    // 设置重置属性
    jls_score = 0
    gkname = gkArr[0]
    // 重新排列
    ys()
    // scene.gotoPage(1);
    scene.gotoAndPause(0, 0)
    // sb_cw(); 
};
// 点击翻牌
var tempFP = ""
window.c=function(m){
  if(m.name == tempFP){
    // 不做响应
    return
  }
  tempFP = m.name
  dj_start(m);
};

window.check_up=function(){
  /*
  var fs_p = scene.getObjectByName("score").text;
  if(fs_p>5){
  }else{
    swal({ 
      title: '分数过低', 
      text: '分数达到6分即可抽奖', 
      type: 'warning',
      showCancelButton: true, 
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: '重玩', 
      cancelButtonText: '取消', 
    }).then(function(){
    scene.prevPage();
      sb_cw();
    });
    
  }*/
};

window.tishi=function(page){
  Mugine.Utils.Toast.info('活动已经结束',{type:'toast'});
}

