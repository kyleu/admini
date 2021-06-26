import "./client.css"
import {menuInit} from "./menu";
import {flashInit} from "./flash";
import {editorInit} from "./editor";
import {sortableEdit, sortableInit} from "./sortable";
import {linkInit} from "./link";
import {themeInit, themeReset} from "./theme";
import {setSiblingToNull} from "./form";

export function init(): void {
  (window as any).admini = {
    "sortableEdit": sortableEdit,
    "themeReset": themeReset,
    "setSiblingToNull": setSiblingToNull
  };
  menuInit();
  flashInit();
  linkInit();
  editorInit();
  sortableInit();
  themeInit();
}

document.addEventListener("DOMContentLoaded", init);
