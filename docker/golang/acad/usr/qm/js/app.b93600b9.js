(function(t){function e(e){for(var o,i,r=e[0],c=e[1],l=e[2],m=0,u=[];m<r.length;m++)i=r[m],Object.prototype.hasOwnProperty.call(a,i)&&a[i]&&u.push(a[i][0]),a[i]=0;for(o in c)Object.prototype.hasOwnProperty.call(c,o)&&(t[o]=c[o]);d&&d(e);while(u.length)u.shift()();return n.push.apply(n,l||[]),s()}function s(){for(var t,e=0;e<n.length;e++){for(var s=n[e],o=!0,r=1;r<s.length;r++){var c=s[r];0!==a[c]&&(o=!1)}o&&(n.splice(e--,1),t=i(i.s=s[0]))}return t}var o={},a={app:0},n=[];function i(e){if(o[e])return o[e].exports;var s=o[e]={i:e,l:!1,exports:{}};return t[e].call(s.exports,s,s.exports,i),s.l=!0,s.exports}i.m=t,i.c=o,i.d=function(t,e,s){i.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:s})},i.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},i.t=function(t,e){if(1&e&&(t=i(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var s=Object.create(null);if(i.r(s),Object.defineProperty(s,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)i.d(s,o,function(e){return t[e]}.bind(null,o));return s},i.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return i.d(e,"a",e),e},i.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},i.p="/usr/qm/";var r=window["webpackJsonp"]=window["webpackJsonp"]||[],c=r.push.bind(r);r.push=e,r=r.slice();for(var l=0;l<r.length;l++)e(r[l]);var d=c;n.push([0,"chunk-vendors"]),s()})({0:function(t,e,s){t.exports=s("56d7")},"56d7":function(t,e,s){"use strict";s.r(e);s("ac1f"),s("841c"),s("1276"),s("e260"),s("e6cf"),s("cca6"),s("a79d");var o=s("2b0e"),a=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("v-app",{attrs:{id:"inspire"}},[s("v-progress-linear",{directives:[{name:"show",rawName:"v-show",value:t.$store.state.progress,expression:"$store.state.progress"}],attrs:{indeterminate:"",color:"cyan"}}),s("Toolbar",{directives:[{name:"show",rawName:"v-show",value:t.$store.state.login.status,expression:"$store.state.login.status"}]}),"admin_user"==t.$store.state.page?s("AdminUserPage"):t._e(),"admin_agent"==t.$store.state.page?s("AdminAgentPage"):t._e(),"agent_user"==t.$store.state.page?s("AgentUserPage"):t._e(),s("Login",{directives:[{name:"show",rawName:"v-show",value:!t.$store.state.login.status,expression:"!$store.state.login.status"}]})],1)},n=[],i=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",[s("v-app",{attrs:{id:"inspire"}},[s("v-content",[s("v-container",{attrs:{fluid:"","fill-height":""}},[s("v-layout",{attrs:{"align-center":"","justify-center":""}},[s("v-flex",{attrs:{xs12:"",sm8:"",md4:""}},[s("v-card",{staticClass:"elevation-12"},[s("v-toolbar",{attrs:{color:"primary",dark:"",flat:""}},[s("v-toolbar-title",[t._v("管理员登陆")]),s("v-spacer")],1),s("v-card-text",[s("v-form",[s("v-select",{attrs:{items:t.loginclass,label:"登陆系统"},model:{value:t.form.class,callback:function(e){t.$set(t.form,"class",e)},expression:"form.class"}}),s("v-text-field",{attrs:{label:"手机",name:"phone",type:"text"},model:{value:t.form.phone,callback:function(e){t.$set(t.form,"phone",e)},expression:"form.phone"}}),s("v-text-field",{attrs:{label:"密码",name:"password",type:"password"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}})],1)],1),s("v-card-actions",[s("v-spacer"),s("v-btn",{attrs:{color:"primary"},on:{click:t.login}},[t._v("登陆")])],1)],1)],1)],1)],1)],1)],1)],1)},r=[],c={data:function(){return{loginclass:["管理员","代理人"],form:{class:"",phone:"",password:""},condition:!1}},methods:{login:function(){var t=this,e="",s="";if("管理员"==this.form.class&&(e=this.$domain+"/api/qm/auth/admin/login",s="admin"),"代理人"==this.form.class&&(e=this.$domain+"/api/qm/auth/agent/login",s="agent"),""!=e){var o={phone:this.form.phone,password:this.form.password};this.$store.commit("setProgress",!0),this.$post(e,o).then((function(e){window.console.log(e),1==e.code?(t.$cookies.set("qm_jwt_token",e.data,7),t.$cookies.set("qm_jwt_token_class",s,7),t.$store.commit("setLogin",{status:!0,class:s}),"agent"==s&&t.$store.commit("setAgentPhone",t.form.phone)):window.alert(e.msg+":"+e.err),t.condition=!1,t.$store.commit("setProgress",!1)})).catch((function(e){window.alert(e),t.condition=!1}))}else window.alert("请选择登陆系统")}}},l=c,d=s("2877"),m=s("6544"),u=s.n(m),p=s("7496"),h=s("8336"),v=s("b0af"),f=s("99d9"),g=s("a523"),w=s("a75b"),b=s("0e8f"),x=s("4bd4"),_=s("a722"),$=s("b974"),I=s("2fa4"),k=s("8654"),V=s("71d9"),P=s("2a7f"),y=Object(d["a"])(l,i,r,!1,null,null,null),T=y.exports;u()(y,{VApp:p["a"],VBtn:h["a"],VCard:v["a"],VCardActions:f["a"],VCardText:f["b"],VContainer:g["a"],VContent:w["a"],VFlex:b["a"],VForm:x["a"],VLayout:_["a"],VSelect:$["a"],VSpacer:I["a"],VTextField:k["a"],VToolbar:V["a"],VToolbarTitle:P["b"]});var C=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("v-container",[s("v-toolbar",[s("v-toolbar-title",[t._v("仪表盘")]),s("v-spacer"),s("v-toolbar-items",["admin"==t.$store.state.login.class?s("v-btn",{attrs:{text:""},on:{click:t.s_agent}},[t._v("代理管理")]):t._e(),"agent"==t.$store.state.login.class?s("v-btn",{attrs:{text:""},on:{click:t.s_password}},[t._v("修改密码")]):t._e(),s("v-btn",{attrs:{text:""},on:{click:t.s_user}},[t._v("用户管理")])],1),t.$vuetify.breakpoint.smAndUp?[s("v-btn",{attrs:{icon:""},on:{click:t.logout}},[s("v-icon",[t._v("exit_to_app")])],1)]:t._e()],2),s("v-dialog",{attrs:{"max-width":"500px"},model:{value:t.repassword.dialog,callback:function(e){t.$set(t.repassword,"dialog",e)},expression:"repassword.dialog"}},[s("v-card",[s("v-card-title",[s("span",{staticClass:"headline"},[t._v(t._s(t.repassword.formTitle))])]),s("v-card-text",[s("v-container",[s("v-row",[s("v-col",{attrs:{cols:"12",sm:"12"}},[s("v-text-field",{attrs:{label:"新密码"},model:{value:t.repassword.passwordA,callback:function(e){t.$set(t.repassword,"passwordA",e)},expression:"repassword.passwordA"}})],1),s("v-col",{attrs:{cols:"12",sm:"12"}},[s("v-text-field",{attrs:{label:"确认新密码"},model:{value:t.repassword.passwordB,callback:function(e){t.$set(t.repassword,"passwordB",e)},expression:"repassword.passwordB"}})],1)],1)],1)],1),s("v-card-actions",[s("v-spacer"),s("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:t.close}},[t._v("取消")]),s("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:t.save}},[t._v("保存")])],1)],1)],1)],1)},j=[],O={name:"Toolbar",data:function(){return{agent:!1,repassword:{dialog:!1,formTitle:"修改密码",passwordA:"",passwordB:""}}},created:function(){},methods:{logout:function(){this.$cookies.set("qm_jwt_token","",-1),this.$cookies.set("qm_jwt_token_class","",-1),this.$store.commit("setLogin",{status:!1,class:""}),this.$store.commit("setAgentPhone",""),this.$store.commit("setPage","")},s_agent:function(){this.$store.commit("setPage","admin_agent")},s_password:function(){this.repassword.dialog=!0},close:function(){this.repassword.dialog=!1},save:function(){var t=this;if(this.repassword.passwordA.length<6)alert("密码不足6位");else if(this.repassword.passwordA==this.repassword.passwordB){var e=this.$domain+"/api/qm/password?password="+this.repassword.passwordB;this.$put(e).then((function(e){window.console.log(e),1==e.code?(t.repassword.dialog=!1,window.alert("修改成功")):window.alert(e.msg+":"+e.err),t.$store.commit("setProgress",!1)})).catch((function(e){window.alert(e),t.$store.commit("setProgress",!1)}))}else alert("两次密码不一致")},s_user:function(){"admin"==this.$store.state.login.class?this.$store.commit("setPage","admin_user"):this.$store.commit("setPage","agent_user")}}},A=O,q=s("62ad"),S=s("169a"),B=s("132d"),D=s("0fd9"),z=Object(d["a"])(A,C,j,!1,null,null,null),L=z.exports;u()(z,{VBtn:h["a"],VCard:v["a"],VCardActions:f["a"],VCardText:f["b"],VCardTitle:f["c"],VCol:q["a"],VContainer:g["a"],VDialog:S["a"],VIcon:B["a"],VRow:D["a"],VSpacer:I["a"],VTextField:k["a"],VToolbar:V["a"],VToolbarItems:P["a"],VToolbarTitle:P["b"]});var U=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("v-container",[s("v-data-table",{staticClass:"elevation-1",attrs:{headers:t.headers,items:t.desserts},scopedSlots:t._u([{key:"top",fn:function(){return[s("v-toolbar",{attrs:{flat:"",color:"white"}},[s("v-toolbar-title",[t._v("我的用户")]),s("v-divider",{staticClass:"mx-4",attrs:{inset:"",vertical:""}}),s("v-spacer"),s("v-dialog",{attrs:{"max-width":"500px"},scopedSlots:t._u([{key:"activator",fn:function(e){var o=e.on;return[s("v-btn",t._g({staticClass:"mb-2",attrs:{color:"primary",dark:""}},o),[t._v("新增用户")])]}}]),model:{value:t.dialog,callback:function(e){t.dialog=e},expression:"dialog"}},[s("v-card",[s("v-card-title",[s("span",{staticClass:"headline"},[t._v(t._s(t.formTitle))])]),s("v-card-text",[s("v-container",[s("v-row",[s("v-col",{attrs:{cols:"12",sm:"12"}},[s("v-text-field",{attrs:{label:"代理人手机"},model:{value:t.editedItem.agent_phone,callback:function(e){t.$set(t.editedItem,"agent_phone",e)},expression:"editedItem.agent_phone"}})],1),s("v-col",{attrs:{cols:"12",sm:"12"}},[s("v-text-field",{attrs:{label:"姓名"},model:{value:t.editedItem.name,callback:function(e){t.$set(t.editedItem,"name",e)},expression:"editedItem.name"}})],1),s("v-col",{attrs:{cols:"12",sm:"12"}},[s("v-text-field",{attrs:{label:"手机"},model:{value:t.editedItem.phone,callback:function(e){t.$set(t.editedItem,"phone",e)},expression:"editedItem.phone"}})],1)],1)],1)],1),s("v-card-actions",[s("v-spacer"),s("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:t.close}},[t._v("取消")]),s("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:t.save}},[t._v("保存")])],1)],1)],1)],1)]},proxy:!0},{key:"item.action",fn:function(e){var o=e.item;return[s("v-icon",{staticClass:"mr-2",attrs:{small:""},on:{click:function(e){return t.editItem(o)}}},[t._v(" edit ")]),s("v-icon",{attrs:{small:""},on:{click:function(e){return t.deleteItem(o)}}},[t._v(" delete ")])]}},{key:"no-data",fn:function(){return[s("v-btn",{attrs:{color:"primary"},on:{click:t.initialize}},[t._v("刷新")])]},proxy:!0}])})],1)},F=[],E=(s("c975"),s("a434"),s("b0c0"),{data:function(){return{dialog:!1,headers:[{text:"姓名",align:"start",value:"name"},{text:"代理人手机",align:"center",value:"agent_phone"},{text:"手机",align:"center",value:"phone"},{text:"编辑",align:"end",value:"action",sortable:!1}],desserts:[],editedIndex:-1,editedItem:{name:"",agent_phone:"",phone:""},defaultItem:{name:"",agent_phone:"",phone:""}}},computed:{formTitle:function(){return-1===this.editedIndex?"新增用户":"变更代理人"}},watch:{dialog:function(t){t||this.close()}},created:function(){this.initialize()},methods:{initialize:function(){var t=this,e=this.$domain+"/api/qm/user?agent_phone=&offset=0&limit=10000";this.$get(e).then((function(e){window.console.log(e),1==e.code?t.desserts=e.data:window.alert(e.msg+":"+e.err),t.$store.commit("setProgress",!1)})).catch((function(e){window.alert(e),t.$store.commit("setProgress",!1)}))},editItem:function(t){this.editedIndex=this.desserts.indexOf(t),this.editedItem=Object.assign({},t),this.dialog=!0},deleteItem:function(t){var e=this,s=this.desserts.indexOf(t);if(confirm("确认要删除此用户?")){this.$store.commit("setProgress",!0);var o=this.$domain+"/api/qm/admin/user?phone="+t.phone;this.$delete(o).then((function(t){window.console.log(t),1==t.code?e.desserts.splice(s,1):window.alert(t.msg+":"+t.err),e.$store.commit("setProgress",!1)})).catch((function(t){window.alert(t),e.$store.commit("setProgress",!1)}))}},close:function(){var t=this;this.dialog=!1,setTimeout((function(){t.editedItem=Object.assign({},t.defaultItem),t.editedIndex=-1}),300)},save:function(){var t=this;if(this.editedIndex>-1){var e=this.$domain+"/api/qm/admin/user?phone="+this.editedItem.phone+"&agent_phone="+this.editedItem.agent_phone;this.$put(e).then((function(e){window.console.log(e),1==e.code?Object.assign(t.desserts[t.editedIndex],t.editedItem):window.alert(e.msg+":"+e.err),t.$store.commit("setProgress",!1)})).catch((function(e){window.alert(e),t.$store.commit("setProgress",!1)}))}else{var s=this.$domain+"/api/qm/admin/user",o={agent_phone:this.editedItem.agent_phone,name:this.editedItem.name,phone:this.editedItem.phone,password:"001122"};this.$store.commit("setProgress",!0),this.$post(s,o).then((function(e){window.console.log(e),1==e.code?t.desserts.push(e.data):window.alert(e.msg+":"+e.err),t.$store.commit("setProgress",!1)})).catch((function(e){window.alert(e),t.$store.commit("setProgress",!1)}))}this.close()}}}),R=E,M=s("8fea"),N=s("ce7e"),J=Object(d["a"])(R,U,F,!1,null,null,null),G=J.exports;u()(J,{VBtn:h["a"],VCard:v["a"],VCardActions:f["a"],VCardText:f["b"],VCardTitle:f["c"],VCol:q["a"],VContainer:g["a"],VDataTable:M["a"],VDialog:S["a"],VDivider:N["a"],VIcon:B["a"],VRow:D["a"],VSpacer:I["a"],VTextField:k["a"],VToolbar:V["a"],VToolbarTitle:P["b"]});var H=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("v-container",[s("v-data-table",{staticClass:"elevation-1",attrs:{headers:t.headers,items:t.desserts,"items-per-page":t.limit,page:t.offset+1},scopedSlots:t._u([{key:"top",fn:function(){return[s("v-toolbar",{attrs:{flat:"",color:"white"}},[s("v-toolbar-title",[t._v("代理管理")]),s("v-divider",{staticClass:"mx-4",attrs:{inset:"",vertical:""}}),s("v-spacer"),s("v-dialog",{attrs:{"max-width":"500px"},scopedSlots:t._u([{key:"activator",fn:function(e){var o=e.on;return[s("v-btn",t._g({staticClass:"mb-2",attrs:{color:"primary",dark:""}},o),[t._v("新增代理")])]}}]),model:{value:t.dialog,callback:function(e){t.dialog=e},expression:"dialog"}},[s("v-card",[s("v-card-title",[s("span",{staticClass:"headline"},[t._v(t._s(t.formTitle))])]),s("v-card-text",[s("v-container",[s("v-row",[s("v-col",{attrs:{cols:"12",sm:"12"}},[s("v-text-field",{attrs:{label:"姓名"},model:{value:t.editedItem.name,callback:function(e){t.$set(t.editedItem,"name",e)},expression:"editedItem.name"}})],1),s("v-col",{attrs:{cols:"12",sm:"12"}},[s("v-text-field",{attrs:{label:"手机"},model:{value:t.editedItem.phone,callback:function(e){t.$set(t.editedItem,"phone",e)},expression:"editedItem.phone"}})],1),s("v-col",{attrs:{cols:"12",sm:"12"}},[s("v-text-field",{attrs:{label:"密码"},model:{value:t.editedItem.password,callback:function(e){t.$set(t.editedItem,"password",e)},expression:"editedItem.password"}})],1),s("v-col",{attrs:{cols:"12",sm:"12"}},[s("v-text-field",{attrs:{label:"最多可添加条数"},model:{value:t.editedItem.count,callback:function(e){t.$set(t.editedItem,"count",e)},expression:"editedItem.count"}})],1)],1)],1)],1),s("v-card-actions",[s("v-spacer"),s("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:t.close}},[t._v("取消")]),s("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:t.save}},[t._v("保存")])],1)],1)],1)],1)]},proxy:!0},{key:"item.action",fn:function(e){var o=e.item;return[s("v-icon",{staticClass:"mr-2",attrs:{small:""},on:{click:function(e){return t.searchUser(o)}}},[t._v(" search ")]),s("v-icon",{attrs:{small:""},on:{click:function(e){return t.deleteItem(o)}}},[t._v(" delete ")])]}},{key:"no-data",fn:function(){return[s("v-btn",{attrs:{color:"primary"},on:{click:t.initialize}},[t._v("刷新")])]},proxy:!0}])})],1)},K=[],Q=(s("e25e"),{data:function(){return{limit:15,offset:0,dialog:!1,headers:[{text:"姓名",align:"start",value:"name"},{text:"手机",align:"start",value:"phone"},{text:"密码",align:"start",value:"password"},{text:"可添加用户数",align:"start",value:"Count"},{text:"编辑",align:"end",value:"action",sortable:!1}],desserts:[],editedIndex:-1,editedItem:{name:"",phone:"",password:"",count:0},defaultItem:{name:"",phone:"",password:"",count:0}}},computed:{formTitle:function(){return-1===this.editedIndex?"新增用户":"编辑用户"}},watch:{dialog:function(t){t||this.close()}},created:function(){this.initialize()},methods:{initialize:function(){var t=this,e=this.$domain+"/api/qm/agent?offset=0&limit=10000";this.$get(e).then((function(e){window.console.log(e),1==e.code?t.desserts=e.data:window.alert(e.msg+":"+e.err),t.$store.commit("setProgress",!1)})).catch((function(e){window.alert(e),t.$store.commit("setProgress",!1)}))},editItem:function(t){this.editedIndex=this.desserts.indexOf(t),this.editedItem=Object.assign({},t),this.dialog=!0},searchUser:function(t){this.$store.commit("setAgentPhone",t.phone),this.$store.commit("setPage","agent_user")},deleteItem:function(t){var e=this,s=this.desserts.indexOf(t);if(confirm("确认要删除此用户?")){this.$store.commit("setProgress",!0);var o=this.$domain+"/api/qm/admin/agent?phone="+t.phone;this.$delete(o).then((function(t){window.console.log(t),1==t.code?e.desserts.splice(s,1):window.alert(t.msg+":"+t.err),e.$store.commit("setProgress",!1)})).catch((function(t){window.alert(t),e.$store.commit("setProgress",!1)}))}},close:function(){var t=this;this.dialog=!1,setTimeout((function(){t.editedItem=Object.assign({},t.defaultItem),t.editedIndex=-1}),300)},save:function(){var t=this;if(this.editedIndex>-1)Object.assign(this.desserts[this.editedIndex],this.editedItem);else{var e=this.$domain+"/api/qm/admin/agent",s={name:this.editedItem.name,phone:this.editedItem.phone,password:this.editedItem.password,count:window.parseInt(this.editedItem.count)};this.$store.commit("setProgress",!0),this.$post(e,s).then((function(e){window.console.log(e),1==e.code?t.desserts.push(e.data):window.alert(e.msg+":"+e.err),t.$store.commit("setProgress",!1)})).catch((function(e){window.alert(e),t.$store.commit("setProgress",!1)}))}this.close()}}}),W=Q,X=Object(d["a"])(W,H,K,!1,null,null,null),Y=X.exports;u()(X,{VBtn:h["a"],VCard:v["a"],VCardActions:f["a"],VCardText:f["b"],VCardTitle:f["c"],VCol:q["a"],VContainer:g["a"],VDataTable:M["a"],VDialog:S["a"],VDivider:N["a"],VIcon:B["a"],VRow:D["a"],VSpacer:I["a"],VTextField:k["a"],VToolbar:V["a"],VToolbarTitle:P["b"]});var Z=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("v-container",[s("v-data-table",{staticClass:"elevation-1",attrs:{headers:t.headers,items:t.desserts},scopedSlots:t._u([{key:"top",fn:function(){return[s("v-toolbar",{attrs:{flat:"",color:"white"}},[s("v-toolbar-title",[t._v("代理用户")]),s("v-divider",{staticClass:"mx-4",attrs:{inset:"",vertical:""}}),s("v-spacer"),s("v-dialog",{attrs:{"max-width":"500px"},scopedSlots:t._u([{key:"activator",fn:function(e){var o=e.on;return[s("v-btn",t._g({staticClass:"mb-2",attrs:{color:"primary",dark:""}},o),[t._v("新增用户")])]}}]),model:{value:t.dialog,callback:function(e){t.dialog=e},expression:"dialog"}},[s("v-card",[s("v-card-title",[s("span",{staticClass:"headline"},[t._v(t._s(t.formTitle))])]),s("v-card-text",[s("v-container",[s("v-row",[s("v-col",{attrs:{cols:"12",sm:"12"}},[s("v-text-field",{attrs:{label:"姓名"},model:{value:t.editedItem.name,callback:function(e){t.$set(t.editedItem,"name",e)},expression:"editedItem.name"}})],1),s("v-col",{attrs:{cols:"12",sm:"12"}},[s("v-text-field",{attrs:{label:"手机"},model:{value:t.editedItem.phone,callback:function(e){t.$set(t.editedItem,"phone",e)},expression:"editedItem.phone"}})],1)],1)],1)],1),s("v-card-actions",[s("v-spacer"),s("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:t.close}},[t._v("取消")]),s("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:t.save}},[t._v("保存")])],1)],1)],1)],1)]},proxy:!0},{key:"item.action",fn:function(e){var o=e.item;return[s("v-icon",{attrs:{small:""},on:{click:function(e){return t.deleteItem(o)}}},[t._v(" delete ")])]}},{key:"no-data",fn:function(){return[s("v-btn",{attrs:{color:"primary"},on:{click:t.initialize}},[t._v("刷新")])]},proxy:!0}])})],1)},tt=[],et={data:function(){return{dialog:!1,headers:[{text:"姓名",align:"start",value:"name"},{text:"手机",align:"start",value:"phone"},{text:"是否答题",align:"start",value:"write_off"},{text:"答题得分",align:"center",value:"score"},{text:"编辑",align:"end",value:"action",sortable:!1}],desserts:[],editedIndex:-1,editedItem:{name:"",agent_phone:"",phone:""},defaultItem:{name:"",agent_phone:"",phone:"",write_off:"",score:0}}},computed:{formTitle:function(){return-1===this.editedIndex?"新增用户":"编辑用户"}},watch:{dialog:function(t){t||this.close()}},created:function(){this.initialize()},methods:{initialize:function(){var t=this,e=this.$domain+"/api/qm/user?agent_phone="+this.$store.state.agent_phone+"&offset=0&limit=10000";this.$get(e).then((function(e){window.console.log(e),1==e.code?t.desserts=e.data:window.alert(e.msg+":"+e.err),t.$store.commit("setProgress",!1)})).catch((function(e){window.alert(e),t.$store.commit("setProgress",!1)}))},editItem:function(t){this.editedIndex=this.desserts.indexOf(t),this.editedItem=Object.assign({},t),this.dialog=!0},deleteItem:function(t){var e=this,s=this.desserts.indexOf(t);if(confirm("确认要删除此用户?")){this.$store.commit("setProgress",!0);var o=this.$domain+"/api/qm/agent/user?phone="+t.phone;this.$delete(o).then((function(t){window.console.log(t),1==t.code?e.desserts.splice(s,1):window.alert(t.msg+":"+t.err),e.$store.commit("setProgress",!1)})).catch((function(t){window.alert(t),e.$store.commit("setProgress",!1)}))}},close:function(){var t=this;this.dialog=!1,setTimeout((function(){t.editedItem=Object.assign({},t.defaultItem),t.editedIndex=-1}),300)},save:function(){var t=this;if(this.editedIndex>-1)Object.assign(this.desserts[this.editedIndex],this.editedItem);else{var e=this.$domain+"/api/qm/agent/user",s={agent_phone:this.editedItem.agent_phone,name:this.editedItem.name,phone:this.editedItem.phone,password:"001122"};this.$store.commit("setProgress",!0),this.$post(e,s).then((function(e){window.console.log(e),1==e.code?t.desserts.push(e.data):window.alert(e.msg+":"+e.err),t.$store.commit("setProgress",!1)})).catch((function(e){window.alert(e),t.$store.commit("setProgress",!1)}))}this.close()}}},st=et,ot=Object(d["a"])(st,Z,tt,!1,null,null,null),at=ot.exports;u()(ot,{VBtn:h["a"],VCard:v["a"],VCardActions:f["a"],VCardText:f["b"],VCardTitle:f["c"],VCol:q["a"],VContainer:g["a"],VDataTable:M["a"],VDialog:S["a"],VDivider:N["a"],VIcon:B["a"],VRow:D["a"],VSpacer:I["a"],VTextField:k["a"],VToolbar:V["a"],VToolbarTitle:P["b"]});var nt={name:"App",components:{Login:T,Toolbar:L,AdminUserPage:G,AdminAgentPage:Y,AgentUserPage:at},data:function(){return{}},created:function(){this.$cookies.get("qm_jwt_token")&&this.$cookies.get("qm_jwt_token_class")&&this.$store.commit("setLogin",{status:!0,class:this.$cookies.get("qm_jwt_token_class")}),this.$store.commit("setProgress",!1)}},it=nt,rt=s("8e36"),ct=Object(d["a"])(it,a,n,!1,null,null,null),lt=ct.exports;u()(ct,{VApp:p["a"],VProgressLinear:rt["a"]});s("d1e78");var dt=s("f309");o["a"].use(dt["a"]);var mt=new dt["a"]({icons:{iconfont:"md"}}),ut=s("bc3a"),pt=s.n(ut),ht=s("a78e"),vt=s.n(ht);s("d3b7");function ft(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};return new Promise((function(s,o){pt.a.get(t,{params:e}).then((function(t){s(t.data)})).catch((function(t){o(t)}))}))}function gt(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};return new Promise((function(s,o){pt.a.post(t,e).then((function(t){s(t.data)}),(function(t){o(t)}))}))}function wt(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};return new Promise((function(s,o){pt.a.delete(t,e).then((function(t){s(t.data)}),(function(t){o(t)}))}))}function bt(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};return new Promise((function(s,o){pt.a.put(t,e).then((function(t){s(t.data)}),(function(t){o(t)}))}))}pt.a.defaults.timeout=5e3,pt.a.defaults.headers={"Content-Type":"application/json"},pt.a.interceptors.request.use((function(t){var e=vt.a.get("qm_jwt_token");return t.headers={qm_jwt_token:e},t}),(function(t){return Promise.reject(t)})),pt.a.interceptors.response.use((function(t){return t}),(function(t){return Promise.reject(t)}));var xt=s("2f62");o["a"].use(xt["a"]);var _t=new xt["a"].Store({state:{login:{status:!1,class:""},page:"",progress:!0,agent_phone:""},mutations:{setLogin:function(t,e){t.login=e},setProgress:function(t,e){t.progress=e},setPage:function(t,e){t.page=e},setAgentPhone:function(t,e){t.agent_phone=e}},actions:{},modules:{}});s.d(e,"search",(function(){return $t})),o["a"].config.productionTip=!1;var $t=function(t){for(var e=window.location.search.substring(1).split("&"),s=0;s<e.length;s++){var o=e[s].split("=");if(o[0]===t)return decodeURIComponent(o[1])}return!1};o["a"].prototype.$domain="https://iuu.pub",o["a"].prototype.$search=$t,o["a"].prototype.$axios=pt.a,o["a"].prototype.$cookies=vt.a,o["a"].prototype.$get=ft,o["a"].prototype.$post=gt,o["a"].prototype.$put=bt,o["a"].prototype.$delete=wt,new o["a"]({vuetify:mt,store:_t,render:function(t){return t(lt)}}).$mount("#app")}});
//# sourceMappingURL=app.b93600b9.js.map