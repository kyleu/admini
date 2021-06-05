import "./client.css"
import {menuInit} from "./menu";
import {flashInit} from "./flash";
import {editorInit} from "./editor";
import {sortableEdit, sortableInit} from "./sortable";
import {linkInit} from "./link";

export function init(): void {
  (window as any).admini = {
    "sortableEdit": sortableEdit
  };
  menuInit();
  flashInit();
  linkInit();
  editorInit();
  sortableInit();
}

document.addEventListener("DOMContentLoaded", init);
