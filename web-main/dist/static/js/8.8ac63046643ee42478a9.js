webpackJsonp([8],{176:function(t,e,a){"use strict";var r=a(177),n=a.n(r),i=a(199);a.n(i);e.a={name:"header",data:function(){return{msg:"Welcome to Your Vue.js PWA",currentAuth:{},q:""}},watch:{q:n()(function(t){if(!t){var e=document.querySelector(".search");e.classList.toggle("is-opened"),e.classList.remove("fadeInRight"),e.classList.add("fadeOutRight"),setTimeout(function(){return e.classList.remove("fadeOutRight")},100)}this.$bus.$emit("onKeyupSearch",t)},500)},methods:{signOut:function(t){localStorage.removeItem("token"),window.location.reload()},onClickClearSearch:function(t){this.q=""},onClickSearch:function(t){var e=t.target.closest(".search"),a=e.querySelector("input");e.querySelector(".material-icons");e.classList.contains("is-opened")?(e.classList.toggle("is-opened"),e.classList.remove("fadeInRight"),e.classList.add("fadeOutRight"),setTimeout(function(){return e.classList.remove("fadeOutRight")},100)):(e.classList.toggle("is-opened"),e.classList.remove("fadeOutRight"),e.classList.add("fadeInRight"),setTimeout(function(){return a.focus()},500))}},destroyed:function(){this.$bus.$off("currentAuth")},mounted:function(){var t=this,e=new mdc.drawer.MDCTemporaryDrawer(document.querySelector(".mdc-drawer--temporary"));document.querySelector("img.logo").addEventListener("click",function(){e.open=!0}),this.$bus.$on("currentAuth",function(e){e.photo=e.photo?e.photo:"data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs=",e.photo.indexOf("data:image/gif")<0&&(e.photo=e.photo.replace("https://","").replace("http://",""),e.photo="//images.weserv.nl/?il&q=100&w=64&h=64&t=square&shape=circle&url="+e.photo),t.currentAuth=e})}}},177:function(t,e,a){function r(t,e,a){function r(e){var a=h,r=b;return h=b=void 0,x=e,g=t.apply(r,a)}function c(t){return x=t,B=setTimeout(f,e),_?r(t):g}function m(t){var a=t-k,r=t-x,n=e-a;return y?s(n,w-r):n}function l(t){var a=t-k,r=t-x;return void 0===k||a>=e||a<0||y&&r>=w}function f(){var t=i();if(l(t))return C(t);B=setTimeout(f,m(t))}function C(t){return B=void 0,D&&h?r(t):(h=b=void 0,g)}function u(){void 0!==B&&clearTimeout(B),x=0,h=k=b=B=void 0}function p(){return void 0===B?g:C(i())}function v(){var t=i(),a=l(t);if(h=arguments,b=this,k=t,a){if(void 0===B)return c(k);if(y)return B=setTimeout(f,e),r(k)}return void 0===B&&(B=setTimeout(f,e)),g}var h,b,w,g,B,k,x=0,_=!1,y=!1,D=!0;if("function"!=typeof t)throw new TypeError(d);return e=o(e)||0,n(a)&&(_=!!a.leading,y="maxWait"in a,w=y?A(o(a.maxWait)||0,e):w,D="trailing"in a?!!a.trailing:D),v.cancel=u,v.flush=p,v}var n=a(50),i=a(194),o=a(195),d="Expected a function",A=Math.max,s=Math.min;t.exports=r},191:function(t,e,a){"use strict";function r(t){a(192)}Object.defineProperty(e,"__esModule",{value:!0});var n=a(176),i=a(200),o=a(11),d=r,A=o(n.a,i.a,!1,d,"data-v-25f6bd68",null);e.default=A.exports},192:function(t,e,a){var r=a(193);"string"==typeof r&&(r=[[t.i,r,""]]),r.locals&&(t.exports=r.locals);a(44)("2e7f88e1",r,!0,{})},193:function(t,e,a){e=t.exports=a(43)(!0),e.push([t.i,"form.login .mdc-button[data-v-25f6bd68],form.login .mdc-text-field[data-v-25f6bd68]{width:100%}form.login .logo[data-v-25f6bd68]{width:5rem;margin:0 auto 4rem;display:block}form.login .mdc-form-field[data-v-25f6bd68]{display:block}.powered[data-v-25f6bd68]{position:absolute;bottom:1rem;left:50%;-webkit-transform:translateX(-50%);transform:translateX(-50%);width:12.025rem;font-size:.75rem}.powered img[data-v-25f6bd68]{margin:0 .5rem}header[data-v-25f6bd68]{position:relative;background-color:#ed235c;overflow:hidden}img.logo[data-v-25f6bd68]{width:1.5rem;padding:1rem;-webkit-filter:drop-shadow(1px 3px 1px rgba(0,0,0,.5));filter:drop-shadow(1px 3px 1px rgba(0,0,0,.5))}span.search[data-v-25f6bd68]{color:#fff;position:absolute;top:50%;width:12.5rem;left:calc(100% - 2rem);-webkit-transform:translateY(-50%);transform:translateY(-50%)}span.search.is-opened[data-v-25f6bd68]{color:#000;right:1rem;left:auto;top:20%;width:12rem}span.search.is-opened input[data-v-25f6bd68]{visibility:visible;width:9rem}span.search.is-opened .material-icons[data-v-25f6bd68]{position:absolute;top:50%;-webkit-transform:translateY(-50%);transform:translateY(-50%);left:.5rem;z-index:1}span.search .remove[data-v-25f6bd68]{position:absolute;top:50%;-webkit-transform:translateY(-50%);transform:translateY(-50%);right:1rem;font-size:1.25rem}span.search .material-icons[data-v-25f6bd68]{vertical-align:middle}span.search input[data-v-25f6bd68]{visibility:hidden;padding:.5rem .5rem .5rem 2.5rem;width:8rem;border:none;border-radius:1rem}.mdc-drawer .mdc-drawer__header-content[data-v-25f6bd68]{display:block}.mdc-drawer .mdc-drawer__header-content>[data-v-25f6bd68]{width:320px;display:block;line-height:1.25rem;color:#fff}.mdc-drawer .mdc-drawer__header-content img[data-v-25f6bd68]{width:4rem;height:4rem;margin-bottom:1rem}.mdc-drawer .mdc-drawer__content[data-v-25f6bd68]{color:#9b51e0}.mdc-drawer .content[data-v-25f6bd68]{position:absolute;top:60%;left:50%;-webkit-transform:translate(-50%,-50%);transform:translate(-50%,-50%);width:80%;text-align:center;font-size:.675rem}.mdc-drawer img[data-v-25f6bd68]{width:8rem;border-radius:50%}.mdc-drawer .mdc-list-divider[data-v-25f6bd68],.mdc-drawer .mdc-list-item[data-v-25f6bd68]{position:absolute;width:100%}.mdc-drawer .mdc-list-item[data-v-25f6bd68]{bottom:0}.mdc-drawer .mdc-list-item .material-icons[data-v-25f6bd68],.mdc-drawer .mdc-list-item[data-v-25f6bd68]{color:#ed235c}.mdc-drawer .mdc-list-divider[data-v-25f6bd68]{bottom:3rem}","",{version:3,sources:["/project/src/components/Header.vue"],names:[],mappings:"AACA,oFACE,UAAY,CACb,AACD,kCACE,WAAY,AACZ,mBAAyB,AACzB,aAAe,CAChB,AACD,4CACE,aAAe,CAChB,AACD,0BACE,kBAAmB,AACnB,YAAa,AACb,SAAU,AACV,mCAAoC,AAC5B,2BAA4B,AACpC,gBAAiB,AACjB,gBAAkB,CACnB,AACD,8BACI,cAAgB,CACnB,AACD,wBACE,kBAAmB,AACnB,yBAA0B,AAC1B,eAAiB,CAClB,AACD,0BACE,aAAc,AACd,aAAc,AACd,uDAA4D,AACpD,8CAAoD,CAC7D,AACD,6BACE,WAAY,AAEZ,kBAAmB,AACnB,QAAS,AACT,cAAe,AACf,uBAAwB,AACxB,mCAAoC,AAC5B,0BAA4B,CACrC,AACD,uCAEI,WAAa,AACb,WAAY,AACZ,UAAW,AACX,QAAS,AACT,WAAa,CAChB,AACD,6CACM,mBAAoB,AACpB,UAAY,CACjB,AACD,uDACM,kBAAmB,AACnB,QAAS,AACT,mCAAoC,AAC5B,2BAA4B,AACpC,WAAY,AACZ,SAAW,CAChB,AACD,qCACI,kBAAmB,AACnB,QAAS,AACT,mCAAoC,AAC5B,2BAA4B,AACpC,WAAY,AACZ,iBAAmB,CACtB,AACD,6CACI,qBAAuB,CAC1B,AACD,mCACI,kBAAmB,AAInB,iCAAqB,AACrB,WAAY,AACZ,YAAa,AACb,kBAAoB,CACvB,AACD,yDACE,aAAe,CAEhB,AACD,0DACI,YAAa,AACb,cAAe,AACf,oBAAqB,AACrB,UAAY,CACf,AACD,6DACI,WAAY,AACZ,YAAa,AACb,kBAAoB,CACvB,AACD,kDACE,aAAe,CAChB,AACD,sCACE,kBAAmB,AACnB,QAAS,AACT,SAAU,AACV,uCAAyC,AACjC,+BAAiC,AACzC,UAAW,AACX,kBAAmB,AACnB,iBAAmB,CACpB,AACD,iCACE,WAAY,AACZ,iBAAmB,CACpB,AACD,2FACE,kBAAmB,AACnB,UAAY,CACb,AACD,4CACE,QAAU,CACX,AACD,wGACI,aAAe,CAClB,AACD,+CACE,WAAa,CACd",file:"Header.vue",sourcesContent:["\nform.login .mdc-text-field[data-v-25f6bd68], form.login .mdc-button[data-v-25f6bd68] {\n  width: 100%;\n}\nform.login .logo[data-v-25f6bd68] {\n  width: 5rem;\n  margin: 0 auto 4rem auto;\n  display: block;\n}\nform.login .mdc-form-field[data-v-25f6bd68] {\n  display: block;\n}\n.powered[data-v-25f6bd68] {\n  position: absolute;\n  bottom: 1rem;\n  left: 50%;\n  -webkit-transform: translateX(-50%);\n          transform: translateX(-50%);\n  width: 12.025rem;\n  font-size: .75rem;\n}\n.powered img[data-v-25f6bd68] {\n    margin: 0 .5rem;\n}\nheader[data-v-25f6bd68] {\n  position: relative;\n  background-color: #ED235C;\n  overflow: hidden;\n}\nimg.logo[data-v-25f6bd68] {\n  width: 1.5rem;\n  padding: 1rem;\n  -webkit-filter: drop-shadow(1px 3px 1px rgba(0, 0, 0, 0.5));\n          filter: drop-shadow(1px 3px 1px rgba(0, 0, 0, 0.5));\n}\nspan.search[data-v-25f6bd68] {\n  color: #fff;\n  /*text-shadow: 2px 2px 4px rgba(0, 0, 0, .5);*/\n  position: absolute;\n  top: 50%;\n  width: 12.5rem;\n  left: calc(100% - 2rem);\n  -webkit-transform: translateY(-50%);\n          transform: translateY(-50%);\n}\nspan.search.is-opened[data-v-25f6bd68] {\n    /*color: var(--mdc-theme-primary, #6200ee);*/\n    color: black;\n    right: 1rem;\n    left: auto;\n    top: 20%;\n    width: 12rem;\n}\nspan.search.is-opened input[data-v-25f6bd68] {\n      visibility: visible;\n      width: 9rem;\n}\nspan.search.is-opened .material-icons[data-v-25f6bd68] {\n      position: absolute;\n      top: 50%;\n      -webkit-transform: translateY(-50%);\n              transform: translateY(-50%);\n      left: .5rem;\n      z-index: 1;\n}\nspan.search .remove[data-v-25f6bd68] {\n    position: absolute;\n    top: 50%;\n    -webkit-transform: translateY(-50%);\n            transform: translateY(-50%);\n    right: 1rem;\n    font-size: 1.25rem;\n}\nspan.search .material-icons[data-v-25f6bd68] {\n    vertical-align: middle;\n}\nspan.search input[data-v-25f6bd68] {\n    visibility: hidden;\n    padding-top: .5rem;\n    padding-right: .5rem;\n    padding-bottom: .5rem;\n    padding-left: 2.5rem;\n    width: 8rem;\n    border: none;\n    border-radius: 1rem;\n}\n.mdc-drawer .mdc-drawer__header-content[data-v-25f6bd68] {\n  display: block;\n  /*padding-top: 3rem;*/\n}\n.mdc-drawer .mdc-drawer__header-content > *[data-v-25f6bd68] {\n    width: 320px;\n    display: block;\n    line-height: 1.25rem;\n    color: #fff;\n}\n.mdc-drawer .mdc-drawer__header-content img[data-v-25f6bd68] {\n    width: 4rem;\n    height: 4rem;\n    margin-bottom: 1rem;\n}\n.mdc-drawer .mdc-drawer__content[data-v-25f6bd68] {\n  color: #9B51E0;\n}\n.mdc-drawer .content[data-v-25f6bd68] {\n  position: absolute;\n  top: 60%;\n  left: 50%;\n  -webkit-transform: translate(-50%, -50%);\n          transform: translate(-50%, -50%);\n  width: 80%;\n  text-align: center;\n  font-size: .675rem;\n}\n.mdc-drawer img[data-v-25f6bd68] {\n  width: 8rem;\n  border-radius: 50%;\n}\n.mdc-drawer .mdc-list-item[data-v-25f6bd68], .mdc-drawer .mdc-list-divider[data-v-25f6bd68] {\n  position: absolute;\n  width: 100%;\n}\n.mdc-drawer .mdc-list-item[data-v-25f6bd68] {\n  bottom: 0;\n}\n.mdc-drawer .mdc-list-item[data-v-25f6bd68], .mdc-drawer .mdc-list-item .material-icons[data-v-25f6bd68] {\n    color: #ED235C;\n}\n.mdc-drawer .mdc-list-divider[data-v-25f6bd68] {\n  bottom: 3rem;\n}\n"],sourceRoot:""}])},194:function(t,e,a){var r=a(49),n=function(){return r.Date.now()};t.exports=n},195:function(t,e,a){function r(t){if("number"==typeof t)return t;if(i(t))return o;if(n(t)){var e="function"==typeof t.valueOf?t.valueOf():t;t=n(e)?e+"":e}if("string"!=typeof t)return 0===t?t:+t;t=t.replace(d,"");var a=s.test(t);return a||c.test(t)?m(t.slice(2),a?2:8):A.test(t)?o:+t}var n=a(50),i=a(196),o=NaN,d=/^\s+|\s+$/g,A=/^[-+]0x[0-9a-f]+$/i,s=/^0b[01]+$/i,c=/^0o[0-7]+$/i,m=parseInt;t.exports=r},196:function(t,e,a){function r(t){return"symbol"==typeof t||i(t)&&n(t)==o}var n=a(51),i=a(52),o="[object Symbol]";t.exports=r},199:function(t,e,a){function r(t,e,a){var r=!0,d=!0;if("function"!=typeof t)throw new TypeError(o);return i(a)&&(r="leading"in a?!!a.leading:r,d="trailing"in a?!!a.trailing:d),n(t,e,{leading:r,maxWait:e,trailing:d})}var n=a(177),i=a(50),o="Expected a function";t.exports=r},200:function(t,e,a){"use strict";var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"header"}},[a("header",[a("img",{staticClass:"logo",attrs:{src:"static/img/logo.svg"}}),a("span",{staticClass:"search animated",on:{click:function(e){t.onClickSearch(e)}}},[a("i",{staticClass:"material-icons"},[t._v("search")]),a("input",{directives:[{name:"model",rawName:"v-model.trim",value:t.q,expression:"q",modifiers:{trim:!0}}],attrs:{type:"text",name:"q"},domProps:{value:t.q},on:{input:function(e){e.target.composing||(t.q=e.target.value.trim())},blur:function(e){t.$forceUpdate()}}}),a("div",{staticClass:"remove",on:{click:function(e){t.onClickClearSearch(e)}}},[t._v("x")])])]),a("aside",{staticClass:"mdc-drawer mdc-drawer--temporary mdc-typography"},[a("nav",{staticClass:"mdc-drawer__drawer"},[a("header",{staticClass:"mdc-drawer__header"},[a("div",{staticClass:"mdc-drawer__header-content"},[a("div",{staticClass:"photo"},[a("img",{attrs:{src:t.currentAuth.photo}})]),a("div",{staticClass:"name"},[t._v(t._s(t.currentAuth.name))]),a("div",{staticClass:"email"},[t._v(t._s(t.currentAuth.email))])])]),a("nav",{staticClass:"mdc-drawer__content mdc-list",attrs:{id:"icon-with-text-demo"}},[t._m(0),a("div",{staticClass:"mdc-list-divider",attrs:{role:"separator"}}),a("a",{staticClass:"mdc-list-item",attrs:{href:"#"},on:{click:function(e){e.preventDefault(),t.signOut(e)}}},[a("i",{staticClass:"material-icons mdc-list-item__graphic",attrs:{"aria-hidden":"true"}},[t._v("power_settings_new")]),t._v("Just sign me out !")])])])])])},n=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"content"},[a("img",{attrs:{src:"static/img/uc.gif"}}),a("br"),a("br"),t._v("we are still under construction ~~"),a("br"),t._v("or perhaps you just want to sign out?")])}],i={render:r,staticRenderFns:n};e.a=i}});
//# sourceMappingURL=8.8ac63046643ee42478a9.js.map