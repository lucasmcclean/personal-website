"use strict";

const tabs = Array.from(document.getElementsByClassName("tab"));
const links = Array.from(document.getElementsByClassName("link"));
let mostRecent = null;

window.onload = function() {
  tabs.forEach((tab) => {
    htmx.on("#" + tab.id, "htmx:beforeRequest", (e) => {
      if (e.srcElement == mostRecent) {
        e.preventDefault();
      }
      mostRecent = e.srcElement;
    });
  });
}

