(()=>{function m(){for(let e of Array.from(document.querySelectorAll(".menu-container .final")))e.scrollIntoView({block:"nearest"})}function f(){let e=document.getElementById("flash-container");if(e===null)return;let n=e.querySelectorAll(".flash");n.length>0&&setTimeout(()=>{for(let t of n){let r=t;r.style.opacity="0",setTimeout(()=>r.remove(),500)}},3e3)}var u="--selected";function d(){let e={},n={};for(let t of Array.from(document.querySelectorAll(".editor"))){let r=t,s=()=>{e={},n={};for(let a of r.elements){let o=a;if(o.name.length>0)if(o.name.endsWith(u))n[o.name]=o;else{(o.type!=="radio"||o.checked)&&(e[o.name]=o.value);let l=()=>{let i=n[o.name+u];i&&(i.checked=e[o.name]!==o.value)};o.onchange=l,o.onkeyup=l}}};r.onreset=s,s()}}function p(){if(window.Sortable)for(let e of Array.from(document.getElementsByClassName("sortable")))y(e)}function E(e){for(;e.parentElement&&!e.classList.contains("drag-container");)e=e.parentElement;e.classList.remove("readonly"),y(e)}function y(e){let n=window.Sortable;if(n){let t=e.querySelector(".l");t||(t=e);let s={group:{name:"nested"},handle:".handle",onAdd:o=>{let l=o.item;new n(l.querySelector(".container"),s),l.querySelector(".remove").onclick=function(){g(e,l)},c(e)},onUpdate:()=>c(e),animation:150,fallbackOnBody:!0,swapThreshold:.65};for(let o of Array.from(t.getElementsByClassName("container")))new n(o,s);for(let o of Array.from(t.getElementsByClassName("remove")))o.onclick=function(){g(e,o.parentElement?.parentElement)};let a=e.querySelector(".r");if(a){let o={group:{name:"nested",pull:"clone",put:!1},handle:".handle",animation:150,fallbackOnBody:!0,swapThreshold:.65,sort:!1};for(let l of Array.from(a.getElementsByClassName("container")))new n(l,o)}c(e)}}function g(e,n){n.remove(),c(e)}function c(e){let n=document.querySelector(".drag-state");if(!n)return;let t=document.querySelector(".drag-state-original"),r=e.querySelector(".tracked"),[s,a]=h(r),o=JSON.stringify(s);if(t){t.value.length===0&&(t.value=o);let l=document.querySelector(".drag-actions");l&&(t.value===o?l.classList.add("no-changes"):l.classList.remove("no-changes"));let i=document.querySelector(".drag-tracked-size");i&&(a===1?i.innerText=a.toString(10)+(i.dataset.sing?" "+i.dataset.sing:""):i.innerText=a.toString(10)+(i.dataset.plur?" "+i.dataset.plur:""))}n.value=o}function h(e){let n=0,t=[];for(let r of Array.from(e.children))if(r.classList.contains("item")){let[s,a]=T(r);s&&t.push(s),n+=a}return[t,n]}function T(e){let n=1,t={k:e.dataset.key,p:e.dataset.originalPath};for(let r of Array.from(e.children))if(r.classList.contains("container")){let[s,a]=h(r);s.length>0&&(t.c=s),n+=a}return[t,n]}function I(){for(let e of Array.from(document.getElementsByClassName("link-confirm"))){let n=e;n.onclick=function(){let t=n.dataset.message;return t&&t.length===0&&(t="Are you sure?"),confirm(t)}}}var k=[];function L(){let e=document.querySelectorAll(".color-var");if(e.length>0)for(let n of Array.from(e)){let t=n,r=t.dataset.var;k.push(r),!(!r||r.length===0)&&(t.oninput=function(){document.documentElement.style.setProperty("--"+r,t.value)})}}function v(){for(let e of k)document.documentElement.style.removeProperty("--"+e)}function A(){window.admini={sortableEdit:E,themeReset:v},m(),f(),I(),d(),p(),L()}document.addEventListener("DOMContentLoaded",A);})();
//# sourceMappingURL=client.js.map
