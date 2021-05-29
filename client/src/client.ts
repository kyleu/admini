import "./client.css"
import {appName} from "./util";
import {menuInit} from "./menu";
import {flashInit} from "./flash";
import {editorInit} from "./editor";
import {sortableInit} from "./sortable";
import {linkInit} from "./link";

export function init(): void {
  console.log("[" + appName + "]");
  menuInit();
  flashInit();
  linkInit();
  editorInit();
  sortableInit();
}

document.addEventListener("DOMContentLoaded", init);
