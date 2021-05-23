export function sortableInit() {
  const Sortable = (window as any).Sortable;
  if (Sortable) {
    for (const dd of Array.from(document.getElementsByClassName("drag-container"))) {
      let l = dd.querySelector(".l");
      if (!l) {
        l = dd;
      }
      const lOpts = {group: {name: 'nested'}, onAdd: onAdd, animation: 150, fallbackOnBody: true, swapThreshold: 0.65};
      function onAdd(ev: Event) {
        const i = (ev as any).item as HTMLElement;
        new Sortable(i.querySelector(".container"), lOpts);
        (i.querySelector(".remove") as HTMLElement).onclick = function() { remove(dd, i); };
        update(dd);
      }
      for (const c of Array.from(l.getElementsByClassName('container'))) {
        new Sortable(c, lOpts);
      }
      for (const rem of Array.from(l.getElementsByClassName("remove"))) {
        (rem as HTMLElement).onclick = function() { remove(dd, rem.parentElement?.parentElement!); };
      }

      const r = dd.querySelector(".r");
      if (r) {
        const rOpts = {group: {name: 'nested', pull: "clone", put: false}, animation: 150, fallbackOnBody: true, swapThreshold: 0.65, sort: false};
        for (const c of Array.from(r.getElementsByClassName('container'))) {
          new Sortable(c, rOpts);
        }
      }
      update(dd)
    }
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
  const js = JSON.stringify(readContainer(trackedEl));
  if (origEl) {
    if (origEl.value.length === 0) {
      origEl.value = js;
    }
    const aEl = document.querySelector(".drag-actions") as HTMLElement;
    if (origEl.value === js) {
      aEl.classList.add('no-changes');
    } else {
      aEl.classList.remove('no-changes');
    }
  }

  sEl.value = js;
}

interface Item {
  k: string
  p: string
  c?: Item[]
}

function readContainer(c: Element): Item[] {
  const ret: Item[] = [];
  for (const i of Array.from(c.children)) {
    if (i.classList.contains("item")) {
      const item = readItem(i as HTMLElement);
      if (item) {
        ret.push(item);
      }
    }
  }
  return ret;
}

function readItem(i: HTMLElement): Item | undefined {
  var ret: Item = {
    k: i.dataset.key as string,
    p: i.dataset.originalPath as string
  };
  for (const x of Array.from(i.children)) {
    if (x.classList.contains("container")) {
      const kids = readContainer(x);
      if (kids.length > 0) {
        ret.c = kids;
      }
    }
  }
  return ret;
}
