interface SortableEvent {
  item: HTMLElement;
}

interface SortableOptions {
  group: {
    name: string;
    pull?: "clone";
    put?: boolean;
  };
  handle: string;
  onAdd?: (ev: SortableEvent) => void;
  onUpdate?: () => void;
  animation: number;
  fallbackOnBody: boolean;
  swapThreshold: number;
  sort?: boolean;
}

type SortableConstructor = new (el: Element, options: SortableOptions) => unknown;

declare global {
  interface Window {
    Sortable?: SortableConstructor;
  }
}

export function sortableInit() {
  if (window.Sortable) {
    (window.admini as Record<string, unknown>).sortableEdit = sortableEdit;
    for (const dd of Array.from(document.getElementsByClassName("sortable"))) {
      sortableCreate(dd);
    }
  }
}

export function sortableEdit(el: Element) {
  while (el.parentElement && !el.classList.contains("drag-container")) {
    el = el.parentElement;
  }
  el.classList.remove("readonly");
  sortableCreate(el);
}

function sortableCreate(dd: Element) {
  const { Sortable } = window;
  if (Sortable) {
    let l = dd.querySelector(".l");
    if (!l) {
      l = dd;
    }
    const onAdd = (ev: SortableEvent) => {
      const i = ev.item;
      const container = i.querySelector(".container");
      if (container) {
        new Sortable(container, lOpts);
      }
      const removeButton = i.querySelector(".remove");
      if (removeButton instanceof HTMLElement) {
        removeButton.onclick = function () {
          remove(dd, i);
        };
      }
      update(dd);
    };
    const lOpts: SortableOptions = {
      group: { name: "nested" },
      handle: ".handle",
      onAdd: onAdd,
      onUpdate: () => update(dd),
      animation: 150,
      fallbackOnBody: true,
      swapThreshold: 0.65
    };
    for (const c of Array.from(l.getElementsByClassName("container"))) {
      new Sortable(c, lOpts);
    }
    for (const rem of Array.from(l.getElementsByClassName("remove"))) {
      if (rem instanceof HTMLElement) {
        rem.onclick = function () {
          const item = rem.parentElement?.parentElement;
          if (item) {
            remove(dd, item);
          }
        };
      }
    }

    const r = dd.querySelector(".r");
    if (r) {
      const rOpts: SortableOptions = {
        group: { name: "nested", pull: "clone", put: false },
        handle: ".handle",
        animation: 150,
        fallbackOnBody: true,
        swapThreshold: 0.65,
        sort: false
      };
      for (const c of Array.from(r.getElementsByClassName("container"))) {
        new Sortable(c, rOpts);
      }
    }
    update(dd);
  }
}

function remove(dd: Element, rem: Element) {
  rem.remove();
  update(dd);
}

function update(dd: Element) {
  const sEl = document.querySelector(".drag-state") as HTMLInputElement;
  if (!sEl) {
    return;
  }
  const origEl = document.querySelector(".drag-state-original") as HTMLInputElement;
  const trackedEl = dd.querySelector(".tracked") as HTMLElement;
  const [i, count] = readContainer(trackedEl);
  const js = JSON.stringify(i);
  if (origEl) {
    const changed = origEl.value !== js;
    if (origEl.value.length === 0) {
      origEl.value = js;
    }
    const aEl = document.querySelector(".drag-actions") as HTMLElement;
    if (aEl) {
      if (changed) {
        aEl.classList.remove("no-changes");
      } else {
        aEl.classList.add("no-changes");
      }
    }
    const sEl = document.querySelector(".drag-tracked-size") as HTMLElement;
    if (sEl) {
      if (count === 1) {
        sEl.innerText = count.toString(10) + (sEl.dataset.sing ? " " + sEl.dataset.sing : "");
      } else {
        sEl.innerText = count.toString(10) + (sEl.dataset.plur ? " " + sEl.dataset.plur : "");
      }
    }
    if (changed) {
      window.onbeforeunload = function () {
        return true;
      };
    } else {
      window.onbeforeunload = null;
    }
  }

  sEl.value = js;
}

interface Item {
  k: string;
  t: string;
  p: string;
  c?: Item[];
}

function readContainer(c: Element): [Item[], number] {
  if (c.children.length === 0) {
    return [[], 0];
  }
  let count = 0;
  const ret: Item[] = [];
  for (const i of Array.from(c.children)) {
    if (i.classList.contains("item")) {
      const [item, c] = readItem(i as HTMLElement);
      if (item) {
        ret.push(item);
      }
      count += c;
    }
  }
  return [ret, count];
}

function readItem(i: HTMLElement): [Item | undefined, number] {
  let count = 1;
  const ret: Item = {
    k: i.dataset.key as string,
    t: i.dataset.title as string,
    p: i.dataset.originalPath as string
  };
  for (const x of Array.from(i.children)) {
    if (x.classList.contains("container")) {
      const [kids, c] = readContainer(x);
      if (kids.length > 0) {
        ret.c = kids;
      }
      count += c;
    }
  }
  return [ret, count];
}
