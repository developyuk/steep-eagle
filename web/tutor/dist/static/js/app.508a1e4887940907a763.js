webpackJsonp([4],{13:function(t,e,n){"use strict";var u=n(12),a=n.n(u);e.a={routerBeforeEach:function(t,e,n,u){localStorage.getItem("token")?a.a.defaults.headers.common.Authorization="Bearer "+localStorage.getItem("token"):a.a.defaults.headers.common.Authorization="",a.a.get("http://35.187.229.48:90/auth").then(function(e){setTimeout(function(){return t.app.$bus.$emit("currentAuth",e.data)},297),setTimeout(function(){return t.app.$bus.$emit("currentAuth",e.data)},792),setTimeout(function(){return t.app.$bus.$emit("currentAuth",e.data)},1287),setTimeout(function(){return t.app.$bus.$emit("currentAuth",e.data)},2079),setTimeout(function(){return t.app.$bus.$emit("currentAuth",e.data)},3366),setTimeout(function(){return t.app.$bus.$emit("currentAuth",e.data)},5445)}).catch(function(t){u({path:"/sign"})}),u()}}},14:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var u=n(3),a=n(17),r=n(21),o=n(13),i=n(41);n.n(i);window.mdc=i,n(42),u.a.config.productionTip=!1,r.a.beforeEach(function(t,e,n){return o.a.routerBeforeEach(r.a,t,e,n)}),u.a.prototype.$bus=new u.a({}),new u.a({el:"#app",router:r.a,template:"<App/>",components:{App:a.a}})},17:function(t,e,n){"use strict";function u(t){n(18)}var a=n(5),r=n(20),o=n(11),i=u,s=o(a.a,r.a,!1,i,null,null);e.a=s.exports},18:function(t,e){},20:function(t,e,n){"use strict";var u=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"mdc-layout-grid",attrs:{id:"app"}},[n("router-view")],1)},a=[],r={render:u,staticRenderFns:a};e.a=r},21:function(t,e,n){"use strict";var u=n(3),a=n(22);u.a.use(a.a),e.a=new a.a({routes:[{path:"/",name:"Schedules",component:function(){return n.e(1).then(n.bind(null,45))},meta:{requiresAuth:!0}},{path:"/students",name:"Students",component:function(){return n.e(0).then(n.bind(null,46))},meta:{requiresAuth:!0}},{path:"/sign",name:"Sign",component:function(){return n.e(2).then(n.bind(null,47))}}]})},42:function(t,e){},5:function(t,e,n){"use strict";e.a={name:"app",mounted:function(){var t=this,e=new WebSocket("ws://35.187.229.48:90/ws/students"),n=new WebSocket("ws://35.187.229.48:90/ws/");this.$bus.$on("reqCreatedWs",function(){t.$bus.$emit("onCreatedWs",{wsHome:n,wsStudents:e})})}}}},[14]);
//# sourceMappingURL=app.508a1e4887940907a763.js.map