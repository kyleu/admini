import {appName} from "./util";
import "./client.css"

function scrollInit() {
  const x = document.querySelectorAll(".menu-container .final")
  for (const n of x) {
    n.scrollIntoView({block: "nearest"});
  }
}

export function init(): void {
  console.log(appName + "!!!");
  scrollInit()
}

document.addEventListener("DOMContentLoaded", init);
