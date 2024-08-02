"use strict";

window.onload = function() {
  const links = Array.from(document.getElementsByClassName("link"));
  let mostRecent = "";

  links.forEach((link) => {
    link.addEventListener("focus", select);
    link.addEventListener("mouseover", select);

    // If the description is already active don't send a network request
    htmx.on("#" + link.id, "htmx:beforeRequest", (e) => {
      let id = e.srcElement.id;
      if (id == mostRecent) {
        e.preventDefault();
      } else {
        mostRecent = id;
      }
    });
  });

  // Initialize page
  select({ srcElement: { id: "tl" } })
}

let previous = null;

function select(e) {
  let linkName = e.srcElement.id;
  if (linkName != previous) {
    deselect(previous);
    document.getElementById(linkName).style.animation = ".3s ease both select-link";
    document.getElementById(linkName + "-tab").style.animation = ".3s ease both select-tab";
    previous = linkName;
  }
}

function deselect(linkName) {
  if (linkName != null) {
    document.getElementById(linkName).style.animation = "";
    document.getElementById(linkName + "-tab").style.animation = "";
  }
}


