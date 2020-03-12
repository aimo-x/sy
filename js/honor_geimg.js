mugeda.addEventListener("renderready", function(){
  var vConsole = new VConsole();
  console.log('Hello world');
})
var up_img_status = false
up_img = () => {
  up_img_status = true
  var scene = mugeda.scene
  scene.getObjectByName("upBtn").scene.gotoAndPause(1)
}

var up_qrcode_status = false
up_qrcode = () => {
  up_qrcode_status = true
  var scene = mugeda.scene
  scene.getObjectByName("upQrcode").scene.gotoAndPause(1)
}

// 生成短图
ge_img = (em, img_name, page) => {
  var scene = mugeda.scene
  var head_img = scene.getObjectByName("头像")
  var name = scene.getObjectByName("姓名").text
  var company = scene.getObjectByName("公司").text
  var community = scene.getObjectByName("地址").text
  var phone = scene.getObjectByName("电话").text
  var time =  scene.getObjectByName("时效").text
  var qrcode = scene.getObjectByName("二维码")
  var qrcode_desc = scene.getObjectByName("二维码备注").text
  // /*
  if(!up_img_status){
    swal ( "" ,  "请上传您的照片!" ,  "error" )
    return
  }
  // up_qrcode_status
  if(!up_qrcode_status){
    swal ( "" ,  "请上传您的二维码!" ,  "error" )
    return
  }
  if(name.length<1){
    swal ( "" ,  "请输入您的姓名!" ,  "error" )
    return
  }
  if(company.length<1){
    swal ( "" ,  "请输入您的公司!" ,  "error" )
    return
  }
  if(community.length<1){
    swal ( "" ,  "请输入您的地址!" ,  "error" )
    return
  }
  if(phone.length<1){
    swal ( "" ,  "请输入您的电话!" ,  "error" )
    return
  }
  if(time.length<1){
    swal ( "" ,  "请输入物流时效!" ,  "error" )
    return
  }
  // */
  var app_width = 750
  var app_height = 1500
  bgArr = [scene.getObjectByName("短图#0").src]
  if(img_name == "长图"){
    app_height = 4702
    bgArr = []
    for(var i = 0; i < 4; i++){
      bgArr[i] = scene.getObjectByName("长图#"+i).src
    }
  }
  var boxDIV = document.createElement("div")
  boxDIV.style.width = app_width/2+"px"
  boxDIV.style.height = app_height/2+"px"
  boxDIV.innerHTML = ""

  var appC = scene.getObjectByName(img_name+"容器").dom
  appC.innerHTML = ""
  
  var name_p = {x:362, y:402, size:24}
  var company_p = {x:362, y:479, size:24}
  var community_p = {x:362, y:557, size:24}
  var phone_p = {x:362, y:634, size:24}
  var time_p = {x:385, y:995, size:30}

  app = new PIXI.Application({
      width: app_width, height: app_height, transparent: false,  forceCanvas: true, 
  })

  app.view.style.width = app_width/2+"px"
  app.view.style.height = app_height/2+"px"
  
  boxDIV.appendChild(app.view)
  appC.appendChild(boxDIV)

  // 背景
  var app_bgTexture = []
  var app_bg = []
  for(var i = 0; i < bgArr.length; i++){
    app_bgTexture[i] = PIXI.Texture.from(bgArr[i])
    app_bg[i] = new PIXI.Sprite(app_bgTexture[i])
    app_bg[i].x = 0
    app_bg[i].y = 0+i*1500
    app.stage.addChild(app_bg[i])
  }

  // 头像
  var app_head_imgTexture = PIXI.Texture.from(head_img.src)
  var app_head_img = new PIXI.Sprite(app_head_imgTexture)
  app_head_img.width = head_img.width
  app_head_img.height = head_img.height
  app_head_img.x = 140
  app_head_img.y = 380
  app.stage.addChild(app_head_img)
  
   // 姓名
   var app_nameStyle = new PIXI.TextStyle({
    fontFamily: 'Arial',
    fontSize: name_p.size,
    fill:'#582f0e',
    align:'left',
    fontWeight: 'bold',
  });
  var nameText = "" // \n           ss
  if(name.length>8){
    for(var i = 0; i < parseInt(name.length/8+1);i++){
      nameText += name.substr(i*8,8)+ "\n           "
    }
  }else{
    nameText = name
  }
  var app_name = new PIXI.Text('姓名：'+nameText, app_nameStyle); // 改变此值
  app_name.x = name_p.x
  app_name.y = name_p.y
  app.stage.addChild(app_name)

  // 公司
  var app_companyStyle = new PIXI.TextStyle({
    fontFamily: 'Arial',
    fontSize: company_p.size,
    fill:'#582f0e',
    align:'left',
    fontWeight: 'bold',
  });
  var companyText = "" // \n           ss
  if(company.length>8){
    for(var i = 0; i < parseInt(company.length/8+1);i++){
      companyText += company.substr(i*8,8)+ "\n           "
    }
  }else{
    companyText = company
  }
  
  var app_company = new PIXI.Text('公司：'+companyText, app_companyStyle); // 改变此值
  app_company.x = company_p.x
  app_company.y = company_p.y
  app.stage.addChild(app_company)

  // 社区
  var app_communityStyle = new PIXI.TextStyle({
    fontFamily: 'Arial',
    fontSize: community_p.size,
    fill:'#582f0e',
    align:'left',
    fontWeight: 'bold',
  });
  var communityText = "" // \n           ss
  if(community.length>8){
    for(var i = 0; i < parseInt(community.length/8+1);i++){
      communityText += community.substr(i*8,8)+ "\n           "
    }
  }else{
    communityText = community
  }
  
  var app_community = new PIXI.Text('社区：'+communityText, app_communityStyle); // 改变此值
  app_community.x = community_p.x
  app_community.y = community_p.y
  app.stage.addChild(app_community)

  // 电话
  var app_phoneStyle = new PIXI.TextStyle({
    fontFamily: 'Arial',
    fontSize: phone_p.size,
    fill:'#582f0e',
    align:'left',
    fontWeight: 'bold',
  });
  var phoneText = ""
  phoneText = phone
  /*
   // \n           ss
  if(phone.length>8){
    for(var i = 0; i < parseInt(phone.length/8+1);i++){
      phoneText += phone.substr(i*8,8)+ "\n           "
    }
  }else{
    phoneText = phone
  }
  */
  var app_phone = new PIXI.Text('电话：'+phoneText, app_phoneStyle); // 改变此值
  app_phone.x = phone_p.x
  app_phone.y = phone_p.y
  app.stage.addChild(app_phone)

  // 物流
  var app_timeStyle = new PIXI.TextStyle({
    fontFamily: 'Arial',
    fontSize: time_p.size,
    fill:'#582f0e',
    align:'left',
    fontWeight: 'bold',
  });
  var timeText = "" // \n           ss
  if(time.length>4){
    for(var i = 0; i < parseInt(time.length/4+1);i++){
      timeText += time.substr(i*4,4)+ "\n                  "
    }
  }else{
    timeText = time
  }
  
  var app_time = new PIXI.Text('物流时效：'+timeText, app_timeStyle); // 改变此值
  app_time.x = time_p.x
  app_time.y = time_p.y
  app.stage.addChild(app_time)


  // 二维码
  var app_qrcode_Texture = PIXI.Texture.from(qrcode.src)
  var app_qrcode = new PIXI.Sprite(app_qrcode_Texture)
  app_qrcode.width = qrcode.width
  app_qrcode.height = qrcode.height
  if(img_name == "长图"){
    app_qrcode.x = 494
    app_qrcode.y = 4344
  }else{
    app_qrcode.x = 457
    app_qrcode.y = 1164
  }

  app.stage.addChild(app_qrcode)

  // 二维码备注
  var app_qrcode_descStyle = new PIXI.TextStyle({
    fontFamily: 'Arial',
    fontSize: 24,
    fill:'#ffffff',
    align:'left',
    fontWeight: 'bold',
  });

  var app_qrcode_desc = new PIXI.Text(qrcode_desc, app_qrcode_descStyle); // 改变此值
  if(img_name == "长图"){
    app_qrcode_desc.x = 521
    app_qrcode_desc.y = 4596
  }else{
    app_qrcode_desc.x = 484
    app_qrcode_desc.y = 1410
  }
  app.stage.addChild(app_qrcode_desc)

  app.ticker.add((delta) => {
    // rotate the container[i]!
    // use delta to create frame-independent transform
  })
  // 
 
  
  // toDataMugeda();
  scene.gotoAndPause(0, page)
}

toDataMugeda = () => {
  s = mugeda.scene
  // app.view.style.display = "block"
  rc1 = s.getObjectByName("短图")
  rc2 = s.getObjectByName("长图")
  rc1.src = app.view.toDataURL()
  rc2.src = app.view.toDataURL()
}