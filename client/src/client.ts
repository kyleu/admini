import {appName} from "./util";

export function init(): void {
  console.log(appName + "!!!");
}

document.addEventListener("DOMContentLoaded", init);
