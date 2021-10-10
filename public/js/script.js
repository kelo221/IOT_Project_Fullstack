"use strict";

let systemIsAutomatic = true


window.addEventListener('DOMContentLoaded', (event) => {

    const homeDiv = document.getElementById("homeContent")
    const graphDiv = document.getElementById("graphContent")
    graphDiv.style.display = "none"
    const homeButton = document.getElementById("homeButton")
    const graphButton = document.getElementById("graphButton")
    const graphContainer =   document.getElementById("graphObject")
    const clearFanData = document.getElementById("clearFanData")

    //  Home button handling
    homeButton.addEventListener("click", () => {
        console.log("homeButton clicked.")
        if (homeDiv.style.display === "none") {
            homeDiv.style.display = "block"
            graphDiv.style.display = "none"
        }
    });
    //  Home button handling END

    // Database button
    clearFanData.addEventListener("click", () => {
        console.log("database button pressed")
        reloadGraph()
    });
    // Database button END

    //  Graph  handling
    function yourFunction(){
        reloadGraph()

        setTimeout(yourFunction, 5000);
    }
    yourFunction();

    function getGraphSize(){
        return (((window.innerWidth)/(document.getElementById('graphContent').clientHeight)).toFixed(2)).toString()
    }

    function setGraphSize(){
        graphContainer.style.transform="scale("+getGraphSize() +")"
    }


    function reloadGraph(){
        graphContainer.contentWindow.location.reload(true);
        //  console.log(getGraphSize())
        console.log("reloaded graph")
    }

    graphContainer.style.transform="scale(0.1)"
    graphContainer.style.transform="scale("+getGraphSize() +")"

    window.onresize = setGraphSize
    //  Graph  handling END


    //  Graph button handling
    graphButton.addEventListener("click", () => {
        console.log("graphButton clicked.")
        if (graphDiv.style.display === "none") {
            homeDiv.style.display = "none"
            graphDiv.style.display = "block"
        }
        graphContainer.style.transform="scale("+getGraphSize() +")"
    });
    //  Graph button handling END


});

window.onload = function () {

    const gaugeSvg = document.getElementById('fanGauge').contentDocument
    const pointer = gaugeSvg.getElementById('pointer')
    // const pointer = document.getElementById('pressurePointer')
    const modeSwitch = document.getElementById("switchImage")

    const errorMessageButton = document.getElementById("errorMessage")
    const errorContainer = document.getElementById("errorContainer")

    // Fan Error Message
    errorMessageButton.addEventListener("click", () => {
        console.log("errorMessageButton clicked.")
        // errorContainer.style.display = "none"
        errorContainer.classList.toggle('fade')
        errorContainer.style.display = "none"

    });
    // Fan Error Message END


    pointer.style.transition = "all 0.25s"

    console.log(pointer)

    // Mode switch button
    modeSwitch.addEventListener("click", () => {
        console.log("switchImage clicked.")
        if (systemIsAutomatic) {
            modeSwitch.src = "img/switchMpink.png"
            systemIsAutomatic = false
        } else {
            modeSwitch.src = "img/switchA.png"
            systemIsAutomatic = true
        }
    });
    // Mode switch button END


};

