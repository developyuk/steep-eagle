webpackJsonp([0],{49:function(t,e,a){"use strict";function n(t){a(51)}Object.defineProperty(e,"__esModule",{value:!0});var d=a(50),i=a(53),r=a(3),o=n,s=r(d.a,i.a,!1,o,"data-v-e1afdd56",null);e.default=s.exports},50:function(t,e,a){"use strict";e.a={name:"sign",data:function(){return{msg:"Welcome to Your Vue.js PWA",username:""}},methods:{sign:function(){var t=this,e={username:this.username.toLowerCase()};axios.post("http://35.187.229.48:90/sign",e).then(function(e){localStorage.setItem("token",e.data.token),t.$router.push(t.$route.query.redirect)}).catch(function(e){console.log(e);var a=e.response,n=a.status,d=a.data;200!==n?t.errMsg=d.message+".<br/> Re-check your authentication.":console.log(e)})}},mounted:function(){window.mdc.autoInit()}}},51:function(t,e,a){var n=a(52);"string"==typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);a(48)("07572bb0",n,!0,{})},52:function(t,e,a){e=t.exports=a(47)(!0),e.push([t.i,'form.login .mdc-button[data-v-e1afdd56],form.login .mdc-text-field[data-v-e1afdd56]{width:100%}form.login .logo[data-v-e1afdd56]{width:5rem;margin:0 auto 4rem;display:block}form.login .mdc-form-field[data-v-e1afdd56]{display:block}.powered[data-v-e1afdd56]{position:absolute;bottom:1rem;left:50%;-webkit-transform:translateX(-50%);transform:translateX(-50%);width:12.025rem;font-size:.75rem}.powered img[data-v-e1afdd56]{margin:0 .5rem}.mdc-list-divider[data-v-e1afdd56]{border-bottom-color:#dcdcdc}.hide[data-v-e1afdd56]{display:none}.clearfix[data-v-e1afdd56]:after,.clearfix[data-v-e1afdd56]:before{content:"";clear:both;display:table}.mdc-snackbar[data-v-e1afdd56]{z-index:9}#sign[data-v-e1afdd56]{padding:0;height:100vh}main[data-v-e1afdd56]{height:100%}.left[data-v-e1afdd56]{position:relative}.right[data-v-e1afdd56]{background-image:url(//images.weserv.nl/?output=jpg&q=21&url=dl.dropboxusercontent.com/s/asuxh2qp5zh79uu/Rectangle-compressor.png);position:relative}.right .container[data-v-e1afdd56]{margin:2rem;position:absolute;top:50%;-webkit-transform:translateY(-50%);transform:translateY(-50%)}.right .container__text[data-v-e1afdd56]{color:#fff;font-size:2rem;font-weight:700}',"",{version:3,sources:["/project/src/pages/Sign.vue"],names:[],mappings:"AACA,oFACE,UAAY,CACb,AACD,kCACE,WAAY,AACZ,mBAAyB,AACzB,aAAe,CAChB,AACD,4CACE,aAAe,CAChB,AACD,0BACE,kBAAmB,AACnB,YAAa,AACb,SAAU,AACV,mCAAoC,AAC5B,2BAA4B,AACpC,gBAAiB,AACjB,gBAAkB,CACnB,AACD,8BACI,cAAgB,CACnB,AACD,mCACE,2BAA6B,CAC9B,AACD,uBACE,YAAc,CACf,AACD,mEACE,WAAY,AACZ,WAAY,AACZ,aAAe,CAChB,AACD,+BACE,SAAW,CACZ,AACD,uBACE,UAAW,AACX,YAAc,CACf,AACD,sBACE,WAAa,CACd,AACD,uBACE,iBAAmB,CACpB,AACD,wBACE,mIAAoI,AACpI,iBAAmB,CACpB,AACD,mCAEI,YAAa,AACb,kBAAmB,AACnB,QAAS,AACT,mCAAoC,AAC5B,0BAA4B,CACvC,AACD,yCACM,WAAY,AACZ,eAAgB,AAChB,eAAiB,CACtB",file:"Sign.vue",sourcesContent:['\nform.login .mdc-text-field[data-v-e1afdd56], form.login .mdc-button[data-v-e1afdd56] {\n  width: 100%;\n}\nform.login .logo[data-v-e1afdd56] {\n  width: 5rem;\n  margin: 0 auto 4rem auto;\n  display: block;\n}\nform.login .mdc-form-field[data-v-e1afdd56] {\n  display: block;\n}\n.powered[data-v-e1afdd56] {\n  position: absolute;\n  bottom: 1rem;\n  left: 50%;\n  -webkit-transform: translateX(-50%);\n          transform: translateX(-50%);\n  width: 12.025rem;\n  font-size: .75rem;\n}\n.powered img[data-v-e1afdd56] {\n    margin: 0 .5rem;\n}\n.mdc-list-divider[data-v-e1afdd56] {\n  border-bottom-color: #dcdcdc;\n}\n.hide[data-v-e1afdd56] {\n  display: none;\n}\n.clearfix[data-v-e1afdd56]::before, .clearfix[data-v-e1afdd56]::after {\n  content: "";\n  clear: both;\n  display: table;\n}\n.mdc-snackbar[data-v-e1afdd56] {\n  z-index: 9;\n}\n#sign[data-v-e1afdd56] {\n  padding: 0;\n  height: 100vh;\n}\nmain[data-v-e1afdd56] {\n  height: 100%;\n}\n.left[data-v-e1afdd56] {\n  position: relative;\n}\n.right[data-v-e1afdd56] {\n  background-image: url(//images.weserv.nl/?output=jpg&q=21&url=dl.dropboxusercontent.com/s/asuxh2qp5zh79uu/Rectangle-compressor.png);\n  position: relative;\n}\n.right .container[data-v-e1afdd56] {\n    /*padding: 4rem;*/\n    margin: 2rem;\n    position: absolute;\n    top: 50%;\n    -webkit-transform: translateY(-50%);\n            transform: translateY(-50%);\n}\n.right .container__text[data-v-e1afdd56] {\n      color: #fff;\n      font-size: 2rem;\n      font-weight: 700;\n}\n'],sourceRoot:""}])},53:function(t,e,a){"use strict";var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"mdc-layout-grid",attrs:{id:"sign"}},[a("main",{staticClass:"mdc-layout-grid__inner"},[a("section",{staticClass:"left mdc-layout-grid__cell mdc-layout-grid__cell--span-4"},[a("form",{staticClass:"login",on:{submit:function(e){return e.preventDefault(),t.sign(e)}}},[a("img",{staticClass:"logo",attrs:{src:"https://images.weserv.nl/?crop=110,95,88,107&url=dl.dropboxusercontent.com/s/psvta5uwq4s0m5y/logo2.jpg"}}),a("div",{staticClass:"mdc-form-field"},[a("div",{staticClass:"mdc-text-field",attrs:{"data-mdc-auto-init":"MDCTextField"}},[a("input",{directives:[{name:"model",rawName:"v-model.trim",value:t.username,expression:"username",modifiers:{trim:!0}}],staticClass:"mdc-text-field__input",attrs:{id:"id",name:"username",type:"text",required:"required"},domProps:{value:t.username},on:{input:function(e){e.target.composing||(t.username=e.target.value.trim())},blur:function(e){t.$forceUpdate()}}}),a("label",{staticClass:"mdc-text-field__label",attrs:{for:"id"}},[t._v("Yourname")]),a("div",{staticClass:"mdc-line-ripple"})])]),t._m(0)]),t._m(1)]),t._m(2)])])},d=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"mdc-form-field"},[a("button",{staticClass:"mdc-button mdc-button--raised",attrs:{type:"submit"}},[t._v("login")])])},function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"powered"},[t._v("powered by"),a("img",{attrs:{src:"https://images.weserv.nl/?h=10&url=dl.dropboxusercontent.com/s/htl2v26j5imxgxa/Group.png"}})])},function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("section",{staticClass:"right mdc-layout-grid__cell mdc-layout-grid__cell--span-8"},[a("div",{staticClass:"container"},[a("div",{staticClass:"container__text"},[t._v("Scheduling never been this easy.")]),a("div",{staticClass:"container__text"},[t._v("Manage with M.")])])])}],i={render:n,staticRenderFns:d};e.a=i}});
//# sourceMappingURL=0.4581ef1f3989b32b2341.js.map