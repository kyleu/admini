import "./client.css"
import {appName} from "./util";
import {editorInit} from "./editor";

function menuInit() {
  const x = document.querySelectorAll(".menu-container .final")
  for (const n of x) {
    n.scrollIntoView({block: "nearest"});
  }
}

export function init(): void {
  console.log(appName + "!!!");
  menuInit();
  editorInit();
}

document.addEventListener("DOMContentLoaded", init);
