(function(t){function e(e){for(var a,s,i=e[0],c=e[1],u=e[2],d=0,v=[];d<i.length;d++)s=i[d],Object.prototype.hasOwnProperty.call(r,s)&&r[s]&&v.push(r[s][0]),r[s]=0;for(a in c)Object.prototype.hasOwnProperty.call(c,a)&&(t[a]=c[a]);l&&l(e);while(v.length)v.shift()();return o.push.apply(o,u||[]),n()}function n(){for(var t,e=0;e<o.length;e++){for(var n=o[e],a=!0,i=1;i<n.length;i++){var c=n[i];0!==r[c]&&(a=!1)}a&&(o.splice(e--,1),t=s(s.s=n[0]))}return t}var a={},r={app:0},o=[];function s(e){if(a[e])return a[e].exports;var n=a[e]={i:e,l:!1,exports:{}};return t[e].call(n.exports,n,n.exports,s),n.l=!0,n.exports}s.m=t,s.c=a,s.d=function(t,e,n){s.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:n})},s.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},s.t=function(t,e){if(1&e&&(t=s(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var n=Object.create(null);if(s.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var a in t)s.d(n,a,function(e){return t[e]}.bind(null,a));return n},s.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return s.d(e,"a",e),e},s.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},s.p="/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],c=i.push.bind(i);i.push=e,i=i.slice();for(var u=0;u<i.length;u++)e(i[u]);var l=c;o.push([0,"chunk-vendors"]),n()})({0:function(t,e,n){t.exports=n("cd49")},"034f":function(t,e,n){"use strict";n("85ec")},2546:function(t,e,n){"use strict";n("72a3")},"65ff":function(t,e,n){"use strict";n("cbea")},"689d":function(t,e,n){"use strict";n("b1d7")},"72a3":function(t,e,n){},"85ec":function(t,e,n){},b1d7:function(t,e,n){},cbea:function(t,e,n){},cd49:function(t,e,n){"use strict";n.r(e);var a={};n.r(a),n.d(a,"API",(function(){return f})),n.d(a,"default",(function(){return m}));n("e260"),n("e6cf"),n("cca6"),n("a79d");var r=n("2b0e"),o=n("f309");r["a"].use(o["a"]);var s=new o["a"]({}),i=(n("d5e8"),n("5363"),n("8c4f")),c=(n("a15b"),n("96cf"),n("1da1")),u=n("d4ec"),l=n("bee2"),d=n("bc3a"),v=n.n(d),p=window.APIPort;void 0===p&&(p=9e3),v.a.defaults.baseURL="http://localhost:"+p+"/",r["a"].prototype.$http=v.a;var f=function(){function t(){Object(u["a"])(this,t)}return Object(l["a"])(t,[{key:"getWordList",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(){var e;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.get("wordlist");case 2:return e=t.sent,t.abrupt("return",e.data.wordList);case 4:case"end":return t.stop()}}),t)})));function e(){return t.apply(this,arguments)}return e}()},{key:"getMessages",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(){var e;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.get("messages");case 2:return e=t.sent,t.abrupt("return",e.data.messages);case 4:case"end":return t.stop()}}),t)})));function e(){return t.apply(this,arguments)}return e}()},{key:"getContacts",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(){var e;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.get("contacts");case 2:return e=t.sent,t.abrupt("return",e.data.contacts);case 4:case"end":return t.stop()}}),t)})));function e(){return t.apply(this,arguments)}return e}()},{key:"deleteContact",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.delete("/contacts",{data:e});case 2:case"end":return t.stop()}}),t)})));function e(e){return t.apply(this,arguments)}return e}()},{key:"updateContact",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.put("/contacts",e);case 2:case"end":return t.stop()}}),t)})));function e(e){return t.apply(this,arguments)}return e}()},{key:"sendMessage",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.post("/sendmessage",e);case 2:case"end":return t.stop()}}),t)})));function e(e){return t.apply(this,arguments)}return e}()},{key:"newContact",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e,n){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.post("/newcontact",{Words:e.join(" "),Alias:n});case 2:case"end":return t.stop()}}),t)})));function e(e,n){return t.apply(this,arguments)}return e}()},{key:"newFollowing",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e,n){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.post("/newfollowing",{Words:e.join(" "),Alias:n});case 2:case"end":return t.stop()}}),t)})));function e(e,n){return t.apply(this,arguments)}return e}()},{key:"deleteAccount",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.delete("/accounts",{data:e});case 2:case"end":return t.stop()}}),t)})));function e(e){return t.apply(this,arguments)}return e}()},{key:"updateAccount",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.put("/accounts",e);case 2:case"end":return t.stop()}}),t)})));function e(e){return t.apply(this,arguments)}return e}()},{key:"createAccount",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.post("/accounts",e);case 2:case"end":return t.stop()}}),t)})));function e(e){return t.apply(this,arguments)}return e}()},{key:"getNewAccounts",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.get("/accounts/new");case 2:return t.abrupt("return",t.sent.data.accounts);case 3:case"end":return t.stop()}}),t)})));function e(){return t.apply(this,arguments)}return e}()},{key:"getAccounts",value:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,v.a.get("/accounts");case 2:return t.abrupt("return",t.sent.data.accounts);case 3:case"end":return t.stop()}}),t)})));function e(){return t.apply(this,arguments)}return e}()}]),t}(),m={install:function(t,e){t.api=new f}},h=n("2f62");r["a"].use(h["a"]);var w,b=new h["a"].Store({state:{messages:[],accounts:[],contacts:[],wordList:[]},mutations:{},actions:{loadWordList:function(){var t=this;return Object(c["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,r["a"].api.getWordList();case 2:t.state.wordList=e.sent;case 3:case"end":return e.stop()}}),e)})))()},loadAccounts:function(){var t=this;return Object(c["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,r["a"].api.getAccounts();case 2:t.state.accounts=e.sent;case 3:case"end":return e.stop()}}),e)})))()},loadContacts:function(){var t=this;return Object(c["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,r["a"].api.getContacts();case 2:t.state.contacts=e.sent;case 3:case"end":return e.stop()}}),e)})))()}}}),g=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-app",[n("div",{staticClass:"menubar d-flex justify-center align-center"},[n("mainmenu"),n("v-tabs",{staticClass:"d-flex justify-center align-center"},[n("v-tab",{attrs:{to:"/"}},[t._v("Postbox")])],1)],1),n("router-view")],1)},x=[],y=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-menu",{attrs:{bottom:"","min-width":"200px",rounded:"","offset-y":"","close-on-content-click":!1},scopedSlots:t._u([{key:"activator",fn:function(e){var a=e.on;return[n("v-btn",t._g({staticClass:"ma-1 ml-5",attrs:{icon:"","x-large":""}},a),[n("v-avatar",{attrs:{color:"teal",size:"48"}},[n("span",{staticClass:"white--text headline"})])],1)]}}])},[n("v-card",[n("v-list-item-content",{staticClass:"justify-center"},[n("div",{staticClass:"mx-auto text-center"},[t._l(t.accounts,(function(e){return n("div",{key:e.address},[n("h3",[t._v(t._s(e.alias))]),n("p",{staticClass:"caption mt-1 my-1 px-2"},[n("copyable",[t._v(" "+t._s(e.address)+" ")])],1),n("v-divider",{staticClass:"my-2"})],1)})),n("manage-accounts"),n("v-divider",{staticClass:"my-2"}),n("manage-contacts"),n("v-divider",{staticClass:"my-2"}),n("v-btn",{attrs:{depressed:"",rounded:"",text:""}},[t._v(" Quit ")])],2)])],1)],1)},k=[],_=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"copyable",class:{copied:t.copied},attrs:{tabindex:"0"},on:{click:t.copy}},[n("div",{staticClass:"background"}),t._t("default"),n("div",{ref:"tooltip",staticClass:"tooltip"},[t._v("Copied")])],2)},C=[],V=(n("4d63"),n("ac1f"),n("25f0"),n("5319"),n("262e")),j=n("2caf"),O=n("9ab4"),R=n("2fe1"),A={copy:function(t){var e=document.createElement("textarea");e.value=t,e.setAttribute("readonly",""),e.style.contain="strict",e.style.position="absolute",e.style.left="-9999px",e.style.fontSize="12pt";var n=document.getSelection(),a=n.rangeCount>0?n.getRangeAt(0):null,r=!document.activeElement||"INPUT"!==document.activeElement.nodeName&&"TEXTAREA"!==document.activeElement.nodeName?null:document.activeElement;document.body.append(e),e.select(),e.selectionStart=0,e.selectionEnd=t.length;var o=!1;try{o=document.execCommand("copy")}catch(s){}return e.remove(),r?r.focus():a&&(n.removeAllRanges(),n.addRange(a)),o}},T=r["a"].extend({props:{text:String}}),$=w=function(t){Object(V["a"])(n,t);var e=Object(j["a"])(n);function n(){var t;return Object(u["a"])(this,n),t=e.apply(this,arguments),t.copied=!1,t._copiedResetTimeout=null,t}return Object(l["a"])(n,[{key:"copy",value:function(){var t=this,e=this.text;if(!e){var n=this.$refs.tooltip.textContent;e=this.$el.innerText.replace(new RegExp("\\s*".concat(n,"$")),"")}A.copy(e),window.clearTimeout(this._copiedResetTimeout),this.copied=!0,this._copiedResetTimeout=window.setTimeout((function(){t.copied=!1}),w.DISPLAY_TIME)}},{key:"onKeyDown",value:function(t){" "!==t.key&&"Enter"!==t.key||this.copy()}},{key:"created",value:function(){this.onKeyDown=this.onKeyDown.bind(this)}},{key:"mounted",value:function(){this.$el.addEventListener("keydown",this.onKeyDown)}},{key:"beforeDestroy",value:function(){this.$el.removeEventListener("keydown",this.onKeyDown)}}]),n}(T);$.DISPLAY_TIME=800,$=w=Object(O["a"])([Object(R["b"])({name:"Copyable"})],$);var I=$,L=I,S=(n("689d"),n("2877")),D=Object(S["a"])(L,_,C,!1,null,"2e3f6e3a",null),E=D.exports,M=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-dialog",{attrs:{"max-width":"1000"},scopedSlots:t._u([{key:"activator",fn:function(e){var a=e.on;return[n("v-btn",t._g({attrs:{depressed:"",rounded:"",text:""}},a),[t._v("Edit Accounts")])]}}]),model:{value:t.open,callback:function(e){t.open=e},expression:"open"}},[n("v-card",[n("v-card-title",[n("span",{staticClass:"headline"},[t._v("Edit Accounts")])]),n("v-card-text",[n("v-list",[n("v-subheader",[t._v("Accounts")]),n("v-list-item-group",{attrs:{color:"primary"}},t._l(t.accounts,(function(e,a){return n("v-list-item",{key:a},[n("v-dialog",{attrs:{"max-width":"400"},scopedSlots:t._u([{key:"activator",fn:function(a){var r=a.on;return[n("v-list-item-content",{staticClass:"d-flex flex-row"},[n("v-list-item-title",[t._v(t._s(e.alias)+" - "),n("span",{staticClass:"caption grey--text",domProps:{textContent:t._s(e.address)}})])],1),n("v-list-item-action",[n("v-btn",t._g({attrs:{icon:""}},r),[n("v-icon",{attrs:{color:"grey lighten-1"}},[t._v("mdi-pencil")])],1)],1)]}}],null,!0),model:{value:t.editopen[e.address],callback:function(n){t.$set(t.editopen,e.address,n)},expression:"editopen[account.address]"}},[n("v-card",[n("v-card-title",[n("span",{staticClass:"headline"},[t._v("Edit Account")])]),n("v-card-text",[n("v-text-field",{attrs:{label:"Name"},model:{value:e.alias,callback:function(n){t.$set(e,"alias",n)},expression:"account.alias"}})],1),n("v-card-actions",[n("v-dialog",{attrs:{"max-width":"290"},scopedSlots:t._u([{key:"activator",fn:function(e){var a=e.on,r=e.attrs;return[n("v-btn",t._g(t._b({attrs:{color:"red darken-1",text:""}},"v-btn",r,!1),a),[t._v(" Delete ")])]}}],null,!0),model:{value:t.sureDialog[e.address],callback:function(n){t.$set(t.sureDialog,e.address,n)},expression:"sureDialog[account.address]"}},[n("v-card",[n("v-card-title",{staticClass:"headline"},[t._v(" Are you Sure? ")]),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{color:"green darken-1",text:""},on:{click:function(n){t.sureDialog[e.address]=!1}}},[t._v(" No ")]),n("v-btn",{attrs:{color:"green darken-1",text:""},on:{click:function(n){t.sureDialog[e.address]=!1,t.deleteAccount(e)}}},[t._v(" Yes ")])],1)],1)],1),n("v-spacer"),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(n){t.editopen[e.address]=!1}}},[t._v(" Close ")]),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(n){t.editopen[e.address]=!1,t.updateAccount(e)}}},[t._v(" Save ")])],1)],1)],1)],1)})),1)],1),n("new-account")],1),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(e){t.open=!1}}},[t._v(" Close ")])],1)],1)],1)},N=[],W=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-dialog",{attrs:{"max-width":"700"},scopedSlots:t._u([{key:"activator",fn:function(e){var a=e.on;return[n("v-btn",t._g({attrs:{rounded:"",outlined:"",elevation:"1"}},a),[t._v("+ Account")])]}}]),model:{value:t.open,callback:function(e){t.open=e},expression:"open"}},[n("v-card",[n("v-card-title",[n("span",{staticClass:"headline"},[t._v("+ Account")])]),n("v-card-text",[n("v-container",[n("v-row",[n("v-col",{attrs:{cols:"12"}},[n("v-text-field",{attrs:{label:"Name"},model:{value:t.alias,callback:function(e){t.alias=e},expression:"alias"}})],1)],1)],1)],1),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(e){t.open=!1}}},[t._v(" Close ")]),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:t.newAccount}},[t._v(" Add ")])],1)],1)],1)},P=[],F=r["a"].extend({components:{},props:{},data:function(){return{open:!1,words:[],alias:""}},methods:{newAccount:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return this.open=!1,t.next=3,r["a"].api.createAccount({alias:this.alias});case 3:this.$store.dispatch("loadAccounts");case 4:case"end":return t.stop()}}),t,this)})));function e(){return t.apply(this,arguments)}return e}()}}),B=F,K=n("6544"),U=n.n(K),Y=n("8336"),z=n("b0af"),G=n("99d9"),H=n("62ad"),J=n("a523"),Q=n("169a"),X=n("0fd9"),q=n("2fa4"),Z=n("8654"),tt=Object(S["a"])(B,W,P,!1,null,"33fdef92",null),et=tt.exports;U()(tt,{VBtn:Y["a"],VCard:z["a"],VCardActions:G["a"],VCardText:G["b"],VCardTitle:G["c"],VCol:H["a"],VContainer:J["a"],VDialog:Q["a"],VRow:X["a"],VSpacer:q["a"],VTextField:Z["a"]});var nt=r["a"].extend({components:{NewAccount:et},props:{},computed:Object(h["b"])(["accounts"]),data:function(){return{open:!1,sureDialog:[],editopen:[]}},created:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:case"end":return t.stop()}}),t)})));function e(){return t.apply(this,arguments)}return e}(),methods:{updateAccount:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,r["a"].api.updateAccount(e);case 2:this.$store.dispatch("loadAccounts");case 3:case"end":return t.stop()}}),t,this)})));function e(e){return t.apply(this,arguments)}return e}(),deleteAccount:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,r["a"].api.deleteAccount(e);case 2:this.$store.dispatch("loadAccounts");case 3:case"end":return t.stop()}}),t,this)})));function e(e){return t.apply(this,arguments)}return e}()}}),at=nt,rt=n("132d"),ot=n("8860"),st=n("da13"),it=n("1800"),ct=n("5d23"),ut=n("1baa"),lt=n("e0c7"),dt=Object(S["a"])(at,M,N,!1,null,"047bc807",null),vt=dt.exports;U()(dt,{VBtn:Y["a"],VCard:z["a"],VCardActions:G["a"],VCardText:G["b"],VCardTitle:G["c"],VDialog:Q["a"],VIcon:rt["a"],VList:ot["a"],VListItem:st["a"],VListItemAction:it["a"],VListItemContent:ct["b"],VListItemGroup:ut["a"],VListItemTitle:ct["d"],VSpacer:q["a"],VSubheader:lt["a"],VTextField:Z["a"]});var pt=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-dialog",{attrs:{"max-width":"1000"},scopedSlots:t._u([{key:"activator",fn:function(e){var a=e.on;return[n("v-btn",t._g({attrs:{depressed:"",rounded:"",text:""}},a),[t._v("Manage Contacts")])]}}]),model:{value:t.open,callback:function(e){t.open=e},expression:"open"}},[n("v-card",[n("v-card-title",[n("span",{staticClass:"headline"},[t._v("Manage Contacts")])]),n("v-card-text",[n("v-list",[n("v-subheader",[t._v("Contacts")]),n("v-list-item-group",{attrs:{color:"primary"}},t._l(t.contacts,(function(e,a){return n("v-list-item",{key:a},[n("v-dialog",{attrs:{"max-width":"400"},scopedSlots:t._u([{key:"activator",fn:function(a){var r=a.on;return[n("v-list-item-content",{staticClass:"d-flex flex-row"},[n("v-list-item-title",[t._v(t._s(e.alias)+" - "),n("span",{staticClass:"caption grey--text",domProps:{textContent:t._s(e.words)}})])],1),n("v-list-item-action",[n("v-btn",t._g({attrs:{icon:""}},r),[n("v-icon",{attrs:{color:"grey lighten-1"}},[t._v("mdi-pencil")])],1)],1)]}}],null,!0),model:{value:t.editopen[e.words],callback:function(n){t.$set(t.editopen,e.words,n)},expression:"editopen[contact.words]"}},[n("v-card",[n("v-card-title",[n("span",{staticClass:"headline"},[t._v("Edit Contact")])]),n("v-card-text",[n("v-text-field",{attrs:{label:"Name"},model:{value:e.alias,callback:function(n){t.$set(e,"alias",n)},expression:"contact.alias"}})],1),n("v-card-actions",[n("v-dialog",{attrs:{"max-width":"290"},scopedSlots:t._u([{key:"activator",fn:function(e){var a=e.on,r=e.attrs;return[n("v-btn",t._g(t._b({attrs:{color:"red darken-1",text:""}},"v-btn",r,!1),a),[t._v(" Delete ")])]}}],null,!0),model:{value:t.sureDialog[e.words],callback:function(n){t.$set(t.sureDialog,e.words,n)},expression:"sureDialog[contact.words]"}},[n("v-card",[n("v-card-title",{staticClass:"headline"},[t._v(" Are you Sure? ")]),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{color:"green darken-1",text:""},on:{click:function(n){t.sureDialog[e.words]=!1}}},[t._v(" No ")]),n("v-btn",{attrs:{color:"green darken-1",text:""},on:{click:function(n){t.sureDialog[e.words]=!1,t.deleteContact(e)}}},[t._v(" Yes ")])],1)],1)],1),n("v-spacer"),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(n){t.editopen[e.words]=!1}}},[t._v(" Close ")]),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(n){t.editopen[e.words]=!1,t.updateContact(e)}}},[t._v(" Save ")])],1)],1)],1)],1)})),1)],1),n("new-contact")],1),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(e){t.open=!1}}},[t._v(" Close ")])],1)],1)],1)},ft=[],mt=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-dialog",{attrs:{"max-width":"500"},scopedSlots:t._u([{key:"activator",fn:function(e){var a=e.on;return[n("v-btn",t._g({attrs:{rounded:"",small:"",outlined:"",elevation:"1"}},a),[t._v("+ New Contact")])]}}]),model:{value:t.open,callback:function(e){t.open=e},expression:"open"}},[n("v-card",[n("v-card-title",[n("span",{staticClass:"headline"},[t._v("+ New Contact")])]),n("v-card-text",[n("v-container",[n("v-row",[n("v-col",{attrs:{cols:"12"}},[n("words-autocomplete",{model:{value:t.words,callback:function(e){t.words=e},expression:"words"}}),n("v-text-field",{attrs:{label:"Name"},model:{value:t.addressAlias,callback:function(e){t.addressAlias=e},expression:"addressAlias"}})],1)],1)],1)],1),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(e){t.open=!1}}},[t._v(" Close ")]),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:t.newContact}},[t._v(" Add ")])],1)],1)],1)},ht=[],wt=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("div",{staticClass:"words d-flex align-center"},t._l(t.words,(function(e,a){return n("WordsAutocompleteWord",{key:a,ref:"word"+a,refInFor:!0,attrs:{index:a},on:{done:function(e){return t.wordCompleted(a)},overflow:t.pushOverflowToNextWord},model:{value:t.words[a],callback:function(e){t.$set(t.words,a,e)},expression:"words[index]"}})})),1)])},bt=[],gt=(n("99af"),n("1b40")),xt=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"word mx-1 px-2 d-flex align-center"},[n("v-text-field",{directives:[{name:"show",rawName:"v-show",value:!t.done,expression:"!done"}],ref:"input",staticClass:"py-1",attrs:{dense:"","hide-details":""},on:{input:t.inputCheck,blur:t.update,keydown:t.checkKeydown},model:{value:t.word,callback:function(e){t.word=e},expression:"word"}}),n("span",{directives:[{name:"show",rawName:"v-show",value:t.done,expression:"done"}],on:{click:t.editField}},[t._v(t._s(t.word))]),0===t.filteredItems.length||0===t.word.length||t.done?t._e():n("v-card",{staticClass:"autocomplete-modal"},[n("v-list",{attrs:{dense:""}},t._l(t.filteredItems,(function(e,a){return n("v-list-item",{key:e,class:{active:a==t.suggestIndex},on:{click:function(n){return t.selectWord(e)}}},[t._v(t._s(e))])})),1)],1)],1)},yt=[],kt=(n("4de4"),n("caad"),n("a9e3"),n("2532"),n("1276"),n("2ca0"),n("498a"),function(t){Object(V["a"])(n,t);var e=Object(j["a"])(n);function n(){var t;return Object(u["a"])(this,n),t=e.apply(this,arguments),t.word="",t.editing=!0,t.suggestIndex=0,t}return Object(l["a"])(n,[{key:"created",value:function(){this.word=this.value}},{key:"update",value:function(){this.$emit("input",this.word)}},{key:"focus",value:function(){this.input.focus()}},{key:"insert",value:function(t){this.word=t,this.validate()}},{key:"inputCheck",value:function(t){console.log(t)}},{key:"editField",value:function(){var t=this;this.editing=!0,this.$nextTick((function(){t.focus()}))}},{key:"selectWord",value:function(t){var e=this;this.word=t,this.$nextTick((function(){e.validate()}))}},{key:"suggestDown",value:function(){this.suggestIndex++}},{key:"suggestUp",value:function(){this.suggestIndex--}},{key:"validate",value:function(){this.word=this.word.toLowerCase(),this.word=this.word.trim();var t=this.$store.state.wordList,e=this.word.split(" ");if(e.length>1){this.word=e[0],e.shift();var n=e.join(" ");" "!==n&&this.$emit("overflow",{text:n,index:this.index})}console.log(t,this.word),t.includes(this.word)&&(this.$emit("input",this.word),this.$emit("done"),this.editing=!1)}},{key:"checkKeydown",value:function(t){var e=this;switch(console.log(t.which),t.which){case 13:case 9:t.preventDefault(),this.word=this.filteredItems[this.suggestIndex],this.$nextTick((function(){e.validate()}));break;case 32:this.validate();break;case 38:this.suggestUp();break;case 40:this.suggestDown();break}}},{key:"done",get:function(){return this.valid&&!this.editing}},{key:"valid",get:function(){var t=this.$store.state.wordList;return t.includes(this.word)}},{key:"filteredItems",get:function(){var t=this,e=this.$store.state.wordList;return e.filter((function(e){return e.startsWith(t.word)})).filter((function(t,e){return e<3}))}}]),n}(gt["d"]));Object(O["a"])([Object(gt["b"])({type:String,default:""})],kt.prototype,"value",void 0),Object(O["a"])([Object(gt["b"])({type:Number,default:0})],kt.prototype,"index",void 0),Object(O["a"])([Object(gt["c"])("input")],kt.prototype,"input",void 0),kt=Object(O["a"])([Object(gt["a"])({})],kt);var _t=kt,Ct=_t,Vt=(n("fb13"),Object(S["a"])(Ct,xt,yt,!1,null,"7e8f6dc1",null)),jt=Vt.exports;U()(Vt,{VCard:z["a"],VList:ot["a"],VListItem:st["a"],VTextField:Z["a"]});var Ot=function(t){Object(V["a"])(n,t);var e=Object(j["a"])(n);function n(){var t;return Object(u["a"])(this,n),t=e.apply(this,arguments),t.words=[],t}return Object(l["a"])(n,[{key:"wordsWatcher",value:function(t){this.$emit("input",t)}},{key:"created",value:function(){this.words=this.value;while(this.words.length<5)this.words=this.words.concat([""])}},{key:"wordCompleted",value:function(t){console.log("word"+(t+1));var e=this.$refs["word"+(t+1)];e&&e[0].focus()}},{key:"pushOverflowToNextWord",value:function(t){var e=t.index,n=t.text;this.$refs["word"+(e+1)][0].insert(n)}}]),n}(gt["d"]);Object(O["a"])([Object(gt["b"])({type:Array,default:[]})],Ot.prototype,"value",void 0),Object(O["a"])([Object(gt["c"])("word1")],Ot.prototype,"word1",void 0),Object(O["a"])([Object(gt["c"])("word2")],Ot.prototype,"word2",void 0),Object(O["a"])([Object(gt["c"])("word3")],Ot.prototype,"word3",void 0),Object(O["a"])([Object(gt["c"])("word4")],Ot.prototype,"word4",void 0),Object(O["a"])([Object(gt["c"])("word5")],Ot.prototype,"word5",void 0),Object(O["a"])([Object(gt["e"])("words")],Ot.prototype,"wordsWatcher",null),Ot=Object(O["a"])([Object(gt["a"])({components:{WordsAutocompleteWord:jt}})],Ot);var Rt=Ot,At=Rt,Tt=(n("e3d9"),Object(S["a"])(At,wt,bt,!1,null,"91be55e4",null)),$t=Tt.exports,It=$t,Lt=r["a"].extend({components:{WordsAutocomplete:It},props:{},data:function(){return{open:!1,words:[],addressAlias:""}},methods:{newContact:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return this.open=!1,t.next=3,r["a"].api.newContact(this.words,this.addressAlias);case 3:this.$store.dispatch("loadContacts");case 4:case"end":return t.stop()}}),t,this)})));function e(){return t.apply(this,arguments)}return e}()}}),St=Lt,Dt=Object(S["a"])(St,mt,ht,!1,null,"8b1ffbc4",null),Et=Dt.exports;U()(Dt,{VBtn:Y["a"],VCard:z["a"],VCardActions:G["a"],VCardText:G["b"],VCardTitle:G["c"],VCol:H["a"],VContainer:J["a"],VDialog:Q["a"],VRow:X["a"],VSpacer:q["a"],VTextField:Z["a"]});var Mt=r["a"].extend({components:{NewContact:Et},props:{},computed:Object(h["b"])(["contacts"]),data:function(){return{open:!1,sureDialog:[],editopen:[]}},created:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:case"end":return t.stop()}}),t)})));function e(){return t.apply(this,arguments)}return e}(),methods:{updateContact:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,r["a"].api.updateContact(e);case 2:this.$store.dispatch("loadContacts");case 3:case"end":return t.stop()}}),t,this)})));function e(e){return t.apply(this,arguments)}return e}(),deleteContact:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(e){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,r["a"].api.deleteContact(e);case 2:this.$store.dispatch("loadContacts");case 3:case"end":return t.stop()}}),t,this)})));function e(e){return t.apply(this,arguments)}return e}()}}),Nt=Mt,Wt=Object(S["a"])(Nt,pt,ft,!1,null,"605a54ca",null),Pt=Wt.exports;U()(Wt,{VBtn:Y["a"],VCard:z["a"],VCardActions:G["a"],VCardText:G["b"],VCardTitle:G["c"],VDialog:Q["a"],VIcon:rt["a"],VList:ot["a"],VListItem:st["a"],VListItemAction:it["a"],VListItemContent:ct["b"],VListItemGroup:ut["a"],VListItemTitle:ct["d"],VSpacer:q["a"],VSubheader:lt["a"],VTextField:Z["a"]});var Ft=r["a"].extend({components:{Copyable:E,ManageAccounts:vt,ManageContacts:Pt},props:{},computed:Object(h["b"])(["accounts"]),data:function(){return{}},methods:{}}),Bt=Ft,Kt=n("8212"),Ut=n("ce7e"),Yt=n("e449"),zt=Object(S["a"])(Bt,y,k,!1,null,"6ce99022",null),Gt=zt.exports;U()(zt,{VAvatar:Kt["a"],VBtn:Y["a"],VCard:z["a"],VDivider:Ut["a"],VListItemContent:ct["b"],VMenu:Yt["a"]});var Ht=r["a"].extend({name:"App",components:{mainmenu:Gt},data:function(){return{}},created:function(){this.$store.dispatch("loadWordList"),this.$store.dispatch("loadAccounts"),this.$store.dispatch("loadContacts")}}),Jt=Ht,Qt=(n("034f"),n("7496")),Xt=n("71a3"),qt=n("fe57"),Zt=Object(S["a"])(Jt,g,x,!1,null,null,null),te=Zt.exports;U()(Zt,{VApp:Qt["a"],VTab:Xt["a"],VTabs:qt["a"]});var ee=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-main",[n("v-container",[n("div",{staticClass:"messages"},[n("new-message"),n("h2",{staticClass:"mb-1 ml-4"},[t._v("Messages")]),n("v-list",t._l(t.messages,(function(e){return n("div",{key:e.id},[n("v-dialog",{attrs:{"max-width":"1200"},scopedSlots:t._u([{key:"activator",fn:function(a){var r=a.on;return[n("v-list-item",t._g({attrs:{link:""}},r),[n("v-list-item-content",[n("v-list-item-title",[t._v(t._s(e.topic))]),n("v-list-item-subtitle",[n("div",[n("b",[t._v(t._s(e.name))]),t._v(" - "+t._s(e.content))])])],1),n("v-list-item-action",[n("v-list-item-action-text",[t._v(t._s(e.time))])],1)],1)]}}],null,!0)},[n("v-card",[n("v-card-title",{staticClass:"headline grey lighten-2"},[t._v(" "+t._s(e.topic)+" ")]),n("v-card-text",[t._v(" "+t._s(e.content)+" ")]),n("v-divider")],1)],1)],1)})),0)],1)])],1)},ne=[],ae=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-dialog",{attrs:{"max-width":"1200"},scopedSlots:t._u([{key:"activator",fn:function(e){var a=e.on;return[n("v-btn",t._g({staticClass:"writebtn",attrs:{rounded:"",outlined:"",elevation:"1"}},a),[t._v("+ Write")])]}}]),model:{value:t.open,callback:function(e){t.open=e},expression:"open"}},[n("v-card",[n("v-card-title",[n("span",{staticClass:"headline"},[t._v("Write")])]),n("v-card-text",[n("v-container",[n("v-row",[n("v-col",{attrs:{cols:"12",sm:"6",md:"4"}},[n("new-contact"),n("br"),n("br"),n("v-autocomplete",{attrs:{label:"Reciever","hide-selected":"",multiple:"",dense:"",chips:"","deletable-chips":"",items:t.contacts,"search-input":t.recieverSearch},on:{"update:searchInput":function(e){t.recieverSearch=e},"update:search-input":function(e){t.recieverSearch=e},change:function(e){t.recieverSearch=""}},model:{value:t.newmessage.reciever,callback:function(e){t.$set(t.newmessage,"reciever",e)},expression:"newmessage.reciever"}})],1),n("v-col",{attrs:{cols:"12"}},[n("v-text-field",{attrs:{label:"Title"},model:{value:t.newmessage.title,callback:function(e){t.$set(t.newmessage,"title",e)},expression:"newmessage.title"}}),n("v-textarea",{attrs:{label:"Text"},model:{value:t.newmessage.content,callback:function(e){t.$set(t.newmessage,"content",e)},expression:"newmessage.content"}})],1)],1)],1)],1),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:function(e){t.open=!1}}},[t._v(" Close ")]),n("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:t.sendMessage}},[t._v(" Send ")])],1)],1)],1)},re=[],oe=(n("d81d"),r["a"].extend({components:{NewContact:Et},props:{},data:function(){return{open:!1,contacts:[],recieverSearch:"",newmessage:{reciever:[],content:"",title:""},reloadTimer:null}},created:function(){var t=this;this.reloadTimer=setInterval(Object(c["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,r["a"].api.getContacts();case 2:t.contacts=e.sent,t.contacts=t.contacts.map((function(t){return t.alias}));case 4:case"end":return e.stop()}}),e)}))),1e4)},beforeDestroy:function(){clearInterval(this.reloadTimer)},methods:{sendMessage:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return this.open=!1,t.next=3,r["a"].api.sendMessage(this.newmessage);case 3:case"end":return t.stop()}}),t,this)})));function e(){return t.apply(this,arguments)}return e}()}})),se=oe,ie=(n("2546"),n("c6a6")),ce=n("a844"),ue=Object(S["a"])(se,ae,re,!1,null,"9c16fb82",null),le=ue.exports;U()(ue,{VAutocomplete:ie["a"],VBtn:Y["a"],VCard:z["a"],VCardActions:G["a"],VCardText:G["b"],VCardTitle:G["c"],VCol:H["a"],VContainer:J["a"],VDialog:Q["a"],VRow:X["a"],VSpacer:q["a"],VTextField:Z["a"],VTextarea:ce["a"]});var de=r["a"].extend({name:"Postbox",components:{NewMessage:le},data:function(){return{messages:[],reloadTimer:null}},created:function(){var t=Object(c["a"])(regeneratorRuntime.mark((function t(){var e=this;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,r["a"].api.getMessages();case 2:this.messages=t.sent,this.reloadTimer=setInterval(Object(c["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,r["a"].api.getMessages();case 2:e.messages=t.sent;case 3:case"end":return t.stop()}}),t)}))),1e4);case 4:case"end":return t.stop()}}),t,this)})));function e(){return t.apply(this,arguments)}return e}(),beforeDestroy:function(){clearInterval(this.reloadTimer)}}),ve=de,pe=n("f6c4"),fe=Object(S["a"])(ve,ee,ne,!1,null,"a9d77ab0",null),me=fe.exports;U()(fe,{VCard:z["a"],VCardText:G["b"],VCardTitle:G["c"],VContainer:J["a"],VDialog:Q["a"],VDivider:Ut["a"],VList:ot["a"],VListItem:st["a"],VListItemAction:it["a"],VListItemActionText:ct["a"],VListItemContent:ct["b"],VListItemSubtitle:ct["c"],VListItemTitle:ct["d"],VMain:pe["a"]});var he=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-main",[n("v-container",[n("div",{staticClass:"messages"},[n("v-btn",{staticClass:"writebtn",attrs:{rounded:"",outlined:"",elevation:"1"}},[t._v("+ Publish")]),n("h2",{staticClass:"mb-1 ml-4"},[t._v("New")]),n("v-list",t._l(5,(function(e){return n("div",{key:e},[n("v-list-item",{key:e,attrs:{link:""}},[n("v-list-item-content",[n("v-list-item-title",[t._v("Test Message Title #"+t._s(e))]),n("v-list-item-subtitle",[n("div",[n("b",[t._v("Max Mustermann")]),t._v(" - Hey Bob I just wanted to write you a message because ...")])])],1),n("v-list-item-action",[n("v-list-item-action-text",[t._v("Tue 13:20")])],1)],1)],1)})),0),n("h2",{staticClass:"mb-1 ml-4"},[t._v("Seen")]),n("v-list",t._l(5,(function(e){return n("div",{key:e},[n("v-list-item",{key:e,attrs:{link:""}},[n("v-list-item-content",[n("v-list-item-title",[t._v("Test Message Title #"+t._s(e))]),n("v-list-item-subtitle",[n("div",[n("b",[t._v("Max Mustermann")]),t._v(" - Hey Bob I just wanted to write you a message because ...")])])],1),n("v-list-item-action",[n("v-list-item-action-text",[t._v("Tue 13:20")])],1)],1)],1)})),0)],1)])],1)},we=[],be=r["a"].extend({name:"Newsletter",components:{},data:function(){return{}}}),ge=be,xe=(n("65ff"),Object(S["a"])(ge,he,we,!1,null,"dab80fa8",null)),ye=xe.exports;U()(xe,{VBtn:Y["a"],VContainer:J["a"],VList:ot["a"],VListItem:st["a"],VListItemAction:it["a"],VListItemActionText:ct["a"],VListItemContent:ct["b"],VListItemSubtitle:ct["c"],VListItemTitle:ct["d"],VMain:pe["a"]}),r["a"].use(i["a"]),r["a"].use(m),r["a"].config.productionTip=!1;var ke=new i["a"]({mode:"history",routes:[{path:"/",component:me},{path:"/newsletter",component:ye}]});new r["a"]({router:ke,vuetify:s,store:b,render:function(t){return t(te)}}).$mount("#app")},d863:function(t,e,n){},e3d9:function(t,e,n){"use strict";n("e7be")},e7be:function(t,e,n){},fb13:function(t,e,n){"use strict";n("d863")}});
//# sourceMappingURL=app.e345bc57.js.map