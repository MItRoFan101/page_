document.addEventListener("DOMContentLoaded", function() {
    var columns = document.querySelectorAll('.row.mb-2 .col-md-6');
    columns.forEach(function(column) {
        column.classList.add('show');
    });
});



/* document.getElementById('showMessageBox').addEventListener('mouseenter', function() {
    var rectangle = document.getElementById('messageBox');
    var box = this.getBoundingClientRect();
    
    rectangle.style.width = box.width + 'px';
    rectangle.style.height = box.height + 'px';
    rectangle.style.opacity = 1;
}); */


