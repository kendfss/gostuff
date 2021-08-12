// SPDX-License-Identifier: Unlicense OR MIT

(function() {
// START OMIT
var root = document.getElementById("list");

var list = document.createElement("div");
list.style = "overflow-y: scroll; max-height: 200px;"
root.appendChild(list);

var factor = 1;

function addRows() {
	for (var i = 0; i < 1000; i++) {
		var row = document.createElement("p");
		row.innerText = "Row #" + i*factor;
		list.appendChild(row);
	}
}
addRows()

var btn = document.createElement("button");
btn.innerText = "Multiply by 10";
btn.addEventListener("click", function() {
	while (list.firstChild) {
		list.firstChild.remove();
	}
	factor *= 10;
	addRows();
});
root.appendChild(btn);
//END OMIT
})()
