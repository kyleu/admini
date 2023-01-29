(()=>{function U(){for(let e of Array.from(document.querySelectorAll(".menu-container .final")))e.scrollIntoView({block:"nearest"})}var H="mode-light",x="mode-dark";function F(){for(let e of Array.from(document.getElementsByClassName("mode-input"))){let n=e;n.onclick=function(){switch(n.value){case"":document.body.classList.remove(H),document.body.classList.remove(x);break;case"light":document.body.classList.add(H),document.body.classList.remove(x);break;case"dark":document.body.classList.remove(H),document.body.classList.add(x);break}}}}function B(e,n,t){let o=document.getElementById("flash-container");o===null&&(o=document.createElement("div"),o.id="flash-container",document.body.insertAdjacentElement("afterbegin",o));let r=document.createElement("div");r.className="flash";let i=document.createElement("input");i.type="radio",i.style.display="none",i.id="hide-flash-"+e,r.appendChild(i);let s=document.createElement("label");s.htmlFor="hide-flash-"+e;let c=document.createElement("span");c.innerHTML="\xD7",s.appendChild(c),r.appendChild(s);let l=document.createElement("div");l.className="content flash-"+n,l.innerText=t,r.appendChild(l),o.appendChild(r),j(r)}function $(){let e=document.getElementById("flash-container");if(e===null)return B;let n=e.querySelectorAll(".flash");if(n.length>0)for(let t of n)j(t);return B}function j(e){setTimeout(()=>{e.style.opacity="0",setTimeout(()=>e.remove(),500)},5e3)}function R(){for(let e of Array.from(document.getElementsByClassName("link-confirm"))){let n=e;n.onclick=function(){let t=n.dataset.message;return t&&t.length===0&&(t="Are you sure?"),confirm(t)}}}function h(e,n){let t;n?t=n.querySelectorAll(e):t=document.querySelectorAll(e);let o=[];return t.forEach(r=>{o.push(r)}),o}function I(e,n){let t=h(e,n);switch(t.length){case 0:return;case 1:return t[0];default:console.warn(`found [${t.length}] elements with selector [${e}], wanted zero or one`)}}function p(e,n){let t=I(e,n);if(!t)throw`no element found for selector [${e}]`;return t}function T(e,n,t="block"){return typeof e=="string"&&(e=p(e)),e.style.display=n?t:"none",e}function _(){return h(".reltime").forEach(e=>{w(e.dataset.time||"",e)}),w}function w(e,n){let t=(e||"").replace(/-/g,"/").replace(/[TZ]/g," ")+" UTC",o=new Date(t),r=(new Date().getTime()-o.getTime())/1e3,i=Math.floor(r/86400),s=o.getFullYear(),c=o.getMonth()+1,l=o.getDate();if(isNaN(i)||i<0||i>=31)return s.toString()+"-"+(c<10?"0"+c.toString():c.toString())+"-"+(l<10?"0"+l.toString():l.toString());let a="",u=0;return i==0?r<5?(u=1,a="just now"):r<60?(u=1,a=Math.floor(r)+" seconds ago"):r<120?(u=10,a="1 minute ago"):r<3600?(u=30,a=Math.floor(r/60)+" minutes ago"):r<7200?(u=60,a="1 hour ago"):(u=60,a=Math.floor(r/3600)+" hours ago"):i==1?(u=600,a="yesterday"):i<7?(u=600,a=i+" days ago"):(u=6e3,a=Math.ceil(i/7)+" weeks ago"),n&&(n.innerText=a,setTimeout(()=>w(e,n),u*1e3)),a}function W(){return se}function se(e,n,t,o,r){if(!e)return;let i=e.id+"-list",s=document.createElement("datalist"),c=document.createElement("option");c.value="",c.innerText="Loading...",s.appendChild(c),s.id=i,e.parentElement?.prepend(s),e.setAttribute("autocomplete","off"),e.setAttribute("list",i);let l={},a="";function u(m){let d=n;return d.includes("?")?d+"&t=json&"+t+"="+encodeURIComponent(m):d+"?t=json&"+t+"="+encodeURIComponent(m)}function N(m){let d=l[m];!d||!d.frag||(a=m,s.replaceChildren(d.frag.cloneNode(!0)))}function re(){let m=e.value;if(m.length===0)return;let d=u(m),M=!m||!a;if(!M){let f=l[a];f&&(M=!f.data.find(E=>m===r(E)))}if(!!M){if(l[m]&&l[m].url===d){N(m);return}fetch(d).then(f=>f.json()).then(f=>{if(!f)return;let E=Array.isArray(f)?f:[f],q=document.createDocumentFragment(),D=10;for(let k=0;k<E.length&&D>0;k++){let O=r(E[k]),ie=o(E[k]);if(O){let b=document.createElement("option");b.value=O,b.innerText=ie,q.appendChild(b),D--}}l[m]={url:d,data:E,frag:q,complete:!1},N(m)})}}e.oninput=ae(re,250),console.log("managing ["+e.id+"] autocomplete")}function ae(e,n){let t=0;return function(...o){t!==0&&window.clearTimeout(t),t=window.setTimeout(function(){e(null,...o)},n)}}function G(){document.addEventListener("keydown",e=>{e.key==="Escape"&&document.location.hash.startsWith("#modal-")&&(document.location.hash="")})}function C(e,n){return`<svg class="icon" style="width: ${n}px; height: ${n}px;"><use xlink:href="#svg-${e}"></use></svg>`}function P(e){let n=p("input.result",e),t=p(".tags",e),o=n.value.split(",").map(i=>i.trim()).filter(i=>i!=="");T(n,!1),t.innerHTML="";for(let i of o)t.appendChild(z(i,e));I(".add-item",e)?.remove();let r=document.createElement("div");r.className="add-item",r.innerHTML=C("plus",22),r.onclick=function(){me(t,e)},e.insertBefore(r,p(".clear",e))}function Y(){for(let e of h(".tag-editor"))P(e);return P}function ce(e,n){return e.parentElement!==n.parentElement?null:e===n?0:e.compareDocumentPosition(n)&Node.DOCUMENT_POSITION_FOLLOWING?-1:1}var v;function z(e,n){let t=document.createElement("div");t.className="item",t.draggable=!0,t.ondragstart=function(s){s.dataTransfer?.setDragImage(document.createElement("div"),0,0),t.classList.add("dragging"),v=t},t.ondragover=function(){let s=ce(t,v);if(!s)return;let c=s===-1?t:t.nextSibling;v.parentElement?.insertBefore(v,c),S(n)},t.ondrop=function(s){s.preventDefault()},t.ondragend=function(s){t.classList.remove("dragging"),s.preventDefault()};let o=document.createElement("div");o.innerText=e,o.className="value",o.onclick=function(){J(t)},t.appendChild(o);let r=document.createElement("input");r.className="editor",t.appendChild(r);let i=document.createElement("div");return i.innerHTML=C("times",13),i.className="close",i.onclick=function(){le(t)},t.appendChild(i),t}function le(e){let n=e.parentElement?.parentElement;e.remove(),n&&S(n)}function me(e,n){let t=z("",n);e.appendChild(t),J(t)}function J(e){let n=p(".value",e),t=p(".editor",e);t.value=n.innerText;let o=function(){if(t.value===""){e.remove();return}n.innerText=t.value,T(n,!0),T(t,!1);let r=e.parentElement?.parentElement;r&&S(r)};t.onblur=o,t.onkeydown=function(r){if(r.code==="Enter")return r.preventDefault(),o(),!1},T(n,!1),T(t,!0),t.focus()}function S(e){let n=[],t=h(".item .value",e);for(let r of t)n.push(r.innerText);let o=p("input.result",e);o.value=n.join(", ")}var K="--selected";function ue(e){let n=e.parentElement?.parentElement?.querySelector("input");if(!n)throw"no associated input found";n.value="\u2205"}function A(e){e.onreset=()=>A(e);let n={},t={};for(let o of e.elements){let r=o;if(r.name.length>0)if(r.name.endsWith(K))t[r.name]=r;else{(r.type!=="radio"||r.checked)&&(n[r.name]=r.value);let i=()=>{let s=t[r.name+K];s&&(s.checked=n[r.name]!==r.value)};r.onchange=i,r.onkeyup=i}}}function Q(){for(let e of Array.from(document.querySelectorAll("form.editor")))A(e);return[ue,A]}var de=[];function V(){let e=document.querySelectorAll(".color-var");if(e.length>0)for(let n of Array.from(e)){let t=n,o=t.dataset.var,r=t.dataset.mode;de.push(o),!(!o||o.length===0)&&(t.oninput=function(){fe(r,o,t.value)})}}function fe(e,n,t){let o=document.querySelector("#mockup-"+e);if(!o){console.error("can't find mockup for mode ["+e+"]");return}switch(n){case"color-foreground":g(o,".mock-main",t);break;case"color-background":y(o,".mock-main",t);break;case"color-foreground-muted":g(o,".mock-main .mock-muted",t);break;case"color-background-muted":y(o,".mock-main .mock-muted",t);break;case"color-link-foreground":g(o,".mock-main .mock-link",t);break;case"color-link-visited-foreground":g(o,".mock-main .mock-link-visited",t);break;case"color-nav-foreground":g(o,".mock-nav",t),g(o,".mock-nav .mock-link",t);break;case"color-nav-background":y(o,".mock-nav",t);break;case"color-menu-foreground":g(o,".mock-menu",t),g(o,".mock-menu .mock-link",t);break;case"color-menu-background":y(o,".mock-menu",t);break;case"color-menu-selected-foreground":g(o,".mock-menu .mock-link-selected",t);break;case"color-menu-selected-background":y(o,".mock-menu .mock-link-selected",t);break;default:console.error("invalid key ["+n+"]")}}function Z(e,n,t){let o=e.querySelectorAll(n);if(o.length==0)throw"empty query selector ["+n+"]";o.forEach(r=>t(r))}function y(e,n,t){Z(e,n,o=>o.style.backgroundColor=t)}function g(e,n,t){Z(e,n,o=>o.style.color=t)}function X(){if(window.Sortable){window.admini.sortableEdit=ge;for(let e of Array.from(document.getElementsByClassName("sortable")))ee(e)}}function ge(e){for(;e.parentElement&&!e.classList.contains("drag-container");)e=e.parentElement;e.classList.remove("readonly"),ee(e)}function ee(e){let n=window.Sortable;if(n){let t=e.querySelector(".l");t||(t=e);let r={group:{name:"nested"},handle:".handle",onAdd:s=>{let c=s.item;new n(c.querySelector(".container"),r),c.querySelector(".remove").onclick=function(){te(e,c)},L(e)},onUpdate:()=>L(e),animation:150,fallbackOnBody:!0,swapThreshold:.65};for(let s of Array.from(t.getElementsByClassName("container")))new n(s,r);for(let s of Array.from(t.getElementsByClassName("remove")))s.onclick=function(){te(e,s.parentElement?.parentElement)};let i=e.querySelector(".r");if(i){let s={group:{name:"nested",pull:"clone",put:!1},handle:".handle",animation:150,fallbackOnBody:!0,swapThreshold:.65,sort:!1};for(let c of Array.from(i.getElementsByClassName("container")))new n(c,s)}L(e)}}function te(e,n){n.remove(),L(e)}function L(e){let n=document.querySelector(".drag-state");if(!n)return;let t=document.querySelector(".drag-state-original"),o=e.querySelector(".tracked"),[r,i]=ne(o),s=JSON.stringify(r);if(t){let c=t.value!==s;t.value.length===0&&(t.value=s);let l=document.querySelector(".drag-actions");l&&(c?l.classList.remove("no-changes"):l.classList.add("no-changes"));let a=document.querySelector(".drag-tracked-size");a&&(i===1?a.innerText=i.toString(10)+(a.dataset.sing?" "+a.dataset.sing:""):a.innerText=i.toString(10)+(a.dataset.plur?" "+a.dataset.plur:"")),c?window.onbeforeunload=function(){return!0}:window.onbeforeunload=null}n.value=s}function ne(e){if(e.children.length===0)return[[],0];let n=0,t=[];for(let o of Array.from(e.children))if(o.classList.contains("item")){let[r,i]=pe(o);r&&t.push(r),n+=i}return[t,n]}function pe(e){let n=1,t={k:e.dataset.key,t:e.dataset.title,p:e.dataset.originalPath};for(let o of Array.from(e.children))if(o.classList.contains("container")){let[r,i]=ne(o);r.length>0&&(t.c=r),n+=i}return[t,n]}function oe(){X()}function Ee(){let[e,n]=Q();window.admini={relativeTime:_(),autocomplete:W(),setSiblingToNull:e,initForm:n,flash:$(),tags:Y()},U(),F(),R(),G(),V(),oe()}document.addEventListener("DOMContentLoaded",Ee);})();
//# sourceMappingURL=client.js.map
