webpackJsonp([1],{188:function(t,e,s){"use strict";var n=s(12),i=s.n(n);e.a={name:"students",components:{spinner:function(){return s.e(6).then(s.bind(null,182))},"tab-bottom":function(){return s.e(5).then(s.bind(null,184))},header1:function(){return s.e(7).then(s.bind(null,186))},"form-rate-review":function(){return s.e(8).then(s.bind(null,220))},empty:function(){return s.e(9).then(s.bind(null,252))}},data:function(){return{msg:"Welcome to Your Vue.js PWA",sessions:null,currentAuth:null,currentView:{},lastId:""}},methods:{onClickList:function(t){var e=t.target.closest(".mdc-list-item").nextSibling.nextSibling;console.log("clicked",e);var s=void 0,n=void 0,i=e.getAttribute("id");if(this.lastId){var a=this.lastId.split("-");s=a[0],n=a[1],this.$set(this.currentView[s],n,"empty")}s=e.getAttribute("sid"),n=e.getAttribute("uid"),"empty"===i&&(this.$set(this.currentView[s],n,"form-rate-review"),this.lastId=s+"-"+n),console.log("clicked",s,n,this.lastId)},getStudentsSessions:function(){var t=this,e="http://35.187.229.48:81/tutors/"+this.currentAuth.id+"/sessions";i.a.get(e).then(function(e){t.sessions=e.data._embedded.items;var s=[];t.sessions?t.sessions.forEach(function(e,n,i){e.students&&e.students.forEach(function(a,c,r){console.log(a),s[e.id]||(s[e.id]=[]),s[e.id][a.id]="empty";var o=a.photo?a.photo:"https://image.flaticon.com/icons/png/128/201/201818.png";o=o.replace("https://","").replace("http://",""),o="//images.weserv.nl/?output=png&il&q=100&w=96&h=96&t=square&url="+o,t.$set(i[n].students[c],"photo",o)})}):t.sessions=[],t.currentView=s}).catch(function(t){return console.log(t)})}},mounted:function(){var t=this;this.$bus.$on("currentAuth",function(e){t.currentAuth=e,t.getStudentsSessions()}),this.$bus.$on("onAfterSubmitRateReview",function(e){t.getStudentsSessions()})}}},218:function(t,e,s){var n=s(219);"string"==typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);s(43)("46a43266",n,!0,{})},219:function(t,e,s){e=t.exports=s(42)(!0),e.push([t.i,"#students[data-v-e9bb0cc2]{position:relative;height:100vh}#form-rate-review[data-v-e9bb0cc2]{-webkit-transition:height .3s cubic-bezier(0,0,.2,1) 0ms;transition:height .3s cubic-bezier(0,0,.2,1) 0ms}.empty[data-v-e9bb0cc2]{text-align:center;position:absolute;top:50%;-webkit-transform:translateY(-70%);transform:translateY(-70%);width:100%}.mdc-list-item[data-v-e9bb0cc2]{padding:.5rem 1rem;background-color:#fff;z-index:2}.mdc-list-item__text[data-v-e9bb0cc2]{text-transform:capitalize}.mdc-list-item__graphic[data-v-e9bb0cc2],.mdc-list-item__graphic img[data-v-e9bb0cc2]{width:40px;height:40px}span.module[data-v-e9bb0cc2]{text-transform:uppercase}span.branch[data-v-e9bb0cc2],span.day-time[data-v-e9bb0cc2]{text-transform:capitalize}span.day-time[data-v-e9bb0cc2]{float:right}","",{version:3,sources:["/home/andry/Documents/steep-eagle/web-main/src/pages/Students.vue"],names:[],mappings:"AACA,2BACE,kBAAmB,AACnB,YAAc,CACf,AACD,mCACE,yDAAgE,AAChE,gDAAwD,CACzD,AACD,wBACE,kBAAmB,AACnB,kBAAmB,AACnB,QAAS,AACT,mCAAoC,AAC5B,2BAA4B,AACpC,UAAY,CACb,AACD,gCACE,mBAAoB,AACpB,sBAAuB,AACvB,SAAW,CACZ,AACD,sCACE,yBAA2B,CAC5B,AACD,sFACE,WAAY,AACZ,WAAa,CACd,AACD,6BACE,wBAA0B,CAC3B,AACD,4DACE,yBAA2B,CAC5B,AACD,+BACE,WAAa,CACd",file:"Students.vue",sourcesContent:["\n#students[data-v-e9bb0cc2] {\n  position: relative;\n  height: 100vh;\n}\n#form-rate-review[data-v-e9bb0cc2] {\n  -webkit-transition: height 300ms 0ms cubic-bezier(0, 0, 0.2, 1);\n  transition: height 300ms 0ms cubic-bezier(0, 0, 0.2, 1);\n}\n.empty[data-v-e9bb0cc2] {\n  text-align: center;\n  position: absolute;\n  top: 50%;\n  -webkit-transform: translateY(-70%);\n          transform: translateY(-70%);\n  width: 100%;\n}\n.mdc-list-item[data-v-e9bb0cc2] {\n  padding: .5rem 1rem;\n  background-color: #fff;\n  z-index: 2;\n}\n.mdc-list-item__text[data-v-e9bb0cc2] {\n  text-transform: capitalize;\n}\n.mdc-list-item__graphic[data-v-e9bb0cc2], .mdc-list-item__graphic img[data-v-e9bb0cc2] {\n  width: 40px;\n  height: 40px;\n}\nspan.module[data-v-e9bb0cc2] {\n  text-transform: uppercase;\n}\nspan.branch[data-v-e9bb0cc2], span.day-time[data-v-e9bb0cc2] {\n  text-transform: capitalize;\n}\nspan.day-time[data-v-e9bb0cc2] {\n  float: right;\n}\n"],sourceRoot:""}])},256:function(t,e,s){"use strict";var n=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"mdc-typography",attrs:{id:"students"}},[s("header1"),t.sessions?t._e():s("spinner"),t.sessions&&!t.sessions.length?s("div",{staticClass:"empty"},[s("img",{attrs:{src:"https://images.weserv.nl/?w=300&url=dl.dropboxusercontent.com/s/6wrsi1smelwdhj7/face-2.png"}}),t._m(0)]):t._e(),s("div",{staticClass:"mdc-list-group"},[t._l(t.sessions,function(e,n){return[s("h3",{staticClass:"mdc-list-group__subheader"},[s("span",{staticClass:"module"},[t._v(t._s(e.class._embedded.module.name))]),t._v(" - "),s("span",{staticClass:"branch"},[t._v(t._s(e.class._embedded.branch.name))]),t._v("  "),s("span",{staticClass:"day-time"},[t._v(t._s(e.class.start_at)+" - "+t._s(e.class.finish_at))])]),s("ul",{staticClass:"mdc-list"},[t._l(e.students,function(n,i){return[s("li",{staticClass:"mdc-list-item",on:{click:function(e){t.onClickList(e)}}},[s("div",{staticClass:"mdc-list-item__graphic"},[s("img",{attrs:{src:n.photo}})]),s("span",{staticClass:"mdc-list-item__text"},[t._v(t._s(n.name))])]),s("hr",{staticClass:"mdc-list-divider"}),s(t.currentView[""+e.id][""+n.id],{tag:"component",attrs:{sid:e.id,uid:n.id,name:n.name}})]})],2)]})],2),s("tab-bottom")],1)},i=[function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("p",[t._v("It’s holiday time, man! "),s("br"),t._v("Pls enjoy yout time by being away from your phone.")])}],a={render:n,staticRenderFns:i};e.a=a},45:function(t,e,s){"use strict";function n(t){s(218)}Object.defineProperty(e,"__esModule",{value:!0});var i=s(188),a=s(256),c=s(11),r=n,o=c(i.a,a.a,!1,r,"data-v-e9bb0cc2",null);e.default=o.exports}});
//# sourceMappingURL=1.74de937dc71c6fd08916.js.map