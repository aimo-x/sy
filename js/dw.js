//百度统计
var _hmt = _hmt || [];
(function() {
  var hm = document.createElement("script");
  hm.src = "https://hm.baidu.com/hm.js?f28f11103fc3d69ecd33c645ce781d54";
  var s = document.getElementsByTagName("script")[0]; 
  s.parentNode.insertBefore(hm, s);
})();

/*
// 通用登陆代码开始
const domain = "https://minis.iuu.pub"
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
Authorization_token = getCookie("authorization_token_dw")
if(search("test") == "true"){
    Authorization_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjaWQiOiIxIiwiZXhwIjoxNzE0ODgyNzAxLCJuYmYiOjE1ODUyODI3MDEsIm9pZCI6Im9kVll5MVp1MURDbWZBcWpkRGNfSm1SbXR1aUkiLCJ1aWQiOiIzIn0.BaxBzo-NhveX3K3gRPZd682Sw3gME_4BC52ZMAMK-VE"
    // Authorization_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBpZCI6Ind4NmU1ODVmMWI4MzlmMjVjZCIsImV4cCI6MTcxMzcyMzA5OSwibmJmIjoxNTg0MTIzMDk5LCJuaWNrbmFtZSI6IumZhuS6kSIsIm9wZW5pZCI6Im9kVll5MVp1MURDbWZBcWpkRGNfSm1SbXR1aUkiLCJ1bmlvbmlkIjoiIn0.-krpfBpdJhja139EzYwRTuc90AyLL7pbic2KdjJLw1s"
}
wechat_oauth = () => {
    // var reuri = window.location.origin+window.location.pathname;
    var reuri = window.location.href;
    fetch(domain+"/zs/api/bphd/oauthurl?state="+encodeURIComponent(reuri))
    .then((res)=>res.json())
    .then((res)=>{ console.log("authorization_token_dw",res.data)
        if(res.errCode == "success"){
            window.location.href = res.data;  
        }else{
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
    fetch(domain+"/zs/api/bphd/authorization_token?code="+search("code"))
    .then((res)=>res.json())
    .then((res)=>{ console.log(res.data)
        if(res.errCode == "success"){
            var token = res.data
            setCookie("authorization_token_dw",token,1/48);
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


getUserInfo = () => {
    layer.open({
        type: 2
        ,shadeClose: false
        ,content: '加载中...'
      });
      fetch(interfaceAddr +"getUserInfo",{
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
        if(res.errCode == "success"){
          nickname = scene.getObjectByName("微信昵称")
          nickname.text = res.data.User.NickName
          nc = scene.getObjectByName("nc")
          nc.text = res.data.User.NickName
          wxtx = scene.getObjectByName("头像")
          wxtx.src = res.data.User.HeadImg  
        }else{
          Mugine.Utils.Toast.info(res.msg, {type:'info'});
          console.log(res)
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
*/

ip2city = () =>{
    var scene = mugeda.scene
    ip2cityuri = "https://minis.iuu.pub/zs/api/bphd/ip2city"
    fetch(ip2cityuri)
    .then((res)=>res.json())
    .then((res)=>{
        da = JSON.parse(res.data)
        if(da.status == 0){
            wz = da.result.ad_info.city
            
            if(da.result.ad_info.city == "" ){
                wz = da.result.ad_info.province
                if(da.result.ad_info.province == ""){
                    wz = da.result.ad_info.nation
                }
            }
            city = scene.getObjectByName("定位")
            city.text = wz
            dw = scene.getObjectByName("dw")
            dw.text = wz
        }else{
            Mugine.Utils.Toast.info(da.message, {type:'info'});
        }
    })
    .catch((err)=>{
        Mugine.Utils.Toast.info(err, {type:'info'});
    })
}