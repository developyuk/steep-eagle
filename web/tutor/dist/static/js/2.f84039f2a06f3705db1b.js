webpackJsonp([2],{449:function(t,e,s){"use strict";var a=s(12),n=s.n(a);e.a={name:"sign",data:function(){return{username:"",errMsg:""}},mounted:function(){window.mdc.autoInit()},methods:{sign:function(){var t=this,e={username:this.username.toLowerCase()};n.a.post("https://mtutor.codingcamp.id:90/sign",e).then(function(e){localStorage.setItem("token",e.data.token),t.$router.push("/")}).catch(function(e){console.log(e);var s=e.response,a=s.status,n=s.data;200!==a?t.errMsg=n.message+".<br/> Re-check your authentication.":console.log(e)})}}}},49:function(t,e,s){"use strict";function a(t){s(534)}Object.defineProperty(e,"__esModule",{value:!0});var n=s(449),r=s(536),i=s(11),o=a,c=i(n.a,r.a,!1,o,"data-v-3d54204d",null);e.default=c.exports},534:function(t,e,s){var a=s(535);"string"==typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);s(45)("6f2f378e",a,!0,{})},535:function(t,e,s){e=t.exports=s(44)(!0),e.push([t.i,"form.login[data-v-3d54204d]{position:absolute;top:50%;left:50%;-webkit-transform:translate(-50%,-55%);transform:translate(-50%,-55%);max-width:30rem;width:calc(100% - 5rem)}button[data-v-3d54204d]{margin-top:2rem}","",{version:3,sources:["/mnt/data/steep-eagle/web/tutor/src/pages/Sign.vue"],names:[],mappings:"AACA,4BACE,kBAAmB,AACnB,QAAS,AACT,SAAU,AACV,uCAAyC,AACjC,+BAAiC,AACzC,gBAAiB,AACjB,uBAAyB,CAC1B,AACD,wBACE,eAAiB,CAClB",file:"Sign.vue",sourcesContent:["\nform.login[data-v-3d54204d] {\n  position: absolute;\n  top: 50%;\n  left: 50%;\n  -webkit-transform: translate(-50%, -55%);\n          transform: translate(-50%, -55%);\n  max-width: 30rem;\n  width: calc(100% - 5rem);\n}\nbutton[data-v-3d54204d] {\n  margin-top: 2rem;\n}\n"],sourceRoot:""}])},536:function(t,e,s){"use strict";var a=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"mdc-typography",attrs:{id:"sign"}},[s("form",{staticClass:"login",on:{submit:function(e){e.preventDefault(),t.sign(e)}}},[s("img",{staticClass:"logo",attrs:{src:"https://images.weserv.nl/?crop=110,95,88,107&url=dl.dropboxusercontent.com/s/psvta5uwq4s0m5y/logo2.jpg"}}),s("div",{staticClass:"mdc-form-field"},[s("div",{staticClass:"mdc-text-field",attrs:{"data-mdc-auto-init":"MDCTextField"}},[s("input",{directives:[{name:"model",rawName:"v-model.trim",value:t.username,expression:"username",modifiers:{trim:!0}}],staticClass:"mdc-text-field__input",attrs:{id:"id",name:"username",type:"text",required:"required"},domProps:{value:t.username},on:{input:function(e){e.target.composing||(t.username=e.target.value.trim())},blur:function(e){t.$forceUpdate()}}}),s("label",{staticClass:"mdc-text-field__label",attrs:{for:"id"}},[t._v("Yourname")]),s("div",{staticClass:"mdc-line-ripple"})])]),s("div",{staticClass:"mdc-form-field"},[t.errMsg?s("div",{staticClass:"errMsg",domProps:{innerHTML:t._s(t.errMsg)}}):t._e(),s("button",{staticClass:"mdc-button mdc-button--raised",attrs:{type:"submit","data-mdc-auto-init":"MDCRipple"}},[t._v("Login")])])]),t._m(0)])},n=[function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"powered"},[t._v("powered by"),s("img",{attrs:{src:"https://images.weserv.nl/?h=10&url=dl.dropboxusercontent.com/s/htl2v26j5imxgxa/Group.png"}})])}],r={render:a,staticRenderFns:n};e.a=r}});
//# sourceMappingURL=2.f84039f2a06f3705db1b.js.map