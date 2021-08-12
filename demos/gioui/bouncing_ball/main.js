// SPDX-License-Identifier: Unlicense OR MIT

(function() {
// START OMIT
var root = document.getElementById("bounce");
root.style = "display: flex; position: relative; height: 200px";

var ball = document.createElement("div");
ball.style = "background-color: green; position: absolute; width: 40px; height: 40px; border-radius: 50%";
ball.style.bottom = 0;

var btn = document.createElement("button");
btn.style = "margin: auto";
btn.innerText = "Throw ball";
root.appendChild(btn);

var ballVisible = false;
var now;
var animHandle;

btn.addEventListener("click", function() {
	ballVisible = !ballVisible;
	if (ballVisible) {
		root.appendChild(ball);
		now = performance.now();
		btn.innerText = "Hide ball";
		animate(now);
	} else {
		cancelAnimationFrame(animHandle);
		btn.innerText = "Throw ball";
		ball.remove();
	}
});

function animate(time) {
	var v0 = 50.0;
	var g = -9.81;
	var t = (time - now)/1000;
	t *= 10;
	var t1 = -v0 / (.5 * g);
	var t = t % t1;
	var y = v0*t + .5*g*t*t;

	ball.style.bottom = y+"px";

	animHandle = window.requestAnimationFrame(animate);
}
//END OMIT
})()
