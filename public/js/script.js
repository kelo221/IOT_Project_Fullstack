"use strict";

let systemIsAutomatic = true

function convertEpochToSpecificTimezone(timeEpoch, offset) {
    let d = new Date(timeEpoch);
    let utc = d.getTime() + (d.getTimezoneOffset() * 60000);  //This converts to UTC 00:00
    let nd = new Date(utc + (3600000 * offset));
    return nd.toLocaleString();
}

async function clearDB() {
    try {
        return await axios({
            url: 'http://localhost:8080/clearDatabase',
            method: 'Delete',
            timeout: 8000,
            headers: {
                'Content-Type': 'application/json',
            }
        })
    } catch (err) {
        console.error(err);
    }
}


async function getGaugeData(gaugeType) {
    try {
        let res = await axios({
            url: 'http://localhost:8080/getGaugeData',
            method: 'get',
            timeout: 8000,
            headers: {
                'Content-Type': 'application/json',
            }
        })
        if (Object.keys(res.data).length !== 0) {

            //console.log(Object.keys(res.data).length)

            if (gaugeType === "speed")
                return res.data.speed
            if (gaugeType === "pressure")
                return res.data.pressure
            else
                return 0
        } else {
            return 0
        }

    } catch (err) {
        console.error(err);
    }
}


function generateTable() {
    // console.log("asking for users")

    axios({
        method: "get",
        url: "http://localhost:8080/userLogs",
        data: null,
    })
        .then(function (response) {
            //handle success

            let string1 = JSON.stringify(response);
            let parsed = JSON.parse(string1);

            let teachTotal=0
            let xTotal=0
            let vTotal=0
            let currentCount=0

            for (let i = 0; i < parsed.data.length; i++) {

                // This could have been in the database
                if (parsed.data[i].user === "teach"){
                    teachTotal++
                    currentCount = teachTotal
                }
                if (parsed.data[i].user === "x"){
                    xTotal++
                    currentCount = xTotal
                }
                if (parsed.data[i].user === "v"){
                    vTotal++
                    currentCount = vTotal
                }

                let tr = document.createElement('tr')

                let td1 = document.createElement('th')
                let td2 = document.createElement('td')
                let td3 = document.createElement('td')
                let td4 = document.createElement('td')
                let text1 = document.createTextNode((i+1).toString())
                let text2 = document.createTextNode(parsed.data[i].user)
                let text3 = document.createTextNode(convertEpochToSpecificTimezone(parsed.data[i].time, +3))
                let text4 = document.createTextNode(currentCount.toString())
                td1.appendChild(text1)
                td2.appendChild(text2)
                td3.appendChild(text3)
                td4.appendChild(text4)
                tr.appendChild(td1)
                tr.appendChild(td2)
                tr.appendChild(td3)
                tr.appendChild(td4)

                document.getElementById("logContent").appendChild(tr)

            }

        })
        .catch(function (response) {
            //handle error
            console.log(response);
        });


}

