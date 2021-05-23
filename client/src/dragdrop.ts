export function dragDropInit() {
  const lmdd = (window as any).lmdd;
  if (lmdd) {
    for (const dd of Array.from(document.getElementsByClassName("drag-container"))) {
      var opts = { containerClass: "container", draggableItemClass: "item", handleClass: "", nativeScroll: true };
      if (dd.getElementsByClassName("handle").length > 0) {
        opts.handleClass = "handle";
      }
      lmdd.set(dd, opts);
      dd.addEventListener("lmddend", function (ev) { update(dd, ev); }, false);
      for (const rem of Array.from(dd.getElementsByClassName("remove"))) {
        const el = rem.parentElement?.parentElement!;
        if (!el.classList.contains("lmdd-clonner")) {
          (rem as HTMLElement).onclick = function() { remove(dd, el); };
        }
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
    if (to && to.container && to.container.children) {
      const el = to.container.children.item(to.index) as HTMLElement
      for (const rem of Array.from(el.getElementsByClassName("remove"))) {
        if ((rem as HTMLElement).onclick === null) {
          (rem as HTMLElement).onclick = function() { remove(dd, el); };
        }
      }
    }
  }
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
