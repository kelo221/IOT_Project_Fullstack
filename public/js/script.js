window.addEventListener('DOMContentLoaded', (event) => {

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

});