window.addEventListener('DOMContentLoaded', (event) => {

    const homeDiv = document.getElementById("homeContent")
    const graphDiv = document.getElementById("graphContent")
    const loginDiv = document.getElementById("loginContent")
    graphDiv.style.display = "none"
    loginDiv.style.display = "none"


    const homeButton = document.getElementById("homeButton")
    const graphButton = document.getElementById("graphButton")
    const loginsButton = document.getElementById("loginButton")
    const graphContainer = document.getElementById("graphObject")
    const clearFanData = document.getElementById("clearFanData")

    const pressureInput = document.getElementById("pressureInputBox")
    const fanSpeedInput = document.getElementById("fanSpeedInputBox")

    const pressureInputButton = document.getElementById("pressureDataButton")
    const fanSpeedInputButton = document.getElementById("fanDataButton")
    const modeSwitch = document.getElementById("switchImage")

    const errorMessageButton = document.getElementById("errorMessage")
    const errorContainer = document.getElementById("errorContainer")

    const gaugeGraphId = document.getElementsByTagName('canvas')

    // Gauge Handling

    const pressureGauge = Gauge(
        document.getElementById("pressureGauge"),
        {
            min: 0,
            max: 120,
            dialStartAngle: 180,
            dialEndAngle: 0,
            value: 81,
            viewBox: "0 0 100 57",
            color: function (value) {
                if (value < 80) {
                    return "#78c985";
                } else if (value < 100) {
                    return "#f5aa49";
                } else if (value > 100) {
                    return "#e48894";
                }
            }
        }
    );

    const fanGauge = Gauge(
        document.getElementById("fanGauge"),
        {
            min: 0,
            max: 100,
            dialStartAngle: 180,
            dialEndAngle: 0,
            value: 81,
            viewBox: "0 0 100 57",
            color: function (value) {
                if (value < 60) {
                    return "#78c985";
                } else if (value < 80) {
                    return "#f5aa49";
                } else if (value > 80) {
                    return "#e48894";
                }
            }
        }
    );


    function gaugeUpdater() {
        getGaugeData("speed")
            .then(res =>
                fanGauge.setValueAnimated(res, 3)
            )
        getGaugeData("pressure")
            .then(res =>
                pressureGauge.setValueAnimated(res, 3)
            )

        setTimeout(gaugeUpdater, 5000);
    }

    gaugeUpdater()


    // Gauge Handling END


    function animationReset() {
        pressureInput.style.animation = 'none';
        pressureInput.offsetHeight; /* trigger reflow */
        pressureInput.style.animation = null;

        fanSpeedInput.style.animation = 'none';
        fanSpeedInput.offsetHeight; /* trigger reflow */
        fanSpeedInput.style.animation = null;
    }


    //  Home button handling
    homeButton.addEventListener("click", () => {
        console.log("homeButton clicked.")
        homeDiv.style.display = "block"
        graphDiv.style.display = "none"
        loginDiv.style.display = "none"
    });
    //  Home button handling END

    //  Graph button handling
    graphButton.addEventListener("click", () => {
        console.log("graphButton clicked.")
        homeDiv.style.display = "none"
        graphDiv.style.display = "block"
        loginDiv.style.display = "none"
        //   graphContainer.style.transform = "scale(" + getGraphSize() + ")"
        //  console.log(getGraphSize())
    });
    //  Graph button handling END

    //  login button handling
    loginsButton.addEventListener("click", () => {
        console.log("loginsButton clicked.")
        homeDiv.style.display = "none"
        graphDiv.style.display = "none"
        loginDiv.style.display = "block"
    });
    //  Home button handling END

    // Database button
    clearFanData.addEventListener("click", () => {
        console.log("database button pressed")
        clearDB().then(r => console.log(r))
        reloadGraph()
    });

    // Database button END


    function reloadGraph() {
       // console.log(gaugeGraphId[0])
       // gaugeGraphId[0]
       // let context = gaugeGraphId[0].getContext('2d');
      //  context.clearRect(0, 0, gaugeGraphId[0].width, gaugeGraphId[0].height); //clear html5 canvas
       // console.log(document.getElementById("zr_0"))
        //document.getElementById("graphCont").contentWindow.location.reload(true)
    }

    function graphUpdater() {
        reloadGraph()
        setTimeout(graphUpdater, 5000);
    }

    graphUpdater()


    // Mode switch button
    modeSwitch.addEventListener("click", () => {

        console.log("switchImage clicked.")
        if (systemIsAutomatic) {
            modeSwitch.src = "public/img/switchMpink.png"
            systemIsAutomatic = false
            sendUserSettings(null, null)
            pressureInputButton.setAttribute('disabled', null);
            pressureInput.setAttribute('disabled', null);
            fanSpeedInput.removeAttribute('disabled')
            fanSpeedInputButton.removeAttribute('disabled')
        } else {
            modeSwitch.src = "public/img/switchA.png"
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
        if (value >= 0 && value <= 100) {
            sendUserSettings(value)
            fanSpeedInput.className = "input is-success"
        } else {
            fanSpeedInput.className = "input is-danger"
            fanSpeedInput.style.animation = "shake 0.5s"
            setTimeout(animationReset, 500)
        }

        fanSpeedInput.value = null
    });
    // Fan speed input END

    // Pressure speed input
    pressureInputButton.addEventListener("click", () => {
        const value = parseInt(pressureInput.value)
        console.log(value)
        if (value >= 0 && value <= 120) {
            sendUserSettings(value)
            pressureInput.className = "input is-success"
        } else {
            pressureInput.className = "input is-danger"
            pressureInput.style.animation = "shake 0.5s"
            setTimeout(animationReset, 500)
        }


        pressureInput.value = null
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


    generateTable()

});
//const gaugeSvg = document.getElementById('pressureGauge').contentDocument
// const pointer = gaugeSvg.getElementById('pointer')
// const pointer = document.getElementById('pressurePointer')
//pointer.style.transform ="translate(221.245px,207.370756px) rotate(-600deg)"
// pointer.style.transition = "all 0.25s"
// pointer.style.transform console.log(pointer)


function sendUserSettings(value = 0) {
    let json;

    console.log("sending user settings")

    // When system is in an automatic state only pressure is sent, and in manual speed is.
    if (systemIsAutomatic) {
        json = JSON.stringify({auto: systemIsAutomatic, pressure: value});
    } else {
        json = JSON.stringify({auto: systemIsAutomatic, speed: value});
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

