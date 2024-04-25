document.addEventListener("DOMContentLoaded", function(event) { 
    var panels = document.querySelectorAll('.panel');

    panels.forEach(function(panel, index) {
      setTimeout(function() {
        panel.style.display = 'block';
        panel.style.right = '50%';
      }, 500 * index); // Задержка увеличивается для каждого следующего дива
    });
  });
