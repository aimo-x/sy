const domain = "iuu.pub"
const domain_test = "iuu.pub"
get_scene=()=>{
   return mugeda.scene;
}
create = () => {
  var scene = mugeda.scene;
  var mailto = "NorthAsia@intralinks.com"
  var subject = "WeChat Inbound Inquiry"
  var title = ""
  var desc = ""
  var urlP = "?mailto="+mailto+"&subject="+subject+"&title="+title+"&desc="+desc

  var name = scene.getObjectByName("姓名").text
  var position = scene.getObjectByName("职位").text
  var company = scene.getObjectByName("公司").text
  var email = scene.getObjectByName("Email").text
  if(name.length<1){
    swal("请输入姓名","", "error");
    return
  }
  if(position.length<1){
    swal("请输入职位","", "error");
    return
  }
  if(company.length<1){
    swal("请输入公司","", "error");
    return
  }
  if(email.length<1){
    swal("请输入Email","", "error");
    return
  }
  var data = {
    mugeda_form_id:1,
    name:name,
    position:position,
    company:company,
    email:email
  }

  fetch('https://'+domain+'/v2/api/mugeda_form_content'+urlP,{
        headers: { 
            "Content-Type": "application/json"
        },
        method: 'POST',
        mode: 'cors', // no-cors, cors, *same-origin
        redirect: 'follow', // manual, *follow, error
        referrer: 'no-referrer', // *client, no-referrer
        body: JSON.stringify(data)
    })
    .then(function(response) {
        return response.json();
    })
    .then(function(res) {
        if(res.code == "success" ){
          // swal("提交成功","", "success");
          scene.gotoAndPause(1, 8); // 跳转到相对于某页的某帧并暂停

        }else{
          swal(res.msg, res.err, "error");
        }
    })
    .catch((err)=>{
        swal("发生错误", err, "error");
    });
    
}


create_test = () => {
  
  var mailto = "494745409@qq.com"
  var subject = "邮件主题"
  var title = "邮件副标题"
  var desc = "邮件描述"
  var urlP = "?mailto="+mailto+"&subject="+subject+"&title="+title+"&desc="+desc

  var name = "姓名"
  var position = "职位"
  var company = "公司"
  var email = "Email"

  var data = {
    mugeda_form_id:1,
    name:name,
    position:position,
    company:company,
    email:email
  }

  fetch('https://'+domain_test+'/v2/api/mugeda_form_content'+urlP,{
        headers: { 
            "Content-Type": "application/json"
        },
        method: 'POST',
        body: JSON.stringify(data)
    })
    .then(function(response) {
        return response.json();
    })
    .then(function(res) {
        if(res.code == "success" ){
          swal("提交成功","", "success");
        }else{
          swal(res.msg, res.err, "error");
        }
    })
    .catch((err)=>{
        swal("发生错误", err, "error");
    });
    
}