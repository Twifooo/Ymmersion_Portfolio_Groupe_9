window.addEventListener('scroll', function() {
    var scrollPosition = window.scrollY;
    var overlay = document.querySelector('.parallax-overlay');
    var opacity = scrollPosition / (document.documentElement.scrollHeight - window.innerHeight);
    overlay.style.opacity = Math.min(opacity, 1);
});
