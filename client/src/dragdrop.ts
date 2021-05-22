export function dragDropInit() {
  const lmdd = (window as any).lmdd;
  if (lmdd) {
    for (const dd of Array.from(document.getElementsByClassName("drag-container"))) {
      var opts = { containerClass: "container", draggableItemClass: "item", handleClass: "" };
      if (dd.getElementsByClassName("handle").length > 0) {
        opts.handleClass = "handle";
      }
      lmdd.set(dd, opts);
      dd.addEventListener("lmddend", function (ev) { update(dd, ev); }, false);
      for (const rem of Array.from(dd.getElementsByClassName("remove"))) {
        (rem as HTMLElement).onclick = function() { remove(dd, rem.parentElement?.parentElement!); };
      }

      update(dd, null);
    }
  }
}

function remove(dd: Element, rem: Element) {
  rem.remove();
  update(dd, null);
}

function update(dd: Element, ev: Event | null) {
  if(ev) {
    const to = (ev as any).detail.to;
    if (to) {
      console.log(to.index, to.container.children.item(to.index));
      const el = to.container.children.item(to.index) as HTMLElement
      for (const rem of Array.from(el.getElementsByClassName("remove"))) {
        if ((rem as HTMLElement).onclick === null) {
          (rem as HTMLElement).onclick = function() { remove(dd, el); };
        }
      }
    }
  }
  const stateEls = dd.getElementsByClassName("drag-state");
  if (stateEls.length !== 1) {
    return;
  }
  const sEl = (stateEls.item(0) as HTMLInputElement);
  const origEls = dd.getElementsByClassName("drag-state-original");
  const tracked = dd.getElementsByClassName("tracked");
  if (tracked.length > 1) {
    throw "too many tracked drag/drops";
  }
  if (tracked.length === 1) {
    const el = tracked.item(0) as HTMLElement;
    const js = JSON.stringify(readContainer(el));
    if (origEls.length === 1) {
      const oEl = (origEls.item(0) as HTMLInputElement);
      if (oEl.value.length === 0) {
        oEl.value = js;
      }
      const buttonContainers = dd.getElementsByClassName("action-buttons");
      if (buttonContainers.length === 1) {
        (buttonContainers.item(0) as HTMLElement).style.display = oEl.value === js ? "none" : "block";
      }
    }

    sEl.value = js;
  }
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
