const selected = "--selected";

export function editorInit() {
  let editorCache: { [key: string]: string; } = {};
  let selectedCache: { [key: string]: HTMLInputElement; } = {};
  const x = document.querySelectorAll(".editor");
  for (const n of x) {
    const frm = n as HTMLFormElement;
    const buildCache = () => {
      editorCache = {};
      selectedCache = {};
      for (const el of frm.elements) {
        const input = el as HTMLInputElement;
        if (input.name.length > 0) {
          if (input.name.endsWith(selected)) {
            selectedCache[input.name] = input;
          } else {
            if ((input.type != "radio") || input.checked) {
              editorCache[input.name] = input.value;
            }
            const evt = () => {
              selectedCache[input.name + selected].checked = editorCache[input.name] != input.value;
            };
            input.onchange = evt;
            input.onkeyup = evt;
          }
        }
      }
      console.log(editorCache);
    }
    frm.onreset = buildCache;
    buildCache();
  }
}
