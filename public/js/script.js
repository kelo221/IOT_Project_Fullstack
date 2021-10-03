"use strict";

window.addEventListener('DOMContentLoaded', (event) => {
    let homeDiv = document.getElementById("homeContent");
    let graphDiv = document.getElementById("graphContent");
    const homeButton =  document.getElementById("homeButton")
    const graphButton =  document.getElementById("graphButton")
    graphDiv.style.display = "none";

    //  logo fan handling
    const fan =  document.getElementById("fan")
    let currentRot = 0
    fan.style.transition = "all 0.25s";

    function showTime() {
    fan.style.transform= "rotate(" + currentRot + "deg)"
        currentRot+=20
    }

    setInterval(showTime, 100);
    //  logo fan handling


    //  home button handling
    homeButton.addEventListener("click", () => {
        console.log("homeButton clicked.");
        if (homeDiv.style.display === "none") {
            homeDiv.style.display = "block";
            graphDiv.style.display = "none";
        }
    });
    //  home button handling


    //  graph button handling
    graphButton.addEventListener("click", () => {
        console.log("graphButton clicked.");
        if (graphDiv.style.display === "none") {
            homeDiv.style.display = "none";
            graphDiv.style.display = "block";
        }
    });
    //  graph button handling

});

