const domain = "iuu.pub"
const domain_test = "127.0.0.1:8082"
const gacmotor_voice_heart_id = "fa857b22-e92e-52ce-b428-8c2ed217a6ea"
const gacmotor_voice_heart_id_test = "aec30cd2-1a2d-5294-adef-36059154560c"
mugeda.addEventListener("renderready", function(){
  // 当动画准备完成，开始播放前的那一刻引发回调。
  // 换行符 \n 
  s = mugeda.scene
});
var branchArr = [], nameArr = [], phoneArr = [];
team_form_in = () => {
  // branchArr = [], nameArr = [], phoneArr = [];
  var scene = mugeda.scene
  var team_info = ""
  var QZ = "", desc = ""
  for(var i = 0; i < 3; i++){
    switch (i) {
      case 0:
        QZ = "A"
        desc = "队长："
      break;
      case 1:
        QZ = "B"
        desc = "队员："
      break;
      case 2:
        QZ = "C"
        desc = "队员："
      break;
    }
    branchArr[i] = desc + scene.getObjectByName(QZ + "团队部门").text
    nameArr[i] = scene.getObjectByName(QZ + "团队姓名").text
    phoneArr[i] = scene.getObjectByName(QZ + "团队电话").text
    team_info += branchArr[i] + "-" + nameArr[i] + "-" + phoneArr[i] + "\n"
  }
  scene.getObjectByName("团队信息").text = team_info
  console.log(team_info)
  scene.gotoAndPause(1, 4)
}

next_form = () => {
  var scene = mugeda.scene;
  var shape = scene.getObjectByName("形式").text
  var branch = scene.getObjectByName("部门").text
  var name = scene.getObjectByName("姓名").text
  var phone = scene.getObjectByName("联系方式").text
  if(shape!="个人" && shape != "团队"){
    swal("请选择参赛形式","", "error");
    return
  }
  if(branch.length<1 && shape == "个人"){
    swal("请输入参赛部门","", "error");
    return
  }
  if(name.length<1 && shape == "个人"){
    swal("请输入员工姓名","", "error");
    return
  }
  if(phone.length<1 && shape == "个人"){
    swal("请输入联系方式","", "error");
    return
  }
  var team_info = scene.getObjectByName("团队信息").text
  if(team_info<1 && shape == "团队"){
    swal("请输入团队信息","", "error");
    return
  }
  scene.gotoAndPause(3, 4); // 跳转到相对于某页的某帧并暂停
}
var create_status = false
create = () => {
  var scene = mugeda.scene;
  var loading = document.createElement("img")
  loading.style.width = "50%"
  loading.src = scene.getObjectByName("加载动画").src // "/c/user/data/5e367035cec58904ea6d1c30/5e5b0fbde49bc67b98259e44.gif"// 
  //loading.type = "range"
  swal({
    //text:"loading",
    content: loading,
    closeOnClickOutside: false,
    buttons: false,
    // timer: 3000,
  })
  if(create_status){
    
    return
  }
  
  
  /*
  var shape = scene.getObjectByName("参赛形式").text
  var branch = scene.getObjectByName("参赛部门").text
  var name = scene.getObjectByName("员工姓名").text
  var phone = scene.getObjectByName("联系方式").text
  var models = scene.getObjectByName("代言车型").text
  var slogan = scene.getObjectByName("广告语").text
  var originality = scene.getObjectByName("创意说明").text
  */
  var shape = scene.getObjectByName("形式").text
  /*
  var branch = scene.getObjectByName("部门").text
  var name = scene.getObjectByName("姓名").text
  var phone = scene.getObjectByName("联系方式").text
  */
  var mouthpiece = scene.getObjectByName("代言人").text
  var models = scene.getObjectByName("车型").text
  var slogan = scene.getObjectByName("广告语").text
  var originality = scene.getObjectByName("创意说明").text
  
  if(shape == "个人"){
    branchArr = [], nameArr = [], phoneArr = [];
    branchArr[0] = scene.getObjectByName("部门").text
    nameArr[0] = scene.getObjectByName("姓名").text
    phoneArr[0] = scene.getObjectByName("联系方式").text
  }
  var branch="", name="", phone=""
  for(var i = 0; i < branchArr.length; i++){
    branch += branchArr[i]+"\n"
    name += nameArr[i]+"\n"
    phone += phoneArr[i]+"\n"
  }
  if(mouthpiece.length<1){
    swal("请输入代言人姓名","", "error");
    return
  }
  if(models != "全新GS4" && models != "GS4 phev" && models != "GS4 Coupe"){
    swal("请选择代言车型","", "error");
    return
  }
  if(slogan.length<1){
    swal("请输入广告语","", "error");
    return
  }
  if(originality.length<1){
    swal("请输入创意说明","", "error");
    return
  }
  var data = {
    gacmotor_voice_heart_id:gacmotor_voice_heart_id,
    branch:branch,
    shape:shape,
    name:name,
    phone:phone,
    mouthpiece:mouthpiece,
    models:models,
    slogan:slogan,
    originality:originality
  }

  fetch('https://'+domain+'/v2/api/gacmotor_voice_heart',{
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
          swal.close()
          scene.gotoAndPause(4, 4); // 跳转到相对于某页的某帧并暂停
        }else{
          swal(res.msg, res.err, "error");
        }
    })
    .catch((err)=>{
        swal("发生错误", err, "error");
    });
}


test_create = () => {
  var shape = "参赛形式"
  var branch = "参赛部门"
  var name = "员工姓名"
  var phone = "联系方式"
  var models = "代言车型"
  var slogan = "广告语"
  var originality = "创意说明"
  if(branch.length<1){
    swal("请输入参赛部门","", "error");
    return
  }
  if(shape.length<1){
    swal("请输入参赛形式","", "error");
    return
  }
  if(name.length<1){
    swal("请输入员工姓名","", "error");
    return
  }
  if(phone.length<1){
    swal("请输入联系方式","", "error");
    return
  }
  if(models.length<1){
    swal("请输入代言车型","", "error");
    return
  }
  if(slogan.length<1){
    swal("请输入广告语","", "error");
    return
  }
  if(originality.length<1){
    swal("请输入创意说明","", "error");
    return
  }
  var data = {
    gacmotor_voice_heart_id:gacmotor_voice_heart_id_test,
    branch:branch,
    shape:shape,
    name:name,
    phone:phone,
    models:models,
    slogan:slogan,
    originality:originality
  }

  fetch('//'+domain_test+'/v2/api/gacmotor_voice_heart',{
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
          swal("提交成功","", "success");
        }else{
          swal(res.msg, res.err, "error");
        }
    })
    .catch((err)=>{
        swal("发生错误", err, "error");
    });
}
