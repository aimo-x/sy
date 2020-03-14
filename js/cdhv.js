const domain = "https://iuu.pub"
const domain_test = "127.0.0.1:9999"
const uuid = "779bcec2-587f-5cfa-8207-fc0289b750cc"
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
Authorization_token = getCookie("authorization_token_cdhv")
if(search("test") == "true"){
    Authorization_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBpZCI6Ind4NmU1ODVmMWI4MzlmMjVjZCIsImV4cCI6MTcxMzcyMzA5OSwibmJmIjoxNTg0MTIzMDk5LCJuaWNrbmFtZSI6IumZhuS6kSIsIm9wZW5pZCI6Im9kVll5MVp1MURDbWZBcWpkRGNfSm1SbXR1aUkiLCJ1bmlvbmlkIjoiIn0.-krpfBpdJhja139EzYwRTuc90AyLL7pbic2KdjJLw1s"
}
wechat_oauth = () => {
    // var reuri = window.location.origin+window.location.pathname;
    var reuri = window.location.href;
    fetch(domain+"/v2/api/oauth_wechat_h5/oauthurl?state="+encodeURIComponent(reuri))
    .then((res)=>res.json())
    .then((res)=>{ console.log("authorization_token_cdhv",res.data)
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
    fetch(domain+"/v2/api/oauth_wechat_h5/authorization_token?code="+search("code"))
    .then((res)=>res.json())
    .then((res)=>{ console.log(res.data)
        if(res.code == "success"){
            var token = res.data
            setCookie("authorization_token_cdhv",token,1/48);
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

window.cityArr = [
    {"city": "请选择城市", community: [
        "请选择小区", "未选择城市"
    ]},
    {"city": "福州", community: [
        "请选择小区", "北湖苑一区", "北湖苑二区", "北湖苑三区", "北湖苑四区", "领域", "央著"
    ]},
    {"city": "连江", community: [
        "请选择小区", "连江建发领郡"
    ]},
    {"city": "永泰", community: [
        "请选择小区", "福州山外山一期（梧桐山居）", "福州山外山二期（鹿鸣山居）"
    ]},
    {"city": "三明", community: [
        "请选择小区", "三明建发永郡", "三明建发燕郡", "三明建发央墅", "三明建发玺院"
    ]},
    {"city": "建阳", community: [
        "请选择小区", "建阳悦城一区", "建阳悦城二区", "建阳悦府"
    ]},
    {"city": "建瓯", community: [
        "请选择小区", "建瓯悦城一区", "建瓯悦城二区"
    ]}
];

// 设置城市下拉框
setCity = () => {
    var scene = mugeda.scene
    var city = scene.getObjectByName("城市").dom.getElementsByTagName("select")[0]
    var option = ``
    for(var i = 0; i < cityArr.length; i++){
        if(i === 0){
            option += `<option value=`+ cityArr[i].city +` selected="selected">` + cityArr[i].city + `</option>`
        }else{
            option += `<option value=`+ cityArr[i].city +` >` + cityArr[i].city + `</option>`
        }
    }
    city.innerHTML = option
    console.log(option)
}
// 设置小区下拉框
setCommunity = () => {
    var scene = mugeda.scene
    var city = scene.getObjectByName("城市").text
    if(city == "请选择城市"){
        swal("请选择城市", "", "error")
        return
    }
    var community = scene.getObjectByName("小区").dom.getElementsByTagName("select")[0]
    var option = ``
    for(var i = 0; i < cityArr.length; i++){
        if(cityArr[i].city == city ){
            scene.getObjectByName("小区").text = cityArr[i].community[0]
            for(var n = 0; n < cityArr[i].community.length; n++){
                if(n == 0){
                    option += `<option value=`+ cityArr[i].community[n] +` selected="selected">` +  cityArr[i].community[n] + `</option>`
                }else{
                    option += `<option value=`+ cityArr[i].community[n] +`>` +  cityArr[i].community[n] + `</option>`
                }
                
            }
        }
    }
    community.innerHTML = option
    scene.gotoAndPause(0, 7); // 跳转到相对于某页的某帧并暂停
}
// 第二步 姓名页面
next_one = () => {
    var scene = mugeda.scene
    var community = scene.getObjectByName("小区").text
    var unit_no = scene.getObjectByName("单元号").text
    var name = scene.getObjectByName("姓名").text
    var phone = scene.getObjectByName("电话").text
    var address = scene.getObjectByName("地址").text
    if(community == "请选择小区"){
        swal("请选择小区", "", "error")
        return
    }
    if(unit_no == "请输入单元号"){
        swal("请输入单元号", "", "error")
        return
    }
    if(name.length < 1){
        swal("请输入姓名", "", "error")
        return
    }
    if(phone.length < 1){
        swal("请输入电话", "", "error")
        return
    }
    if(address.length < 1){
        swal("请输入地址", "", "error")
        return
    }
    scene.gotoAndPause(0, 8)
}
// updateImage
updateImage = (em) => {
    em.alpha = 1
}

var init_src = ""
// mugeda init
mugeda.addEventListener("renderready", function(){
    scene = mugeda.scene
    setCity()
    init_src = scene.getObjectByName("completeImage").src
    console.log(init_src)
    var vConsole = new VConsole();
});
// 参与活动
JoinVote = () => {
    var scene = mugeda.scene
    var load = document.createElement("img");
    load.src = scene.getObjectByName("加载动画").src
    load.style.width = "60%"
    load.type = "range";
    swal({
        buttons: false,
        content: load,
    });
    
    var name = scene.getObjectByName("姓名").text
    var phone = scene.getObjectByName("电话").text
    var city = scene.getObjectByName("城市").text
    var community = scene.getObjectByName("小区").text
    var unit_no = scene.getObjectByName("单元号").text
    var address = scene.getObjectByName("地址").text
    var dish_name = scene.getObjectByName("菜名").text
    if(community == "请选择小区"){
        swal("请选择小区", "", "error")
        return
    }
    if(unit_no.length < 1){
        swal("请输入单元号", "", "error")
        return
    }
    if(name.length < 1){
        swal("请输入姓名", "", "error")
        return
    }
    if(phone.length < 1){
        swal("请输入电话", "", "error")
        return
    }
    if(address.length < 1){
        swal("请输入地址", "", "error")
        return
    }
    if(dish_name.length < 1){
        swal("请输入菜名", "", "error")
        return
    }
    var complete_image = scene.getObjectByName("completeImage").src
    if(complete_image == init_src){
        swal("请上传成品图", "", "error")
        return
    }
    var process_imagesArr = []
    process_imagesArr[0] = scene.getObjectByName("processImages#0").src
    process_imagesArr[1] = scene.getObjectByName("processImages#1").src
    process_imagesArr[2] = scene.getObjectByName("processImages#2").src
    var process_images = ""
    if(process_imagesArr[0] == init_src && process_imagesArr[1] == init_src && process_imagesArr[2] == init_src){
        swal("请上传制作过程图", "", "error")
        return
    }
    for(var i = 0;i < process_imagesArr.length; i++){
        if(process_imagesArr[i]!= init_src ){
            process_images += process_imagesArr[i] + ","
        }
    }
    process_images = process_images.substring(0, process_images.length-1)
    console.log(process_images)
    var data = {
        name:name,
        phone:phone,
        city:city,
        community:community,
        unit_no:unit_no,
        address:address,
        dish_name:dish_name,
        complete_image:complete_image,
        process_images:process_images
    }
    fetch(domain+"/v2/api/construction_development/h5/vote/data?construction_development_h5_vote_uuid="+uuid,{
        headers: {
            'Authorization': Authorization_token,
        },
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, cors, *same-origin
        redirect: 'follow', // manual, *follow, error
        referrer: 'no-referrer', // *client, no-referrer
        body: JSON.stringify(data)
    })
    .then((res)=>res.json())
    .then((res)=>{ 
    if(res.code == "success"){
        // swal("提交失败", res.msg, "error")
        var bh = scene.getObjectByName("编号")
        bh.text = res.data.number.toString().length == 1 ? "00" + res.data.number.toString() : res.data.number.toString().length == 2 ? "0" + res.data.number.toString() : res.data.number.toString()
        console.log(res.data)
        scene.gotoAndPause(0, 9) // 跳转到相对于某页的某帧并暂停
    }else{
       swal("提交失败", res.msg, "error")
    }
    })
    .catch((err)=>{
        swal("提交失败", err, "error")
    })
}