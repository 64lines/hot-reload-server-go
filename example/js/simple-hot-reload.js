/**
 * Hot reload feature.
 * Version 1.0
 * Author: https://github.com/64lines/
 */
(() => {
  async function runHotReload() {
    const url = `http://localhost:${HOT_RELOAD_PORT || "8080"}/hot-reload`;
    try {
      const response = await fetch(url);

      if (!response.ok) {
        console.log(`Response status: ${response.status}`);
      }

      const json = await response.json();
   
      if (json.update) {
        window.location.reload();
      }
    } catch (error) {
      if (error instanceof TypeError) {
        console.log("Network error:", error.message);
      } else {
        console.log(error.message)
      }
    }
  }

  requestAnimationFrame(() => setInterval(() => {
    runHotReload()
  }, 500));
})();
