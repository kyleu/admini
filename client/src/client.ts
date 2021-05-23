import "./client.css"
import {appName} from "./util";
import {menuInit} from "./menu";
import {flashInit} from "./flash";
import {editorInit} from "./editor";
import {sortableInit} from "./sortable";

export function init(): void {
  console.log("[" + appName + "]");
  menuInit();
  flashInit();
  editorInit();
  sortableInit();
}

document.addEventListener("DOMContentLoaded", init);
