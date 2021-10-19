"use strict";

let systemIsAutomatic = true

window.addEventListener('DOMContentLoaded', (event) => {

    const homeDiv = document.getElementById("homeContent")
    const graphDiv = document.getElementById("graphContent")
    graphDiv.style.display = "none"

    const homeButton = document.getElementById("homeButton")
    const graphButton = document.getElementById("graphButton")
    const graphContainer = document.getElementById("graphObject")
    const clearFanData = document.getElementById("clearFanData")

    const pressureInput = document.getElementById("pressureInputBox")
    const fanSpeedInput = document.getElementById("fanSpeedInputBox")

    const pressureInputButton = document.getElementById("pressureDataButton")
    const fanSpeedInputButton = document.getElementById("fanDataButton")
    const modeSwitch = document.getElementById("switchImage")

    const errorMessageButton = document.getElementById("errorMessage")
    const errorContainer = document.getElementById("errorContainer")


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
    function yourFunction() {
        reloadGraph()

        setTimeout(yourFunction, 5000);
    }

    yourFunction();

    function getGraphSize() {
        return (((window.innerWidth) / (document.getElementById('graphContent').clientHeight)).toFixed(2)).toString()
    }

    function setGraphSize() {
        graphContainer.style.transform = "scale(" + getGraphSize() + ")"
    }


    function reloadGraph() {
        graphContainer.contentWindow.location.reload(true);
        //  console.log(getGraphSize())
        console.log("reloaded graph")
    }

    graphContainer.style.transform = "scale(0.1)"
    graphContainer.style.transform = "scale(" + getGraphSize() + ")"

    window.onresize = setGraphSize
    //  Graph  handling END

    //  Graph button handling
    graphButton.addEventListener("click", () => {
        console.log("graphButton clicked.")
        if (graphDiv.style.display === "none") {
            homeDiv.style.display = "none"
            graphDiv.style.display = "block"
        }
        graphContainer.style.transform = "scale(" + getGraphSize() + ")"
    });
    //  Graph button handling END

    // Mode switch button
    modeSwitch.addEventListener("click", () => {

        console.log("switchImage clicked.")
        if (systemIsAutomatic) {
            modeSwitch.src = "img/switchMpink.png"
            systemIsAutomatic = false

            pressureInputButton.setAttribute('disabled', null);
            pressureInput.setAttribute('disabled', null);
            fanSpeedInput.removeAttribute('disabled')
            fanSpeedInputButton.removeAttribute('disabled')
        } else {
            modeSwitch.src = "img/switchA.png"
            systemIsAutomatic = true
            sendUserSettings(null, null)

            fanSpeedInput.setAttribute('disabled', null);
            fanSpeedInputButton.setAttribute('disabled', null);
            pressureInput.removeAttribute('disabled')
            pressureInputButton.removeAttribute('disabled')
        }
    });
    // Mode switch button END

    // Fan speed input
    fanSpeedInputButton.addEventListener("click", () => {
        const value = parseInt(fanSpeedInput.value)
        console.log(value)
        sendUserSettings(value)
    });
    // Fan speed input END

    // Pressure speed input
    pressureInputButton.addEventListener("click", () => {
        const value = parseInt(pressureInput.value)
        console.log(value)
        sendUserSettings(value)
    });
    // Pressure speed input END

    // Fan Error Message
    errorMessageButton.addEventListener("click", () => {
        console.log("errorMessageButton clicked.")
        // errorContainer.style.display = "none"
        errorContainer.classList.toggle('fade')
        errorContainer.style.display = "none"

    });
    // Fan Error Message END

});

//const gaugeSvg = document.getElementById('pressureGauge').contentDocument
// const pointer = gaugeSvg.getElementById('pointer')
// const pointer = document.getElementById('pressurePointer')
//pointer.style.transform ="translate(221.245px,207.370756px) rotate(-600deg)"
// pointer.style.transition = "all 0.25s"
// pointer.style.transform console.log(pointer)


function sendUserSettings(value) {
    let json;

    // When system is in an automatic state only pressure is sent, and in manual speed is.
    if (systemIsAutomatic) {
        json = JSON.stringify({Auto: systemIsAutomatic, pressure: value});
    } else {
        json = JSON.stringify({Auto: systemIsAutomatic, speed: value});
    }

    axios({
        method: "post",
        url: "http://localhost:8080/getUserSettings",
        data: json,
        headers: {
            "Content-Type": "application/json",
        },

    })
        .then(function (response) {
            //handle success
            console.log(response);
        })
        .catch(function (response) {
            //handle error
            console.log(response);
        });
}

