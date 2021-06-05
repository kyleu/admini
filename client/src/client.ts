import "./client.css"
import {menuInit} from "./menu";
import {flashInit} from "./flash";
import {editorInit} from "./editor";
import {sortableEdit, sortableInit} from "./sortable";
import {linkInit} from "./link";
import {themeInit, themeReset} from "./theme";

export function init(): void {
  (window as any).admini = {
    "sortableEdit": sortableEdit,
    "themeReset": themeReset
  };
  menuInit();
  flashInit();
  linkInit();
  editorInit();
  sortableInit();
  themeInit();
}

document.addEventListener("DOMContentLoaded", init);
