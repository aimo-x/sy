
// <script src="https://cdnjs.cloudflare.com/ajax/libs/pixi.js/4.8.5/pixi.min.js"></script>
// <script src="https://cdnjs.cloudflare.com/ajax/libs/sweetalert/2.1.2/sweetalert.min.js"></script>

var up_qrcode_status = false;
function up_qrcode() {
    up_qrcode_status = true;
    var scene = mugeda.scene;
    scene.getObjectByName("upQrcode").scene.gotoAndPause(1);
}
window.Fsize = 5; 
var app_width = 750;
var app_height = 1500;
function ge_img(){
    var boxDIV = document.createElement("div");
    boxDIV.style.width = app_width/2+"px";
    boxDIV.style.height = app_height/2+"px";
    boxDIV.innerHTML = "";
     
    var scene = mugeda.scene

    var bg = scene.getObjectByName("背景");

    var companyT = scene.getObjectByName("公司").text;
    var shopT = scene.getObjectByName("店铺").text;
    var addressT = scene.getObjectByName("地址").text
   
    var phoneT = scene.getObjectByName("电话").text
    var benefitT = scene.getObjectByName("利益").text;
    var zz = scene.getObjectByName("zz");
    var qrcodeT = scene.getObjectByName("二维码");
    
    // up_qrcode_status
    if(!up_qrcode_status){
        swal ( "" ,  "请上传您的二维码!" ,  "error" );
        return;
    }
    if(companyT.length<1){
        swal ( "" ,  "请输入您的公司!" ,  "error" );
        return;
    }
    if(shopT.length<1){
        swal ( "" ,  "请输入您的店铺!" ,  "error" );
        return;
    }
    if(addressT.length<1){
        swal ( "" ,  "请输入您的地址!" ,  "error" );
        return;
    }
    if(phoneT.length<1){
        swal ( "" ,  "请输入您的电话!" ,  "error" );
        return;
    }
    if(benefitT.length<1){
        swal ( "" ,  "请输入利益，买赠，权益!" ,  "error" );
        return;
    }
    
    var appC =  scene.getObjectByName("容器").dom;
    appC.innerHTML = "";
    var companyP = {x:192, y:32, size: 30 + Fsize, fontWeight:"500",fill:"#494742"};
    var shopP = {x:368, y:1355, size: 20 + Fsize, fontWeight:"500",fill:"#2c2c2c"};
    var addressP = {x:369, y:1379, size: 12 + Fsize, fontWeight:"500",fill:"#2c2c2c"};
    var phoneP = {x:607, y:1358, size: 14 + Fsize, fontWeight:"500",fill:"#ffffff"};
    var benefitP = {x:370, y:1315, size: 16 +Fsize, fontWeight:"500",fill:"#2c2c2c"};
    window.app = new PIXI.Application({
        width: app_width, height: app_height, transparent: false,  forceCanvas: true, 
    });

    app.view.style.width = app_width/2+"px";
    app.view.style.height = app_height/2+"px";
    
    boxDIV.appendChild(app.view);
    appC.appendChild(boxDIV);

    // 背景
    var bgTexture = PIXI.Texture.from(bg.src)
    var bg = new PIXI.Sprite(bgTexture)
    bg.x = 0;
    bg.y = 0;
    app.stage.addChild(bg);
    // 公司
    var companyStyle = new PIXI.TextStyle({
        fontFamily: 'Arial',
        fontSize: companyP.size,
        fill: companyP.fill,
        align: 'left',
        fontWeight: companyP.fontWeight,
    });
    var company = new PIXI.Text(companyT, companyStyle); // 改变此值
    company.x = companyP.x;
    company.y = companyP.y;
    app.stage.addChild(company);
    // 店铺
    var shopStyle = new PIXI.TextStyle({
        fontFamily: 'Arial',
        fontSize: shopP.size,
        fill: shopP.fill,
        align: 'left',
        fontWeight: shopP.fontWeight,
    });
    var shop = new PIXI.Text(shopT, shopStyle); // 改变此值
    shop.x = shopP.x;
    shop.y = shopP.y;
    app.stage.addChild(shop);
    // 地址
    var addressStyle = new PIXI.TextStyle({
        fontFamily: 'Arial',
        fontSize: addressP.size,
        fill: addressP.fill,
        align: 'left',
        fontWeight: addressP.fontWeight,
    });
    var address = new PIXI.Text(addressT, addressStyle); // 改变此值
    address.x = addressP.x;
    address.y = addressP.y;
    app.stage.addChild(address);
    // 手机
    var phoneStyle = new PIXI.TextStyle({
        fontFamily: 'Arial',
        fontSize: phoneP.size,
        fill: phoneP.fill,
        align: 'left',
        fontWeight: phoneP.fontWeight,
    });
    var phone = new PIXI.Text(phoneT, phoneStyle); // 改变此值
    phone.x = phoneP.x;
    phone.y = phoneP.y;
    app.stage.addChild(phone);
     // 利益
     var benefitStyle = new PIXI.TextStyle({
        fontFamily: 'Arial',
        fontSize: benefitP.size,
        fill: benefitP.fill,
        align: 'left',
        fontWeight: benefitP.fontWeight,
    });
    var benefit = new PIXI.Text(benefitT, benefitStyle); // 改变此值
    benefit.x = benefitP.x;
    benefit.y = benefitP.y;
    app.stage.addChild(benefit);
    // 遮挡
    var ZZA = PIXI.Texture.from(zz.src);
    ZZimg = new PIXI.Sprite(ZZA);
    // ZZimg.anchor.set(0.5);
    ZZimg.width = 178;
    ZZimg.height = 178;
    ZZimg.x = 31;
    ZZimg.y = 1291;
    app.stage.addChild(ZZimg)
    // 二维码
    var qrcodeTexture = PIXI.Texture.from(qrcodeT.src);
    var qrcode = new PIXI.Sprite(qrcodeTexture);
    qrcode.width = qrcodeT.width;
    qrcode.height = qrcodeT.height;
    qrcode.x = 31;
    qrcode.y = 1291;

    app.stage.addChild(qrcode);

    app.ticker.add((delta) => {
        // rotate the container[i]!
        // use delta to create frame-independent transform
        qrcode.mask = ZZimg
    });
    scene.gotoAndPause(0, 1);
}
function toDataMugeda(){
    var s = mugeda.scene;
    rc1 = s.getObjectByName("背景");
    rc1.src = app.view.toDataURL();
}
mugeda.defineCallback("up_qrcode",function(em){
    up_qrcode(em);
});
mugeda.defineCallback("ge_img",function(em){
    ge_img(em);
});
mugeda.defineCallback("toDataMugeda",function(em){
    toDataMugeda(em);
});