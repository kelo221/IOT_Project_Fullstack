window.addEventListener('DOMContentLoaded', (event) => {

    //  logo fan handling
    const fan =  document.getElementById("fan")
    let currentRot = 0

    function showTime() {
    fan.style.transform= "rotate(" + currentRot + "deg)"
        currentRot++
    }

    setInterval(showTime, 1);
    //  logo fan handling

});

