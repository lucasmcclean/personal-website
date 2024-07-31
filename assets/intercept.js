"use strict";

const tabs = Array.from(document.getElementsByClassName("tab"));
let mostRecent = null;

window.onload = function() {
  tabs.forEach((tab) => {
    console.log(tab.id);
    htmx.on("#" + tab.id, "htmx:beforeRequest", (e) => {
      console.log(e.srcElement, mostRecent);
      if (e.srcElement == mostRecent) {
        e.preventDefault();
      }
      mostRecent = e.srcElement;
    });
  });
}

