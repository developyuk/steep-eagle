webpackJsonp([1],{166:function(t,s,e){"use strict";var a=e(30),n=e.n(a),i=e(43),c=e(45),r=e(5),o=e.n(r);s.a={name:"students",components:{TabBottom:i.a,header1:c.a},data:function(){return{msg:"Welcome to Your Vue.js PWA",students:[],sessions:[],currentAuth:null}},methods:{getStudentsSessions:function(){var t=this;this.$bus.$on("currentAuth",function(s){var e="http://35.185.184.7:82/tutors/"+s.id+"/sessions?students=null";o.a.get(e).then(function(s){t.sessions=s.data._embedded,t.sessions.forEach(function(s,e){t.$set(t.sessions[e],"branch",{name:""}),t.$set(t.sessions[e],"module",{name:""}),t.$set(t.sessions[e],"students",[]),o.a.get("http://35.185.184.7:82"+s._embedded.class._links.branch.href).then(function(s){t.$set(t.sessions[e],"branch",s.data)}).catch(function(t){return console.log(t)}),o.a.get("http://35.185.184.7:82"+s._embedded.class._links.module.href).then(function(s){t.$set(t.sessions[e],"module",s.data)}).catch(function(t){return console.log(t)}),s._links.students&&s._links.students.forEach(function(s,a){o.a.get("http://35.185.184.7:82"+s.href).then(function(s){var a=t.sessions[e].students.length;t.$set(t.sessions[e].students,a,s.data)}).catch(function(t){return console.log(t)})})})}).catch(function(t){return console.log(t)})})},getPosition:function(t){for(var s=0,e=0;t;){if("BODY"==t.tagName){var a=t.scrollLeft||document.documentElement.scrollLeft,n=t.scrollTop||document.documentElement.scrollTop;s+=t.offsetLeft-a+t.clientLeft,e+=t.offsetTop-n+t.clientTop}else s+=t.offsetLeft-t.scrollLeft+t.clientLeft,e+=t.offsetTop-t.scrollTop+t.clientTop;t=t.offsetParent}return{x:s,y:e}}},mounted:function(){var t=this;this.getStudentsSessions(),this.$bus.$on("currentAuth",function(s){console.log(s),t.currentAuth=s});var s=this;setTimeout(function(){var t=document.querySelector(".action");n()(document.querySelectorAll("#students .mdc-list")).forEach(function(e,a){n()(e.querySelectorAll(".mdc-list-item")).forEach(function(e,n){var i=new Hammer(e),c=!1;i.on("panleft panright panend",function(i){var r=i.target.offsetWidth/3;if("panright"!==i.type&&"panleft"!==i.type||(i.deltaX>0&&(t.className="action action--right",c=!1,i.deltaX>r&&(c=!0)),i.deltaX<0&&(t.className="action action--left",c=!1,i.deltaX<-1*r&&(c=!0)),e.style.transform="translateX("+i.deltaX+"px)",t.style.top=s.getPosition(e).y+"px",t.style.height="4rem"),"panend"===i.type&&(e.style.transform="translateX(0px)",Math.abs(i.deltaX)>r)){var o=s.sessions[a-1];s.$delete(o.students,n),o.students.length||s.$delete(s.sessions,a-1),t.style.height="0",console.log("lewat center",a-1,n),c=!1}})})}),console.log("add hammer event")},999)}}},167:function(t,s,e){"use strict";Object.defineProperty(s,"__esModule",{value:!0});var a=e(20),n=e(170),i=e(174),c=e(5),r=e.n(c),o=e(235);e.n(o);a.a.config.productionTip=!1,e(236),e(237),window.mdc=o,i.a.beforeEach(function(t,s,e){r.a.defaults.headers.common.Authorization="Bearer "+localStorage.getItem("token"),r.a.get("http://35.185.184.7:82/auth").then(function(t){setTimeout(function(){return i.a.app.$bus.$emit("currentAuth",t.data)},792),e()}).catch(function(s){t.matched.some(function(t){return t.meta.requiresAuth})?e({path:"/sign",query:{redirect:t.path}}):e()})}),a.a.prototype.$bus=new a.a({}),new a.a({el:"#app",router:i.a,template:"<App/>",components:{App:n.a}})},170:function(t,s,e){"use strict";function a(t){e(171)}var n=e(22),i=e(173),c=e(3),r=a,o=c(n.a,i.a,!1,r,null,null);s.a=o.exports},171:function(t,s){},173:function(t,s,e){"use strict";var a=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"mdc-layout-grid",attrs:{id:"app"}},[e("router-view")],1)},n=[],i={render:a,staticRenderFns:n};s.a=i},174:function(t,s,e){"use strict";var a=e(20),n=e(175),i=e(176),c=e(197),r=e(232);a.a.use(n.a),s.a=new n.a({routes:[{path:"/",name:"Schedules",component:c.a,meta:{requiresAuth:!0}},{path:"/students",name:"Students",component:r.a,meta:{requiresAuth:!0}},{path:"/sign",name:"Sign",component:i.a}]})},176:function(t,s,e){"use strict";function a(t){e(177)}var n=e(23),i=e(196),c=e(3),r=a,o=c(n.a,i.a,!1,r,"data-v-6907a474",null);s.a=o.exports},177:function(t,s){},196:function(t,s,e){"use strict";var a=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"mdc-typography",attrs:{id:"sign"}},[e("form",{on:{submit:function(s){s.preventDefault(),t.sign(s)}}},[e("img",{staticClass:"logo",attrs:{src:"https://images.weserv.nl/?crop=110,100,87,105&url=dl.dropbox.com/s/psvta5uwq4s0m5y/logo2.jpg"}}),e("div",{staticClass:"mdc-form-field"},[e("div",{staticClass:"mdc-text-field mdc-text-field--box mdc-text-field--with-leading-icon",attrs:{"data-mdc-auto-init":"MDCTextField"}},[e("i",{staticClass:"material-icons mdc-text-field__icon",attrs:{tabindex:"0"}},[t._v("account_circle")]),e("input",{directives:[{name:"model",rawName:"v-model.trim",value:t.id,expression:"id",modifiers:{trim:!0}}],staticClass:"mdc-text-field__input",attrs:{id:"id",type:"text",required:"required"},domProps:{value:t.id},on:{input:function(s){s.target.composing||(t.id=s.target.value.trim())},blur:function(s){t.$forceUpdate()}}}),e("label",{staticClass:"mdc-text-field__label mdc-text-field__label--float-above",attrs:{for:"id"}},[t._v("Enter your email")]),e("div",{staticClass:"mdc-line-ripple"})])]),e("div",{staticClass:"mdc-form-field"},[e("div",{staticClass:"mdc-text-field mdc-text-field--box mdc-text-field--with-leading-icon",attrs:{"data-mdc-auto-init":"MDCTextField"}},[e("i",{staticClass:"material-icons mdc-text-field__icon",attrs:{tabindex:"1"}},[t._v("lock_outline")]),e("input",{directives:[{name:"model",rawName:"v-model.trim",value:t.pwd,expression:"pwd",modifiers:{trim:!0}}],staticClass:"mdc-text-field__input",attrs:{id:"pwd",type:"password",required:"required"},domProps:{value:t.pwd},on:{input:function(s){s.target.composing||(t.pwd=s.target.value.trim())},blur:function(s){t.$forceUpdate()}}}),e("label",{staticClass:"mdc-text-field__label mdc-text-field__label--float-above",attrs:{for:"pwd"}},[t._v("Enter your password")]),e("div",{staticClass:"mdc-line-ripple"})])]),e("div",{staticClass:"mdc-form-field"},[t.errMsg?e("div",{staticClass:"errMsg"},[t._v(t._s(t.errMsg))]):t._e(),e("button",{staticClass:"mdc-button mdc-button--raised",attrs:{type:"submit","data-mdc-auto-init":"MDCRipple"}},[t._v("Sign in")])])])])},n=[],i={render:a,staticRenderFns:n};s.a=i},197:function(t,s,e){"use strict";function a(t){e(198)}var n=e(29),i=e(231),c=e(3),r=a,o=c(n.a,i.a,!1,r,"data-v-0ba37d2a",null);s.a=o.exports},198:function(t,s){},22:function(t,s,e){"use strict";s.a={name:"app",mounted:function(){window.mdc.autoInit()}}},225:function(t,s){},226:function(t,s,e){"use strict";var a=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("nav",{staticClass:"mdc-tab-bar mdc-tab-bar--icons-with-text mdc-elevation--z8",attrs:{id:"icon-text-tab-bar","data-mdc-auto-init":"MDCTabBar"}},[e("router-link",{staticClass:"mdc-tab mdc-tab--with-icon-and-text",class:{"mdc-tab--active":t.isActivePath("/")},attrs:{to:"/"}},[e("i",{staticClass:"material-icons mdc-tab__icon",attrs:{"aria-hidden":"true"}},[t._v("home")]),e("span",{staticClass:"mdc-tab__icon-text"},[t._v("Home")])]),e("router-link",{staticClass:"mdc-tab mdc-tab--with-icon-and-text",class:{"mdc-tab--active":t.isActivePath("/students")},attrs:{to:"/students"}},[e("i",{staticClass:"material-icons mdc-tab__icon",attrs:{"aria-hidden":"true"}},[t._v("face")]),e("span",{staticClass:"mdc-tab__icon-text"},[t._v("Student")])])],1)},n=[],i={render:a,staticRenderFns:n};s.a=i},227:function(t,s){},228:function(t,s,e){"use strict";var a=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{attrs:{id:"header"}},[e("header",[e("img",{staticClass:"logo",attrs:{src:"static/img/logo.svg"}}),e("span",{staticClass:"search",on:{click:function(s){t.comingSoon(s)}}},[t._v("search?")])]),e("aside",{staticClass:"mdc-drawer mdc-drawer--temporary mdc-typography"},[e("nav",{staticClass:"mdc-drawer__drawer"},[e("header",{staticClass:"mdc-drawer__header"},[e("div",{staticClass:"mdc-drawer__header-content"},[e("div",{staticClass:"photo"},[e("img",{attrs:{src:"https://images.weserv.nl/?il&q=100&w=64&h=64&t=square&shape=circle&url="+t.currentAuth.photo}})]),e("div",{staticClass:"name"},[t._v(t._s(t.currentAuth.name))]),e("div",{staticClass:"email"},[t._v(t._s(t.currentAuth.email))])])]),e("nav",{staticClass:"mdc-drawer__content mdc-list",attrs:{id:"icon-with-text-demo"}},[t._m(0),e("div",{staticClass:"mdc-list-divider",attrs:{role:"separator"}}),e("a",{staticClass:"mdc-list-item",attrs:{href:"#"},on:{click:function(s){s.preventDefault(),t.signOut(s)}}},[e("i",{staticClass:"material-icons mdc-list-item__graphic",attrs:{"aria-hidden":"true"}},[t._v("power_settings_new")]),t._v("Just sign me out !")])])])])])},n=[function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"content"},[e("img",{attrs:{src:"static/img/uc.gif"}}),e("br"),e("br"),t._v("we are still under construction ~~"),e("br"),t._v("or perhaps you just want to sign out?")])}],i={render:a,staticRenderFns:n};s.a=i},23:function(t,s,e){"use strict";var a=e(5),n=e.n(a);s.a={name:"sign",data:function(){return{msg:"Welcome to Your Vue.js PWA",id:"user3@ex.com",pwd:"asdqwes",errMsg:""}},mounted:function(){mdc.ripple.MDCRipple.attachTo(document.querySelector(".mdc-button"))},methods:{sign:function(){var t=this,s={email:this.id,pwd:this.pwd};n.a.post("http://35.185.184.7:82/sign",s).then(function(s){localStorage.setItem("token",s.data.token),t.$router.push(t.$route.query.redirect)}).catch(function(s){console.log(s);var e=s.response,a=e.status,n=e.data;401===a?t.errMsg=n.message+". Re-check your authentication.":console.log(s)})}}}},230:function(t,s,e){function a(t){return e(n(t))}function n(t){var s=i[t];if(!(s+1))throw new Error("Cannot find module '"+t+"'.");return s}var i={"./af":47,"./af.js":47,"./ar":48,"./ar-dz":49,"./ar-dz.js":49,"./ar-kw":50,"./ar-kw.js":50,"./ar-ly":51,"./ar-ly.js":51,"./ar-ma":52,"./ar-ma.js":52,"./ar-sa":53,"./ar-sa.js":53,"./ar-tn":54,"./ar-tn.js":54,"./ar.js":48,"./az":55,"./az.js":55,"./be":56,"./be.js":56,"./bg":57,"./bg.js":57,"./bm":58,"./bm.js":58,"./bn":59,"./bn.js":59,"./bo":60,"./bo.js":60,"./br":61,"./br.js":61,"./bs":62,"./bs.js":62,"./ca":63,"./ca.js":63,"./cs":64,"./cs.js":64,"./cv":65,"./cv.js":65,"./cy":66,"./cy.js":66,"./da":67,"./da.js":67,"./de":68,"./de-at":69,"./de-at.js":69,"./de-ch":70,"./de-ch.js":70,"./de.js":68,"./dv":71,"./dv.js":71,"./el":72,"./el.js":72,"./en-au":73,"./en-au.js":73,"./en-ca":74,"./en-ca.js":74,"./en-gb":75,"./en-gb.js":75,"./en-ie":76,"./en-ie.js":76,"./en-nz":77,"./en-nz.js":77,"./eo":78,"./eo.js":78,"./es":79,"./es-do":80,"./es-do.js":80,"./es-us":81,"./es-us.js":81,"./es.js":79,"./et":82,"./et.js":82,"./eu":83,"./eu.js":83,"./fa":84,"./fa.js":84,"./fi":85,"./fi.js":85,"./fo":86,"./fo.js":86,"./fr":87,"./fr-ca":88,"./fr-ca.js":88,"./fr-ch":89,"./fr-ch.js":89,"./fr.js":87,"./fy":90,"./fy.js":90,"./gd":91,"./gd.js":91,"./gl":92,"./gl.js":92,"./gom-latn":93,"./gom-latn.js":93,"./gu":94,"./gu.js":94,"./he":95,"./he.js":95,"./hi":96,"./hi.js":96,"./hr":97,"./hr.js":97,"./hu":98,"./hu.js":98,"./hy-am":99,"./hy-am.js":99,"./id":100,"./id.js":100,"./is":101,"./is.js":101,"./it":102,"./it.js":102,"./ja":103,"./ja.js":103,"./jv":104,"./jv.js":104,"./ka":105,"./ka.js":105,"./kk":106,"./kk.js":106,"./km":107,"./km.js":107,"./kn":108,"./kn.js":108,"./ko":109,"./ko.js":109,"./ky":110,"./ky.js":110,"./lb":111,"./lb.js":111,"./lo":112,"./lo.js":112,"./lt":113,"./lt.js":113,"./lv":114,"./lv.js":114,"./me":115,"./me.js":115,"./mi":116,"./mi.js":116,"./mk":117,"./mk.js":117,"./ml":118,"./ml.js":118,"./mr":119,"./mr.js":119,"./ms":120,"./ms-my":121,"./ms-my.js":121,"./ms.js":120,"./mt":122,"./mt.js":122,"./my":123,"./my.js":123,"./nb":124,"./nb.js":124,"./ne":125,"./ne.js":125,"./nl":126,"./nl-be":127,"./nl-be.js":127,"./nl.js":126,"./nn":128,"./nn.js":128,"./pa-in":129,"./pa-in.js":129,"./pl":130,"./pl.js":130,"./pt":131,"./pt-br":132,"./pt-br.js":132,"./pt.js":131,"./ro":133,"./ro.js":133,"./ru":134,"./ru.js":134,"./sd":135,"./sd.js":135,"./se":136,"./se.js":136,"./si":137,"./si.js":137,"./sk":138,"./sk.js":138,"./sl":139,"./sl.js":139,"./sq":140,"./sq.js":140,"./sr":141,"./sr-cyrl":142,"./sr-cyrl.js":142,"./sr.js":141,"./ss":143,"./ss.js":143,"./sv":144,"./sv.js":144,"./sw":145,"./sw.js":145,"./ta":146,"./ta.js":146,"./te":147,"./te.js":147,"./tet":148,"./tet.js":148,"./th":149,"./th.js":149,"./tl-ph":150,"./tl-ph.js":150,"./tlh":151,"./tlh.js":151,"./tr":152,"./tr.js":152,"./tzl":153,"./tzl.js":153,"./tzm":154,"./tzm-latn":155,"./tzm-latn.js":155,"./tzm.js":154,"./uk":156,"./uk.js":156,"./ur":157,"./ur.js":157,"./uz":158,"./uz-latn":159,"./uz-latn.js":159,"./uz.js":158,"./vi":160,"./vi.js":160,"./x-pseudo":161,"./x-pseudo.js":161,"./yo":162,"./yo.js":162,"./zh-cn":163,"./zh-cn.js":163,"./zh-hk":164,"./zh-hk.js":164,"./zh-tw":165,"./zh-tw.js":165};a.keys=function(){return Object.keys(i)},a.resolve=n,t.exports=a,a.id=230},231:function(t,s,e){"use strict";var a=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"mdc-typography",attrs:{id:"schedules"}},[e("header1"),e("ul",{staticClass:"mdc-list"},t._l(t.classes,function(s,a){return e("li",{staticClass:"mdc-list-item"},[e("div",{staticClass:"mdc-list-item__graphic"},[e("img",{attrs:{src:s._embedded.module.image}})]),e("span",{staticClass:"mdc-list-item__text"},[t._v(t._s(s._embedded.module.name)),e("span",{staticClass:"mdc-list-item__secondary-text"},[t._v(t._s(s._embedded.branch.name)+" "+t._s(s.day)+" "+t._s(s.time))])]),"start"===t.buttonStatus(s)?e("button",{staticClass:"mdc-button mdc-button--raised mdc-button--compact",attrs:{"data-mdc-auto-init":"MDCRipple"},on:{click:function(e){t.start(e,s.id,a)}}},[t._v("Start")]):t._e(),"disabled"===t.buttonStatus(s)?e("button",{staticClass:"mdc-button mdc-button--raised mdc-button--compact",attrs:{disabled:"disabled","data-mdc-auto-init":"MDCRipple"},on:{click:function(e){t.start(e,s.id,a)}}},[t._v("Start")]):t._e(),"late"===t.buttonStatus(s)?e("button",{staticClass:"mdc-button mdc-button--raised mdc-button--compact",attrs:{"data-mdc-auto-init":"MDCRipple"},on:{click:function(e){t.start(e,s.id,a)}}},[t._v("Activate")]):t._e(),"ongoing"===t.buttonStatus(s)?e("span",{staticClass:"ongoing"},[t._v("ongoing")]):t._e()])})),t._m(0),t._m(1),e("tab-bottom")],1)},n=[function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("aside",{staticClass:"mdc-dialog",attrs:{id:"my-mdc-dialog",role:"alertdialog","aria-labelledby":"my-mdc-dialog-label","aria-describedby":"my-mdc-dialog-description"}},[e("div",{staticClass:"mdc-dialog__surface"},[e("header",{staticClass:"mdc-dialog__header"},[e("h2",{staticClass:"mdc-dialog__header__title",attrs:{id:"my-mdc-dialog-label"}},[t._v("Start this class?")])]),e("section",{staticClass:"mdc-dialog__body",attrs:{id:"my-mdc-dialog-description"}},[t._v("Class will be start when you click yes.")]),e("footer",{staticClass:"mdc-dialog__footer"},[e("button",{staticClass:"mdc-button mdc-dialog__footer__button mdc-dialog__footer__button--cancel",attrs:{type:"button"}},[t._v("No")]),e("button",{staticClass:"mdc-button mdc-dialog__footer__button mdc-dialog__footer__button--accept",attrs:{type:"button"}},[t._v("Yes")])])]),e("div",{staticClass:"mdc-dialog__backdrop"})])},function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"mdc-snackbar",attrs:{"aria-live":"assertive","aria-atomic":"true","aria-hidden":"true"}},[e("div",{staticClass:"mdc-snackbar__text"}),e("div",{staticClass:"mdc-snackbar__action-wrapper"},[e("button",{staticClass:"mdc-snackbar__action-button",attrs:{type:"button"}})])])}],i={render:a,staticRenderFns:n};s.a=i},232:function(t,s,e){"use strict";function a(t){e(233)}var n=e(166),i=e(234),c=e(3),r=a,o=c(n.a,i.a,!1,r,"data-v-808ba34c",null);s.a=o.exports},233:function(t,s){},234:function(t,s,e){"use strict";var a=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"mdc-typography",attrs:{id:"students"}},[e("header1"),e("div",{staticClass:"mdc-list-group"},[t._l(t.sessions,function(s,a){return[e("h3",{staticClass:"mdc-list-group__subheader"},[e("span",{staticClass:"module"},[t._v(t._s(s.module.name))]),t._v("  "),e("span",{staticClass:"branch"},[t._v(t._s(s.branch.name))])]),e("ul",{staticClass:"mdc-list"},[t._l(s.students,function(s,a){return[e("li",{staticClass:"mdc-list-item"},[e("div",{staticClass:"mdc-list-item__graphic"},[e("img",{attrs:{src:s.photo}})]),e("span",{staticClass:"mdc-list-item__text"},[t._v(t._s(s.name))])]),e("hr",{staticClass:"mdc-list-divider"})]})],2)]})],2),e("div",{staticClass:"action"}),e("tab-bottom")],1)},n=[],i={render:a,staticRenderFns:n};s.a=i},236:function(t,s){},29:function(t,s,e){"use strict";var a=e(30),n=e.n(a),i=e(43),c=e(45),r=e(0),o=e.n(r),d=e(5),l=e.n(d);s.a={name:"schedules",components:{TabBottom:i.a,header1:c.a},data:function(){return{msg:"Welcome to Your Vue.js PWA",classes:[],dialog:null,snackbar:null,thisClass:{id:0},currentAuth:null}},methods:{comingSoon:function(t){t.target.innerText="coming soon"},buttonStatus:function(t){var s=t.ts,e=t._embedded;e=e.last_session?e.last_session.created_at:e.last_session;var a=o()(s),n=o()(),i=a.diff(n,"minutes");if(e){var c=o()(e);if(n.isAfter(c)&&c.diff(n,"minutes")>-120)return"ongoing"}return i<-5?"late":0===a.diff(n,"days")?"start":"disabled"},activate:function(t){var s=this,e="http://35.185.184.7:82/classes/"+t+"/sessions";l.a.post(e,{id:this.currentAuth.id}).then(function(t){s.getSchedules()}).catch(function(t){return console.log(t)})},start:function(t,s,e){this.thisClass=this.classes[e],this.dialog.lastFocusedTarget=t.target,this.dialog.show()},getSchedules:function(){var t=this;arguments.length>0&&void 0!==arguments[0]&&arguments[0];l.a.get("http://35.185.184.7:82/classes?sort=ts.asc").then(function(s){t.classes=s.data._embedded,t.classes.forEach(function(s,e,a){var n=s._embedded.module.image.replace("https://","").replace("http://","");n="//images.weserv.nl/?output=jpg&il&q=100&w=96&h=96&t=square&url="+n,t.$set(a[e]._embedded.module,"image",n)})}).catch(function(t){return console.log(t)})}},mounted:function(){var t=this,s=this;n()(document.querySelectorAll(".mdc-tab, .mdc-button")).forEach(function(t){return mdc.ripple.MDCRipple.attachTo(t)}),this.dialog=mdc.dialog.MDCDialog.attachTo(document.querySelector("#my-mdc-dialog")),this.snackbar=mdc.snackbar.MDCSnackbar.attachTo(document.querySelector(".mdc-snackbar")),this.dialog.listen("MDCDialog:accept",function(){s.activate(s.thisClass.id);var t={message:"Class "+s.thisClass.id+" has been started"};s.snackbar.show(t)}),this.getSchedules(),this.$bus.$on("currentAuth",function(s){t.currentAuth=s})}}},43:function(t,s,e){"use strict";function a(t){e(225)}var n=e(44),i=e(226),c=e(3),r=a,o=c(n.a,i.a,!1,r,"data-v-aa8328ea",null);s.a=o.exports},44:function(t,s,e){"use strict";s.a={name:"tabBottom",data:function(){return{msg:"Welcome to Your Vue.js PWA"}},methods:{isActivePath:function(t){return this.$route.path===t}},mounted:function(){mdc.ripple.MDCRipple.attachTo(document.querySelector(".mdc-tab-bar"))}}},45:function(t,s,e){"use strict";function a(t){e(227)}var n=e(46),i=e(228),c=e(3),r=a,o=c(n.a,i.a,!1,r,"data-v-d0e4480e",null);s.a=o.exports},46:function(t,s,e){"use strict";s.a={name:"header",data:function(){return{msg:"Welcome to Your Vue.js PWA",currentAuth:{}}},methods:{signOut:function(t){localStorage.removeItem("token"),window.location.reload()},comingSoon:function(t){t.target.classList.remove("animated","fadeIn"),"coming soon"!==t.target.innerText.toLowerCase()?t.target.innerText="coming soon":t.target.innerText="search?",t.target.classList.add("animated","fadeIn")}},destroyed:function(){this.$bus.$off("currentAuth")},mounted:function(){var t=this,s=new mdc.drawer.MDCTemporaryDrawer(document.querySelector(".mdc-drawer--temporary"));document.querySelector("img.logo").addEventListener("click",function(){s.open=!0}),this.$bus.$on("currentAuth",function(s){s.photo=s.photo.replace("https://","").replace("http://",""),t.currentAuth=s})}}}},[167]);
//# sourceMappingURL=app.2f0fa3e0b4e12e57cf11.js.map