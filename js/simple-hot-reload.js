/**
 * Hot reload feature.
 * Version 1.0
 * Author: https://github.com/64lines/
 */
setInterval(() => {
  fetch('/')
  .then(response => response.text())
  .then(html => {
    const doc = new DOMParser().parseFromString(html, "text/html");
    document.querySelector('body').innerHTML = doc.querySelector('body').innerHTML;
  })
  .catch(error => {
    console.error('Failed to fetch page: ', error)
  })
}, 1000);