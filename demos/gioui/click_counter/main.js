// SPDX-License-Identifier: Unlicense OR MIT

(function() {
// START OMIT
var count = 0;

var root = document.getElementById("clickcounter");

var label = document.createElement("h3");
label.innerText = "Number of clicks: 0";
root.appendChild(label);

var btn = document.createElement("button");
btn.innerText = "Click Me!";
btn.addEventListener("click", function() {
	count++;
	label.innerText = "Number of clicks:" + count;
});
root.appendChild(btn);
//END OMIT
})()
