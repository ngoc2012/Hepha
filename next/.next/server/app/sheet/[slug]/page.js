(()=>{var e={};e.id=638,e.ids=[638],e.modules={2934:e=>{"use strict";e.exports=require("next/dist/client/components/action-async-storage.external.js")},4580:e=>{"use strict";e.exports=require("next/dist/client/components/request-async-storage.external.js")},5869:e=>{"use strict";e.exports=require("next/dist/client/components/static-generation-async-storage.external.js")},399:e=>{"use strict";e.exports=require("next/dist/compiled/next-server/app-page.runtime.prod.js")},852:e=>{"use strict";e.exports=require("async_hooks")},1017:e=>{"use strict";e.exports=require("path")},7310:e=>{"use strict";e.exports=require("url")},9123:(e,t,r)=>{"use strict";r.r(t),r.d(t,{GlobalError:()=>i.a,__next_app__:()=>p,originalPathname:()=>u,pages:()=>d,routeModule:()=>h,tree:()=>c}),r(9754),r(1112),r(6042);var n=r(4285),a=r(2875),s=r(8190),i=r.n(s),l=r(3289),o={};for(let e in l)0>["default","tree","pages","GlobalError","originalPathname","__next_app__","routeModule"].indexOf(e)&&(o[e]=()=>l[e]);r.d(t,o);let c=["",{children:["sheet",{children:["[slug]",{children:["__PAGE__",{},{page:[()=>Promise.resolve().then(r.bind(r,9754)),"/app/src/app/sheet/[slug]/page.tsx"]}]},{}]},{metadata:{icon:[async e=>(await Promise.resolve().then(r.bind(r,3881))).default(e)],apple:[],openGraph:[],twitter:[],manifest:void 0}}]},{layout:[()=>Promise.resolve().then(r.bind(r,1112)),"/app/src/app/layout.tsx"],"not-found":[()=>Promise.resolve().then(r.t.bind(r,6042,23)),"next/dist/client/components/not-found-error"],metadata:{icon:[async e=>(await Promise.resolve().then(r.bind(r,3881))).default(e)],apple:[],openGraph:[],twitter:[],manifest:void 0}}],d=["/app/src/app/sheet/[slug]/page.tsx"],u="/sheet/[slug]/page",p={require:r,loadChunk:()=>Promise.resolve()},h=new n.AppPageRouteModule({definition:{kind:a.x.APP_PAGE,page:"/sheet/[slug]/page",pathname:"/sheet/[slug]",bundlePath:"",filename:"",appPaths:[]},userland:{loaderTree:c}})},8432:(e,t,r)=>{let n={"3f5d2bc3cff514f7c798c09137773583596c841b":()=>Promise.resolve().then(r.bind(r,7812)).then(e=>e.default)};async function a(e,...t){return(await n[e]()).apply(null,t)}e.exports={"3f5d2bc3cff514f7c798c09137773583596c841b":a.bind(null,"3f5d2bc3cff514f7c798c09137773583596c841b")}},9189:(e,t,r)=>{Promise.resolve().then(r.t.bind(r,5157,23)),Promise.resolve().then(r.t.bind(r,2144,23)),Promise.resolve().then(r.t.bind(r,3582,23)),Promise.resolve().then(r.t.bind(r,9587,23)),Promise.resolve().then(r.t.bind(r,5329,23)),Promise.resolve().then(r.t.bind(r,599,23))},1815:()=>{},8643:(e,t,r)=>{Promise.resolve().then(r.bind(r,3094))},3094:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>o});var n=r(7908),a=r(2178);r(5888);let s=(0,r(8412).$)("3f5d2bc3cff514f7c798c09137773583596c841b"),i=({value:e,onChange:t,onKeyDown:r,placeholder:a="",rows:s=4,cols:i=50})=>n.jsx("textarea",{className:"border border-gray-400 rounded",value:e,onChange:t,onKeyDown:r,placeholder:a,rows:s,cols:i}),l=()=>{let[e,t]=(0,a.useState)(""),[r,l]=(0,a.useState)(""),[o,c]=(0,a.useState)("Python"),[d,u]=(0,a.useState)("3"),p=[{lang:"Python",ver:["2","3"]},{lang:"C",ver:[]}],h=async()=>{l(await s("Code",e,o,d))},g=p.find(e=>e.lang===o),x=g?g.ver:[];return(0,n.jsxs)("div",{className:"border-2 border-gray-400 rounded",children:[(0,n.jsxs)("div",{className:"flex space-x-4",children:[(0,n.jsxs)("div",{children:[n.jsx("label",{htmlFor:"lang",children:"Language:"}),n.jsx("select",{id:"lang",value:o,onChange:e=>{let t=e.target.value;c(t);let r=p.find(e=>e.lang===t);r&&r.ver.length>0?u(r.ver[0]):u("")},children:p.map(e=>n.jsx("option",{value:e.lang,children:e.lang},e.lang))})]}),(0,n.jsxs)("div",{children:[n.jsx("label",{htmlFor:"ver",children:"Version:"}),n.jsx("select",{id:"ver",value:d,onChange:e=>{u(e.target.value)},disabled:0===x.length,children:x.map(e=>n.jsx("option",{value:e,children:e},e))})]})]}),n.jsx("button",{className:"bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded",onClick:h,children:"Run"}),n.jsx(i,{value:e,onChange:e=>t(e.target.value),onKeyDown:r=>{if("Tab"===r.key){r.preventDefault();let n=r.currentTarget.selectionStart;t([e.slice(0,n),"	",e.slice(n)].join("")),r.currentTarget.selectionStart=n+1,r.currentTarget.selectionEnd=n+1}}}),n.jsx("div",{id:"output",children:r})]})};function o({params:e}){let[t,...r]=e.slug.split("-"),a=r.join("-");return(0,n.jsxs)(n.Fragment,{children:[n.jsx("div",{children:t}),(0,n.jsxs)("div",{children:["My Sheet : ",a]}),n.jsx(l,{})]})}},1112:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>c,metadata:()=>o});var n=r(9807),a=r(2520),s=r.n(a),i=r(8966),l=r.n(i);r(5023);let o={title:"Hepha",description:"Engineers calculation sheets"};function c({children:e}){return n.jsx("html",{lang:"en",children:n.jsx("body",{className:`${s().variable} ${l().variable}`,children:e})})}},9754:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>n});let n=(0,r(6592).createProxy)(String.raw`/app/src/app/sheet/[slug]/page.tsx#default`)},7812:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>a});var n=r(3722);async function a(e,t,r,n){let a=await fetch("http://"+process.env.GIN_ADDR+"/new_box",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({user:"hepha",sheet:1,type:e,lang:r,ver:n,data:t,caption:"test"}),cache:"no-cache"});if(!a.ok)throw Error("Failed to fetch data");return a.text()}r(453),(0,r(6309).h)([a]),(0,n.j)("3f5d2bc3cff514f7c798c09137773583596c841b",a)},3881:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>a});var n=r(5394);let a=e=>[{type:"image/x-icon",sizes:"48x48",url:(0,n.fillMetadataSegment)(".",e.params,"favicon.ico")+""}]},5023:()=>{}};var t=require("../../../webpack-runtime.js");t.C(e);var r=e=>t(t.s=e),n=t.X(0,[264,769,515],()=>r(9123));module.exports=n})();