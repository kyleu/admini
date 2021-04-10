import {appName} from "./util";
import "./client.css"

export function init(): void {
  console.log(appName + "!!!");
}

document.addEventListener("DOMContentLoaded", init);
