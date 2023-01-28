// Content managed by Project Forge, see [projectforge.md] for details.
import "./client.css"
import {menuInit} from "./menu";
import {modeInit} from "./mode";
import {flashInit} from "./flash";
import {linkInit} from "./link";
import {timeInit} from "./time";
import {autocompleteInit} from "./autocomplete";
import {modalInit} from "./modal";
import {tagsInit} from "./tags";
import {editorInit} from "./editor";
import {themeInit} from "./theme";
import {appInit} from "./app";

declare global {
  interface Window {
    "admini": {
      relativeTime: (time: string, el?: HTMLElement) => string;
      autocomplete: (el: HTMLInputElement, url: string, field: string, title: (x: any) => string, val: (x: any) => string) => void;
      setSiblingToNull: (el: HTMLElement) => void;
      initForm: (frm: HTMLFormElement) => void;
      flash: (key: string, level: string, msg: string) => void;
      tags: (el: HTMLElement) => void;
    };
  }
}

export function init(): void {
  const [s, i] = editorInit();
  window.admini = {
    relativeTime: timeInit(),
    autocomplete: autocompleteInit(),
    setSiblingToNull: s,
    initForm: i,
    flash: flashInit(),
    tags: tagsInit()
  };
  menuInit();
  modeInit();
  linkInit();
  modalInit();
  themeInit();
  appInit();
}

document.addEventListener("DOMContentLoaded", init);
