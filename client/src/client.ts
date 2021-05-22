import "./client.css"
import {appName} from "./util";
import {menuInit} from "./menu";
import {flashInit} from "./flash";
import {editorInit} from "./editor";
import {dragDropInit} from "./dragdrop";

export function init(): void {
  console.log("[" + appName + "]");
  menuInit();
  flashInit();
  editorInit();
  dragDropInit();
}

document.addEventListener("DOMContentLoaded", init);
