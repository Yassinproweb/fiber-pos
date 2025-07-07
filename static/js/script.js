let orderStatus = document.getElementById('status');
let scolor = document.querySelectorAll('.scolor');
let sback = document.querySelector('.sback');

function themeOrder() {
  if (orderStatus.innerText === " Placed") {
    sback.textContent = "bg-pos_plc";
    scolor.forEach(color => {
      color.textContent = "text-pos_plc";
    });
  } else if (orderStatus.innerText === " Canceled") {
    sback.textContent = "bg-pos_cld";
    scolor.forEach(color => {
      color.textContent = "text-pos_cld";
    });
  } else if (orderStatus.innerText === " Ready") {
    sback.textContent = "bg-pos_rdy";
    scolor.forEach(color => {
      color.textContent = "text-pos_rdy";
    });
  } else if (orderStatus.innerText === " Preparing") {
    sback.textContent = "bg-pos_pdg";
    scolor.forEach(color => {
      color.textContent = "text-pos_pdg";
    });
  } else if (orderStatus.innerText === " Transit") {
    sback.textContent = "bg-pos_trs";
    scolor.forEach(color => {
      color.textContent = "text-pos_trs";
    });
  } else {
    sback.textContent = "bg-pos_dlv";
    scolor.forEach(color => {
      color.textContent = "text-pos_dlv";
    });
  }
}

window.addEventListener("DOMContentLoaded", themeOrder());
