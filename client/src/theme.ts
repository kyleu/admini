const keys: string[] = []

export function themeInit() {
  const x = document.querySelectorAll(".color-var");
  if (x.length > 0) {
    for (const el of Array.from(x)) {
      const i = (el as HTMLInputElement)
      const v = i.dataset["var"] as string;
      const m = i.dataset["mode"] as string;
      keys.push(v);
      console.log(v);
      if (!v || v.length === 0) {
        continue;
      }
      i.oninput = function () {
        console.log(m + "/" + v + "!!!");
        document.documentElement.style.setProperty("--" + v, i.value);
      }
    }
  }
}

export function themeReset() {
  for (const k of keys) {
    document.documentElement.style.removeProperty("--" + k);
  }
}
