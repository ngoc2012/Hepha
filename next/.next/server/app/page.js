(()=>{var e={};e.id=931,e.ids=[931],e.modules={2934:e=>{"use strict";e.exports=require("next/dist/client/components/action-async-storage.external.js")},4580:e=>{"use strict";e.exports=require("next/dist/client/components/request-async-storage.external.js")},5869:e=>{"use strict";e.exports=require("next/dist/client/components/static-generation-async-storage.external.js")},399:e=>{"use strict";e.exports=require("next/dist/compiled/next-server/app-page.runtime.prod.js")},852:e=>{"use strict";e.exports=require("async_hooks")},1017:e=>{"use strict";e.exports=require("path")},7310:e=>{"use strict";e.exports=require("url")},5267:(e,t,r)=>{"use strict";r.r(t),r.d(t,{GlobalError:()=>i.a,__next_app__:()=>u,originalPathname:()=>p,pages:()=>c,routeModule:()=>h,tree:()=>l}),r(9306),r(1112),r(6042);var a=r(4285),n=r(2875),s=r(8190),i=r.n(s),o=r(3289),d={};for(let e in o)0>["default","tree","pages","GlobalError","originalPathname","__next_app__","routeModule"].indexOf(e)&&(d[e]=()=>o[e]);r.d(t,d);let l=["",{children:["__PAGE__",{},{page:[()=>Promise.resolve().then(r.bind(r,9306)),"/app/src/app/page.tsx"],metadata:{icon:[async e=>(await Promise.resolve().then(r.bind(r,3881))).default(e)],apple:[],openGraph:[],twitter:[],manifest:void 0}}]},{layout:[()=>Promise.resolve().then(r.bind(r,1112)),"/app/src/app/layout.tsx"],"not-found":[()=>Promise.resolve().then(r.t.bind(r,6042,23)),"next/dist/client/components/not-found-error"],metadata:{icon:[async e=>(await Promise.resolve().then(r.bind(r,3881))).default(e)],apple:[],openGraph:[],twitter:[],manifest:void 0}}],c=["/app/src/app/page.tsx"],p="/page",u={require:r,loadChunk:()=>Promise.resolve()},h=new a.AppPageRouteModule({definition:{kind:n.x.APP_PAGE,page:"/page",pathname:"/",bundlePath:"",filename:"",appPaths:[]},userland:{loaderTree:l}})},5220:(e,t,r)=>{let a={"8e48100e65dbfae75906df9d78a8a2c63bb9f32b":()=>Promise.resolve().then(r.bind(r,4789)).then(e=>e.default),"69eb91a966024d02e51862c207b71bc8c8687d24":()=>Promise.resolve().then(r.bind(r,6518)).then(e=>e.default)};async function n(e,...t){return(await a[e]()).apply(null,t)}e.exports={"8e48100e65dbfae75906df9d78a8a2c63bb9f32b":n.bind(null,"8e48100e65dbfae75906df9d78a8a2c63bb9f32b"),"69eb91a966024d02e51862c207b71bc8c8687d24":n.bind(null,"69eb91a966024d02e51862c207b71bc8c8687d24")}},9189:(e,t,r)=>{Promise.resolve().then(r.t.bind(r,5157,23)),Promise.resolve().then(r.t.bind(r,2144,23)),Promise.resolve().then(r.t.bind(r,3582,23)),Promise.resolve().then(r.t.bind(r,9587,23)),Promise.resolve().then(r.t.bind(r,5329,23)),Promise.resolve().then(r.t.bind(r,599,23))},1815:()=>{},6464:(e,t,r)=>{Promise.resolve().then(r.bind(r,2284))},2284:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>d});var a=r(7908),n=r(2178),s=r(9877);r(5888);var i=r(8412);(0,i.$)("8e48100e65dbfae75906df9d78a8a2c63bb9f32b");let o=(0,i.$)("69eb91a966024d02e51862c207b71bc8c8687d24");function d(){let e=(0,s.useRouter)(),[t,r]=(0,n.useState)([]),i=async()=>{let t=await o("hepha","No title","No description");if(t.error){console.error(t.error);return}console.log(t),e.push("/sheet/"+t.slug)};return(0,a.jsxs)("div",{children:[a.jsx("button",{className:"bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded",onClick:i,children:"New Sheet"}),a.jsx("h1",{children:"Sheet List"}),a.jsx("ul",{children:t.map(e=>a.jsx("li",{children:(0,a.jsxs)("div",{className:"flex space-x-4",children:[a.jsx("h2",{children:e.title}),a.jsx("p",{children:e.description}),(0,a.jsxs)("p",{children:[" by ",e.owner]})]})},e.id))})]})}},1112:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>l,metadata:()=>d});var a=r(9807),n=r(2520),s=r.n(n),i=r(8966),o=r.n(i);r(5023);let d={title:"Hepha",description:"Engineers calculation sheets"};function l({children:e}){return a.jsx("html",{lang:"en",children:a.jsx("body",{className:`${s().variable} ${o().variable}`,children:e})})}},9306:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>a});let a=(0,r(6592).createProxy)(String.raw`/app/src/app/page.tsx#default`)},4789:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>n});var a=r(3722);async function n(e,t){let r=await fetch("http://"+process.env.GIN_ADDR+"/sheets",{method:"POST",headers:{"Content-Type":"text/plain"},body:JSON.stringify({name:"",length:e,page:t}),cache:"no-cache"});if(!r.ok)throw Error("Failed to fetch data");let a=await r.json();return a.error?(console.error(a.error),[]):a}r(453),(0,r(6309).h)([n]),(0,a.j)("8e48100e65dbfae75906df9d78a8a2c63bb9f32b",n)},6518:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>n});var a=r(3722);async function n(e,t,r){let a=await fetch("http://"+process.env.GIN_ADDR+"/new_sheet",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({owner:e,title:t,description:r}),cache:"no-cache"});if(!a.ok)throw Error("Failed to fetch data");let n=await a.json();return n.error?(console.error(n.error),[]):n}r(453),(0,r(6309).h)([n]),(0,a.j)("69eb91a966024d02e51862c207b71bc8c8687d24",n)},3881:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>n});var a=r(5394);let n=e=>[{type:"image/x-icon",sizes:"48x48",url:(0,a.fillMetadataSegment)(".",e.params,"favicon.ico")+""}]},5023:()=>{}};var t=require("../webpack-runtime.js");t.C(e);var r=e=>t(t.s=e),a=t.X(0,[264,769,515],()=>r(5267));module.exports=a})();