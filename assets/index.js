"use strict";

window.onload = function() {
  const tabs = Array.from(document.getElementsByClassName("tab"));
  let mostRecent = "";

  tabs.forEach((tab) => {
    tab.addEventListener("click", activateSub);
    tab.addEventListener("mouseover", activateSub);

    // If the description is already active don't send a network request
    htmx.on("#" + tab.id, "htmx:beforeRequest", (e) => {
      let id = e.srcElement.id;
      console.log(mostRecent, id);
      if (id == mostRecent) {
        e.preventDefault();
      } else {
        mostRecent = id;
      }
    });
  });

  // Initialize page
  activateSub({ srcElement: { id: "tl" } })
}

let previous = null;

function activateSub(e) {
  let tabName = e.srcElement.id;
  if (tabName != previous) {
    deactivateSub(previous);
    document.getElementById(tabName + "-sub").style.animation = ".3s ease both select-link";
    previous = tabName;
  }
}

function deactivateSub(tabName) {
  if (tabName != null) {
    document.getElementById(tabName + "-sub").style.animation = "";
  }
}


