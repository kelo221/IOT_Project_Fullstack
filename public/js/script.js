"use strict";

let systemIsAutomatic = true

window.onload = function() {
    const homeDiv = document.getElementById("homeContent")
    const graphDiv = document.getElementById("graphContent")
    graphDiv.style.display = "none"
    const homeButton = document.getElementById("homeButton")
    const graphButton = document.getElementById("graphButton")
    const gaugeSvg = document.getElementById('fanGauge').contentDocument
    const pointer = gaugeSvg.getElementById('pointer')
    const modeSwitch = document.getElementById("switchImage")

    //  Logo fan handling
    const fan = document.getElementById("fan")
    let currentRot = 0
    fan.style.transition = "all 0.25s"

    function showTime() {
        fan.style.transform = "rotate(" + currentRot + "deg)"
        pointer.style.transform =  "rotate(" + currentRot + "deg)"
       // pointer.style.transform = "translate(-87px, -67px)"
        currentRot += 20
    }

    setInterval(showTime, 100)
    //  Logo fan handling


    //  Home button handling
    homeButton.addEventListener("click", () => {
        console.log("homeButton clicked.")
        if (homeDiv.style.display === "none") {
            homeDiv.style.display = "block"
            graphDiv.style.display = "none"
        }
    });
    //  Home button handling


    //  Graph button handling
    graphButton.addEventListener("click", () => {
        console.log("graphButton clicked.")
        if (graphDiv.style.display === "none") {
            homeDiv.style.display = "none"
            graphDiv.style.display = "block"
        }
    });
    //  Graph button handling


        // pointer
/*        pointer.setAttribute('transform-origin', '0 20');
        pointer.setAttribute("transform", "rotate(100)");*/



        // pointer
    // pointer.style.transformOrigin = '65% 58%';

    pointer.style.transition = "all 0.25s"
//    pointer.style.transformOrigin = '70% 80%'
    pointer.style.transformOrigin = '310px 275px'

    //pointer.style.transform =  "rotate(10deg)"
    //pointer.style.transform = "translate(-87px, -67px)"
    console.log(pointer)

    // Mode switch button
    modeSwitch.addEventListener("click", () => {
        console.log("switchImage clicked.")
        if (systemIsAutomatic) {
            modeSwitch.src = "img/switchM.png"
            systemIsAutomatic = false
        } else {
            modeSwitch.src = "img/switchA.png"
            systemIsAutomatic = true
        }
    });
    // Mode switch button
};


